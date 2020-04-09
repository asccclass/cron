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
func main() {
   c := cron.New()
   c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
   c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
   c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
   c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
   c.AddFunc("@every 1m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
   c.Start()
   c.AddFunc("@daily", func() { fmt.Println("Every day") })
// inspect(c.Entries())
   defer c.Stop() // Stop the scheduler (does not stop any jobs already running).
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
