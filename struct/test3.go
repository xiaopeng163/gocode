package main

import "encoding/binary"
import "fmt"

// DefaultRouteDistinguisher
type DefaultRouteDistinguisher struct {
	Type  uint16
	Value []byte
}

// DecodeFromBytes
func (rd *DefaultRouteDistinguisher) DecodeFromBytes(data []byte) error {
	rd.Type = binary.BigEndian.Uint16(data[0:2])
	rd.Value = data[2:8]
	return nil
}

type PmsiTunnelType uint8

const (
	PMSI_TUNNEL_TYPE_NO_TUNNEL      PmsiTunnelType = 0
	PMSI_TUNNEL_TYPE_RSVP_TE_P2MP   PmsiTunnelType = 1
	PMSI_TUNNEL_TYPE_MLDP_P2MP      PmsiTunnelType = 2
	PMSI_TUNNEL_TYPE_PIM_SSM_TREE   PmsiTunnelType = 3
	PMSI_TUNNEL_TYPE_PIM_SM_TREE    PmsiTunnelType = 4
	PMSI_TUNNEL_TYPE_BIDIR_PIM_TREE PmsiTunnelType = 5
	PMSI_TUNNEL_TYPE_INGRESS_REPL   PmsiTunnelType = 6
	PMSI_TUNNEL_TYPE_MLDP_MP2MP     PmsiTunnelType = 7
)

func (p PmsiTunnelType) String() string {
	switch p {
	case PMSI_TUNNEL_TYPE_NO_TUNNEL:
		return "no-tunnel"
	case PMSI_TUNNEL_TYPE_RSVP_TE_P2MP:
		return "rsvp-te-p2mp"
	case PMSI_TUNNEL_TYPE_MLDP_P2MP:
		return "mldp-p2mp"
	case PMSI_TUNNEL_TYPE_PIM_SSM_TREE:
		return "pim-ssm-tree"
	case PMSI_TUNNEL_TYPE_PIM_SM_TREE:
		return "pim-sm-tree"
	case PMSI_TUNNEL_TYPE_BIDIR_PIM_TREE:
		return "bidir-pim-tree"
	case PMSI_TUNNEL_TYPE_INGRESS_REPL:
		return "ingress-repl"
	case PMSI_TUNNEL_TYPE_MLDP_MP2MP:
		return "mldp-mp2mp"
	default:
		return fmt.Sprintf("PmsiTunnelType(%d)", uint8(p))
	}
}

// main function
func main() {

	rd := DefaultRouteDistinguisher{}
	fmt.Println(rd)
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	rd.DecodeFromBytes(data)
	fmt.Println(rd)
}
