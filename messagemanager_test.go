package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetBackground(t *testing.T) {
	type args struct {
		f FeedbackEvent
	}
	tests := []struct {
		name string
		args args
		want FeedbackEvent
	}{
		{name: "Satisfied",
			args: args{f: FeedbackEvent{Rating: "Satisfied", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test"}},
			want: FeedbackEvent{Rating: "Satisfied", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test", HipChat: HipChat{Background: "green"}},
		},
		{name: "Neutral",
			args: args{f: FeedbackEvent{Rating: "Neutral", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test"}},
			want: FeedbackEvent{Rating: "Neutral", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test", HipChat: HipChat{Background: "yellow"}},
		},
		{name: "Not Satisfied",
			args: args{f: FeedbackEvent{Rating: "Not Satisfied", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test"}},
			want: FeedbackEvent{Rating: "Not Satisfied", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test", HipChat: HipChat{Background: "red"}},
		},
		{name: "Default",
			args: args{f: FeedbackEvent{Rating: "", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test"}},
			want: FeedbackEvent{Rating: "", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test", HipChat: HipChat{Background: "gray"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBackground(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBackground() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetRoom(t *testing.T) {
	type args struct {
		f   FeedbackEvent
		cfg Configuration
	}
	tests := []struct {
		name string
		args args
		want FeedbackEvent
	}{
		{name: "Mobile",
			args: args{f: FeedbackEvent{ClientInformation: ClientInformation{Build: "release"}, Category: "mobile feedback"}, cfg: Configuration{ReleseApp: "ReleseApp"}},
			want: FeedbackEvent{ClientInformation: ClientInformation{Build: "release"}, Category: "mobile feedback", HipChat: HipChat{Room: "ReleseApp"}},
		},
		{name: "Web",
			args: args{f: FeedbackEvent{ClientInformation: ClientInformation{Build: "release"}, Category: ""}, cfg: Configuration{ReleseWeb: "ReleseWeb"}},
			want: FeedbackEvent{ClientInformation: ClientInformation{Build: "release"}, Category: "", HipChat: HipChat{Room: "ReleseWeb"}},
		},
		{name: "Development",
			args: args{f: FeedbackEvent{ClientInformation: ClientInformation{Build: "developer"}, Category: "mobile feedback"}, cfg: Configuration{Development: "ReleseApp"}},
			want: FeedbackEvent{ClientInformation: ClientInformation{Build: "developer"}, Category: "mobile feedback", HipChat: HipChat{Room: "ReleseApp"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetRoom(tt.args.f, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatMessage(t *testing.T) {
	type args struct {
		f FeedbackEvent
	}
	tests := []struct {
		name string
		args args
		want FeedbackEvent
	}{
		{name: "Basic format",
			args: args{f: FeedbackEvent{Rating: "Satisfied", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test"}},
			want: FeedbackEvent{Rating: "Satisfied", CustomerID: "CustomerID", Source: "feedback", Comment: "My Test", HipChat: HipChat{FormatedMsg: fmt.Sprintf("%s\tCustomerID : %s\t Source : %s\n%s\n\n", "Satisfied", "CustomerID", "feedback", "My Test")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatMessage(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
