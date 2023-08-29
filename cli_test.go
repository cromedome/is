package main

import (
	"testing"

	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

const tmux = "tmux"

func TestCliVersion(t *testing.T) {
	t.Parallel()
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = tmux
		cmd.Version.Op = ops.Ne
		cmd.Version.Val = "1"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = "tmuxzzz"
		cmd.Version.Op = ops.Ne
		cmd.Version.Val = "1"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = tmux
		cmd.Version.Op = "eq"
		cmd.Version.Val = "1"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = tmux
		cmd.Version.Op = "eq"
		cmd.Version.Val = "zzz"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: false}
		cmd := CLICmd{}
		cmd.Version.Name = tmux
		cmd.Version.Op = ops.Unlike
		cmd.Version.Val = "zzz"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: false}
		cmd := CLICmd{}
		cmd.Version.Name = tmux
		cmd.Version.Op = ops.Like
		cmd.Version.Val = ""
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: false}
		cmd := CLICmd{}
		cmd.Version.Name = tmux
		cmd.Version.Op = ops.Like
		cmd.Version.Val = "3.*"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
}

func TestCliAge(t *testing.T) {
	t.Parallel()
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Age.Name = tmux
		cmd.Age.Op = "gt"
		cmd.Age.Val = "1"
		cmd.Age.Unit = "s"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Age.Name = tmux
		cmd.Age.Op = "lt"
		cmd.Age.Val = "100000"
		cmd.Age.Unit = "days"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Age.Name = tmux
		cmd.Age.Op = "lt"
		cmd.Age.Val = "1.1"
		cmd.Age.Unit = "d"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{}
		cmd.Age.Name = "tmuxxx"
		cmd.Age.Op = "lt"
		cmd.Age.Val = "1"
		cmd.Age.Unit = "d"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
}
