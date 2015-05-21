package bowling

import (
	"testing"
)

func rollMany(g *Game, norolls int, pins int) {
	for i := 0; i < norolls; i++ {
		g.Roll(pins)
	}
}

func TestAllGutterBowls(t *testing.T) {
	t.Log("When I roll all gutter balls the score should be 0")
	g := NewGame()
	rollMany(g, 20, 0)
	if g.Score() != 0 {
		t.Fail()
	}
}

func TestAllSinglePinRolls(t *testing.T) {
	t.Log("When I roll all ones the score should be 20")
	g := NewGame()
	rollMany(g, 20, 1)
	if g.Score() != 20 {
		t.Error("got", g.Score(), " expected ", 20)
	}
}

func TestScoreSpare(t *testing.T) {
	t.Log("When I roll a spare, a 3, then gutter balls I score 16")
	g := NewGame()

	g.Roll(6)
	g.Roll(4)

	g.Roll(3)

	rollMany(g, 17, 0)

	if g.Score() != 16 {
		t.Error("got ", g.Score(), " expected ", 16)
	}
}

func TestRollOneStrike(t *testing.T) {
	t.Log("When I roll 1 strike, a 3, and a 4 I score 24")
	g := NewGame()

	g.Roll(10)
	g.Roll(3)
	g.Roll(4)
	rollMany(g, 16, 0)

	if g.Score() != 24 {
		t.Error("got ", g.Score(), " expected ", 24)
	}
}

func TestRollPerfectGame(t *testing.T) {
	t.Log("When I roll all strikes I score 300")
	g := NewGame()
	rollMany(g, 12, 10)
	if g.Score() != 300 {
		t.Error("got ", g.Score(), " expected ", 300)
	}
}
