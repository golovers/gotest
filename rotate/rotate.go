package rotate

func RotateLeft(a []int32, d int32) []int32 {
	l := int32(len(a))
	if l <= 1 {
		return a
	}
	n := d % l
	switch {
	case n == 0:
		return a
	case l == 2 && n%int32(2) == 0:
		return a
	default:
		for i := int32(0); i < n; i++ {
			a = append(a[1:], a[0])
		}
		return a
	}
}

// RotateLeftOrg is a simple solution, but not good performance
func RotateLeftOrg(a []int32, d int32) []int32 {
	for i := int32(0); i < d; i++ {
		a = append(a[1:], a[0])
	}
	return a
}
