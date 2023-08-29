package main

import (
	"CwIndustries/apiHandlers"
	"github.com/aws/aws-lambda-go/lambda"
)



func main() {
	lambda.Start(apiHandlers.Router)
}

