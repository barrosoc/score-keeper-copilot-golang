package main

import (
	"fmt"
)

type ScoreKeeper struct {
	scoreTeamA int
	scoreTeamB int
}

func NewScoreKeeper() *ScoreKeeper {
	return &ScoreKeeper{0, 0}
}

func (sk *ScoreKeeper) scoreTeamA1() {
	sk.scoreTeamA += 1
}

func (sk *ScoreKeeper) scoreTeamB1() {
	sk.scoreTeamB += 1
}

func (sk *ScoreKeeper) scoreTeamA2() {
	sk.scoreTeamA += 2
}

func (sk *ScoreKeeper) scoreTeamB2() {
	sk.scoreTeamB += 2
}

func (sk *ScoreKeeper) scoreTeamA3() {
	sk.scoreTeamA += 3
}

func (sk *ScoreKeeper) scoreTeamB3() {
	sk.scoreTeamB += 3
}

func (sk *ScoreKeeper) getScore() string {
	return fmt.Sprintf("%03d:%03d", sk.scoreTeamA, sk.scoreTeamB)
}
