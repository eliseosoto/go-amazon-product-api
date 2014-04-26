package client

import "testing"

// Your own keys go here.
const (
	keyId        = ""
	keySecret    = ""
	associateTag = ""
)

func TestApiRequest(t *testing.T) {
	apiReq := ApiRequest{keyId, keySecret, associateTag}

	apiRes, err := apiReq.ItemLookup()
	if err != nil {
		t.Error(err)
	}

	if &apiRes == nil {
		t.Error("Nooo")
	}
}
