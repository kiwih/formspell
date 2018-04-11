# formspell

This program makes form-based basic spells for d&d 5e.

It scales damage die, to-hit, saving throws, number of targets, range, and other variables based on the requested CR number, and whether or not you wish for the spell to have effects or no damage.

It also features randomness to make things more interesting :-)

## Example outputs
* Simple damage spells are easy to generate
```
$ ./formspell
(CR 10) Name: The Matron's Water Typhoon
Targets: (+7 to hit) 2 target(s) (Range: 15/90ft)
32 (5d12) Lightning damage

$ ./formspell
(CR 10) Name: The Professor's Spectral Daydream
Targets: (+7 to hit) 1 target(s) (Range: 15ft)
58 (13d8) Cold damage

$ ./formspell
(CR 10) Name: The Dean's Green Curse
Targets: A 25ft cone
45 (7d12) Radiant damage (CON 14 to half)

$ ./formspell
(CR 10) Name: Hammond's Red Bolt
Targets: (+7 to hit) 1 target(s) (range 55 feet)
39 (6d12) Thunder damage

```
* We can also add effects to the spells using the `-e` tag.
```
$ ./formspell -e
(CR 10) Name: Hammond's Weak Decimation
Targets: (+7 to hit) 5 target(s) (Range: 25ft)
14 (4d6) Necrotic damage
A target is 'Fatigued' (CON 15 to avoid) for 3 rounds or until cured via lesser restoration (The target can repeat this saving throw whenever they take damage).
```
* There are also other options, for setting a challenge rating using `-cr`, and to have no damage component, using `-nd`. 
```
$ ./formspell -cr=21 -e
(CR 21) Name: The Matron's Tranquil Shout
Targets: (Range: Self), All targets within a 55 foot radius
77 (22d6) Psychic damage
A target is 'Charmed' (INT 20 to avoid) for 5 rounds or until cured via lesser restoration.

$ ./formspell -cr=21 -e
(CR 21) Name: The Professor's Jaunty Slash
Targets: (+12 to hit) 5 target(s) (Range: 15ft)
16 (3d10) Psychic damage
A target is 'Paralyzed' (WIS 21 to avoid) for 3 rounds or until cured via lesser restoration.

$ ./formspell -cr=15
(CR 15) Name: The Dean's Central Humdinger
Targets: All targets within a 10ft cone
63 (18d6) Cold damage (STR 18 to half)

$ ./formspell -cr=15 -nd -e
(CR 15) Name: Hammond's Air Enhancement
Targets: (Range: Self), All targets within a 20 foot radius
A target is 'Blinded' (DEX 18 to avoid) for 1 rounds or until cured via lesser restoration (The target can repeat this saving throw at the end of their turns).

$ ./formspell -cr=15 -nd -e
(CR 15) Name: The Matron's Gluten-free Explosion
Targets: All targets inside 7 contiguous squares, with first square within 40 feet
A target is 'Poisoned' (CON 17 to avoid) until cured via lesser restoration (The target can repeat this saving throw at the end of their turns).
```

## Installation

```
$ go get github.com/kiwih/formspell
$ cd $GOPATH/src/github.com/kiwih/formspell/main
$ go build -o formspell
```

## Usage
```
$ cd $GOPATH/src/github.com/kiwih/formspell/main
$ formspell [-cr=<##>] [-e] [-nd]
    -cr=<##>    Sets the challenge rating of the monster making the spell. Default is 10.
    -e          Add a secondary effect to the spell.
    -nd         Removed damage from the spell (so it functions as a pure debuff).
```