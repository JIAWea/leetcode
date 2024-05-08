package string

import "testing"

func Test_reverseWords1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				"the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "test2",
			args: args{
				"  hello world  ",
			},
			want: "world hello",
		},
		{
			name: "test3",
			args: args{
				"a good   example",
			},
			want: "example good a",
		},
		{
			name: "test4",
			args: args{
				" ",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords1(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
