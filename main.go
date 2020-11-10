package main

import (
	"log"
	"net/http"
	"strings"
	"sync"
	// "time"
)

var lock sync.Mutex
var existArr = sync.Map{}

func tellIn(oneStr string) bool {
	ret := false
	if _, ok := existArr.Load(oneStr); ok {
		ret = true
	}
	return ret
}
func ClearHandler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	existArr = sync.Map{}
	w.Write([]byte("清理成功"))
	lock.Unlock()
}
func addArr(newArr []string) {
	for _, one := range newArr {
		existArr.Store(one, one)
	}
}
func ArrHandler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	newArr := strings.Split(r.FormValue("newarr"), ",")
	ret := []string{}

	for _, one := range newArr {
		if tellIn(one) {
			ret = append(ret, "true")
		} else {
			ret = append(ret, "false")
		}
	} //每一次调用判断完之后再设置到map里，这样就不会导致第一次传入相同元素时某个值为真的情况
	addArr(newArr)
	// log.Printf("时间戳（毫秒）：%v;newArr:%s",time.Now().UnixNano(),newArr)
	w.Write([]byte(strings.Join(ret, ",")))
	lock.Unlock()
}
func StartAPI() {
	http.HandleFunc("/clzclear", ClearHandler)
	http.HandleFunc("/clzarr", ArrHandler)
	err := http.ListenAndServeTLS("localhost:8888", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("服务端报错：", err.Error())
	}

}
func main() {
	StartAPI()
}
