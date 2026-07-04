package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type ScanResult struct {
	IP string `json:"ip"`
	Ports []int		`json:"ports"`
}

func portscanner(c *gin.Context) {
	// ip := c.Query("ip")
	// start := c.Query("start")
	// end := c.Query("end")
	var req ScanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400,gin.H{"Error":"Invalid request Body"})
		return
	}
	ip := req.IP
	start := req.start
	end := req.end
	// error handling
	if ip == "" || start == "" || end == "" {
		c.JSON(400,gin.H{"Error":"Required ip, start , end"})
		return
	}
	
	OpenPorts := main1(ip, start, end)
	if OpenPorts == nil {
		c.JSON(400,gin.H{"Error":"Invalid Range."})
		
		return
	}
	resulsts := ScanResult{
		IP : ip,
		Ports: OpenPorts,
	}
	c.JSON(200,resulsts)
}

type ScanRequest struct {
	IP string `json:"ip"`
	start string `json:"start"`
	end string		`json:"end"`
}

func main(){
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status":"OK"})
	})
	r.POST("/scan", portscanner)
	r.Run(":8080")
}