package game

type GamePlayStateType int

const (
	STANDBY GamePlayStateType = iota
	FIRING
	EXPLODING
	STABLE
	OVER
)

type GamePlayState struct {
	state GamePlayStateType
}
