package client

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/eliseosoto/go-amazon-product-api"
)

func TestItemLookupErrorResponse(t *testing.T) {
	data := `
    <?xml version="1.0"?>
    <ItemLookupErrorResponse xmlns="http://ecs.amazonaws.com/doc/2011-08-01/">
      <Error>
          <Code>RequestExpired</Code>
          <Message>Request has expired. Timestamp date is 2014-04-26T19:15:21-07:00.</Message>
      </Error>
      <RequestId>2befa87e-0e2e-47c9-9452-8b9132c01a40</RequestId>
    </ItemLookupErrorResponse>
  `

	resp := client.ItemLookupErrorResponse{}

	err := xml.Unmarshal([]byte(data), &resp)

	if err != nil {
		t.Error(err)
	}

	expectedError := client.Error{"RequestExpired",
		"Request has expired. Timestamp date is 2014-04-26T19:15:21-07:00."}

	if !reflect.DeepEqual(expectedError, resp.Error) {
		t.Errorf("expectedError: %v, got %v", expectedError, resp.Error)
	}
}

func TestItemLookupResponse(t *testing.T) {
	data := `
  <?xml version="1.0"?>
  <ItemLookupResponse xmlns="http://webservices.amazon.com/AWSECommerceService/2011-08-01">
      <OperationRequest>
          <HTTPHeaders>
              <Header Name="UserAgent" Value="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.69 Safari/537.36"></Header>
          </HTTPHeaders>
          <RequestId>4aa922ae-5435-4791-b047-0c8292f7a680</RequestId>
          <Arguments>
              <Argument Name="Operation" Value="ItemLookup"></Argument>
              <Argument Name="Service" Value="AWSECommerceService"></Argument>
              <Argument Name="AssociateTag" Value="supervacacom-20"></Argument>
              <Argument Name="Version" Value="2011-08-01"></Argument>
              <Argument Name="Keywords" Value="iñarritu"></Argument>
              <Argument Name="SearchIndex" Value="Books"></Argument>
              <Argument Name="Signature" Value="CHzQvPRqacm0Cha9pbSSApYFukHWvdoEmKQz17960qM="></Argument>
              <Argument Name="AWSAccessKeyId" Value="AKIAIVZ6JESOPVL3CTSA"></Argument>
              <Argument Name="Timestamp" Value="2014-04-26T14:02:13-07:00"></Argument>
          </Arguments>
          <RequestProcessingTime>0.0032950000000000</RequestProcessingTime>
      </OperationRequest>
      <Items>
          <Request>
              <IsValid>False</IsValid>
              <ItemLookupRequest>
                  <IdType>ASIN</IdType>
                  <ResponseGroup>Small</ResponseGroup>
                  <SearchIndex>Books</SearchIndex>
                  <VariationPage>All</VariationPage>
              </ItemLookupRequest>
              <Errors>
                  <Error>
                      <Code>AWS.MissingParameters</Code>
                      <Message>Your request is missing required parameters. Required parameters include ItemId.</Message>
                  </Error>
                  <Error>
                      <Code>AWS.RestrictedParameterValueCombination</Code>
                      <Message>Your request contained a restricted parameter combination.  When IdType equals ASIN, SearchIndex cannot be present.</Message>
                  </Error>
              </Errors>
          </Request>
      </Items>
  </ItemLookupResponse>
`

	var resp client.ItemLookupResponse

	err := xml.Unmarshal([]byte(data), &resp)
	if err != nil {
		t.Error(err)
	}

	expectedIsValid := false

	expectedItemLookupRequest := client.ItemLookupRequest{XMLName: xml.Name{
		Space: "http://webservices.amazon.com/AWSECommerceService/2011-08-01",
		Local: "ItemLookupRequest"}, IdType: "ASIN", ResponseGroup: "Small",
		SearchIndex: "Books", VariationPage: "All"}

	expectedErrors := []client.Error{
		{"AWS.MissingParameters",
			"Your request is missing required parameters. Required parameters include ItemId."},
		{"AWS.RestrictedParameterValueCombination",
			"Your request contained a restricted parameter combination.  When IdType equals ASIN, SearchIndex cannot be present."}}

	if !reflect.DeepEqual(expectedIsValid, resp.Request.IsValid) {
		t.Errorf("expectedIsValid: %v, got %v", expectedIsValid, resp.Request.IsValid)
	}

	if !reflect.DeepEqual(expectedItemLookupRequest, resp.Request.ItemLookupRequest) {
		t.Errorf("expectedItemLookupRequest: %v, got %v", expectedItemLookupRequest, resp.Request.ItemLookupRequest)
	}

	if !reflect.DeepEqual(expectedErrors, resp.Request.Errors) {
		t.Errorf("expectedErrors: %v, got %v", expectedErrors, resp.Request.Errors)
	}
}
