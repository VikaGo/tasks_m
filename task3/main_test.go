package main

import (
	"testing"
	_ "time"
)

func TestConvertToMilitaryTime(t *testing.T) {
	type args struct {
		time12format string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "time1",
			args:    args{time12format: "12:03:16PM"},
			want:    "12:03:16",
			wantErr: true,
		},
		{
			name:    "time2",
			args:    args{time12format: "12:45:13AM"},
			want:    "00:45:13",
			wantErr: true,
		},
		{
			name:    "time3",
			args:    args{time12format: ""},
			want:    "",
			wantErr: false,
		},
		{
			name:    "time4",
			args:    args{time12format: "11:48:54PM"},
			want:    "23:48:54",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToMilitaryTime(tt.args.time12format)
			if got != tt.want {
				t.Errorf("ConvertToMilitaryTime() got = %v, want %v", got, tt.want)
				if (err != nil) != tt.wantErr {
					t.Errorf("ConvertToMilitaryTime() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
