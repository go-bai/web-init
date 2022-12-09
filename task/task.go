package task

import (
	"time"

	"github.com/robfig/cron/v3"
)

var T = cron.New(cron.WithLocation(time.Local), cron.WithSeconds())

func init() {
	// T.AddFunc("*/5 * * * * *", func() {
	// 	log.Println("定时任务")
	// })
}
