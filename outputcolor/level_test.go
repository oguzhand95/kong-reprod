// Copyright 2021-2024 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package outputcolor_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oguzhand95/kong-reprod/outputcolor"
)

func TestKong(t *testing.T) {
	tests := []struct {
		args      []string
		wantLevel outputcolor.Level
		wantErr   string
	}{
		{
			args:    []string{"-cfoo"},
			wantErr: "unknown flag -f",
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			level, err := parse(t, tt.args)

			if tt.wantErr != "" {
				require.Error(t, err, tt.wantErr)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantLevel, *level)
		})
	}
}

func parse(t *testing.T, args []string) (*outputcolor.Level, error) {
	t.Helper()

	var cli struct {
		Color     *outputcolor.Level `short:"c"`
		Leftovers []string           `arg:"" optional:""`
	}

	parser, err := kong.New(&cli, outputcolor.TypeMapper)
	require.NoError(t, err, "failed to create command-line argument parser")

	var parseError *kong.ParseError
	_, err = parser.Parse(args)
	if err == nil || errors.As(err, &parseError) {
		return cli.Color, err
	}
	require.NoError(t, err, "failed to parse command-line arguments")

	return nil, nil
}
