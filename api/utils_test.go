package api

import "testing"

func TestStringToUInt(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{"1", args{"1"}, 1, false},
		{"2", args{"a"}, 0, true},
		{"3", args{""}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToUInt(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToUInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToUInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
