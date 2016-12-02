package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type AuctionReader struct {
	file *os.File
	character string
}

func (r *AuctionReader) succeeded(e error) (bool, error) {
	if e != nil {
		return false, e
	}
	return true, e
}

func (r *AuctionReader) open(path, characterName, extension string) bool {
	r.character = characterName
	file, err := os.Open(path + characterName + extension)
	succeeded, e := r.succeeded(err)
	if succeeded && e == nil {
		r.file = file
		return true
	} else {
		r.file = nil
		panic(e)
	}
	return false
}

func (r *AuctionReader) read(channel chan <- string) {
	fmt.Println("file: ", r.file)

	parser := AuctionParser{}

	totalLines := 0
	auctionLines := 0

	if r.file == nil {
		panic("File has not been set on this instance of reader")
	} else {
		scanner := bufio.NewScanner(r.file)
		//items := make(chan)
		i := 0

		lastLine := ""

		for scanner.Scan() {
			lineString := strings.Replace(scanner.Text(), "You auction", "Kongsong auctions", -1)
			//fmt.Println(strconv.Itoa(i) + " - " + lineString)
			totalLines++
			if(parser.isAuctionLine(lineString)) {
				auctionLines++
				lastLine = lineString
			}
			i++
		}

		if(DEBUG) {
			parser.parseLine(lastLine)
		}
	}
	channel <- r.character + " finished logging their current data"
	fmt.Println("Total lines: " + strconv.Itoa(totalLines) + ", Auction lines: " + strconv.Itoa(auctionLines))
}