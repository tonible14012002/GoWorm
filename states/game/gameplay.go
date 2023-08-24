package game

type GamePlayStateType int

const (
	STANDBY GamePlayStateType = iota
	FIRING
	EXPLODING
	STABLE
)

type GamePlayState struct {
	state GamePlayStateType
}
