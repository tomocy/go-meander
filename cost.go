package meander

import "strings"

type cost int8

const (
	unknownCost cost = iota
	cost1
	cost2
	cost3
	cost4
	cost5
)

var costStrings = map[cost]string{
	cost1: "$",
	cost2: "$$",
	cost3: "$$$",
	cost4: "$$$$",
	cost5: "$$$$$",
}

func (c cost) string() string {
	return costStrings[c]
}

func parseCost(s string) cost {
	for cost, costStr := range costStrings {
		if costStr == s {
			return cost
		}
	}

	return unknownCost
}

type costRange struct {
	from cost
	to   cost
}

func (cr costRange) string() string {
	return cr.from.string() + "..." + cr.to.string()
}

func parseCostRange(s string) *costRange {
	ss := strings.Split(s, "...")
	return &costRange{
		from: parseCost(ss[0]),
		to:   parseCost(ss[1]),
	}
}
