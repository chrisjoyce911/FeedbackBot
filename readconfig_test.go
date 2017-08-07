package main

import (
	"reflect"
	"testing"
)

func Test_createMockConfig(t *testing.T) {
	tests := []struct {
		name string
		want Configuration
	}{
		{name: "Default",
			want: Configuration{
				BotName:      "Slack to HipCat",
				SlackToken:   "SLACK_TOKEN",
				HipToken:     "HIP_TOKEN",
				SlackReport:  "",
				SlackRepTime: 600,
				SlackChannel: "",
				MobHipRoom:   "Mobile Feedback",
				WebHipRoom:   "Web Feedback",
				Channels: []Channel{
					{
						Slack:   "SLACK0101",
						HipChat: "Dev Test Channel"},
					{
						Slack:   "SLACK0123",
						HipChat: "Integration Testing"},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createMockConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createMockConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Configuration
		wantErr bool
	}{
		{name: "Load Default",
			args:    args{filename: "config_test.json"},
			want:    createMockConfig(),
			wantErr: false},
		{name: "Fail to Load",
			args:    args{filename: "configXXXX_test.json"},
			want:    Configuration{},
			wantErr: true},
		{name: "Bad JSON",
			args:    args{filename: "config_bad_test.json"},
			want:    Configuration{},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveConfig(t *testing.T) {
	type args struct {
		c        Configuration
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Save Default",
			args:    args{c: createMockConfig(), filename: "config_test.json"},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveConfig(tt.args.c, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("saveConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want Configuration
	}{
		{name: "Existing File",
			args: args{filename: "config_test.json"},
			want: createMockConfig()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfig(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
