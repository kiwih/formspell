package formspell

import (
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

func newRandomDamageFunction(cr CR) dice.DFunction {
	dices := dice.RepeatDie(dice.DieTypeD12, 3+int(cr)/4+1)
	return dice.DFunction{
		Dice:     dices,
		Constant: 0,
	}
}
