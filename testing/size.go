package size

// Size will return the size of a number
func Size(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a <= 100:
		return "small"
	case a > 100:
		return "big"

	}
	return "unknow"
}
