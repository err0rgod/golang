package main

import (
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
	start := req.START
	end := req.END
	// error handling
	if ip == "" || start == "" || end == "" {
		c.JSON(400,gin.H{"Error":"Required ip, start , end"})
		return
	}
	
	OpenPorts := main1(ip, start, end)
	if OpenPorts == nil {
		// c.JSON(400,gin.H{"Error":"Invalid Range."})
		
		// return
	}
	resulsts := ScanResult{
		IP : ip,
		Ports: OpenPorts,
	}
	record := ScanRecord {
		IP : ip,
		START: start,
		END:  end,
		OPEN_PORTS: OpenPorts,
	}
	AddRecord(record)
	c.JSON(200,resulsts)
}

type ScanRequest struct {
	IP string `json:"ip"`
	START string `json:"start"`
	END string		`json:"end"`
}

func History(c *gin.Context) {
	records := GetRecord()
	c.JSON(200,records)
}

func main(){
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status":"OK"})
	})
	r.GET("/history",History)
	r.POST("/scan", portscanner)
	r.Run(":8080")
}