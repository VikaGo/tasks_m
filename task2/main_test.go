package main

import (
	"errors"
	"strings"
	"testing"
)

func Test_MiniMaxSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		err   error
	}{
		{
			name:  "s",
			args:  args{arr: []int{1, 2, 3}},
			want:  3,
			want1: 5,
			err:   nil,
		},
		{
			name:  "d",
			args:  args{[]int{7, 7, 7, 7, 7}},
			want:  28,
			want1: 28,
			err:   nil,
		},
		{
			name:  "e",
			args:  args{[]int{}},
			want:  0,
			want1: 0,
			err:   errors.New("arr cant be nil"),
		},
		{
			name:  "t",
			args:  args{[]int{8, 7, 6, 5, 5, 4, 3, 1, 2}},
			want:  33,
			want1: 40,
			err:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := MiniMaxSum(tt.args.arr)
			if got != tt.want {
				t.Errorf("MiniMaxSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MiniMaxSum() got1 = %v, want %v", got1, tt.want1)
			}
			if ((err != nil) != (tt.err != nil)) ||
				(tt.err != nil && !strings.Contains(err.Error(), tt.err.Error())) {
				t.Errorf("MiniMaxSum() err = %v, tt.err %v", err, tt.err)
			}
		})
	}
}
