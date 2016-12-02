package main

import (
	"fmt"
	"os"
	//"github.com/alexmk92/stringutil"
	//"strings"
	"strings"
)

func main() {
	if(DEBUG) {
		fmt.Println("Initialising the client reader");
	}

	args := os.Args[1:]

	if(len(args) == 0) {
		fmt.Println("It's required that a character name is provided in the argument list")
	} else {

	}
	r := AuctionReader{}

	if(r.open(strings.TrimSpace(EQ_PATH), strings.TrimSpace(strings.Title(strings.ToLower(EQ_CHARACTER))), FILE_EXTENSION)) {
		reads := make(chan string)
		go r.read(reads)

		read := <-reads
		fmt.Println(read)
	}
}

