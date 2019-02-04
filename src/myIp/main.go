package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
) 

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is what Lambda hands the event.
// The Event format can be found here: https://docs.aws.amazon.com/lambda/latest/dg/eventsources.html#eventsources-api-gateway-request
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (Response, error) {
	// fmt.Println(event)
	// return fmt.Sprintf("Hello %v!", event.RequestContext.Identity.SourceIP), nil

	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"IP": event.RequestContext.Identity.SourceIP,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "myip-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
