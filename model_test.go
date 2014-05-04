package client_test

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

func TestItemLookupResponseInvalid(t *testing.T) {
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

func TestItemLookupResponseValid(t *testing.T) {
	data := `
	<?xml version="1.0"?>
	<ItemLookupResponse xmlns="http://webservices.amazon.com/AWSECommerceService/2011-08-01">
	    <OperationRequest>
	        <HTTPHeaders>
	            <Header Name="UserAgent" Value="Go 1.1 package http"></Header>
	        </HTTPHeaders>
	        <RequestId>3cd38b18-8600-4573-b3fc-69dba88c3ec8</RequestId>
	        <Arguments>
	            <Argument Name="Service" Value="AWSECommerceService"></Argument>
	            <Argument Name="Timestamp" Value="2014-05-03T20:11:25-07:00"></Argument>
	            <Argument Name="Operation" Value="ItemLookup"></Argument>
	            <Argument Name="AssociateTag" Value="supervacacom-20"></Argument>
	            <Argument Name="Version" Value="2011-08-01"></Argument>
	            <Argument Name="Signature" Value="wnUudm6nS0xOEH/mJI+VDisUDM4SMFpaiGFBcCqtqJI="></Argument>
	            <Argument Name="ItemId" Value="059035342X"></Argument>
	            <Argument Name="AWSAccessKeyId" Value="AKIAIVZ6JESOPVL3CTSA"></Argument>
	        </Arguments>
	        <RequestProcessingTime>0.0091620000000000</RequestProcessingTime>
	    </OperationRequest>
	    <Items>
	        <Request>
	            <IsValid>True</IsValid>
	            <ItemLookupRequest>
	                <IdType>ASIN</IdType>
	                <ItemId>059035342X</ItemId>
	                <ResponseGroup>Small</ResponseGroup>
	                <VariationPage>All</VariationPage>
	            </ItemLookupRequest>
	        </Request>
	        <Item>
	            <ASIN>059035342X</ASIN>
	            <DetailPageURL>http://www.amazon.com/Harry-Potter-Sorcerers-Stone-Book/dp/059035342X%3FSubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D165953%26creativeASIN%3D059035342X</DetailPageURL>
	            <ItemLinks>
	                <ItemLink>
	                    <Description>Technical Details</Description>
	                    <URL>http://www.amazon.com/Harry-Potter-Sorcerers-Stone-Book/dp/tech-data/059035342X%3FSubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D386001%26creativeASIN%3D059035342X</URL>
	                </ItemLink>
	                <ItemLink>
	                    <Description>Add To Baby Registry</Description>
	                    <URL>http://www.amazon.com/gp/registry/baby/add-item.html%3Fasin.0%3D059035342X%26SubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D386001%26creativeASIN%3D059035342X</URL>
	                </ItemLink>
	                <ItemLink>
	                    <Description>Add To Wedding Registry</Description>
	                    <URL>http://www.amazon.com/gp/registry/wedding/add-item.html%3Fasin.0%3D059035342X%26SubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D386001%26creativeASIN%3D059035342X</URL>
	                </ItemLink>
	                <ItemLink>
	                    <Description>Add To Wishlist</Description>
	                    <URL>http://www.amazon.com/gp/registry/wishlist/add-item.html%3Fasin.0%3D059035342X%26SubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D386001%26creativeASIN%3D059035342X</URL>
	                </ItemLink>
	                <ItemLink>
	                    <Description>Tell A Friend</Description>
	                    <URL>http://www.amazon.com/gp/pdp/taf/059035342X%3FSubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D386001%26creativeASIN%3D059035342X</URL>
	                </ItemLink>
	                <ItemLink>
	                    <Description>All Customer Reviews</Description>
	                    <URL>http://www.amazon.com/review/product/059035342X%3FSubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D386001%26creativeASIN%3D059035342X</URL>
	                </ItemLink>
	                <ItemLink>
	                    <Description>All Offers</Description>
	                    <URL>http://www.amazon.com/gp/offer-listing/059035342X%3FSubscriptionId%3DAKIAIVZ6JESOPVL3CTSA%26tag%3Dsupervacacom-20%26linkCode%3Dxm2%26camp%3D2025%26creative%3D386001%26creativeASIN%3D059035342X</URL>
	                </ItemLink>
	            </ItemLinks>
	            <ItemAttributes>
	                <Author>J.K. Rowling</Author>
	                <Creator Role="Illustrator">Mary GrandPré</Creator>
	                <Manufacturer>Scholastic</Manufacturer>
	                <ProductGroup>Book</ProductGroup>
	                <Title>Harry Potter and the Sorcerer's Stone (Book 1)</Title>
	            </ItemAttributes>
	        </Item>
	    </Items>
	</ItemLookupResponse>
	`

	var resp client.ItemLookupResponse

	err := xml.Unmarshal([]byte(data), &resp)
	if err != nil {
		t.Error(err)
	}

	expectedIsValid := true
	expectedItemAttributes := client.ItemAttributes{Author: "J.K. Rowling",
		Title: "Harry Potter and the Sorcerer's Stone (Book 1)"}

	if !reflect.DeepEqual(expectedIsValid, resp.Request.IsValid) {
		t.Errorf("expectedIsValid: %v, got %v", expectedIsValid, resp.Request.IsValid)
	}
	if !reflect.DeepEqual(expectedItemAttributes, resp.Items[0].ItemAttributes) {
		t.Errorf("expectedItemAttributes: %v, got %v", expectedItemAttributes, resp.Items[0].ItemAttributes)
	}
}
