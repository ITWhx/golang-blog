/**
*
*@author 吴昊轩
*@create 2020-04-1616:13
 */
package main

import (
	"golang-blog/db"
	"strconv"
	"sync/atomic"
	"time"
)

type FeedMgr struct {
	feedProcessor *FeedProcessor
}

func InitFeedMgr() {

	mgr := FeedMgr{
		feedProcessor: &FeedProcessor{
			redisClient: db.G_redisClient,
			inChan:      make(chan *PushJob, 10),
			outChan:     make(chan *PushJob, 10),
			closeChan:   make(chan struct{}),
		},
	}

	var cnt int64 = 0

	go func() {
		for {
			t := time.Now().Unix()
			tmp := atomic.AddInt64(&cnt, 1)
			job := &PushJob{"uid-1", "uid-1,msg " + strconv.FormatInt(tmp, 10), t}
			time.Sleep(time.Millisecond * 500)

			mgr.feedProcessor.ReadJob(job)
		}
	}()
	//go func() {
	//	for {
	//		t := time.Now().Unix()
	//		tmp := atomic.AddInt64(&cnt, 1)
	//		job := &PushJob{"2", "uid-2,msg " + strconv.FormatInt(tmp, 10), t}
	//		time.Sleep(time.Millisecond * 500)
	//		mgr.feedProcessor.ReadJob(job)
	//	}
	//}()

	go mgr.feedProcessor.DispatchWorkMain()
	go mgr.feedProcessor.WriteWorkerMain()
	time.Sleep(time.Second * 5)
}
