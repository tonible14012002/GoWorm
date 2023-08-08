package common

type Vector struct {
	X int
	Y int
}

type Vectorf struct {
	X float64
	Y float64
}

func (vec1 Vector) IsEqual(vec2 Vector) bool {
	return vec1.X == vec2.X && vec1.Y == vec2.Y
}

func (vec1 Vectorf) IsEqualf(vec2 Vectorf) bool {
	return vec1.X == vec2.X && vec1.Y == vec2.Y
}
