# formspell

This program makes form-based basic spells for d&d 5e.

## Example outputs
```
$ ./formspell
Name: Hammond's Colourful Enhancement
Targets: (+7 to hit) 2 target(s) (range 45 feet)
39 (6d12) Lightning damage

$ ./formspell
Name: Hammond's Water Hailstorm
Targets: 3 contiguous squares, with first square within 15 feet
39 (6d12) Fire damage (DEX 14 to half)

$ ./formspell
Name: Hammond's Weak Typhoon
Targets: 2 5ft-radius sphere(s) within 55 feet
39 (6d12) Force damage (WIS 16 to half)

$ ./formspell
Name: Hammond's Fiendish Daydream
Targets: 2 25ft-radius sphere(s) within 25 feet
39 (6d12) Necrotic damage (CHA 15 to half)

$ ./formspell
Name: Hammond's Happy Hailstorm
Targets: A line 105 feet long
39 (6d12) Poison damage (WIS 14 to half)
```
And with other options:
```
$ ./formspell -cr=21 -e
Name: Hammond's Blue Humdinger
Targets: 2 25ft-radius sphere(s) within 25 feet
58 (9d12) Cold damage (CON 19 to half)
A target is 'Deafened' (CHA 20 to avoid) for 3 rounds or until cured.
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
    -cr=<##>    Sets the challenge rating of the monster making the spell.
    -e          Add a secondary effect to the spell.
    -nd         Removed damage from the spell (so it functions as a pure debuff).
```