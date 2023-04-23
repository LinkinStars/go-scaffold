package cryptor

import (
	"testing"
)

func TestAesSimpleEncrypt(t *testing.T) {
	type args struct {
		data string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				data: "Hello World!",
				key:  "16bit secret key",
			},
			want: "PuMhKY8ZFLnDAwlQ7v/2SQ==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AesSimpleEncrypt(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("AesSimpleEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesSimpleDecrypt(t *testing.T) {
	type args struct {
		data string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				key:  "16bit secret key",
				data: "PuMhKY8ZFLnDAwlQ7v/2SQ==",
			},
			want: "Hello World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AesSimpleDecrypt(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("AesSimpleDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenIVFromKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		args   args
		wantIv string
	}{
		{
			name: "test",
			args: args{
				key: "16bit secret key",
			},
			wantIv: "ba79295cdabd3a86",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIv := GenIVFromKey(tt.args.key); gotIv != tt.wantIv {
				t.Errorf("GenIVFromKey() = %v, want %v", gotIv, tt.wantIv)
			}
		})
	}
}

func TestAesEncrypt(t *testing.T) {
	type args struct {
		data        string
		key         string
		iv          string
		paddingMode PaddingMode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				data:        "Hello World!",
				key:         "16bit secret key",
				iv:          "ba79295cdabd3a86",
				paddingMode: PKCS7,
			},
			want: "PuMhKY8ZFLnDAwlQ7v/2SQ==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AesEncrypt(tt.args.data, tt.args.key, tt.args.iv, tt.args.paddingMode); got != tt.want {
				t.Errorf("AesEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesDecrypt(t *testing.T) {
	type args struct {
		data        string
		key         string
		iv          string
		paddingMode PaddingMode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				data:        "PuMhKY8ZFLnDAwlQ7v/2SQ==",
				key:         "16bit secret key",
				iv:          "ba79295cdabd3a86",
				paddingMode: PKCS7,
			},
			want: "Hello World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AesDecrypt(tt.args.data, tt.args.key, tt.args.iv, tt.args.paddingMode); got != tt.want {
				t.Errorf("AesDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
