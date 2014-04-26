package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"net/url"
	"sort"
	"time"
)

const amzDomain string = "ecs.amazonaws.com"
const amzPath string = "/onca/xml"
const ApiVersion string = "2011-08-01"
const service string = ""

type ApiRequest struct {
	keyId        string
	keySecret    string
	associateTag string
}

type ItemLookupResponse struct {
}

func (self *ApiRequest) ItemLookup() (ItemLookupResponse, error) {
	var lookupResp ItemLookupResponse
	urlParams := self.generateQueryParams("ItemLookup")

	signedUrl, _ := generateSignature(urlParams, self.keySecret)

	log.Println(signedUrl)

	resp, err := http.Get(signedUrl)
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

func (self *ApiRequest) generateQueryParams(operation string) map[string]string {
	params := make(map[string]string)

	params["Service"] = "AWSECommerceService"
	params["Version"] = ApiVersion
	params["AssociateTag"] = self.associateTag
	params["Operation"] = operation
	params["SearchIndex"] = "Books"
	params["Keywords"] = "iñarritu"
	params["Timestamp"] = time.Now().Format(time.RFC3339)
	params["AWSAccessKeyId"] = self.keyId

	return params
}

func generateSignature(queryParams map[string]string, keySecret string) (string, string) {
	const protocol = "http://"

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

	amzUrl := protocol + amzDomain + amzPath + "?" + strBuff.String() + "&Signature=" + signature

	return amzUrl, signature
}
