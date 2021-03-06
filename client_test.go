package client

import "fmt"

var (
	mockKeyId        = "fakeId"
	mockKeySecret    = "fakeSecret"
	mockAssociateTag = "PutYourAssociateTagHere"
)

func ExampleGenerateSignature() {
	// Use a signature generated by:
	// http://associates-amazon.s3.amazonaws.com/signed-requests/helper/index.html
	queryParams := make(map[string]string)
	queryParams["Service"] = "AWSECommerceService"
	queryParams["Version"] = "2011-08-01"
	queryParams["AssociateTag"] = mockAssociateTag
	queryParams["Operation"] = "ItemSearch"
	queryParams["SearchIndex"] = "Books"
	queryParams["Keywords"] = "iñarritu"
	queryParams["Timestamp"] = "2014-04-21T07:03:01.000Z"
	queryParams["AWSAccessKeyId"] = mockKeyId

	_, signature := generateSignature(queryParams, mockKeySecret)
	fmt.Println(signature)
	// Output: dMLseh1BdUcANv8Mjapf/LdnGEBZH/9DAeeTEtjt5Wc=
}
