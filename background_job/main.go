package main

import (
	"fmt"
	"go_bj/schedule"
	"os"
)

func main()  {
	fmt.Printf("Process ID %d \n",os.Getpid())
	s := schedule.NewScheduler(5,10)
	s.ListenAndWork()
	<-schedule.WaitToExit()
	s.Exit()
	fmt.Println("exiting")
}