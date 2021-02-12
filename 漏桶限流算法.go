package main

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	Rate       float64 //固定的出水速率(每秒)
	Capacity   float64 //桶的容量
	Water      float64 //桶中当前的水量
	lastLeakMs int64   //桶上次漏水时间戳(毫秒)
	Lock       sync.Mutex
}

//n表示进水量
//先计算漏水量
//如果桶中当前的水量+n 没有超过桶的容量，返回 true，否则返回 false
func (l *LeakyBucket) Allow(n float64) bool {
	if n < 0 {
		return false
	}
	now := time.Now().UnixNano() / 1e6
	eclipseWater := float64(now-l.lastLeakMs) * l.Rate / 1000
	l.Water = l.Water - eclipseWater
	//如果桶干了
	if l.Water < 0 {
		l.Water = 0
	}
	l.lastLeakMs = now
	//判断水量
	if l.Water + n <= l.Capacity {
		l.Water = l.Water + n
		return true
	}else {
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
	l.lastLeakMs = time.Now().UnixNano() / 1e6
}