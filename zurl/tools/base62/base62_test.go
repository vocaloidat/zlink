package base62

import "testing"

func TestEncodeIntToBase62(t *testing.T) {
	type args struct {
		num uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "基本测试",
			args: args{
				num: 12345,
			},
			want: "3D7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeIntToBase62(tt.args.num); got != tt.want {
				t.Errorf("EncodeIntToBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}
