package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateInvalidCardCodeError(t *testing.T) {
	type args struct {
		cardCode string
	}
	tests := []struct {
		name       string
		args       args
		wantErrMsg string
	}{
		{
			name:       "success",
			args:       args{cardCode: "AS"},
			wantErrMsg: "Card code AS is invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GenerateInvalidCardCodeError(tt.args.cardCode)
			assert.Equal(t, tt.wantErrMsg, err.Error())
		})
	}
}
