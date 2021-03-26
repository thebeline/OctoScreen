package utils

import (
	"time"
	"sync"

	"github.com/gotk3/gotk3/glib"
)

type BackgroundTask struct {
	sync.Mutex

	duration			time.Duration
	task				func()
	close				chan bool
	isRunning			bool
}

func CreateBackgroundTask(
	duration			time.Duration,
	task				func(),
) *BackgroundTask {
	thisInstance := &BackgroundTask{
		task:			task,
		duration: 		duration,
		close: 			make(chan bool, 1),
		isRunning:		false,
	}

	return thisInstance
}

func (this *BackgroundTask) Start() {
	this.Lock()
	defer this.Unlock()

	Logger.Info("New background task started")
	go this.loop()

	this.isRunning = true
}

func (this *BackgroundTask) Close() {
	this.Lock()
	defer this.Unlock()
	if !this.isRunning {
		return
	}

	this.close <- true
	this.isRunning = false
}

func (this *BackgroundTask) loop() {
	this.execute()

	ticker := time.NewTicker(this.duration)
	defer ticker.Stop()
	for {
		select {
			case <-ticker.C:
				this.execute()

			case <-this.close:
				Logger.Info("Background task closed")
				return
		}
	}
}

func (this *BackgroundTask) execute() {
	SourceHandle := glib.IdleAdd(this.task)
	if SourceHandle == nil {
		LogFatalError("common.execute()", "IdleAdd()")
	}
}
