package util

import "testing"

func TestGenerateUUD(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{data: "Juma"}, "juma"},
		{"Test 2", args{data: "Hamis Juma"}, "hamisjuma"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUUID(tt.args.data); got != tt.want {
				t.Errorf("GenerateUUD() = %v, want %v", got, tt.want)
			}
		})
	}
}
