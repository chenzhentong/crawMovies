package models

import "github.com/astaxie/goredis"

var(
	client goredis.Client

)

const(
	URL_QUEUE="url_queue"
	URL_VISIT_SET ="url_visit_set"
)
//连接redis，默认127.0.0.1：6379
func ConnectRedis(addr string){
	client.Addr=addr
}
//加入队列
func PutinQueue(url string){

	client.Lpush(URL_QUEUE,[]byte(url))
}
//弹出对垒
func PopfromQueue() string{
	res,err:=client.Rpop(URL_QUEUE)
	if err!=nil{
		panic(err)
	}else{
		return string(res)
	}
}
//获取队列长度
func GetQueueLength() int{
	len,err:=client.Llen(URL_QUEUE)
	if err!=nil{
		panic(err)
	}else{
		return len
	}
}
func AddToSet(url string){
	client.Sadd(URL_VISIT_SET,[]byte(url))
}
func IsVisit(url string)bool{
	bIsVisit,err:=client.Sismember(URL_VISIT_SET,[]byte(url))
	if err!=nil{
		panic(err)
	}else{
		return bIsVisit
	}

}