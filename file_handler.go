package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

func Update() {
	// loadfile name and initial size to 0
	Filename := "data.txt"
	var last_offset int64 = 0
	// loading initial fine info
	fileinfo, err := os.Stat(Filename)
	if os.IsNotExist(err){
		file,err := os.Create(Filename)
		if err != nil {
			log.Fatal("Error while creating the file.")
		}
		file.Close()
		} else if err != nil {
			log.Fatal(err)
		}
		// setting the size to the actual size of the file
	fileinfo, err = os.Stat(Filename)
	last_offset = fileinfo.Size()



	// creating a fsnotify watcher continously watches the changes
	watcher,err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func(){
		for {
			select {
				// if a change is maded
			case event,ok := <- watcher.Events:
				if !ok {
					return
				}
				// a file change has been made
				if event.Op & fsnotify.Write == fsnotify.Write {
					fileInfo,err := os.Stat(Filename)
					if err != nil {
						log.Printf("Error %v", err)
						continue	
					}
					// if the size changes fom the previous ones
					if fileInfo.Size() < last_offset {
						last_offset = 0
					}else if fileInfo.Size() > last_offset {
						file,err := os.Open(Filename)
						if err != nil {
							log.Printf("Error %v", err)
							continue
						}
						// setting the delimiter to the position from where it needs to read
						_,err = file.Seek(last_offset,0)  
						if err != nil {
							log.Printf("Error %v", err)
							file.Close()
							continue
						}
						// making a new slice to hold the new changes and then print them in string
						new_data := make([]byte, fileInfo.Size() - last_offset)
						n,err := file.Read(new_data)
						if err != nil {
							log.Printf("Error %v", err)
							file.Close()
							continue
						}
						last_offset += int64(n)
						broadcast <- Data {
							Type : websocket.TextMessage,
							Data :new_data,
						}
						// fmt.Print(string(new_data))
						file.Close()
					}
				}
			// if some error occured
			case err,ok := <- watcher.Errors:
				if !ok {return }
				log.Printf("ERROR %v", err)
			}
		}
	}()
	
	err = watcher.Add(Filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	<-done
	
}