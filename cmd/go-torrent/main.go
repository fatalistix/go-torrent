package main

import (
	"fmt"
	"github.com/fatalistix/go-torrent/internal/file"
)

func main() {
	tf, err := file.Read("/Users/vyacheslav/Documents/cult-of-the-lamb-2022.torrent")
	if err != nil {
		panic(err)
	}

	fmt.Println(tf)
}
