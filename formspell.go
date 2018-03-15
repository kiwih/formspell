package formspell

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/kiwih/formspell/dice"
)

//DmgType is for indicating that a string indicates a type of damage
type DmgType string

//DmgTypeXXXX indicates each possible damage type
const (
	DmgTypeAcid        = DmgType("Acid")
	DmgTypeBludgeoning = DmgType("Bludgeoning")
	DmgTypeCold        = DmgType("Cold")
	DmgTypeFire        = DmgType("Fire")
	DmgTypeForce       = DmgType("Force")
	DmgTypeLightning   = DmgType("Lightning")
	DmgTypeNecrotic    = DmgType("Necrotic")
	DmgTypePiercing    = DmgType("Piercing")
	DmgTypePoison      = DmgType("Poison")
	DmgTypePsychic     = DmgType("Psychic")
	DmgTypeRadiant     = DmgType("Radiant")
	DmgTypeSlashing    = DmgType("Slashing")
	DmgTypeThunder     = DmgType("Thunder")
)

var magicDmgTypes = []DmgType{
	DmgTypeAcid,
	//DmgTypeBludgeoning,
	DmgTypeCold,
	DmgTypeFire,
	DmgTypeForce,
	DmgTypeLightning,
	DmgTypeNecrotic,
	//DmgTypePiercing,
	DmgTypePoison,
	DmgTypePsychic,
	DmgTypeRadiant,
	//DmgTypeSlashing,
	DmgTypeThunder,
}

func randomMagicDmgType() DmgType {
	return magicDmgTypes[rand.Intn(len(magicDmgTypes))]
}

//StatType refers to one of the six base stats of d&d
type StatType string

//StatTypeXXXX indicates the stat types
const (
	StatTypeStr = StatType("Str")
	StatTypeCon = StatType("Con")
	StatTypeDex = StatType("Dex")
	StatTypeInt = StatType("Int")
	StatTypeWis = StatType("Wis")
	StatTypeCha = StatType("Cha")
)

var statTypes = []StatType{
	StatTypeStr,
	StatTypeCon,
	StatTypeDex,
	StatTypeInt,
	StatTypeWis,
	StatTypeCha,
}

func randomStatType() StatType {
	return statTypes[rand.Intn(len(statTypes))]
}

//ConditionType refers to one of the many 5e condition types
type ConditionType string

//ConditionTypeXXXX indicates all the possible d&d 5e condition types
const (
	ConditionTypeBlinded       = ConditionType("Blinded")
	ConditionTypeCharmed       = ConditionType("Charmed")
	ConditionTypeDeafened      = ConditionType("Deafened")
	ConditionTypeFatigued      = ConditionType("Fatigued")
	ConditionTypeFrightened    = ConditionType("Frightened")
	ConditionTypeGrappled      = ConditionType("Grappled")
	ConditionTypeIncapacitated = ConditionType("Incapacitated")
	ConditionTypeInvisible     = ConditionType("Invisible")
	ConditionTypeParalyzed     = ConditionType("Paralyzed")
	ConditionTypePetrified     = ConditionType("Petrified")
	ConditionTypePoisoned      = ConditionType("Poisoned")
	ConditionTypeProne         = ConditionType("Prone")
	ConditionTypeRestrained    = ConditionType("Restrained")
	ConditionTypeStunned       = ConditionType("Stunned")
	ConditionTypeUnconscious   = ConditionType("Unconscious")
)

var negativeConditionTypes = []ConditionType{
	ConditionTypeBlinded,
	ConditionTypeCharmed,
	ConditionTypeDeafened,
	ConditionTypeFatigued,
	ConditionTypeFrightened,
	ConditionTypeGrappled,
	ConditionTypeIncapacitated,
	//ConditionTypeInvisible, //this is usually beneficial, so does not feature in negativeConditionTypes
	ConditionTypeParalyzed,
	ConditionTypePetrified,
	ConditionTypePoisoned,
	ConditionTypeProne,
	ConditionTypeRestrained,
	ConditionTypeStunned,
	ConditionTypeUnconscious,
}

func randomNegativeConditonType() ConditionType {
	return negativeConditionTypes[rand.Intn(len(negativeConditionTypes))]
}

//TargetsInfo provides information about how an attack finds its targets
type TargetsInfo struct {
	ToHit       int
	Range       int  //can be 0 feet (i.e. self) to N feet (i.e. ranged)
	DisAdvRange int  //0 indicates no DisAdvRange, otherwise, must be greater than Range, indicating a range at which the caster could use it at disadvantage
	NumTargets  int  //indicates number of targets, 0 indicates self-target only
	AoESquares  bool //indicates if numTargets is actually number of 5-foot squares if AOE squares

	AoERadius   int  //indicates radius length of damage if AOE circle
	AoESphere   bool //is it a sphere
	AoECylinder bool //is it a cylinder
	AoECone     bool //is it a cone
	AoELine     bool //is it a Line
}

func (t TargetsInfo) IsAoE() bool {
	return t.AoESquares || t.AoESphere || t.AoECone || t.AoECylinder || t.AoELine
}

