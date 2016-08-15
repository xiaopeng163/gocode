package main

import "fmt"

const WITHDRAW_LABEL = uint32(0x800000)
const ZERO_LABEL = uint32(0) // some platform uses this as withdraw label

type MPLSLabelStack struct {
	Labels []uint32
}

func (l *MPLSLabelStack) DecodeFromBytes(data []byte) error {
	labels := []uint32{}
	foundBottom := false
	for len(data) >= 3 {
		label := uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2])
		if label == WITHDRAW_LABEL || label == ZERO_LABEL {
			l.Labels = []uint32{label}
			return nil
		}
		data = data[3:]
		labels = append(labels, label>>4)
		if label&1 == 1 {
			foundBottom = true
			break
		}
	}
	if foundBottom == false {
		l.Labels = []uint32{}
		return nil
	}
	l.Labels = labels
	return nil
}

func main() {

	mpls_lable := MPLSLabelStack{}
	raw_data := []byte{0, 1, 0x41}
	err := mpls_lable.DecodeFromBytes(raw_data)
	if err == nil {
		fmt.Println(mpls_lable)
	}
}
