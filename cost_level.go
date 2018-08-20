package meander

import "strings"

type cost int8

const (
	_ cost = iota
	cost1
	cost2
	cost3
	cost4
	cost5
)

var costStrings = map[string]cost{
	"$":     cost1,
	"$$":    cost2,
	"$$$":   cost3,
	"$$$$":  cost4,
	"$$$$$": cost5,
}

func (c cost) string() string {
	for k, v := range costStrings {
		if v == c {
			return k
		}
	}

	return "UNKNOWN"
}

func parseCost(s string) cost {
	return costStrings[s]
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
