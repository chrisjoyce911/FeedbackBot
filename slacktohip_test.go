package main

import (
	"reflect"
	"testing"
)

func Test_forwardMessage(t *testing.T) {
	type args struct {
		slackChannel string
		message      string
		inthis       []Channel
	}
	tests := []struct {
		name               string
		args               args
		wantHipchatChannel string
		wantBackground     string
		wantForward        bool
	}{
		{name: "Will forward",
			args:               args{slackChannel: "SLACK0101", message: "Match oneText", inthis: createMockConfig().Channels},
			wantHipchatChannel: "Match oneChannel",
			wantBackground:     "gray",
			wantForward:        true},
		{name: "Will forward Background",
			args:               args{slackChannel: "SLACK0101", message: "Match oneText Rating: Satisfied", inthis: createMockConfig().Channels},
			wantHipchatChannel: "Match oneChannel",
			wantBackground:     "green",
			wantForward:        true},
		{name: "No forward",
			args:               args{slackChannel: "SLACK", message: "Should not forward", inthis: createMockConfig().Channels},
			wantHipchatChannel: "",
			wantBackground:     "gray",
			wantForward:        false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHipchatChannel, gotBackground, gotForward := forwardMessage(tt.args.slackChannel, tt.args.message, tt.args.inthis)
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

func Test_loadcfg(t *testing.T) {
	type args struct {
		configfile string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg Configuration
	}{
		{name: "Good config",
			args:    args{configfile: "config_test.json"},
			wantCfg: createMockConfig()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCfg := loadcfg(tt.args.configfile); !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("loadcfg() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}
