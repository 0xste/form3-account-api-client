package uuid

import (
	"testing"
)

func TestFromStringV4(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name       string
		args       args
		want       UUID
		wantErr    bool
		wantErrStr string
	}{
		{
			name: "valid uuid v4",
			args: args{
				uuid: "ace2892f-b086-4ea1-9214-f70ca9f9db94",
			},
			want:    UUID("ace2892f-b086-4ea1-9214-f70ca9f9db94"),
			wantErr: false,
		},
		{
			name:       "nil uuid v4",
			args:       args{},
			want:       "",
			wantErr:    true,
			wantErrStr: "invalid uuid provided ",
		},
		{
			name: "invalid uuid v4",
			args: args{
				uuid: "123123",
			},
			want:       "",
			wantErr:    true,
			wantErrStr: "invalid uuid provided 123123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromStringV4(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("\ngotErr  :%v\nwantErr :%v", err.Error(), tt.wantErrStr)
			}
			if err != nil {
				if err.Error() != tt.wantErrStr {
					t.Errorf("\ngotErrStr  :%v\nwantErrStr :%v", err.Error(), tt.wantErrStr)
				}
			}
			if got != tt.want {
				t.Errorf("\ngot  :%v\nwant :%v", got, tt.want)
			}
		})
	}
}

func TestIsUUIDv4(t *testing.T) {
	tests := []struct {
		uid  string
		want bool
	}{
		{uid: "0d3ccd19-f704-4803-9a73-a92225a76ce2", want: true},
		{uid: "33ac2632-5e39-4b7e-b9bb-d9cc859f993e", want: true},
		{uid: "c6f2f991-8c65-4aae-bb6f-e129c1d5b321", want: true},
		{uid: "b25372ae-a463-4e20-8687-78e649dc1e89", want: true},
		{uid: "7a9c0a54-fb95-4dd4-832b-cd5784439852", want: true},
		{uid: "e0c74a47-9f1a-4ee5-85f8-7bbda0922597", want: true},
		{uid: "e12eaf23-b551-4c44-90db-e72ac8ad3d46", want: true},
		{uid: "ab175916-41a2-44b5-b212-265873aa1e69", want: true},
		{uid: "ed42aecf-71bb-4d74-9131-8734909870f2", want: true},
		{uid: "0F29FCA5-7B1C-4767-8805-CC000CB769EB", want: true},
		{uid: "00000000-0000-0000-0000-000000000000", want: true},
		{uid: "", want: false},
		{uid: "33ac2632", want: false},
		{uid: "8c65", want: false},
		{uid: "412452646", want: false},
		{uid: "xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", want: true},
		{uid: "00000000-0000-0000-0000-000000000000", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.uid, func(t *testing.T) {
			if got := IsUUIDv4(tt.uid); got != tt.want {
				t.Errorf("uid: %s \ngot  :%v\nwant :%v", tt.uid, got, tt.want)
			}
		})
	}
}

func TestNewV4(t *testing.T) {
	tests := []struct {
		name   string
		nToGen int
	}{
		{
			name:   "generates a valid UUIDv4",
			nToGen: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.nToGen; i++ {
				got, err := NewV4()
				if !IsUUIDv4(got.String()) {
					t.Errorf("invalid uuid generated %s", got.String())
				}
				if err != nil {
					t.Errorf("err should be nil but got %v", err)
				}
			}
		})
	}
}

func TestUUID_String(t *testing.T) {
	tests := []struct {
		name string
		u    UUID
		want string
	}{
		{
			name: "ok",
			u:    UUID("811bd024-c18c-4b8f-9602-07e7cb60ac25"),
			want: "811bd024-c18c-4b8f-9602-07e7cb60ac25",
		},
		{
			name: "ok nil",
			u:    UUID("00000000-0000-0000-0000-000000000000"),
			want: "00000000-0000-0000-0000-000000000000",
		},
		{
			name: "nil",
			u:    UUID(""),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.String(); got != tt.want {
				t.Errorf("\ngot  :%v\nwant :%v", got, tt.want)
			}
		})
	}
}
