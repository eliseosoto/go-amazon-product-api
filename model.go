package client

import "encoding/xml"

// ItemLookup

type ItemLookupResponse struct {
	XMLName xml.Name `xml:"ItemLookupResponse"`
	/*	Items   []Item   `xml:"Items>Item"`*/
	Request Request `xml:"Items>Request"`
}

type Error struct {
	Code, Message string
}

type ItemLookupErrorResponse struct {
	Error     Error
	RequestId string
}

type Request struct {
	XMLName           xml.Name          `xml:"Request"`
	IsValid           bool              `xml:"IsValid"`
	ItemLookupRequest ItemLookupRequest `xml:"ItemLookupRequest"`
	Errors            []Error           `xml:"Errors>Error"`
}

type ItemLookupRequest struct {
	XMLName       xml.Name `xml:"ItemLookupRequest"`
	IdType        string
	ResponseGroup string
	SearchIndex   string
	VariationPage string
}
