package main



func main(){
	go Update()
	go Server()
	select{}
}