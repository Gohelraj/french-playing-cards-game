package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicatesFromStringSlice(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success",
			args: args{[]string{"AS", "KH", "QC", "KH", "AD", "2D", "AS"}},
			want: []string{"KH", "AS"},
		},
		{
			name: "success - no duplicates",
			args: args{[]string{"AS", "KH", "QC", "AD", "2D"}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindDuplicatesFromStringSlice(tt.args.values)
			assert.Equal(t, tt.want, got)
		})
	}
}
