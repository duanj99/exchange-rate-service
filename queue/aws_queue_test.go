package queue

import (
	"testing"
)

//export AWS_ACCESS_KEY_ID=AKIAYJ3LGLWDFDXT5TOL
//export AWS_SECRET_ACCESS_KEY=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
//export AWS_DEFAULT_REGION=us-west-2

func TestSQSQueueCreation(t *testing.T) {
	queue := InitAwsQueue()

	//queue.CreateQueue("darryl_jiang_exchange_service")
	////time.Sleep(10 * time.Second)
	queue.ListQueue()
}

func TestSQSQueueList(t *testing.T) {
	queue := InitAwsQueue()

	queue.CreateQueue("darryl_jiang_exchange_service")
}