func newRandomTargetsInfo(cr CR) TargetsInfo {
	t := TargetsInfo{}

	self := false
	if rand.Intn(10) > 8 {
		self = true
	}

	//if we're not targeting self, decide how we will target
	if self == false {
		t.ToHit = cr.GetToHit()
		t.Range = 15 + rand.Intn(5)*10

		if rand.Intn(10) > 7 {
			t.DisAdvRange = t.Range + 50 + rand.Intn(10)*5
		}

		t.NumTargets = rand.Intn(int(cr)/2) + 1
	}

	//now lets work out if we're aoe

	if rand.Intn(10) > 5 {
		//this is AoE

		switch rand.Intn(5) {
		case 0:
			t.AoESquares = true
			t.NumTargets += 2
		case 1:
			t.AoESphere = true
			t.NumTargets = rand.Intn(int(cr)/8+1) + 1
		case 2:
			t.AoECylinder = true
			t.NumTargets = rand.Intn(int(cr)/8+1) + 1
		case 3:
			t.AoECone = true
			t.Range += 10
			t.NumTargets = 1
		default:
			t.AoELine = true
			t.Range += 50
			t.NumTargets = 1
		}

		if t.AoESquares == false {
			t.AoERadius = rand.Intn(int(cr))*5 + 5
		}
	}

	if self == true && t.AoERadius == 0 && t.AoESquares == false {
		t.AoERadius = rand.Intn(int(cr)/2)*5 + 5
	}

	return t
}

//String satisfies Stringer interface
func (t TargetsInfo) String() string {
	if t.Range == 0 && !t.AoESquares {
		return fmt.Sprintf("(Self), All targets within %d foot range", t.AoERadius)
	}

	if t.AoESquares {
		return fmt.Sprintf("%d contiguous squares, with first square within %d feet", t.NumTargets, t.Range)
	}
	if t.AoESphere {
		return fmt.Sprintf("%d %dft-radius sphere(s) within %d feet", t.NumTargets, t.AoERadius, t.Range)
	}
	if t.AoECylinder {
		return fmt.Sprintf("%d %dft-radius cylinder(s) of height 100 feet within %d feet", t.NumTargets, t.AoERadius, t.Range)
	}
	if t.AoECone {
		return fmt.Sprintf("A %dft cone", t.Range)
	}
	if t.AoELine {
		return fmt.Sprintf("A %dft line", t.Range)
	}

	if t.DisAdvRange != 0 {
		return fmt.Sprintf("(+%d to hit) %d target(s) (range %d/%d feet)", t.ToHit, t.NumTargets, t.Range, t.DisAdvRange)
	}
	return fmt.Sprintf("(+%d to hit) %d target(s) (range %d feet)", t.ToHit, t.NumTargets, t.Range)
}

//A Save is used to indicate how a targeted thing might avoid damage or an effect
type Save struct {
	StatType StatType
	SaveDC   int
}

func (s Save) String() string {
	return fmt.Sprintf("%s %d", strings.ToUpper(string(s.StatType)), s.SaveDC)
}

func newRandomSave(cr CR) *Save {
	return &Save{
		StatType: randomStatType(),
		SaveDC:   cr.GetSpellSaveDC() + rand.Intn(3) - 1, //add some randomness to the save DC
	}
}

type Effect struct {
	EffectType        ConditionType
	SaveAvoid         *Save //optional, if null, it is not avoidable
	NumRoundsSelfCure int   //if 0, it is "permanent"
}

func (e Effect) String() string {
	str := fmt.Sprintf("A target is '%s'", e.EffectType)
	if e.SaveAvoid != nil {
		str += fmt.Sprintf(" (%s to avoid)", e.SaveAvoid)
	}
	if e.NumRoundsSelfCure == 0 {
		str += fmt.Sprintf(" until cured.")
	} else {
		str += fmt.Sprintf(" for %d rounds or until cured.", e.NumRoundsSelfCure)
	}
	return str
}

func newRandomEffect(cr CR) Effect {
	return Effect{
		EffectType:        randomNegativeConditonType(),
		SaveAvoid:         newRandomSave(cr),
		NumRoundsSelfCure: rand.Intn(2) * 3,
	}
}

type Damage struct {
	DmgType        DmgType
	ToHit          *int  //optional, if null it always hits
	SaveHalfDC     *Save //optional, if null there is no save dc for half
	SaveFullDC     *Save //optional, if null there is no save dc for full
	DamageFunction dice.DFunction
}

func newRandomDamage(cr CR, isAoE bool) Damage {
	d := Damage{
		DmgType:        randomMagicDmgType(),
		SaveHalfDC:     nil,
		SaveFullDC:     nil,
		DamageFunction: newRandomDamageFunction(cr),
	}
	if isAoE {
		d.SaveHalfDC = newRandomSave(cr)
	}
	return d
}

func (d Damage) String() string {
	str := fmt.Sprintf("%s %s damage", d.DamageFunction, d.DmgType)
	if d.SaveHalfDC != nil {
		str += fmt.Sprintf(" (%s to half)", d.SaveHalfDC)
	}
	if d.SaveFullDC != nil {
		str += fmt.Sprintf(" (%s to negate)", d.SaveFullDC)
	}
	return str
}

type Spell struct {
	Name    string
	Targets TargetsInfo
	Damages []Damage
	Effects []Effect
}

//NewSpell creates a damage-based spell for a given CR
func NewSpell(cr CR, addDamage bool, addEffect bool) Spell {
	ds := Spell{
		Name:    newRandomName(),
		Targets: newRandomTargetsInfo(cr),
	}
	if addDamage {
		ds.Damages = []Damage{newRandomDamage(cr, ds.Targets.IsAoE())}
	}
	if addEffect {
		ds.Effects = []Effect{newRandomEffect(cr)}
	}
	return ds
}

func (s Spell) String() string {
	str := fmt.Sprintf("Name: %s\nTargets: %s\n", s.Name, s.Targets)
	for _, d := range s.Damages {
		str += fmt.Sprintf("%s\n", d)
	}
	for _, e := range s.Effects {
		str += fmt.Sprintf("%s\n", e)
	}
	return str
}
