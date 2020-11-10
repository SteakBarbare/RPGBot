package game

type DuelPreparation struct {
	Id              string
	SelectingPlayer string
	IsReady         int
	IsOver          bool
}

type DuelPlayer struct {
	PreparationId  string
	Challenger     string
	Challenged     string
	ChallengerChar string
	ChallengedChar string
}
