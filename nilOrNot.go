package main

import (
	"log"
	"reflect"
)

func nilOrNot(obj interface{}) bool {
	return obj == nil
}

type ArrStruct struct {
	arr []string
}
type OneStruct struct{}

func mytest() {
	x := &OneStruct{}
	log.Println(x == nil, reflect.TypeOf(x))
	log.Println(nilOrNot(x))

	y := &ArrStruct{
		arr: []string{},
	}
	y.arr = append(y.arr, "1")
	y.arr = append(y.arr, "2")
	y.arr = append(y.arr, "3")
	log.Println(y)
	j := append(y.arr, "4")
	k := append(y.arr, "5")
	l := append(y.arr, "6")
	log.Println(j, k, l)
}
