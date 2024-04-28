package test

import (
	errgroup "go_errgroup"
	"testing"

)

func TestWaitGroup(t *testing.T) {
	wait := errgroup.WaitGroup()
	<- wait
	t.Log("finish")
}