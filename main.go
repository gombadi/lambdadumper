package main

/*
Purpose
=======

This app will take all inputs to the Lambda function and pass them to SNS so you can see the trigger events
given to lambda.

*/

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// set by ldflag during compile
var snsTopic string

func main() {

	var b bytes.Buffer

	b.WriteString("AWS Lambda Function input dumper\n\n")

	b.WriteString("os.Args ============\n")

	// print out the raw commandline arguments
	for k, v := range os.Args {
		b.WriteString(fmt.Sprintf("os.Arg.%v:\n%s\n\n", k, v))
	}

	b.WriteString("\nos.Environ ============\n")

	for k, v := range os.Environ() {
		switch {
		case strings.HasPrefix(v, "AWS_SECRET_ACCESS_KEY"):
			b.WriteString(fmt.Sprintf("os.Env.%v:\nAWS_SECRET_ACCESS_KEY=iforgot.isthataproblem?\n\n", k))
		case strings.HasPrefix(v, "AWS_SESSION_TOKEN"):
			b.WriteString(fmt.Sprintf("os.Env.%v:\nAWS_SESSION_TOKEN=Sorryilostit.Didyoulookunderthesofa?\n\n", k))
		default:
			b.WriteString(fmt.Sprintf("os.Env.%v:\n%s\n\n", k, v))
		}
	}

	// if we have an snsTopic set then send to it otherwise print to stdout
	if snsTopic != "" {
		snsparams := &sns.PublishInput{
			Message:  aws.String(b.String()),                                                          // Report output
			Subject:  aws.String(fmt.Sprintf("Lambda Function Event Dumper - %v", time.Now().Unix())), // Message Subject for emails
			TopicArn: aws.String(snsTopic),
		}

		// push the report to SNS for distribution
		snssvc := sns.New(session.New())
		publishResp, err := snssvc.Publish(snsparams)
		if err != nil {
			fmt.Printf("error publishing to AWS SNS: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Published to AWS SNS: %s\n", *publishResp.MessageId)
	} else {
		fmt.Printf("%s\n", b.String())
	}
}

/*

 */
