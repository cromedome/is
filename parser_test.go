package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLIVersion(t *testing.T) {
	ctx := &Context{}
	assert.Equal(t, "1.20.4", cliVersion(ctx, "go", "go version go1.20.4 darwin/amd64"))
	assert.Equal(t, "3.3a", cliVersion(ctx, "tmux", "tmux 3.3a"))
	assert.Equal(t, "0.0.24", cliVersion(ctx, "ubi", "ubi 0.0.24"))
	assert.Equal(t, "v5.36.0", cliVersion(
		ctx,
		"perl",
		`This is perl 5, version 36, subversion 0 (v5.36.0) built for darwin-2level`),
	)
}

func TestCLIOutput(t *testing.T) {
	ctx := &Context{}
	{
		o, err := (cliOutput(ctx, "tmux"))
		assert.NoError(t, err)
		assert.NotEmpty(t, o)
	}

	{
		o, err := (cliOutput(ctx, "tmuxxx"))
		assert.Error(t, err)
		assert.Empty(t, o)
	}
}
