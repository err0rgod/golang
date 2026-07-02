package main

import (
	"fmt"
	"net"
	"sync"

	// "os"
	"strconv"
	"time"
)

// add ip support also later

func ScanPort(port int, wg *sync.WaitGroup) bool {
	defer wg.Done()
	address := "192.168.1.1:"+strconv.Itoa(port)
	conn,err := net.DialTimeout("tcp", address,1*time.Second)
	if err==nil {
		conn.Close()
		return true
	} else {
		return false
	}
}


func main()  {
	var wg sync.WaitGroup
	start := time.Now()
	for i := range(10000) {
		wg.Add(1)

		go ScanPort(i,&wg)
	}
	wg.Wait()
	fmt.Print("The time taken is :" ,time.Since(start).Seconds())
}
