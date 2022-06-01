package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	commands := []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}
	values := [][]int{{}, {1}, {2}, {}, {3}, {}}
	for idx := 0; idx < b.N; idx++ {
		Run(commands, values)
	}
}
func TestRun(t *testing.T) {
	type args struct {
		commands []string
		values   [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{
				commands: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
				values:   [][]int{{}, {1}, {2}, {}, {3}, {}},
			},
			want: []string{"null", "null", "null", "1.5", "null", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.commands, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
