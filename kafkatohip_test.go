package main

import (
	"testing"
)

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

func Test_forwardMessage(t *testing.T) {
	type args struct {
		message string
		inthis  []Channel
	}
	tests := []struct {
		name               string
		args               args
		wantHipchatChannel string
		wantBackground     string
		wantForward        bool
	}{
		{name: "Will forward Background",
			args:               args{message: "Match oneText Rating: Satisfied", inthis: createMockConfig().Channels},
			wantHipchatChannel: "Match oneChannel",
			wantBackground:     "green",
			wantForward:        true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHipchatChannel, gotBackground, gotForward := forwardMessage(tt.args.message, tt.args.inthis)
			if gotHipchatChannel != tt.wantHipchatChannel {
				t.Errorf("forwardMessage() gotHipchatChannel = %v, want %v", gotHipchatChannel, tt.wantHipchatChannel)
			}
			if gotBackground != tt.wantBackground {
				t.Errorf("forwardMessage() gotBackground = %v, want %v", gotBackground, tt.wantBackground)
			}
			if gotForward != tt.wantForward {
				t.Errorf("forwardMessage() gotForward = %v, want %v", gotForward, tt.wantForward)
			}
		})
	}
}
