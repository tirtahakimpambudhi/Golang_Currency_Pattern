package test

import (
	"context"
	errgroup "go_errgroup"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	wait := errgroup.WaitGroup()
	<- wait
	t.Log("finish")
}

func TestErrGroup(t *testing.T) {
	ctx := context.Background()
	wait := errgroup.ErrGroup(ctx)
	<- wait
	t.Log("finish")
}