package client

import (
	"log"
	"testing"

	"github.com/eliseosoto/go-amazon-product-api"
)

// Your own keys go here.
const (
	keyId        = ""
	keySecret    = ""
	associateTag = ""
)

func TestApiRequest(t *testing.T) {
	apiReq := client.ApiRequest{keyId, keySecret, associateTag}

	apiRes, err := apiReq.ItemLookup()

	log.Printf("%+v", apiRes)

	if err != nil {
		t.Error(err)
	}

	if &apiRes == nil {
		t.Error("Nooo")
	}

	if len(apiRes.Request) < 0 {
		t.Error("expected at least 1 Item")
	}
}
