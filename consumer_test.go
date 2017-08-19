package main

import "testing"

func TestFeedBackConsumer(t *testing.T) {
	type args struct {
		cfg      Configuration
		messages chan string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FeedBackConsumer(tt.args.cfg, tt.args.messages)
		})
	}
}
