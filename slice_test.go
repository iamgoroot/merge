package merge

import (
	"reflect"
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
			if equal := reflect.DeepEqual(tt.args.want, tt.args.left); !equal {
				t.Fatal("Expected: ", tt.args.want, "Actual: ", tt.args.left)
			}
		})
	}
}
