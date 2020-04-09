package CRON 

import (
   "github.com/robfig/cron"
)

// Jub Que
type CronJobQue struct {
   EntryID		cron.EntryID	`json:"entryID"`
   CycleStatus		string		`json:"setting"`	// 設定方式
   LastExecuteTime	string		`json:"lastexecute"`	// 最後執行時間
}

// cron job struct
type CronJob struct {
   Cronjob 	*cron.Cron
   CronQue	[]CronJobQue
}

func NewCronJob() (*CronJob) {
   c := cron.New()
   d := []CronJobQue{}

   return &CronJob {
      Cronjob: c,
      CronQue: d,
   }
}

// AddFunc adds a func to the Cron to be run on the given schedule. 
func(core *CronJob) AddJob(timestring string, f cron.Job) (error) {
   id, err := core.Cronjob.AddJob(timestring, f)
   if err != nil {
      return err
   }
   que := CronJobQue {
      EntryID: id,
      CycleStatus: timestring,
      LastExecuteTime: "",
   }
   core.CronQue = append(core.CronQue, que)
   return nil
}

// AddJob adds a Job to the Cron to be run on the given schedule.
func(core *CronJob) AddFunc(timestring string, f func()) (cron.EntryID, error) {
   id, err := core.Cronjob.AddFunc(timestring, f)
   if err != nil {
      return 0, err
   }
   que := CronJobQue {
      EntryID: id,
      CycleStatus: timestring,
      LastExecuteTime: "",
   }
   core.CronQue = append(core.CronQue, que)
   return 0, nil
}

func(core *CronJob) Start() {
   core.Cronjob.Start()
}

// Remove an entry from being run in the future.
func(core *CronJob) Remove(id cron.EntryID) {
   for _, job := range core.CronQue {
      if job.EntryID == id {
         core.Cronjob.Remove(job.EntryID)
      }
   }
}

func(core *CronJob) Stop() {
   core.Cronjob.Stop()
}
