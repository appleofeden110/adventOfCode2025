package firstday

import (
	"strconv"
	"strings"
)

type (
	direction rune

	Rotation struct {
		Direction direction
		NumChange int
	}

	Password struct {
		ActualPassword int
		LastRotation   Rotation
	}
)

const (
	right direction = 'r'
	left  direction = 'l'
)

var DialStart = 50

func GetRotation(input string) *Rotation {
	input = strings.ToLower(input)
	r, isRight := strings.CutPrefix(input, "r")
	l, isLeft := strings.CutPrefix(input, "l")

	if isRight {
		numChange, _ := strconv.Atoi(r)
		return &Rotation{right, numChange}
	}
	if isLeft {
		numChange, _ := strconv.Atoi(l)
		return &Rotation{left, numChange}
	}
	return nil
}

func FindZero(start int, r Rotation) int {
	change := r.NumChange
	switch r.Direction {
	case right:
		for {
			if start >= 100 {
				start -= 100
			}
			if change == 0 {
				break
			}
			if change+start > 100 {
				change -= 100 - start
				start = 100
				continue
			}
			start += change
			change = 0
		}
		return start

	case left:
		for {
			if change == 0 {
				break
			}

			if change > 100 {
				change -= 100
				continue
			}

			start -= change
			if start < 0 {
				start += 100
			}
			change = 0
		}
		return start
	}
	return 0
}

func FindZeroMethod0x434C49434B(start int, r Rotation) (int, int) {
	change := r.NumChange
	var atZeroTimes int
	switch r.Direction {
	case right:
		for {
			if change == 0 {
				break
			}

			start++
			if start == 100 {
				atZeroTimes++
				start = 0
			}
			change--
		}
	case left:
		for {
			if change == 0 {
				break
			}

			start--
			if start == -1 {
				atZeroTimes++
				start = 99
			}
			change--
		}
	}
	return start, atZeroTimes
}
