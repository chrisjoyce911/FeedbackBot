package main

import "testing"

func TestSendToHipChat(t *testing.T) {
	type args struct {
		f     FeedbackEvent
		cfg   Configuration
		token Token
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendToHipChat(tt.args.f, tt.args.cfg, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("SendToHipChat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
