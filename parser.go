package main

import (
	"fmt"
	"strings"
	"encoding/json"
	"regexp"
	"github.com/alexmk92/stringutil"
	"time"
)

type AuctionParser struct {}

func (p *AuctionParser) parseLine(line string) (bool, string) {
	p.extractSaleInformation(line)
	return true, ""
}

func (p *AuctionParser) isAuctionLine(line string) bool {
	isValid, _ := regexp.MatchString("(\\[)([a-zA-Z0-9: ]+)(] ([A-Za-z]+) ((auction)|(auctions)))", line)
	return isValid
}

func (p *AuctionParser) extractSaleInformation(line string) {
	fmt.Println("Extracting information from: ", line)
	line = "[Wed Nov 23 23:26:10 2016] Kongsong auctions, 'WTS Jagged Blade of Mourning || Wavecrasher || Holgresh Elder Beads, WtB Holgresh Elder Beads, Amulet of Necropotence | Lodizal Shield && Turtle Bone Bracer'"

	auction := Auction {}

	// Split the auction string so we have date on the left and auctions on the right
	parts := strings.Split(line, "]")

	layout := "Mon Jan 2 15:04:05 2006"
	date := strings.TrimSpace(strings.Replace(parts[0], "[", "", -1))
	t, err := time.Parse(layout, date)

	if(err != nil) {
		fmt.Println("ERROR WHEN PARSING DATE: ", err)
	}

	// Explode this array so we are left with the seller on the left and items on the right
	auctionParts := strings.Split(parts[1], "auctions,")

	seller := auctionParts[0]

	// Sale data is always encapsulated in single quotes, taking a substring removes these
	items := strings.TrimSpace(auctionParts[1])[1:len(auctionParts[1])-1]
	items = regexp.MustCompile(`(?i)wts`).ReplaceAllLiteralString(items, "")

	// Discard the WTB portion of the string
	wtbIndex := stringutil.CaseInsensitiveIndexOf(items, "WTB")
	if(wtbIndex > -1) {
		items = items[0:wtbIndex]
	} else {
		fmt.Println("SALES ONLY")
	}

	itemList := strings.FieldsFunc(items, func(r rune) bool {
		switch r {
		case '|', ',':
			return true;
		}
		return false
	})

	fmt.Println("Seller is: ", seller)
	fmt.Println("Sale data is: ", itemList)

	i := Item {
		name: "ALEX",
	}
	p.fetchItemDataFromWiki(&i)

	auction.auctioned_at = t

	fmt.Println(json.Marshal(auction))
	//return nil
}

func (p *AuctionParser) fetchItemDataFromWiki(item *Item) {
	item.name = "GARY"
}
