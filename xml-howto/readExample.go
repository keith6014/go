package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type catalog struct {
	XMLName xml.Name `xml:"catalog"`
	Books   []Book   `xml:"book"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Id      string   `xml:"id,attr"`
	Author  string   `xml:"author"`
	Title   string   `xml:"title"`
	Genre   string   `xml:"genre"`
	Price   float32  `xml:"price"`
}

func main() {
	var q catalog
	xmlFile, _ := ioutil.ReadFile("src/xml-howto/data/books.xml")
	err := xml.Unmarshal(xmlFile, &q)
	if err != nil {
		panic(err)
	}

	fmt.Println("%v", q)
}
