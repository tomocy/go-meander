package meander_test

import (
	"reflect"
	"testing"

	"github.com/tomocy/meander"
)

func TestCostValue(t *testing.T) {
	thens := getCostValueTestCases()
	for _, then := range thens {
		have := int(then.in)
		if have != then.want {
			t.Errorf("have %d, but want %d", have, then.want)
		}
	}
}

func getCostValueTestCases() []struct {
	in   meander.Cost
	want int
} {
	return []struct {
		in   meander.Cost
		want int
	}{
		{
			in:   meander.Cost1,
			want: 1,
		},
		{
			in:   meander.Cost2,
			want: 2,
		},
		{
			in:   meander.Cost3,
			want: 3,
		},
		{
			in:   meander.Cost4,
			want: 4,
		},
		{
			in:   meander.Cost5,
			want: 5,
		},
	}
}

func TestCostString(t *testing.T) {
	thens := getCostStringTestCases()
	for _, then := range thens {
		have := then.in.String()
		if have != then.want {
			t.Errorf("have %s, but want %s", have, then.want)
		}
	}
}

func getCostStringTestCases() []struct {
	in   meander.Cost
	want string
} {
	return []struct {
		in   meander.Cost
		want string
	}{
		{
			in:   meander.Cost1,
			want: "$",
		},
		{
			in:   meander.Cost2,
			want: "$$",
		},
		{
			in:   meander.Cost3,
			want: "$$$",
		},
		{
			in:   meander.Cost4,
			want: "$$$$",
		},
		{
			in:   meander.Cost5,
			want: "$$$$$",
		},
	}
}

func TestParseCost(t *testing.T) {
	thens := getParseCostTestCases()
	for _, then := range thens {
		have := meander.ParseCost(then.in)
		if have != then.want {
			t.Errorf("have %d, but want %d", have, then.want)
		}
	}
}

func getParseCostTestCases() []struct {
	in   string
	want meander.Cost
} {
	return []struct {
		in   string
		want meander.Cost
	}{
		{
			in:   "$",
			want: meander.Cost1,
		},
		{
			in:   "$$",
			want: meander.Cost2,
		},
		{
			in:   "$$$",
			want: meander.Cost3,
		},
		{
			in:   "$$$$",
			want: meander.Cost4,
		},
		{
			in:   "$$$$$",
			want: meander.Cost5,
		},
	}
}

func TestCostRangeString(t *testing.T) {
	thens := getCostRangeStringTestCases()
	for _, then := range thens {
		have := then.in.String()
		if have != then.want {
			t.Errorf("have %s, but want %s", have, then.want)
		}
	}
}

func getCostRangeStringTestCases() []struct {
	in   meander.CostRange
	want string
} {
	return []struct {
		in   meander.CostRange
		want string
	}{
		{
			in: meander.CostRange{
				From: meander.Cost3,
				To:   meander.Cost4,
			},
			want: "$$$...$$$$",
		},
		{
			in: meander.CostRange{
				From: meander.Cost2,
				To:   meander.Cost5,
			},
			want: "$$...$$$$$",
		},
	}
}

func TestParseCostRange(t *testing.T) {
	thens := getParseCostRangeTestCases()
	for _, then := range thens {
		have := meander.ParseCostRange(then.in)
		if !reflect.DeepEqual(*have, then.want) {
			t.Errorf("have %#v, but want %#v", *have, then.want)
		}
	}
}

func getParseCostRangeTestCases() []struct {
	in   string
	want meander.CostRange
} {
	return []struct {
		in   string
		want meander.CostRange
	}{
		{
			in: "$$...$$$",
			want: meander.CostRange{
				From: meander.Cost2,
				To:   meander.Cost3,
			},
		},
		{
			in: "$...$$$$$",
			want: meander.CostRange{
				From: meander.Cost1,
				To:   meander.Cost5,
			},
		},
	}
}
