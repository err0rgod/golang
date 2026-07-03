package main

import (
	// "fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// add ip support also later

func ScanPort(address string) bool {
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err == nil {
		conn.Close()
		return true
	} else {
		return false
	}
}

func scanner(start int, end int, ip string) []int{
	var wg sync.WaitGroup
	var OpenPorts []int
	var mu sync.Mutex
	for ;start < end; start++ {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()
			address := ip + ":" + strconv.Itoa(port)
			if ScanPort(address) {
				mu.Lock()
				OpenPorts = append(OpenPorts, port)
				mu.Unlock()
			}
		}(start)
	}
	wg.Wait()
	return OpenPorts
}
func main1(ip string, start string , end string) []int {
	// ip := os.Args[1]
	// start := os.Args[2]
	// end := os.Args[3]
	endports, err1 := strconv.Atoi(end)
	startports, err2 := strconv.Atoi(start)
	if err2 !=nil  || err1 != nil || endports > 65001 || startports < 1 {
		// fmt.Println("Enter a valid range.")
		return nil
	}
	OpenPorts := scanner(startports,endports, ip)
	return OpenPorts
}
