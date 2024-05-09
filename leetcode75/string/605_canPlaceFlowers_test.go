package string

import "testing"

func Test_canPlaceFlowers1(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 1},
				n:         2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers1(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPlaceFlowers2(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 1},
				n:         2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers2(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
