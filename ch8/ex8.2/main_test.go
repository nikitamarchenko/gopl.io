/*
Exercise 8.2: Implement a concurrent File Transfer Protocol (FTP) server. The
server should interpret commands from each client such as cd to change
directory, ls to list a directory, get to send the contents of a file, and
close to close the connection. You can use the standard ftp command as the
client, or write your own.
*/

package main

import (
	"testing"
)

func Test_parsePort(t *testing.T) {
	type args struct {
		param1 string
	}
	tests := []struct {
		name    string
		args    args
		wantR   string
		wantErr bool
	}{
		{"invalid", args{"0,0,0,p,123,5"}, "", true},
		{"127,0,0,1,203,113", args{"127,0,0,1,203,113"}, "127.0.0.1:52081", false},
		{"127,1,2,3,203,113", args{"127,1,2,3,203,113"}, "127.1.2.3:52081", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := parsePort(tt.args.param1)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("parsePort() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
