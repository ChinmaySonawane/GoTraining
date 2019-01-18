package main

func f1(rec chan<- string, msg string) {
	rec <- msg
	//rec <- msg
}

func f2(rec chan<- string, send <-chan string) {
	msg := <-send
	rec <- msg
	rec <- msg
}

func main() {

	chan1 := make(chan string, 1)
	chan2 := make(chan string, 2)
	f1(chan1, "hello")
	f2(chan2, chan1)
	<-chan2
	<-chan2

}
