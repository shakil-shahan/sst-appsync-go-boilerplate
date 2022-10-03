package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Event struct {
	Arguments struct {
		Input map[string]string `json:"input"`
	}
	Identity string `json:"identity"`
	Info     struct {
		FieldName           string            `json:"fieldName"`
		ParentTypeName      string            `json:"parentTypeName"`
		SelectionSetGraphQL string            `json:"selectionSetGraphQL"`
		SelectionSetList    []string          `json:"selectionSetList"`
		Variables           map[string]string `json:"variables"`
	}
	Prev    string `json:"prev"`
	Request struct {
		Headers struct {
			Accept                    string `json:"accept"`
			AcceptEncoding            string `json:"accept-encoding"`
			AcceptLanguage            string `json:"accept-language"`
			CloudfrontForwardedProto  string `json:"cloudfront-forwarded-proto"`
			CloudfrontIsDesktopViewer string `json:"cloudfront-is-desktop-viewer"`
			CloudfrontIsMobileViewer  string `json:"cloudfront-is-mobile-viewer"`
			CloudfrontIsSmarttvViewer string `json:"cloudfront-is-smarttv-viewer"`
			CloudfrontViewerCountry   string `json:"cloudfront-viewer-country"`
			CloudfrontIsTabletViewer  string `json:"cloudfront-is-tablet-viewer"`
			ContentLength             string `json:"content-length"`
			ContentType               string `json:"content-type"`
			Host                      string `json:"host"`
			Hrigin                    string `json:"origin"`
			Referer                   string `json:"Referer"`
			SecFetchDest              string `json:"sec-fetch-dest"`
			SecFetchMode              string `json:"sec-fetch-mode"`
			SecFetchSite              string `json:"sec-fetch-site"`
			UserAgent                 string `json:"user-agent"`
			Via                       string `json:"via"`
			XAmzCfID                  string `json:"x-amz-cf-id"`
			XAmzUserAgent             string `json:"x-amz-user-agent"`
			XAmznTraceID              string `json:"x-amzn-trace-id"`
			XApiKey                   string `json:"x-api-key"`
			XForwardedFor             string `json:"x-forwarded-for"`
			XForwardedPort            string `json:"x-forwarded-port"`
			XForwardedProto           string `json:"x-forwarded-proto"`
		}
	}
	Source string            `json:"source"`
	Stash  map[string]string `json:"stash"`
}

type Response struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type ErrorHandler struct {
	Type    string `json:"error_type"`
	Message string `json:"error_message"`
}

var svc dynamodb.DynamoDB

const tableName string = "test-sharebus-appsync-Notes"

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc = *dynamodb.New(session)

	lambda.Start(Handler)
}

func Handler(ctx context.Context, iEvent interface{}) (Response, error) {
	fmt.Printf("---------------{LAMBDA ctx}---------------\n")
	// Event context
	lc, _ := lambdacontext.FromContext(ctx)
	fmt.Println("AwsRequestID:", lc.AwsRequestID)
	fmt.Println("Identity:", lc.Identity)
	fmt.Println("InvokedFunctionArn:", lc.InvokedFunctionArn)
	fmt.Println("ClientContext:", lc.ClientContext)
	fmt.Println("ClientContext.Client:", lc.ClientContext.Client)
	fmt.Println("CognitoIdentityID:", lc.Identity.CognitoIdentityID)
	fmt.Println("CognitoIdentityPoolID:", lc.Identity.CognitoIdentityPoolID)

	fmt.Printf("---------------{Event Decode}---------------\n")
	var event Event
	eventJsonm, _ := json.MarshalIndent(iEvent, "", "  ")
	eventReader := bytes.NewReader([]byte(eventJsonm))
	json.NewDecoder(eventReader).Decode(&event)

	fmt.Println("Arguments:", event.Arguments)
	fmt.Println("Identity:", event.Identity)
	fmt.Println("Info.FieldName:", event.Info.FieldName)
	fmt.Println("Info.ParentTypeName:", event.Info.ParentTypeName)
	fmt.Println("Info.SelectionSetGraphQL:", event.Info.SelectionSetGraphQL)

	fmt.Println("Payload:", event.Arguments.Input)
	fmt.Println("Payload id:", event.Arguments.Input["id"])
	// Iterating map using for rang loop
	fmt.Printf("---------------{Payload by loop}---------------\n")
	for id, val := range event.Arguments.Input {
		fmt.Println(id, val)
	}
	// Convert to JSON
	fmt.Printf("---------------{Payload converted into JSON Object}---------------\n")
	inputJsonm, _ := json.Marshal(event.Arguments.Input)
	fmt.Println(string(inputJsonm))

	switch event.Info.FieldName {
	case "getOne":
		return getOne(event)
	default:
		return Response{}, errors.New("Wrong method: " + event.Info.FieldName)
	}
}

func getOne(event Event) (Response, error) {
	dbResult, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(event.Arguments.Input["id"]),
			},
		},
	})
	if err != nil {
		fmt.Println("Failed to marshall request")
		return Response{}, err
	}

	if dbResult.Item == nil {
		msg := "Could not find '" + event.Arguments.Input["id"] + "'"
		return Response{}, errors.New(msg)
	}

	result := Response{}

	err = dynamodbattribute.UnmarshalMap(dbResult.Item, &result)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Item:", result)

	return result, nil
}
