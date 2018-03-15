package dice

import (
	"fmt"
	"strconv"
)

//DieType stores the type of a die
type DieType int

//DieTypeXXXX refers to the max value on each dice
const (
	DieTypeD2  DieType = 2
	DieTypeD4  DieType = 4
	DieTypeD6  DieType = 6
	DieTypeD8  DieType = 8
	DieTypeD10 DieType = 10
	DieTypeD12 DieType = 12
)

//RepeatDie makes a nice slice of DieTypes for us
func RepeatDie(t DieType, c int) []DieType {
	d := make([]DieType, c)
	for i := 0; i < c; i++ {
		d[i] = t
	}
	return d
}

//DFunction is for a function of Die that can be of different types
type DFunction struct {
	Dice     []DieType
	Constant int
}

//Evaluate returns the sum of the average value of all Die in the DFunction
// If dice function was 8d10, it would return 45 (average of 1d8 is 4.5)
func (d DFunction) Evaluate() int {
	sum := 0
	for i := 0; i < len(d.Dice); i++ {
		sum += int(d.Dice[i]) + 1
	}
	sum = sum / 2
	return sum + d.Constant
}

//String satisfies the Stringer interface by converting the DFunction to a string
// This will output something like "12 (3d6+3)"
func (d DFunction) String() string {
	lastType := DieType(-1)
	curCount := -1
	buildStr := ""

	addPlus := false

	for i := 0; i < len(d.Dice); i++ {
		if d.Dice[i] != lastType {
			if curCount > 0 {
				//because the dice type has changed, we can add the current accumulated dice type to the string
				if addPlus {
					buildStr += " + "
				}
				buildStr += strconv.Itoa(int(curCount)) + "d" + strconv.Itoa(int(lastType))
				addPlus = true
			}
			curCount = 1
			lastType = d.Dice[i]
		} else {
			curCount++
		}
	}

	//We need the last line added manually because there's no change after the last die, but we still need it added
	if addPlus {
		buildStr += " + "
	}

	buildStr += strconv.Itoa(int(curCount)) + "d" + strconv.Itoa(int(lastType))
	if d.Constant != 0 {
		return fmt.Sprintf("%d (%s + %d)", d.Evaluate(), buildStr, d.Constant)
	}

	return fmt.Sprintf("%d (%s)", d.Evaluate(), buildStr)
}
