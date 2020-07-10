package problem02

import "testing"

func Test_pseudoEncrypt(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"-10", args{-10}, -1270576520},
		{"0", args{0}, 1777613459},
		{"1", args{1}, 561465857},
		{"2", args{2}, 436885871},
		{"3", args{3}, 576481439},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoEncrypt(int32(tt.args.value)); got != int32(tt.want) {
				t.Errorf("pseudoEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateUsername(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name         string
		args         args
		wantUsername string
		wantErr      bool
	}{
		{"negative_index_return_error", args{-10}, "", true},
		{"user_10", args{10}, "user_792482838", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsername, err := generateUsername(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUsername != tt.wantUsername {
				t.Errorf("generateUsername() gotUsername = %v, want %v", gotUsername, tt.wantUsername)
			}
		})
	}
}