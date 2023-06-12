package triangle

type triangleType uint8

const (
	UnknownTriangle = triangleType(iota)
	InvalidTriangle
	RightTriangle
	AcuteTriangle
	ObtuseTriangle
)

func isTriangle(a, b, c int) bool {
	if a+b <= c {
		return false
	}
	if b+c <= a {
		return false
	}
	if a+c <= b {
		return false
	}
	return true
}

func getTriangleType(a, b, c int) triangleType {
	switch {
	case a > 30000:
		return UnknownTriangle
	case b > 20000:
		return UnknownTriangle
	case c > 10000:
		return UnknownTriangle
	case a <= 0:
		return UnknownTriangle
	case b <= 0:
		return UnknownTriangle
	case c <= 0:
		return UnknownTriangle
	}
	if !isTriangle(a, b, c) {
		return InvalidTriangle
	}
	if a*a == b*b + c*c {
		return RightTriangle
	} else if a*a < b*b + c*c {
		return AcuteTriangle
	} else {
		return ObtuseTriangle
	}
}
