package fanin_and_fanout

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

// Merge for multiple channel into one channel
func Merge(chs ...chan []string) chan []string {
	var wg sync.WaitGroup
	out := make(chan []string)
	send := func (ch chan []string)  {
		defer wg.Done()
		for c := range ch {
			out <- c
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go send(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func MergeWithOutWait(chs ...chan []string) chan []string {
	wait := make(chan struct{},len(chs)) 
	out := make(chan []string)
	send := func (ch chan []string)  {
		defer func ()  {
			wait <- struct{}{}
		}()
		for c := range ch {
			out <- c
		}
	}
	for _, ch := range chs {
		go send(ch)
	}
	go func() {
		<- wait
		close(out)
	}()
	return out
}
//BreakUp for One Channel into multiple channel
func BreakUp(worker string, ch chan []string) chan []string {
	out := make(chan []string)
	go func() {
		for value := range ch {
			fmt.Printf("%s %v \n",worker,value)
		}
		close(out)
	}()
	return out
}

func ReadCSV(filename string) (chan []string,error){
	reader, err := os.Open(filename)
	if err != nil {
		return nil,err
	}
	csr := csv.NewReader(reader)
	ch := make(chan []string)
	go func() {
		for {
			line, err := csr.Read()
			if err == io.EOF {
				close(ch)
				return
			}
			ch <- line
		}
	}()
	return ch,nil
}