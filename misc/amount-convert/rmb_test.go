package amount

import (
	"testing"
)

func TestFenToYuan(t *testing.T) {

	type arg struct {
		input int64
		want  string
	}

	testCases := []arg{
		{
			input: 1,
			want:  "0.01",
		},
		{
			input: 0,
			want:  "0.00",
		},
		{
			input: 100,
			want:  "1.00",
		},
		{
			input: 101,
			want:  "1.01",
		},
		{
			input: 105,
			want:  "1.05",
		},
		{
			input: 58805,
			want:  "588.05",
		}, {
			input: 50005,
			want:  "500.05",
		},
	}

	for _, tc := range testCases {
		output := FenToYuan(tc.input)
		if output != tc.want {
			t.Errorf("input:%d want:%s output:%s", tc.input, tc.want, output)
		} else {
			t.Logf("input:%d want:%s output:%s", tc.input, tc.want, output)
		}
	}
}

func TestYuanToFen(t *testing.T) {
	type arg struct {
		input float64
		want  int64
	}

	testCases := []arg{
		{
			input: 0.01,
			want:  1,
		},
		{
			input: 0.00,
			want:  0,
		},
		{
			input: 1.00,
			want:  100,
		},
		{
			input: 1.01,
			want:  101,
		},
		{
			input: 1.05,
			want:  105,
		},
		{
			input: 588.05,
			want:  58805,
		}, {
			input: 500.05,
			want:  50005,
		}, {
			input: 500.55,
			want:  50055,
		}, {
			input: 500.54,
			want:  50054,
		}, {
			input: 500.54,
			want:  50054,
		}, {
			input: 4.19,
			want:  419,
		},
	}

	for _, tc := range testCases {
		output := YuanToFen(tc.input)
		if output != tc.want {
			t.Errorf("input:%f want:%d output:%d", tc.input, tc.want, output)
		} else {
			t.Logf("input:%f want:%d output:%d", tc.input, tc.want, output)
		}
	}
}
