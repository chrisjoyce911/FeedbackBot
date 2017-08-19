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
				BotName:     "Kafka to HipCat",
				Broker:      "127.0.0.1:9092",
				Topic:       "my_topic",
				GroupID:     "feedback_to_hipchat",
				ReleseApp:   "Integration Testing",
				ReleseWeb:   "Integration Testing",
				Development: "Integration Testing",
			},
		},
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
		configfile string
	}
	tests := []struct {
		name    string
		args    args
		want    Configuration
		wantErr bool
	}{
		{name: "Load Default",
			args:    args{configfile: "test_configs/test_config.json"},
			want:    createMockConfig(),
			wantErr: false},
		{name: "Fail to Load",
			args:    args{configfile: "test_configs/test_configXXXX.json"},
			want:    Configuration{},
			wantErr: true},
		{name: "Bad JSON",
			args:    args{configfile: "test_configs/test_config_bad.json"},
			want:    Configuration{},
			wantErr: true},
		{name: "No file",
			args:    args{configfile: ""},
			want:    Configuration{},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.configfile)
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
			args:    args{c: createMockConfig(), filename: "test_configs/test_config.json"},
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

func TestLoadToken(t *testing.T) {

	type args struct {
		tokenfilename string
	}
	tests := []struct {
		name    string
		args    args
		want    Token
		wantErr bool
	}{
		{name: "Load Test Token",
			args:    args{tokenfilename: "test_configs/test_token.json"},
			want:    Token{HipToken: "HIP_TOKEN"},
			wantErr: false},
		{name: "Fail to Load",
			args:    args{tokenfilename: "test_configs/test_tokenXXXX.json"},
			want:    Token{HipToken: "HipCat-Token"},
			wantErr: false},
		{name: "Bad JSON",
			args:    args{tokenfilename: "test_configs/test_config_bad.json"},
			want:    Token{},
			wantErr: true},
		{name: "No file",
			args:    args{tokenfilename: ""},
			want:    Token{},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadToken(tt.args.tokenfilename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveToken(t *testing.T) {
	type args struct {
		c        Token
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Save Default",
			args:    args{c: createMockToken(), filename: "test_configs/test_token.json"},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveToken(tt.args.c, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("saveToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createMockToken(t *testing.T) {
	tests := []struct {
		name string
		want Token
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createMockToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createMockToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
