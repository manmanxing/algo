package limit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	Rate         int64 //固定的token放入速率(每秒)
	Capacity     int64 //桶的容量
	Tokens       int64 //桶中当前的token数量
	LastTokenSec int64 //桶上次放 token 的时间戳(秒)
	Lock         sync.Mutex
}

//先计算放入的token数量,需要判断token是否溢出
//再计算剩余的token数量，如果 >= 0 返回 true，否则返回 false
func (t *TokenBucket) Allow() bool {
	now := time.Now().Unix()
	addTokens := (now - t.LastTokenSec) * t.Rate
	t.LastTokenSec = now
	t.Tokens = t.Tokens + addTokens
	if t.Tokens > t.Capacity {
		t.Tokens = t.Capacity
	}
	if t.Tokens > 0 {
		t.Tokens--
		return true
	} else {
		//不足以获取这么多的 token
		return false
	}
}

func (t *TokenBucket) Set(rate, capacity int64) {
	if rate < 0 || capacity < 0 {
		return
	}
	t.Rate = rate
	t.Capacity = capacity
	t.Tokens = 0
	t.LastTokenSec = time.Now().Unix()
}
