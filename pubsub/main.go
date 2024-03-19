package main

import (
	"fmt"
	"go_pubsub/pkg"
)

func main() {
	mps := pkg.NewMapPubSub[string]()
	defer mps.Close()
	mps.Subscribe("topic1")
	mps.Publish("topic1","halo")
	mps.Publish("topic1","dunia")
	mps.Subscribe("topic2")
	mps.Publish("topic2","halo")
	fmt.Scanln()

}

// func main() {
// 	ps := pkg.NewPubSub[string]()
// 	wg := sync.WaitGroup{}
// 	// subscribe one
// 	wg.Add(1)
// 	subOne := ps.Subscribe()
// 	go func() {
// 		for {
// 			select {
// 			case value , ok := <- subOne:
// 				if !ok {
// 					fmt.Println("exiting sub one")
// 					wg.Done()
// 					return
// 				}
// 				fmt.Printf("value : %s \n",value)
// 			}
// 		}
// 	}()

// 	ps.Publish("hello")
// 	ps.Publish("world")
// 	ps.Close()
// 	wg.Wait()
// 	fmt.Println("finish")
// }