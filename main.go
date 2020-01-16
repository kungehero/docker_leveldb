package main

import "github.com/syndtr/goleveldb/leveldb"

import "fmt"

const (
	Add    = iota + 1
	Delete = 4
	Update
	Get = iota
	All
)

func main() {
	fmt.Println(Delete, Update, Get, All)
	leveldb.OpenFile("", nil)
}
