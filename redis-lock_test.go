/* 绝密 TOP SECRET, COPYRIGHT © AFMOBI GROUP */
package redis_multi_lock

import (
	"testing"
	"github.com/go-redis/redis"
	"fmt"
)

func Test_Lockstruct_Lock(t *testing.T) {
	RedisInit(&redis.Options{Addr:"127.0.0.1:6379"})
	lck := NewLockstruct("a", "b", "c")
	defer lck.UnLock()
	lck.Lock()
	if lck.Error != nil{
		fmt.Println(lck.Error.Error())
	}
	lck.Lock()
	if lck.Error != nil{
		fmt.Println(lck.Error.Error())
	}
}