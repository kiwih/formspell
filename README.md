# formspell

This program makes form-based basic spells for d&d 5e.

## Installation

```
$ go get github.com/kiwih/formspell
$ cd $GOPATH$/src/github.com/kiwih/formspell/main
$ go build -o formspell
```

## Usage
```
$ cd $GOPATH$/src/github.com/kiwih/formspell/main
$ formspell [-cr=<##>] [-e] [-nd]
    -cr=<##>    Sets the challenge rating of the monster making the spell.
    -e          Add a secondary effect to the spell.
    -nd         Removed damage from the spell (so it functions as a pure debuff).
```
