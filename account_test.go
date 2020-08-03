package accountapi_client

import (
	"reflect"
	"testing"
)

func TestNewAccountClient_ok(t *testing.T) {
	client, err := NewAccountClient("http", "localhost", 9000, 1)
	if err != nil {
		t.Errorf("got %s want nil", err.Error())
	}
	expected := "http://localhost:9000/1"
	if client.baseUrl.String() != expected {
		t.Errorf("String() = %v, want %v", client.baseUrl.String(), expected)
	}
}

func TestNewAccountClient(t *testing.T) {
	type args struct {
		protocol string
		host     string
		port     int
		version  int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				protocol: "https",
				host:     "localhost",
				port:     9000,
				version:  1,
			},
			want:    "https://localhost:9000/1",
			wantErr: false,
		},
		{
			name: "invalid protocol",
			args: args{
				protocol: "aaaa",
				host:     "localhost",
				port:     9000,
				version:  1,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid port",
			args: args{
				protocol: "http",
				host:     "localhost",
				port:     0,
				version:  1,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccountClient(tt.args.protocol, tt.args.host, tt.args.port, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccountClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			actual := got.baseUrl
			var actualStr string
			if actual == nil {
				actualStr = ""
			} else {
				actualStr = actual.String()
			}
			if !reflect.DeepEqual(actualStr, tt.want) {
				t.Errorf("NewAccountClient() got = %v, want %v", actualStr, tt.want)
			}
		})
	}
}

func Test_isValidPort(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				port: 8080,
			},
			wantErr: false,
		},
		{
			name: "low",
			args: args{
				port: 0,
			},
			wantErr: true,
		},
		{
			name: "high",
			args: args{
				port: 999999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidPort(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("isValidPort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isValidProtocol(t *testing.T) {
	type args struct {
		protocol string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok http",
			args: args{
				protocol: "http",
			},
			wantErr: false,
		},
		{
			name: "ok https",
			args: args{
				protocol: "https",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				protocol: "tcp",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidProtocol(tt.args.protocol); (err != nil) != tt.wantErr {
				t.Errorf("isValidProtocol() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
