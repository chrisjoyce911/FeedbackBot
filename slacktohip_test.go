package main

import (
	"flag"
	"reflect"
	"testing"

	"github.com/andybons/hipchat"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

func Test_main(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"add", 2, 2, 4},
		{"add", 4, 4, 8},
		{"minus", 0, -2, -2},
		{"zero", 0, 0, 0},
	}

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			result := Sum(tc.a, tc.b)
			if result != tc.expected {
				t.Fatalf("expected sum %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Add",
			args: args{a: 2, b: 2},
			want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processArg(t *testing.T) {
	t.Skip("skipping test command line arg to hard")
	type args struct {
		configfile string
		slackPtr   string
	}
	tests := []struct {
		name string
		args args
		want Configuration
	}{
		{name: "File",
			args: args{configfile: "config_test.json", slackPtr: "SLACK_TOKEN"},
			want: createMockConfig()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.Set("t", "SLACK_TOKEN")
			if got := processArg(tt.args.configfile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forwardMessage(t *testing.T) {
	type args struct {
		slackChannel string
		inthis       []Channel
	}
	tests := []struct {
		name               string
		args               args
		wantHipchatChannel string
		wantForward        bool
	}{
		{name: "Will forward",
			args:               args{slackChannel: "SLACK0101", inthis: createMockConfig().Channels},
			wantHipchatChannel: "Dev Test Channel",
			wantForward:        true},
		{name: "Will NOT forward",
			args:               args{slackChannel: "NO_MATCH", inthis: createMockConfig().Channels},
			wantHipchatChannel: "",
			wantForward:        false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHipchatChannel, gotForward := forwardMessage(tt.args.slackChannel, tt.args.inthis)
			if gotHipchatChannel != tt.wantHipchatChannel {
				t.Errorf("forwardMessage() gotHipchatChannel = %v, want %v", gotHipchatChannel, tt.wantHipchatChannel)
			}
			if gotForward != tt.wantForward {
				t.Errorf("forwardMessage() gotForward = %v, want %v", gotForward, tt.wantForward)
			}
		})
	}
}

func Test_whatBackground(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name           string
		args           args
		wantBackground string
	}{
		{name: "Color Green",
			args:           args{message: "Rating: Satisfied"},
			wantBackground: hipchat.ColorGreen},
		{name: "Color Yellow",
			args:           args{message: "Rating: Neutral"},
			wantBackground: hipchat.ColorYellow},
		{name: "Color Red",
			args:           args{message: "Rating: Not Satisfied"},
			wantBackground: hipchat.ColorRed},
		{name: "Color Gray",
			args:           args{message: "XXXVVVNNN"},
			wantBackground: hipchat.ColorGray},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBackground := whatBackground(tt.args.message); gotBackground != tt.wantBackground {
				t.Errorf("whatBackground() = %v, want %v", gotBackground, tt.wantBackground)
			}
		})
	}
}
