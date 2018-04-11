package formspell

import (
	"math/rand"

	"github.com/kiwih/formspell/dice"
)

//CR is challenge rating, and is used to compute a bunch of different numbers
type CR int

//NewCR converts integer to CR
func NewCR(c int) CR {
	return CR(c)
}

//GetProficiencyBonus returns the proficiency bonus appropriate to that challenge rating
func (c CR) GetProficiencyBonus() int {
	return (int(c) / 4) + 2
}

//GetAbilityBaseStat returns the ability score improvement appropriate to that challenge rating
func (c CR) GetAbilityBaseStat() int {
	return (int(c)+3)/8 + 2
}

//GetToHit returns the ToHit modifier appropriate for this challenge rating
func (c CR) GetToHit() int {
	return c.GetProficiencyBonus() + c.GetAbilityBaseStat()
}

//GetSpellSaveDC returns the save DC appropriate for this challenge rating
func (c CR) GetSpellSaveDC() int {
	return 8 + c.GetToHit()
}

func newRandomDamageFunction(cr CR, numTargets int) dice.DFunction {
	die := dice.DieTypeD12

	if numTargets < 1 {
		numTargets = 1
	}

	numDie := ((4*(3+int(cr)/4+rand.Intn(2)))/(3*numTargets) + 1)
	switch rand.Intn(4) {
	case 0:
		die = dice.DieTypeD10
	case 1:
		die = dice.DieTypeD8
		numDie = (3 * numDie) / 2
	case 2:
		die = dice.DieTypeD6
		numDie = numDie * 2
	default:
	}
	dices := dice.RepeatDie(die, numDie)
	return dice.DFunction{
		Dice:     dices,
		Constant: 0,
	}
}
