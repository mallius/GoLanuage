package main

import (
	"fmt"
)

type book struct {
	name   string
	author string
	isdn   string
	date   string
}

func main() {
	var ret bool
	var Books []book
	Books = []book{
		{"Go01", "Alan", "ISDN0001", "2011"},
		{"Go02", "Bob", "ISDN0002", "2012"},
		{"Go03", "Cox", "ISDN0003", "2013"},
	}

	ret = Books[0].findName("Alan")
	fmt.Println(ret)

}

func (recv book) findAll(Books []book) (ret bool) {
}

func (recv book) findName(name string) (ret bool) {
	if name == recv.name {
		return true
	} else {
		return false
	}
}

func (recv book) findAuthor(author string) (ret bool) {
	if author == recv.author {
		return true
	} else {
		return false
	}
}

func (recv book) findIsdn(isdn string) (ret bool) {
	if isdn == recv.isdn {
		return true
	} else {
		return false
	}
}
