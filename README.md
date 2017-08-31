````
  RedisInit(&redis.Options{Addr:"127.0.0.1:6379"})
    lck := NewLockstruct("a", "b", "c")
    defer lck.UnLock()
    lck.Lock()
    if lck.Error != nil{
      fmt.Println(lck.Error.Error())//nil
    }
    lck.Lock()
    if lck.Error != nil{
      fmt.Println(lck.Error.Error())//could Not obtain lock
    }
````
