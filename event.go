package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func reqHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var body = `{"commands": [{"type":"com.okta.assertion.patch","value":[{"op":"add","path":"/claims/foo","value":{"attributes":{"NameFormat":"urn:oasis:names:tc:SAML:2.0:attrname-format:basic"},"attributeValues": [{"attributes":{"xsi:type":"xs:string"},"value":"bearer"}]}}]}]}`
	var body = `{"commands": [{"type":"com.okta.assertion.patch","value":[{"op":"replace","path":"/authentication/authnContext","value":{"authnContextClassRef":"Something:different?"}}]}]}`
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(reqHandler)
}
