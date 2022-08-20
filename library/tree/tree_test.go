package tree

import (
	"reflect"
	"testing"
)

func TestBfs(t *testing.T) {
	type args struct {
		treeMap map[int][]Node
	}
	tests := []struct {
		name      string
		args      args
		wantItems []Node
	}{
		// TODO: Add test cases.
		{
			name: "test 01",
			args: args{
				treeMap: map[int][]Node{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItems := Bfs(tt.args.treeMap); !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("Bfs() = %v, want %v", gotItems, tt.wantItems)
			}
		})
	}
}