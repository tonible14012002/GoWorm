package common

import (
	"math"
)

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

func (vec1 Vectorf) ToArcTan2f() float64 {
	return math.Atan2(vec1.Y, vec1.X)
}

func (vec1 Vectorf) Add(vec2 Vectorf) Vectorf {
	return Vectorf{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
}
func (vec1 Vectorf) Minus(vec2 Vectorf) Vectorf {
	return Vectorf{
		X: vec1.X - vec2.X,
		Y: vec1.Y - vec2.Y,
	}
}

func (vec1 Vectorf) Dot(vec2 Vectorf) float64 {
	return vec1.X*vec2.X + vec1.Y*vec2.Y
}

func (vec1 Vectorf) Multi(factor float64) Vectorf {
	return Vectorf{
		vec1.X * factor,
		vec1.Y * factor,
	}
}

type MovingDirection int

const (
	Up MovingDirection = iota
	Down
)
