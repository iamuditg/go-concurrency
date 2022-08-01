package main

/*
The problem describes two processes, the producer and the consumer,
who share a common, fixed-size buffer used as a queue. The producer's
job is to generate data, put it into the buffer, and start again.
At the same time, the consumer is consuming the data (i.e., removing it
from the buffer), one piece at a time. The problem is to make sure that
the producer won't try to add data into the buffer if it's full and that
the consumer won't try to remove data from an empty buffer. The solution
for the producer is to either go to sleep or discard data if the buffer is full.
The next time the consumer removes an item from the buffer, it notifies the
producer, who starts to fill the buffer again. In the same way,
the consumer can go to sleep if it finds the buffer empty.
The next time the producer puts data into the buffer, it wakes up the
sleeping consumer.
*/

type Consumer struct {
	msgs *chan int
}

// NewConsumer creates a Consumer
func NewConsumer(msgs *chan int) *Consumer {
	return &Consumer{msgs: msgs}
}

func (c *Consumer) consume() {
	fmt.Println("consumer: Stared")
	for {
		msg := <-*c.msgs
		fmt.Println("consume received:", msg)
	}
}

type Producer struct {
	msgs *chan int
	done *chan bool
}

// NewProducer creates a Producer
func NewProducer(msgs *chan int, done *chan bool) *Producer {
	return &Producer{msgs: msgs, done: done}
}

func (p *Producer) produce(max int) {
	fmt.Println("produce started")
	for i := 0; i < max; i++ {
		*p.msgs <- i
	}
	*p.done <- true
	fmt.Println("produce: Done")
}
