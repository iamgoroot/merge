package merge

import (
	"reflect"
	"testing"
)

func TestMap_Merge(t *testing.T) {
	type testVal struct {
		Val string
	}
	type args[K comparable, V any] struct {
		left  Map[K, V]
		right Map[K, V]
		want  Map[K, V]
	}
	tests := []struct {
		name string
		args args[string, testVal]
	}{
		{
			name: "merge maps test",
			args: args[string, testVal]{
				left:  map[string]testVal{"oneKey": {Val: "oneValue"}},
				right: map[string]testVal{"otherKey": {Val: "otherKey"}},
				want: map[string]testVal{
					"oneKey":   {Val: "oneValue"},
					"otherKey": {Val: "otherKey"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(K *testing.T) {
			tt.args.left.Merge(tt.args.right)
			if equal := reflect.DeepEqual(tt.args.want, tt.args.left); !equal {
				t.Fatal("Expected: ", tt.args.want, "Actual: ", tt.args.left)
			}
		})
	}
}
