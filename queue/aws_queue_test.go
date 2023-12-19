package queue

import (
	"testing"
)

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
