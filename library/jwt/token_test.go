package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeB64Header(t *testing.T) {
	data := "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
	header := base64.StdEncoding.EncodeToString([]byte(data))
	if header != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" {
		t.Fatal("encode err")
	}
}

func TestDecodeB64Header(t *testing.T) {
	// "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		t.Fatal(err)
	}
	ob := map[string]interface{}{}
	err = json.Unmarshal([]byte(data), &ob)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%q e:%v ob:%v\n", data, err, ob["uid"])
}

func TestToMd5(t *testing.T) {
	type args struct {
		encode string
	}
	tests := []struct {
		name       string
		args       args
		wantDecode string
	}{
		// TODO: Add test cases.
		{
			name: "md5 test",
			args: args{
				encode: "12345678",
			},
			wantDecode: "25d55ad283aa400af464c76d713c07ad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDecode := ToMd5(tt.args.encode); gotDecode != tt.wantDecode {
				t.Errorf("ToMd5() = %v, want %v", gotDecode, tt.wantDecode)
			}
		})
	}
}

func TestSecSecret(t *testing.T) {
	uid := "2"
	salt := "test"
	secret := SecSecret(uid, salt)
	t.Log(secret)
}

func TestCreateToken(t *testing.T) {
	type args struct {
		uid    string
		secret string
		exp    int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int64
		wantErr bool
	}{
		{
			name: "test ok",
			args: args{
				uid:    "1",
				secret: "test",
				exp:    33179538730,
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjozMzE3OTUzODczMCwiaXNzIjoiMSJ9.NCAkDlRtG4uqI3SlTtrnGn_SrnsUttzO_AHa6WvZAoA",
		},
		{
			name: "test exp",
			args: args{
				uid:    "1",
				secret: "test",
				exp:    1609497130,
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjA5NDk3MTMwLCJpc3MiOiIxIn0.0MMoqdLT40vMzk3p-uRTAdKolbIayiOAfhtZ5oNu0x8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := SecSecret(tt.args.uid, tt.args.secret)
			t.Logf("secret:%v", secret)
			got, err := CreateToken(tt.args.uid, secret, tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, \nwantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = \n%v, \nwant %v", got, tt.want)
			}
		})
	}
}

func TestAuthToken(t *testing.T) {
	type args struct {
		signedToken string
		secret      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "authToken ",
			args: args{
				signedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjozMzE3OTUzODczMCwiaXNzIjoiMSJ9.NCAkDlRtG4uqI3SlTtrnGn_SrnsUttzO_AHa6WvZAoA",
				secret:      "test",
			},
			want:    "1",
			wantErr: false,
		},
		{
			// 密码错误
			name: "authToken err sercret",
			args: args{
				signedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjA5NDk3MTMwLCJpc3MiOiIxIn0.0MMoqdLT40vMzk3p-uRTAdKolbIayiOAfhtZ5oNu0x8",
				secret:      "test",
			},
			want:    "",
			wantErr: true,
		},
	}
	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := SecSecret(tt.want, tt.args.secret)
			t.Logf("secret:%v", secret)
			got, err := AuthToken(tt.args.signedToken, secret)
			t.Log("index:", index, "err:", err, "got:", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthToken() error = %v, wantErr %v", err, tt.wantErr)
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("AuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetUid(t *testing.T) {
	type args struct {
		signedToken string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test get uid",
			args: args{
				signedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjA5NDk3MTMwLCJpc3MiOiIxIn0.0MMoqdLT40vMzk3p-uRTAdKolbIayiOAfhtZ5oNu0x8",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "6",
			args: args{
				signedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2IiwiZXhwIjoxMDAwLCJpc3MiOiI2In0.vcuNmTF5egtzRwqFMuUrN2ZmST8p3WYPtt0MVfMX-yI",
			},
			want:    "6",
			wantErr: false,
		},
	}
	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUid(tt.args.signedToken)
			t.Log("index:", index, "err:", err, "got:", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthToken() error = %v, wantErr %v", err, tt.wantErr)
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("AuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
