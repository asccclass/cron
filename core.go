package main

import (
   "github.com/robfig/cron"
)

// Jub Que
type CronJobQue struct {
   EntryID		cron.EntryID	`json:"entryID"`
   CycleStatus		string		`json:"setting"`	// 設定方式
   LastExecuteTime	string		'json:"lastexecute"`	// 最後執行時間
}

// cron job struct
type CronJob struct {
   cronjob 	*cron.Cron
   CronQue	*[]CronJobQue
}

func NewCronJob() (*CronJob) {
   c := cron.New()

   return &CronJob {
      cronjob: c,
   }
}

// AddFunc adds a func to the Cron to be run on the given schedule. 
func(core *CronJob) AddJob(timestring string, f func()) (error) {
   id, err := core.cronjob.AddJob(timestring, f)
   if err != nil {
      return err
   }
   que := &CronJob{
      EntryID: id,
      CycleStatus: timestring,
      LastExecureTime: "",
   }
   core.CronQue = append(core.CronQue, que)
   return nil
}

// AddJob adds a Job to the Cron to be run on the given schedule.
func(core *CronJob) AddFunc(timestring string, f func()) (cron.EntryID, error) {
   id, err := core.cronjob.AddFunc(timestring, f)
   if err != nil {
      return err
   }
   que := &CronJob{
      EntryID: id,
      CycleStatus: timestring,
      LastExecureTime: "",
   }
   core.CronQue = append(core.CronQue, que)
   return nil
}

func(core *CronJob) Start() {
   core.cronjob.Start()
}

// Remove an entry from being run in the future.
func(core *CronJob) Remove(id int) {
   for _, job := range core.CronQue {
      if job.EntryID == id {
         core.cronjob.Remove(job.EntryID)
      }
   }
}

func(core *CronJob) Stop() {
   core.cronjob.Stop()
}
