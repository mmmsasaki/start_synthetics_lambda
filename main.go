package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/synthetics"
)

const region = "ap-northeast-1"

func main() {
	lambda.Start(startSynthetics)
}

func startSynthetics() error {
	name := os.Getenv("CANARY_NAME")
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Fatal(err)
	}

	svc := synthetics.New(sess)
	input := &synthetics.StartCanaryInput{
		Name: aws.String(name),
	}
	_, err = svc.StartCanary(input)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
