package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"sort"
)

const amzUrl string = "http://webservices.amazon.com/onca/xml"
const amzDomain string = "ecs.amazonaws.com"
const amzPath string = "/onca/xml"
const API_VERSION string = "2011-08-01"
const service string = ""

type ApiRequest struct {
	keyId        string
	keySecret    string
	associateTag string
}

type ItemLookupResponse struct {
}

func (self ApiRequest) ItemLookup() (ItemLookupResponse, error) {
	var lookupResp ItemLookupResponse
	urlParams := make(map[string]string)
	urlParams["Version"] = API_VERSION

	resp, err := http.Get(amzUrl)
	if err != nil {
		return lookupResp, err
	}
	if resp.StatusCode != http.StatusOK {
		/*		contents, _ := ioutil.ReadAll(resp.Body)
				log.Println(string(contents))*/
		return lookupResp, errors.New("not cool")
	}

	return lookupResp, err
}

func generateSignature(queryParams map[string]string, keySecret string) string {
	keys := make([]string, 0, len(queryParams))
	for k, _ := range queryParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Generate the sorted field1=val1&field2=val2 string
	lastIndex := len(keys) - 1
	var strBuff bytes.Buffer
	for i, k := range keys {
		strBuff.WriteString(k)
		strBuff.WriteString("=")
		strBuff.WriteString(url.QueryEscape(queryParams[k]))

		if i != lastIndex {
			strBuff.WriteString("&")
		}
	}

	// Content to sign
	data := "GET\n" + amzDomain + "\n" + amzPath + "\n" + strBuff.String()
	hash := hmac.New(sha256.New, []byte(keySecret))
	hash.Write([]byte(data))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return signature
}
