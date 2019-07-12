package rotate

func RotateLeft(a []int32, d int32) []int32 {
	n := int32(len(a))
	if n == 0 {
		return a
	}
	rotate := d % n
	switch {
	case rotate == 0:
		return a
	case len(a) <= 1:
		return a
	case len(a) == 2 && rotate%int32(2) == 0:
		return a
	default:
		for i := int32(0); i < rotate; i++ {
			a = append(a[1:], a[0])
		}
		return a
	}
}

func RotateLeftOrg(a []int32, d int32) []int32 {
	for i := int32(0); i < d; i++ {
		a = append(a[1:], a[0])
	}
	return a
}
