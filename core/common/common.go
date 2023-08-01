package common

type Vector struct {
	X int
	Y int
}

func IsEqual(vec1 Vector, vec2 Vector) bool {
	return vec1.X == vec2.X && vec1.Y == vec2.Y
}
