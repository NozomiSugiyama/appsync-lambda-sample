package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func hander(arg Post) (Post, error) {

	fmt.Print("Arg.ID", arg.ID)

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(arg.ID),
			},
			"title": {
				S: aws.String(arg.Title),
			},
		},
		TableName: aws.String("Post"),
	}

	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return Post{}, err
	}

	svc := dynamodb.New(session)

	_, err = svc.PutItem(input)

	if err != nil {
		return Post{}, err
	}

	return arg, nil
}

func main() {
	lambda.Start(hander)
}
