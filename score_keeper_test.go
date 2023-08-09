/*
Implement the ScoreKeeper as a Senior Go Lang Developer following a TDD approach

Context:
We need software to deliver the proper data to the scoreboard for a basketball team.
They need six buttons (each team can score either 1, 2 or 3 points with a single shot).

Write a class ScoreKeeper which offers following methods:
	scoreTeamA1()
	scoreTeamA2()
	scoreTeamA3()
	scoreTeamB1()
	scoreTeamB2()
	scoreTeamB3()
	getScore()

Rules:
The returned String always has seven characters. An example would be 000:000
*/

package main

import (
	"testing"
)

func TestScoreKeeper(t *testing.T) {
	// First test
	t.Run("Test scoreTeamA1", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamA1()
		score := scoreKeeper.getScore()
		if score != "001:000" {
			t.Errorf("Expected 001:000, got %s", score)
		}
	})

	// Second test
	t.Run("Test scoreTeamA2", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamA2()
		score := scoreKeeper.getScore()
		if score != "002:000" {
			t.Errorf("Expected 002:000, got %s", score)
		}
	})

	// Third test
	t.Run("Test scoreTeamA3", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamA3()
		score := scoreKeeper.getScore()
		if score != "003:000" {
			t.Errorf("Expected 003:000, got %s", score)
		}
	})

	// Fourth test
	t.Run("Test scoreTeamB1", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamB1()
		score := scoreKeeper.getScore()
		if score != "000:001" {
			t.Errorf("Expected 000:001, got %s", score)
		}
	})

	// Fifth test
	t.Run("Test scoreTeamB2", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamB2()
		score := scoreKeeper.getScore()
		if score != "000:002" {
			t.Errorf("Expected 000:002, got %s", score)
		}
	})

	// Sixth test
	t.Run("Test scoreTeamB3", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamB3()
		score := scoreKeeper.getScore()
		if score != "000:003" {
			t.Errorf("Expected 000:003, got %s", score)
		}
	})

	// Seventh test
	t.Run("Test scoreTeamA1 and scoreTeamB1", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamA1()
		scoreKeeper.scoreTeamB1()
		score := scoreKeeper.getScore()
		if score != "001:001" {
			t.Errorf("Expected 001:001, got %s", score)
		}
	})

	// Eighth test
	t.Run("Test scoreTeamA2 and scoreTeamB2", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamA2()
		scoreKeeper.scoreTeamB2()
		score := scoreKeeper.getScore()
		if score != "002:002" {
			t.Errorf("Expected 002:002, got %s", score)
		}
	})

	// Ninth test
	t.Run("Test scoreTeamA3 and scoreTeamB3", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamA3()
		scoreKeeper.scoreTeamB3()
		score := scoreKeeper.getScore()
		if score != "003:003" {
			t.Errorf("Expected 003:003, got %s", score)
		}
	})

	// Tenth test
	t.Run("Test scoreTeamA1, scoreTeamA2 and scoreTeamA3", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamA1()
		scoreKeeper.scoreTeamA2()
		scoreKeeper.scoreTeamA3()
		score := scoreKeeper.getScore()
		if score != "006:000" {
			t.Errorf("Expected 006:000, got %s", score)
		}
	})

	// Eleventh test
	t.Run("Test scoreTeamB1, scoreTeamB2 and scoreTeamB3", func(t *testing.T) {
		scoreKeeper := NewScoreKeeper()
		scoreKeeper.scoreTeamB1()
		scoreKeeper.scoreTeamB2()
		scoreKeeper.scoreTeamB3()
		score := scoreKeeper.getScore()
		if score != "000:006" {
			t.Errorf("Expected 000:006, got %s", score)
		}
	})

}
