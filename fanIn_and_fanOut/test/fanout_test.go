package test


import (
	fanin_and_fanout "fifo"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFanOut(t *testing.T) {
	csvOne , err := fanin_and_fanout.ReadCSV("One.csv")
	require.NoError(t,err)
	channelTwo := fanin_and_fanout.BreakUp("2",csvOne)
	channelOne := fanin_and_fanout.BreakUp("1",csvOne)

	for {
		if channelOne == nil && channelTwo == nil {
			break
		}
		select {
		case _,ok := <- channelOne :
			if !ok {
				channelOne = nil
			}
		case _,ok := <- channelTwo :
			if !ok {
				channelTwo = nil
			}
		}
	}
	t.Log("Finish")
}