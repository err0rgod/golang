package main

import (
	"fmt"
	"net"
	"sync"

	// "os"
	// "sync"
	"strconv"
	"time"
)

// add ip support also later

func ScanPort(port int) bool {
	address := "192.168.1.1:"+strconv.Itoa(port)
	conn,err := net.DialTimeout("tcp", address,1*time.Second)
	if err==nil {
		conn.Close()
		return true
	} else {
		return false
	}
}


func scanner()  {
	var wg sync.WaitGroup
	for i := range(1000) {
		wg.Add(1)
		var isOpen bool
		go func() {
			defer wg.Done()
			isOpen = ScanPort(i)

			if isOpen {
				fmt.Println("The port ",i," is Open------------.")
			} else {
				// fmt.Println("The port ",i," is Close.")	
			}
		}()
	}
	wg.Wait()
}
func main() {
	scanner()
}