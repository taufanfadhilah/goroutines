package main

import "fmt"

func channels() {
	c := make(chan string)

	go func(n string) {
		c <- n
	}("Taufan Fadhilah")

	var chanFunc = func(n string) {
		c <- n
	}

	go chanFunc("Hello")

	b := <-c
	o := <-c
	fmt.Println(b)
	fmt.Println(o)

	chanKeDua := make(chan string)
	go ChannelCetakNama(chanKeDua, "Hi")

	var ll = <-chanKeDua
	fmt.Println(ll)

	// Buffered Channel
	pp := make(chan int)
	go func() {
		for p := range pp {
			p = <-pp
			fmt.Println("Terima Data ========", p)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Kirim data")
		pp <- i
	}

}

func ChannelCetakNama(c chan string, v string) {
	c <- v
}
