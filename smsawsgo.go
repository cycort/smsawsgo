package smsawsgo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func Send(rcpt string, msg string) (*sns.PublishOutput, error) {
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)
	params := &sns.PublishInput{
		Message:     aws.String(msg),
		PhoneNumber: aws.String(rcpt),
	}
	resp, err := svc.Publish(params)
	return resp, err
}
