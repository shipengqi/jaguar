package gormutil

import (
	"reflect"
	"testing"

	"github.com/AlekSi/pointer"
)

func TestDePointer(t *testing.T) {
	type args struct {
		offset *int64
		limit  *int64
	}
	tests := []struct {
		name string
		args args
		want *LimitAndOffset
	}{
		{
			name: "both offset and limit are not zero",
			args: args{
				offset: pointer.ToInt64(0),
				limit:  pointer.ToInt64(10),
			},
			want: &LimitAndOffset{
				Offset: 0,
				Limit:  10,
			},
		},
		{
			name: "both offset and limit are zero",
			want: &LimitAndOffset{
				Offset: 0,
				Limit:  1000,
			},
		},
		{
			name: "offset not zero and limit zero",
			args: args{
				offset: pointer.ToInt64(2),
			},
			want: &LimitAndOffset{
				Offset: 2,
				Limit:  1000,
			},
		},
		{
			name: "offset zero and limit not zero",
			args: args{
				limit: pointer.ToInt64(10),
			},
			want: &LimitAndOffset{
				Offset: 0,
				Limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DePointer(tt.args.offset, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzDePointer(f *testing.F) {
	testcases := []int64{1, 2, 3, 4, 5}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, in int64) {
		out := DePointer(pointer.ToInt64(0), &in)
		want := &LimitAndOffset{
			Offset: 0,
			Limit:  int(in),
		}
		if !reflect.DeepEqual(out, want) {
			t.Errorf("got: %v, want: %v", out, want)
		}
	})
}
