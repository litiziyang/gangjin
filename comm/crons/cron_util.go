package crons

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var mutex sync.Mutex

//任务列表
var taskList []timeTask

var log = logrus.New()

// 初始化并启动定时任务
func InitProcessTimer() error {
	c := cron.New()
	_, err := c.AddFunc("@every 30s", ProcessTimerTaskHandler)
	if err != nil {
		return err
	}
	c.Start()
	return nil
}

//定时任务类
type timeTask struct {
	id      string
	create  time.Time
	timeOut time.Duration
	cb      ProcessTimerCallback
}

//要执行的回调函数
type ProcessTimerCallback func(id string, err error) error

//插入定时任务
func InserTimerTask(id string, timeout time.Duration, cb ProcessTimerCallback) {
	var task timeTask
	task.id = id
	task.create = time.Now()
	task.timeOut = timeout
	task.cb = cb

	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < len(taskList); i++ {
		if taskList[i].id == task.id {
			taskList[i].create = task.create
			taskList[i].timeOut = task.timeOut
			return
		}
	}
	taskList = append(taskList, task)
}

// 定时处理任务
func ProcessTimerTaskHandler() {
	var task timeTask
	mutex.Lock()
	defer mutex.Unlock()
	if len(taskList) == 0 {
		return
	}
	for i := 0; i < len(taskList); {
		task = taskList[i]
		if time.Now().Sub(task.create) > task.timeOut {
			err := task.cb(task.id, nil)
			if err != nil {
				return
			}
			taskList = append(taskList[:i], taskList[i+1:]...)
			log.Printf("%s任务被触发", task)
		} else {
			log.Printf("%s任务已经过期", task)
			i++

		}
	}
}

func RemoveTimeTask(id string) {
	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < len(taskList); i++ {
		if taskList[i].id == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			taskList = append(taskList, taskList[i])
			return
		}
	}
}
