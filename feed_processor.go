/**
*
*@author 吴昊轩
*@create 2020-04-1611:32
 */
package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

type PushJob struct {
	Uid      string
	FeedId   string
	TimeLine int64
}
type FeedProcessor struct {
	redisClient *redis.Client
	inChan      chan *PushJob
	outChan     chan *PushJob
	closeChan   chan struct{}
}

func (f *FeedProcessor) ReadJob(job *PushJob) {
	f.inChan <- job
	fmt.Printf("读取微博，内容：%s\n", job.FeedId)
}

func (f *FeedProcessor) DispatchWorkMain() {
	var (
		err error
	)
	for {
		select {
		case job := <-f.inChan:
			//发送到用户发送feed集
			if err = f.pushToSendFeed(job); err != nil {

				fmt.Println(err)
				f.Close()
			} else {
				//发送成功,接下来生成，推送送粉丝接收feed的 job
				go f.createPushJob(job)
			}
		case <-f.closeChan:
			break
		}
	}
}

func (f *FeedProcessor) WriteWorkerMain() {
	var (
		err error
	)
	for {
		//fmt.Println("outchan len ", len(f.outChan))
		select {
		case job := <-f.outChan:
			if err = f.pushToRecvFeed(job); err != nil {
				fmt.Println(err)
				f.Close()
			}

		case <-f.closeChan:
			break
		}
	}
}
func (f *FeedProcessor) pushToSendFeed(job *PushJob) (err error) {
	z := &redis.Z{Score: float64(job.TimeLine), Member: job.FeedId}

	if err = f.redisClient.ZAdd("send:"+job.Uid, z).Err(); err != nil {
		return
	}
	fmt.Printf("uid: %s ,发微博: %s,推送发送feed集 \n", job.Uid, job.FeedId)
	return
}

func (f *FeedProcessor) pushToRecvFeed(job *PushJob) (err error) {
	z := &redis.Z{Score: float64(job.TimeLine), Member: job.FeedId}
	if err = f.redisClient.ZAdd("recv:"+job.Uid, z).Err(); err != nil {
		return
	}
	fmt.Printf("给粉丝：%s ,推送动态:%s", job.Uid, job.FeedId)
	return
}

func (f *FeedProcessor) createPushJob(job *PushJob) {
	var (
		err     error
		fansIds []string
	)
	processor := UserProcessor{job.Uid}
	if fansIds, err = processor.GetFansIds(); err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Printf("获取用户 %s  粉丝集合uid:%+v\n", job.Uid, fansIds)
	for _, fansId := range fansIds {
		pushJob := &PushJob{Uid: fansId, FeedId: job.FeedId, TimeLine: job.TimeLine}
		f.outChan <- pushJob
		fmt.Printf("push job to outchan\n")
	}

}
func (f *FeedProcessor) Close() {
	close(f.closeChan)
}
