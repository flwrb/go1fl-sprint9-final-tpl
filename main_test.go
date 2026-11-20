package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	numbers := []int{1, 20, -10, 100, 44, 0}
	for _, data := range numbers {
		sliceNum := generateRandomElements(data)
		wantLen := data
		if data <= 0 {
			wantLen = 0
		}
		if len(sliceNum) != wantLen {
			t.Errorf("generateRandomElements(%d): ожидалась длина %d, получена %d", data, wantLen, len(sliceNum))
		}
	}
}

func TestMaximum(t *testing.T) {
	type testCase struct {
		data []int
		want int
	}
	cases := []testCase{
		{data: nil, want: 0},
		{data: []int{42}, want: 42},
		{data: []int{1, 9, 3}, want: 9},
		{data: []int{9, 1, 3}, want: 9},
		{data: []int{1, 3, 9}, want: 9},
		{data: []int{-5, -2, -9}, want: -2},
	}
	for _, j := range cases {
		out := maximum(j.data)
		if out != j.want {
			t.Errorf("maximum(%v): ожидаемы ответ %d, получена %d", j.data, j.want, out)
		}
	}
}

func TestMaxChunks(t *testing.T) {
	nums := []int{1, 10, 20, 50, -20, 400, 1000, 0, 13, 8}
	for _, i := range nums {
		data := generateRandomElements(i)
		want := maximum(data)
		got := maxChunks(data)
		if want != got {
			t.Errorf("размер %d:maxChunks вернул %d, maximum вернул %d", i, got, want)
		}
	}
	data := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 777}
	want := 777
	got := maxChunks(data)
	if want != got {
		t.Errorf("размер %d:maxChunks вернул %d, maximum вернул %d", len(data), got, want)
	}
}
