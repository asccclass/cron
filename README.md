### Sherry cron job Tool

### Installation
```
GO111MODULE=on go mod download
make run
```

* github.com/robfig/cron

### Google sheet

```
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
```

### Example

```
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
```

* Set time example

```
   c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
   c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
   c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
   c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
   c.AddFunc("@every 1m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
   c.Start()
   c.AddFunc("@daily", func() { fmt.Println("Every day") })
```
