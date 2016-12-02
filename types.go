package main

import "time"

// Anything in this file will be serialized to JSON to be persisted to the Mongo store

/*
 |------------------------------------------------------------------
 | Auction
 |------------------------------------------------------------------
 |
 | The three structs below represent what an auction is comprised of.
 | Each auction has a seller, array of items for sale and each
 | item has an array of its stats.  Its possible that we could
 | substiute the Stat struct in future for a map of string:int
 | but for completeness its declared as its own struct.
 |
 */

type Auction struct {
	seller string
	items []Item
	auctioned_at time.Time
}

type Item struct {
	name string
	price float32
	statistics []Statistic
}

type Statistic struct {
	name string
	value int32 // doesn't need to be uint as we can have negative stats (i.e. on an AoN)
}