package namegrind

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		errStr string
	}{
		{
			"cannot have zero length",
			"",
			"must have nonzero length",
		},
		{
			"cannot be over MaxNameLen characters",
			"longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong",
			"over maximum length",
		},
		{
			"cannot start with a hyphen",
			"-startswithhyphen",
			"cannot start with a hyphen",
		},
		{
			"cannot end with a hyphen",
			"endwithhyphen-",
			"cannot end with a hyphen",
		},
		{
			"cannot have consecutive hyphens",
			"consecutive--hyphens",
			"cannot contain consecutive hyphens",
		},
		{
			"cannot contain unicode",
			"我叫鸣字",
			"invalid character",
		},
		{
			"can only contain certain characters",
			"hello!",
			"invalid character",
		},
		{
			"can only contain certain characters",
			"hello@",
			"invalid character",
		},
		{
			"works with valid names",
			"heres-a-valid-name",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.errStr == "" {
				require.Nil(t, ValidateName(tt.input))
				return
			}

			err := ValidateName(tt.input)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.errStr)
		})
	}
}


func TestRollout(t *testing.T) {
	in := []struct {
		name   string
		week   int
		height int
	}{
		{
			"honk",
			13,
			15120,
		},
		{
			"beep",
			15,
			17136,
		},
		{
			"farb",
			36,
			38304,
		},
	}

	for _, tt := range in {
		hash, err := HashName(tt.name)
		require.NoError(t, err)
		height, week := Rollout(hash)
		require.Equal(t, tt.height, height)
		require.Equal(t, tt.week, week)
	}
}
