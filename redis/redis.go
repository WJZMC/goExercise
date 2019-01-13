package redis

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)
var (
	pool *redis.Pool
)
func init()  {
	pool = &redis.Pool{
		MaxIdle:8,//最大空闲连接数，会自动增长
		MaxActive:0,//最大活跃连接数，0：不限制
		IdleTimeout:100,//最大空闲时间，超过空闲时间会被pool回收连接
		Dial: func() (redis.Conn, error) {//初始化连接
			return  redis.Dial("tcp",":6379")
		},
	}
}
func RedisTest()  {
	//redisOldDemo()

	defer pool.Close()

	conn:=pool.Get()

	defer conn.Close()
	_,err:=conn.Do("set", "test", "hello")
	if err!=nil{
		fmt.Println("Do err:",err)
		return
	}

	res,err :=conn.Do("get", "test")
	if err!=nil{
		fmt.Println("Do err:",err)
		return
	}
	resValue,err:=redis.String(res,err)
	if err!=nil{
		fmt.Println("Do err:",err)
		return
	}
	fmt.Println(resValue)

}

//func newRedisClient()* redis.Client  {
//	redisClient:=redis.Client{}
//}

//传统用法
func redisOldDemo()  {
	conn,err:=redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil{
		fmt.Println("conn err:",err)
		return
	}
	defer  conn.Close()

	_,err=conn.Do("set", "test", "hello")
	if err!=nil{
		fmt.Println("Do err:",err)
		return
	}

	res,err :=conn.Do("get", "test")
	if err!=nil{
		fmt.Println("Do err:",err)
		return
	}
	resValue,err:=redis.String(res,err)
	if err!=nil{
		fmt.Println("Do err:",err)
		return
	}
	fmt.Println(resValue)


}