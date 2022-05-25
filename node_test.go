package graph

import (
	"reflect"
	"testing"
)

type Data struct {
	ID    int
	Value bool
}

func TestNewNode(t *testing.T) {
	type args struct {
		name  string
		value Data
	}
	tests := []struct {
		name string
		args args
		want *Node[string, Data]
	}{
		{
			name: "TestNewNode",
			args: args{
				name:  "A",
				value: Data{1, false},
			},
			want: &Node[string, Data]{
				parents:  []*Node[string, Data]{},
				children: []*Node[string, Data]{},
				weight:   0,
				id:       "A",
				value:    Data{1, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = \n%+#v\n, want \n%+#v", got, tt.want)
			}
		})
	}
}

func TestNode_AddChild(t *testing.T) {
	chld := NewNode("B", Data{2, false})

	type fields struct {
		parents  []*Node[string, Data]
		children []*Node[string, Data]
		weight   int
		id       string
		value    Data
	}
	type args struct {
		node *Node[string, Data]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node[string, Data]
	}{
		{
			name: "add child node",
			fields: fields{
				parents:  []*Node[string, Data]{},
				children: []*Node[string, Data]{},
				weight:   0,
				id:       "A",
				value:    Data{1, false},
			},
			args: args{
				node: chld,
			},
			want: &Node[string, Data]{
				parents: []*Node[string, Data]{},
				children: []*Node[string, Data]{
					chld,
				},
				weight: 0,
				id:     "A",
				value:  Data{1, false},
			},
		},
		{
			name: "add child node to child node",
			fields: fields{
				parents: []*Node[string, Data]{
					chld,
				},
				children: []*Node[string, Data]{},
				weight:   0,
				id:       "A",
				value:    Data{1, false},
			},
			args: args{
				node: chld,
			},
			want: &Node[string, Data]{
				parents: []*Node[string, Data]{
					chld,
				},
				children: []*Node[string, Data]{
					chld,
				},
				weight: 0,
				id:     "A",
				value:  Data{1, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node[string, Data]{
				parents:  tt.fields.parents,
				children: tt.fields.children,
				weight:   tt.fields.weight,
				id:       tt.fields.id,
				value:    tt.fields.value,
			}
			n.AddChild(tt.args.node)
			if !reflect.DeepEqual(n, tt.want) {
				t.Errorf("Node.AddChild() = \n%+#v\n, want \n%+#v", n, tt.want)
			}
		})
	}
}

func TestNode_WalkDFS(t *testing.T) {
	type fields struct {
		rootNode *Node[string, Data]
	}
	type test struct {
		name   string
		fields fields
		want   []*Node[string, Data]
	}
	tests := []test{
		func() test {
			rootNode := NewNode("A", Data{1, false})
			chld2 := NewNode("B", Data{2, false})
			chld3 := NewNode("C", Data{3, false})
			chld4 := NewNode("D", Data{4, false})

			rootNode.AddChild(chld2)
			chld2.AddChild(chld3)
			chld2.AddChild(chld4)

			wantNodeList := []*Node[string, Data]{
				rootNode,
				chld2,
				chld3,
				chld4,
			}

			return test{
				name: "add child node",
				fields: fields{
					rootNode: rootNode,
				},
				want: wantNodeList,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNodeList := []*Node[string, Data]{}
			tt.fields.rootNode.WalkDFS(func(n *Node[string, Data]) {
				gotNodeList = append(gotNodeList, n)
			})
			if !reflect.DeepEqual(gotNodeList, tt.want) {
				t.Errorf("Node.WalkDFS() got \n%+#v\n, want \n%+#v", gotNodeList, tt.want)
			}
		})
	}
}

func TestNode_WalkBFS(t *testing.T) {
	type fields struct {
		rootNode *Node[string, Data]
	}
	type test struct {
		name   string
		fields fields
		want   []*Node[string, Data]
	}
	tests := []test{
		func() test {
			rootNode := NewNode("A", Data{1, false})
			chld2 := NewNode("B", Data{2, false})
			chld3 := NewNode("C", Data{3, false})
			chld4 := NewNode("D", Data{4, false})

			rootNode.AddChild(chld2)
			chld2.AddChild(chld3)
			chld2.AddChild(chld4)

			wantNodeList := []*Node[string, Data]{
				rootNode,
				chld2,
				chld3,
				chld4,
			}

			return test{
				name: "add child node",
				fields: fields{
					rootNode: rootNode,
				},
				want: wantNodeList,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNodeList := []*Node[string, Data]{}
			tt.fields.rootNode.WalkBFS(func(n *Node[string, Data]) {
				gotNodeList = append(gotNodeList, n)
			})
			if !reflect.DeepEqual(gotNodeList, tt.want) {
				t.Errorf("Node.WalkBFS() got \n%+#v\n, want \n%+#v", gotNodeList, tt.want)
			}
		})
	}
}

func TestNode_Order(t *testing.T) {
	type args struct {
		rootNode *Node[string, Data]
	}
	type test struct {
		name string
		args args
		want map[int][]Node[string, Data]
	}
	tests := []test{
		{
			name: "order",
			args: args{
				rootNode: NewNode("A", Data{1, false}),
			},
			want: map[int][]Node[string, Data]{
				0: {
					*NewNode("A", Data{1, false}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.rootNode.Order()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Order() got \n%+#v\n, want \n%+#v", got, tt.want)
			}
		})
	}
}

func TestNode_ID(t *testing.T) {
	type args struct {
		rootNode *Node[string, Data]
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{
			name: "id a",
			args: args{
				rootNode: NewNode("A", Data{1, false}),
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.rootNode.ID()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.ID() got \n%+#v\n, want \n%+#v", got, tt.want)
			}
		})
	}
}

func TestNode_Value(t *testing.T) {
	type args struct {
		rootNode *Node[string, Data]
	}
	type test struct {
		name string
		args args
		want Data
	}
	tests := []test{
		{
			name: "value a",
			args: args{
				rootNode: NewNode("A", Data{1, false}),
			},
			want: Data{1, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.rootNode.Value()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Value() got \n%+#v\n, want \n%+#v", got, tt.want)
			}
		})
	}
}

func TestNode_Weight(t *testing.T) {
	type args struct {
		rootNode *Node[string, Data]
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{
			name: "weight a",
			args: args{
				rootNode: NewNode("A", Data{1, false}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.rootNode.Weight()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Weight() got \n%+#v\n, want \n%+#v", got, tt.want)
			}
		})
	}
}
