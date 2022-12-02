package main

import "testing"

func Test_miniMaxSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "s",
			args:  args{[]int{1, 2, 3}},
			want:  3,
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MiniMaxSum(tt.args.arr)
			if got != tt.want {
				t.Errorf("miniMaxSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("miniMaxSum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
