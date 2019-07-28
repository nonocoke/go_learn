package goTest

import (
	"github.com/google/go-cmp/cmp"
	//"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct{
		input string
		sep string
		want []string
	}{
		"simple": {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"no sep": {input: "abc", sep: "/", want: []string{"abc"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Fatalf("%s expected: %#v, got: %#v", name, tc.want, got)
			//}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
