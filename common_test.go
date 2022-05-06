package slice

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestMinMax(t *testing.T) {
	tests := []struct {
		a, b, c, d int
	}{
		{
			a: 12,
			b: 8,
			c: 8,
			d: 12,
		},
		{
			a: 20,
			b: 25,
			c: 20,
			d: 25,
		},
		{
			a: 30,
			b: 100,
			c: 30,
			d: 100,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i)+"_min", func(t *testing.T) {
			m := Min(test.a, test.b)
			if m != test.c {
				t.Errorf("expected: %v, got: %v", test.c, m)
			}
		})
		t.Run(strconv.Itoa(i)+"_max", func(t *testing.T) {
			m := Max(test.a, test.b)
			if m != test.d {
				t.Errorf("expected: %v, got: %v", test.d, m)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		els  []int
		want int
	}{
		{
			name: "arithmetic progression",
			els:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 55,
		},
		{
			name: "random",
			els:  []int{20, 15, 1, 0, 5},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.els...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name string
		els  []bool
		want bool
	}{
		{
			name: "all true",
			els:  []bool{true, true, true, true},
			want: true,
		},
		{
			name: "all false",
			els:  []bool{false, false, false},
			want: false,
		},
		{
			name: "empty",
			els:  []bool{},
			want: false,
		},
		{
			name: "random",
			els:  []bool{false, false, true, true, false},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.els...); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name string
		els  []bool
		want bool
	}{
		{
			name: "all false",
			els:  []bool{false, false, false},
			want: false,
		},
		{
			name: "all true",
			els:  []bool{true, true, true, true},
			want: true,
		},
		{
			name: "random",
			els:  []bool{true, true, false, true, true},
			want: false,
		},
		{
			name: "empty",
			els:  []bool{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.els...); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		els  []int
		want []int
	}{
		{
			name: "random",
			els:  []int{1, 6, 4, 8, 2, 9, 3, 10, 1, 6, 4},
			want: []int{1, 2, 3, 4, 6, 8, 9, 10},
		},
		{
			name: "equal",
			els:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: []int{1},
		},
		{
			name: "empty",
			els:  []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.els...)

			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
