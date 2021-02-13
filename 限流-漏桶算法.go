package main

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	Rate       float64 //固定的出水速率(每秒)
	Capacity   float64 //桶的容量
	Water      float64 //桶中当前的水量
	LastLeakMs int64   //桶上次漏水时间戳(毫秒)
	Lock       sync.Mutex
}

//先计算漏水量,需判断桶是否干了
//如果桶中当前的水量+n 没有超过桶的容量，返回 true，否则返回 false
func (l *LeakyBucket) Allow() bool {
	now := time.Now().UnixNano() / 1e6
	eclipseWater := float64(now-l.LastLeakMs) * l.Rate / 1000
	l.Water = l.Water - eclipseWater
	//如果桶干了
	if l.Water < 0 {
		l.Water = 0
	}
	l.LastLeakMs = now
	//判断水量
	if l.Water + 1 <= l.Capacity {
		l.Water ++
		return true
	} else {
		return false
	}
}

//初始化
func (l *LeakyBucket) Set(rate, Capacity float64) {
	if rate < 0 || Capacity < 0 {
		return
	}
	l.Rate = rate
	l.Capacity = Capacity
	l.Water = 0
	l.LastLeakMs = time.Now().UnixNano() / 1e6
}