package main

import (
   // "github.com/line/line-bot-sdk-go/linebot"
   // "http"
   "log"
   "os"
   "time"
   "os/signal"
   "cron/core"
)

func main() {
   // client := &http.Client()
   // channelSecret := os.Getenv("channelSecret")
   // channelAccessToken := os.Getenv("channelAccessToken")
   // bot, err := linebot.New(channelSecret, channelAccessToken, linebot.WithHTTPClient(client))
   // if err != nil {
      // log.Println(err)
   // }
   wp, err := openweather.NewOPWeather("https://api.openweathermap.org/data/2.5/forecast", "bd6514afd0f6bd80c3c6a02cb522efd3", "Taipei")
   if err != nil {
      log.Println(err)
      return
   }
   c := core.NewCronJob()
   c.AddJob("* 1 2-23/3 * * *", wp.SendForcast)  // specific minutes
   // c.AddFunc("30 * * * * *", video.ChangeVideo)  // every second for test
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
