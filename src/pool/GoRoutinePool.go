package main

import (
	"fmt"
	"time"
)

type GoRoutinePool struct {
	WorkQueue       chan func() error
	MaximumPoolSize int
	CorePoolSize    int
	Handler         *RejectedExecutionHandler
	result          chan error
	finishCallback  func()
}

func (self *GoRoutinePool) Init(MaximumPoolSize int, CorePoolSize int, Handler *RejectedExecutionHandler) {
	self.CorePoolSize = CorePoolSize
	self.MaximumPoolSize = MaximumPoolSize
	self.Handler = Handler
	self.WorkQueue = make(chan func() error, CorePoolSize)
	self.result = make(chan error, CorePoolSize)
}

func (self *GoRoutinePool) Start() {
	for i := 0; i < self.CorePoolSize; i++ {
		go func() {
			for {
				task, ok := <-self.WorkQueue
				if !ok {
					break
				}
				err := task()
				self.result <- err
			}
		}()
	}
	// 获取每个work的执行结果
	for j := 0; j < self.CorePoolSize; j++ {
		res, ok := <-self.result
		if !ok {
			break
		}
		if nil != res {
			fmt.Println(res)
		}
	}
	// 所有任务都执行完成，回调函数
	if self.finishCallback != nil {
		self.finishCallback()
	}

}

func (self *GoRoutinePool) Stop() {
	close(self.WorkQueue)
	close(self.result)
}

// 添加任务
func (self *GoRoutinePool) AddTask(task func() error) {
	self.WorkQueue <- task
}

// 设置回调
func (self *GoRoutinePool) SetFinishCallback(callback func()) {
	self.finishCallback = callback
}

type RejectedExecutionHandler struct {
}

func main() {
	pool := new(GoRoutinePool)
	pool.Init(2, 2, new(RejectedExecutionHandler))

	for i := 0; i < 2; i++ {

		pool.AddTask(func() error {
			fmt.Println("执行任务中......")
			//time.Sleep(1 * time.Second)
			return nil
		})
	}
	// 设置回调
	isFinish := false
	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			*isFinish = true
		}(&isFinish)
	})

	pool.Start()
	for !isFinish {
		time.Sleep(time.Millisecond * 100)
	}

	pool.Stop()
	fmt.Println("所有操作已完成！")

}
