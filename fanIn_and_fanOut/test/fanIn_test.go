package test

import (
	fanin_and_fanout "fifo"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFanIn(t *testing.T) {
	csvOne , err := fanin_and_fanout.ReadCSV("One.csv")
	require.NoError(t,err)
	csvTwo , err := fanin_and_fanout.ReadCSV("Two.csv")
	require.NoError(t,err)
	exit := make(chan struct{})
	chM := fanin_and_fanout.Merge(csvOne,csvTwo)
	require.NotNil(t,chM)
	go func() {
		for ch := range chM {
			t.Log(ch)
		}
		close(exit)
	}()
	<- exit
	t.Log("Finish")
}


func TestFanInWithOutWait(t *testing.T) {
	csvOne , err := fanin_and_fanout.ReadCSV("One.csv")
	require.NoError(t,err)
	csvTwo , err := fanin_and_fanout.ReadCSV("Two.csv")
	require.NoError(t,err)
	exit := make(chan struct{})
	chM := fanin_and_fanout.MergeWithOutWait(csvOne,csvTwo)
	require.NotNil(t,chM)
	go func() {
		for ch := range chM {
			t.Log(ch)
		}
		close(exit)
	}()
	<- exit
	t.Log("Finish")
}