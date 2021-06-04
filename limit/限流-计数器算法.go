package limit

import (
	"sync"
	"time"
)

type Counter struct {
	Rate  int           //计数周期内最多允许的请求数
	Begin time.Time     //计数开始时间
	Cycle time.Duration //计数周期
	Count int           //计数周期内累计收到的请求数
	Lock  sync.Mutex
}

//有没有超计数周期
//超过就重置count 和 begin,返回 true
//没有超过：
//判断count有没有超过rate -1，超过就不允许请求，返回 false
//没有超过，count + 1，返回 true
func (c *Counter) Allow() bool {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	now := time.Now()
	if now.Sub(c.Begin) > c.Cycle {
		c.Reset(now)
		return true
	} else {
		if c.Count >= c.Rate-1 {
			return false
		} else {
			c.Count++
			return true
		}
	}
}

//重置
func (c *Counter) Reset(t time.Time) {
	c.Begin = t
	c.Count = 0
}

//初始化
func (c *Counter) Set(rate int, cycle time.Duration) {
	if rate < 0 {
		return
	}
	c.Rate = rate
	c.Begin = time.Now()
	c.Cycle = cycle
	c.Count = 0
}
