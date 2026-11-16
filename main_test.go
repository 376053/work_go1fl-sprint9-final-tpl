package main

import "testing"

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "несколько элементов",
			data: []int{1, 5, 3, 9, 2},
			want: 9,
		},
		{
			name: "все одинаковые",
			data: []int{7, 7, 7, 7},
			want: 7,
		},
		{
			name: "максимум в начале",
			data: []int{100, 1, 2, 3},
			want: 100,
		},
		{
			name: "максимум в конце",
			data: []int{1, 2, 3, 100},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "несколько элементов",
			data: []int{1, 5, 3, 9, 2, 8, 4, 7},
			want: 9,
		},
		{
			name: "все одинаковые",
			data: []int{7, 7, 7, 7, 7, 7, 7, 7},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
