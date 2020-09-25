package password

import "testing"

func TestGenerateForCharset2(t *testing.T) {
	type args struct {
		n     int
		chset Charset
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Uppercase",
			args: args{
				n:     8,
				chset: Uppercase,
			},
		},
		{
			name: "Lowercase",
			args: args{
				n:     8,
				chset: Lowercase,
			},
		},
		{
			name: "Numbers",
			args: args{
				n:     8,
				chset: Numbers,
			},
		},
		{
			name: "Symbols",
			args: args{
				n:     8,
				chset: Symbols,
			},
		},
		{
			name: "Uppercase | Lowercase",
			args: args{
				n:     8,
				chset: Uppercase | Lowercase,
			},
		},
		{
			name: "Uppercase | Lowercase | Numbers",
			args: args{
				n:     8,
				chset: Uppercase | Lowercase | Numbers,
			},
		},
		{
			name: "Uppercase | Lowercase | Symbols",
			args: args{
				n:     8,
				chset: Uppercase | Lowercase | Symbols,
			},
		},
		{
			name: "Uppercase | Lowercase | SimpleSymbols",
			args: args{
				n:     8,
				chset: Uppercase | Lowercase | SimpleSymbols,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateForCharset(tt.args.n, tt.args.chset)
			if len(got) != tt.args.n {
				t.Errorf("GenerateForCharset(%v, %v) returned password with length %v", tt.args.n, tt.args.chset, len(got))
			} else if !uses_charset(got, tt.args.chset) {
				t.Errorf("GenerateForCharset(%v, %v) uses unexpected charset", tt.args.n, tt.args.chset)
			}
		})
	}
}

func uses_charset(str string, chset Charset) bool {
	data := []byte(str)
	for i := 0; i < len(data); i++ {
		if chset&Uppercase != 0 {
			if contains(uppercase, data[i]) {
				continue
			}
		}
		if chset&Lowercase != 0 {
			if contains(lowercase, data[i]) {
				continue
			}
		}
		if chset&Numbers != 0 {
			if contains(numbers, data[i]) {
				continue
			}
		}
		if chset&Symbols != 0 {
			if contains(symbols, data[i]) {
				continue
			}
		}
		if chset&SimpleSymbols != 0 {
			if contains(simple_symbols, data[i]) {
				continue
			}
		}
		return false
	}
	return true
}

func contains(a []byte, ch byte) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == ch {
			return true
		}
	}
	return false
}
