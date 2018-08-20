package meander

import (
	"reflect"
	"testing"
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
	in   cost
	want int
} {
	return []struct {
		in   cost
		want int
	}{
		{
			in:   cost1,
			want: 1,
		},
		{
			in:   cost2,
			want: 2,
		},
		{
			in:   cost3,
			want: 3,
		},
		{
			in:   cost4,
			want: 4,
		},
		{
			in:   cost5,
			want: 5,
		},
	}
}

func TestCostString(t *testing.T) {
	thens := getCostStringTestCases()
	for _, then := range thens {
		have := then.in.string()
		if have != then.want {
			t.Errorf("have %s, but want %s", have, then.want)
		}
	}
}

func getCostStringTestCases() []struct {
	in   cost
	want string
} {
	return []struct {
		in   cost
		want string
	}{
		{
			in:   cost1,
			want: "$",
		},
		{
			in:   cost2,
			want: "$$",
		},
		{
			in:   cost3,
			want: "$$$",
		},
		{
			in:   cost4,
			want: "$$$$",
		},
		{
			in:   cost5,
			want: "$$$$$",
		},
	}
}

func TestParseCost(t *testing.T) {
	thens := getParseCostTestCases()
	for _, then := range thens {
		have := parseCost(then.in)
		if have != then.want {
			t.Errorf("have %d, but want %d", have, then.want)
		}
	}
}

func getParseCostTestCases() []struct {
	in   string
	want cost
} {
	return []struct {
		in   string
		want cost
	}{
		{
			in:   "$",
			want: cost1,
		},
		{
			in:   "$$",
			want: cost2,
		},
		{
			in:   "$$$",
			want: cost3,
		},
		{
			in:   "$$$$",
			want: cost4,
		},
		{
			in:   "$$$$$",
			want: cost5,
		},
	}
}

func TestCostRangeString(t *testing.T) {
	thens := getCostRangeStringTestCases()
	for _, then := range thens {
		have := then.in.string()
		if have != then.want {
			t.Errorf("have %s, but want %s", have, then.want)
		}
	}
}

func getCostRangeStringTestCases() []struct {
	in   costRange
	want string
} {
	return []struct {
		in   costRange
		want string
	}{
		{
			in: costRange{
				from: cost3,
				to:   cost4,
			},
			want: "$$$...$$$$",
		},
		{
			in: costRange{
				from: cost2,
				to:   cost5,
			},
			want: "$$...$$$$$",
		},
	}
}

func TestParseCostRange(t *testing.T) {
	thens := getParseCostRangeTestCases()
	for _, then := range thens {
		have := parseCostRange(then.in)
		if !reflect.DeepEqual(*have, then.want) {
			t.Errorf("have %#v, but want %#v", *have, then.want)
		}
	}
}

func getParseCostRangeTestCases() []struct {
	in   string
	want costRange
} {
	return []struct {
		in   string
		want costRange
	}{
		{
			in: "$$...$$$",
			want: costRange{
				from: cost2,
				to:   cost3,
			},
		},
		{
			in: "$...$$$$$",
			want: costRange{
				from: cost1,
				to:   cost5,
			},
		},
	}
}
