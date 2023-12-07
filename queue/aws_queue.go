package queue

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type AwsQueue struct {
	queueService *sqs.SQS
}

func InitAwsQueue() AwsQueue {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("us-west-2"),
		},
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	fmt.Printf("New AWS Session created")
	return AwsQueue{
		queueService: svc,
	}
}

func (q *AwsQueue) ListQueue() {
	result, err := q.queueService.ListQueues(nil)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i, *url)
	}
}

func (q *AwsQueue) CreateQueue(queueNameInput string) {
	createQueueInput := &sqs.CreateQueueInput{
		QueueName: &queueNameInput,
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("20"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	}

	_, err := q.queueService.CreateQueue(createQueueInput)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	//fmt.Printf("URL: %s", *result.QueueUrl)
}

func (q *AwsQueue) Produce(message string, queueUrl string) {
	sendMessageInput := sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(message),
		QueueUrl:     &queueUrl,
	}

	result, err := q.queueService.SendMessage(&sendMessageInput)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	fmt.Print(result.String())
}

func (q *AwsQueue) Consume(queueUrl string) {
	timeOut := int64(10)
	recivedMessageInput := sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeOut,
	}

	result, err := q.queueService.ReceiveMessage(&recivedMessageInput)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	fmt.Print(result.String())
}
