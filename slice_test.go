package merge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice_Merge(t *testing.T) {
	type args[T any] struct {
		left  Slice[T]
		right Slice[T]
		want  Slice[T]
	}
	tests := []struct {
		name string
		args args[string]
	}{
		{
			name: "1",
			args: args[string]{
				left:  []string{"one"},
				right: []string{"two"},
				want:  []string{"one", "two"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.left.Merge(tt.args.right)
			assert.Equal(t, tt.args.want, tt.args.left)
		})
	}
}
