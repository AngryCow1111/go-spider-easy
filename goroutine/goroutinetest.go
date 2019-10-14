package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// 获取goroutineID
func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
func say(word string) {
	fmt.Println(word)
	time.Sleep(100 * time.Millisecond)
	goid := Goid()
	fmt.Println(goid)

}
func main() {
	//go say("hello")
	//say("world")
	now := time.Now()
	formatNow := now.Format("2006-01-02 15:04:00")
	fmt.Println(formatNow)
}
