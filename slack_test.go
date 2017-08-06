/*

mybot - Illustrative Slack bot in Go

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"reflect"
	"testing"

	"golang.org/x/net/websocket"
)

func Test_slackStart(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name      string
		args      args
		wantWsurl string
		wantID    string
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWsurl, gotID, err := slackStart(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("slackStart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWsurl != tt.wantWsurl {
				t.Errorf("slackStart() gotWsurl = %v, want %v", gotWsurl, tt.wantWsurl)
			}
			if gotID != tt.wantID {
				t.Errorf("slackStart() gotId = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func Test_getMessage(t *testing.T) {
	type args struct {
		ws *websocket.Conn
	}
	tests := []struct {
		name    string
		args    args
		wantM   Message
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, err := getMessage(tt.args.ws)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("getMessage() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func Test_postMessage(t *testing.T) {
	type args struct {
		ws *websocket.Conn
		m  Message
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
			if err := postMessage(tt.args.ws, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("postMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_slackConnect(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  *websocket.Conn
		want1 string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := slackConnect(tt.args.token)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slackConnect() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("slackConnect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
