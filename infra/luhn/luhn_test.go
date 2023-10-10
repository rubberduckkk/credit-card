package luhn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidLuhn(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case normal credit card num",
			args: args{digits: []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4}},
			want: true,
		},
		{
			name: "case invalid credit card",
			args: args{digits: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsValidLuhn(tt.args.digits))
		})
	}
}
