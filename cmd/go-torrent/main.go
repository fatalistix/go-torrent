package main

import (
	"fmt"
	"github.com/fatalistix/go-torrent/internal/file"
)

func main() {
	tf, hash, err := file.Read("/home/deck/Documents/Cult_of_the_Lamb_....torrent")
	if err != nil {
		panic(err)
	}

	fmt.Println(tf)
	fmt.Println(hash)
}
