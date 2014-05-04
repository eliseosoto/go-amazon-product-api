package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io/ioutil"
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
	KeyId        string
	KeySecret    string
	AssociateTag string
}

func (self *ApiRequest) ItemLookup(p ItemLookupParams) (ItemLookupResponse, error) {
	lookupResp := ItemLookupResponse{}
	urlParams := self.generateQueryParams("ItemLookup", p)

	signedUrl, _ := generateSignature(urlParams, self.KeySecret)

	/*	log.Println(signedUrl)*/

	resp, err := http.Get(signedUrl)
	if err != nil {
		return lookupResp, err
	}
	contents, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Println(string(contents))
		return lookupResp, errors.New("not cool")
	}
	log.Println(string(contents))

	err = xml.Unmarshal(contents, &lookupResp)
	if err != nil {
		return lookupResp, err
	}

	return lookupResp, err
}

func (self *ApiRequest) generateQueryParams(operation string,
	p Mappable) map[string]string {
	params := p.Map()

	params["Service"] = "AWSECommerceService"
	params["Version"] = ApiVersion
	params["AssociateTag"] = self.AssociateTag
	params["Operation"] = operation
	params["Timestamp"] = time.Now().Round(time.Second).Format(time.RFC3339)
	params["AWSAccessKeyId"] = self.KeyId

	return params
}

func generateSignature(queryParams map[string]string, keySecret string) (
	string, string) {
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

	amzUrl := protocol + amzDomain + amzPath + "?" + strBuff.String() +
		"&Signature=" + url.QueryEscape(signature)

	return amzUrl, signature
}
