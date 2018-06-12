package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name     string
		input    string
		wantCmd  Command
		wantArgs []string
		wantErr  error
	}{
		{"uppercase", "GET key", new(Get), []string{"key"}, nil},
		{"lowercase", "get key", new(Get), []string{"key"}, nil},
		{"not_existing_command", "command_not_exists arg", nil, nil, ErrCommandNotFound},
		{"with_wrong_args_number", "GET key1 key2", nil, nil, ErrWrongArgsNumber},
		{"args_with_space", `SET key "value with space"`, new(Set), []string{"key", "value with space"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, args, err := Parse(tt.input)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.wantCmd, cmd)
			assert.Equal(t, tt.wantArgs, args)
		})
	}
}
