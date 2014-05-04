package client

import "encoding/xml"

// A struct that can be represented as a map[string]string
type Mappable interface {
	Map() map[string]string
}

// ItemLookup
type ItemLookupParams struct {
	IdType string
	ItemId string
}

func (self ItemLookupParams) Map() map[string]string {
	m := make(map[string]string)
	if self.IdType != "" {
		m["IdType"] = self.IdType
	}
	m["ItemId"] = self.ItemId

	return m
}

type ItemLookupRequest struct {
	XMLName       xml.Name `xml:"ItemLookupRequest"`
	IdType        string
	ResponseGroup string
	SearchIndex   string
	VariationPage string
}

type ItemLookupResponse struct {
	XMLName xml.Name `xml:"ItemLookupResponse"`
	Items   []Item   `xml:"Items>Item"`
	//TODO: Request should be part of Items
	Request Request `xml:"Items>Request"`
}

type ItemLookupErrorResponse struct {
	Error     Error
	RequestId string
}

// Shared types
type Request struct {
	XMLName           xml.Name          `xml:"Request"`
	IsValid           bool              `xml:"IsValid"`
	ItemLookupRequest ItemLookupRequest `xml:"ItemLookupRequest"`
	Errors            []Error           `xml:"Errors>Error"`
}

type Error struct {
	Code, Message string
}

type Item struct {
	ItemAttributes ItemAttributes
}

type ItemAttributes struct {
	/*	<Author>J.K. Rowling</Author>
		<Creator Role="Illustrator">Mary GrandPré</Creator>
		<Manufacturer>Scholastic</Manufacturer>
		<ProductGroup>Book</ProductGroup>
		<Title>Harry Potter and the Sorcerer's Stone (Book 1)</Title>*/
	Author string
	Title  string
}
