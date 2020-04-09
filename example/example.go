package main

import (
   "log"
   "os"
   "fmt"
   "time"
   "os/signal"
   "github.com/asccclass/cron"
)

func main() {
   c := CRON.NewCronJob()
   id, err := c.AddFunc("* * * * *", func() { fmt.Println("Every minutes...") })  // specific minutes
   if err != nil {
      fmt.Println(err.Error())
      return
   }
   fmt.Printf("Running job:%d\n", id)
   c.Start()
   defer c.Stop()
   select{} 

   go func(){
      for {
         time.Sleep(time.Second)
         log.Println("application is running.")
      }
   }()
   msgChan:=make(chan os.Signal,1)
   signal.Notify(msgChan, os.Interrupt, os.Kill)
   <-msgChan
}
