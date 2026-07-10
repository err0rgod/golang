package main



func main(){
	go Update()
	go Server()
	go WebsocketHandler()
	select{}
}