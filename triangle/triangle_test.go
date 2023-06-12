package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		{30001, 6, 2, UnknownTriangle},
		{4, 20001, 2, UnknownTriangle},
		{4, 2, 10001, UnknownTriangle},
		{-1, 2, 2, UnknownTriangle},
		{1, -2, 0, UnknownTriangle},
		{1, 2, -1, UnknownTriangle},
		{8, 4, 2, InvalidTriangle},
		{2, 4, 8, InvalidTriangle},
		{4, 8, 2, InvalidTriangle},
		{5, 4, 3, RightTriangle},
		{7, 5, 3, ObtuseTriangle},
		{5, 5, 3, AcuteTriangle},
		// TODO add more tests for 100% test coverage
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
