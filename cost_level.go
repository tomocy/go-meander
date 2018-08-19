package meander

import "strings"

type Cost int8

const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (c Cost) String() string {
	for k, v := range costStrings {
		if v == c {
			return k
		}
	}

	return "UNKNOWN"
}

func ParseCost(s string) Cost {
	return costStrings[s]
}

type CostRange struct {
	From Cost
	To   Cost
}

func (cr CostRange) String() string {
	return cr.From.String() + "..." + cr.To.String()
}

func ParseCostRange(s string) *CostRange {
	ss := strings.Split(s, "...")
	return &CostRange{
		From: ParseCost(ss[0]),
		To:   ParseCost(ss[1]),
	}
}
