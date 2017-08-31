/* 绝密 TOP SECRET, COPYRIGHT © AFMOBI GROUP */
package redis_multi_lock

import (
	"github.com/go-redis/redis"
	"github.com/bsm/redis-lock"
	"errors"
	"time"
)
const Error_Key_Locked = "could Not obtain lock"

var redisClient *redis.Client

func RedisInit(opt *redis.Options){
	client := redis.NewClient(opt)
	redisClient = client
}

func getRedisClient() *redis.Client{
	if redisClient == nil || redisClient.Ping().Err() != nil {
		RedisInit(redisClient.Options())
	}
	return redisClient
}

type Lockstruct struct {
	keys []string
	locks map[int]*lock.Lock
	Error error
}

func NewLockstruct(keys ...string)*Lockstruct{
	return NewLockstructArray(keys)
}

func NewLockstructArray(keys []string)*Lockstruct{
	lck := new(Lockstruct)
	lck.keys = keys
	lck.locks = make(map[int]*lock.Lock)
	return lck
}

func (this *Lockstruct)Lock(){
	rclient := getRedisClient()
	for i, key := range this.keys{
		locka, err := lock.ObtainLock(rclient, key, &lock.LockOptions{
			LockTimeout:time.Minute,
		})
		if err != nil{
			this.Error = err
			return
		}
		if locka == nil{
			this.Error = errors.New(Error_Key_Locked)
			return
		}
		this.locks[i] = locka
	}
	this.Error = nil
}

func (this *Lockstruct)UnLock(){
	for i, lck := range this.locks{
		lck.Unlock()
		delete(this.locks, i)
	}
}