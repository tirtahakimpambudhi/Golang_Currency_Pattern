package errgroup

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func read(filename string) (chan []string,error){
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

func WaitGroup() chan struct{} {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)
	for _, file := range []string{"one.csv","two.csv","three.csv"} {
		fl := file
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch , err := read(fl)
			if err != nil {
				fmt.Printf("error reading %s \n",err.Error())
			}
			for line := range ch {
				fmt.Println(line)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
	
}