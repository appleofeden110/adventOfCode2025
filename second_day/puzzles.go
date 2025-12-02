package second_day

import (
	"strings"
)

type Range struct {
	start, finish string
}

func ClearInput(input string) []Range {
	ranges := strings.Split(input, ",")

	trueRanges := make([]Range, 0)

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		b, a, _ := strings.Cut(r, "-")
		trueRanges = append(trueRanges, Range{b, a})
	}

	return trueRanges
}
