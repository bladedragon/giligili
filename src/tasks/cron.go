package tasks

import (
	"fmt"
	"github.com/robfig/cron"
	"reflect"
	"runtime"
	"time"
)

var Cron *cron.Cron

func Run(job func() error) {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}
	//每年一次 更新排行榜
	Cron.AddFunc("0 0 0 1 1 *", func() { Run(RestartDailyRank) })
	//每日一次 更新投稿数
	Cron.AddFunc("0 0 0 * * *", func() { Run(RestartPucnt) })
	Cron.Start()

	fmt.Println("Cronjob start.....")
}
