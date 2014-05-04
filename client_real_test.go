package client

import (
	"log"
	"testing"
)

// Your own keys go here.
const (
	keyId        = ""
	keySecret    = ""
	associateTag = ""
)

func TestApiRequest(t *testing.T) {
	apiReq := ApiRequest{keyId, keySecret, associateTag}

	params := ItemLookupParams{ItemId: "059035342X"}
	apiRes, err := apiReq.ItemLookup(params)

	log.Printf("%+v", apiRes)

	if err != nil {
		t.Error(err)
	}
}
