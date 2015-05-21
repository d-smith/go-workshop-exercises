package bowling

type BowlingRolls [20]int

type Game struct {
	rolls   BowlingRolls
	rollIdx int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) {
	g.rolls[g.rollIdx] = pins
	g.rollIdx++
}

func (g *Game) Score() int {
	var gameScore int
	var currentRoll int
	for frame := 0; frame < 10; frame++ {
		if isStrike(g.rolls, currentRoll) {
			gameScore += 10 + g.rolls[currentRoll+1] + g.rolls[currentRoll+2]
			currentRoll++
		} else if isSpare(g.rolls, currentRoll) {
			gameScore += 10 + g.rolls[currentRoll+2]
			currentRoll += 2
		} else {
			gameScore = gameScore + g.rolls[currentRoll] + g.rolls[currentRoll+1]
			currentRoll += 2
		}
	}
	return gameScore
}

func isSpare(rolls BowlingRolls, rollIdx int) bool {
	return rolls[rollIdx]+rolls[rollIdx+1] == 10
}

func isStrike(rolls BowlingRolls, rollIdx int) bool {
	return rolls[rollIdx] == 10
}
