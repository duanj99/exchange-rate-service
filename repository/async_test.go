package repository

import (
	"CurrencyExchangeService/logger"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	mockLogger := logger.NewLogger()
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup
	// counter for each.
	for i := 1; i <= 5; i++ {

		wg.Add(1)
		// Avoid re-use of the same `i` value in each goroutine closure.
		// See [the FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		// for more details.
		i := i
		mockLogger.Info(fmt.Sprintf("Worker %d Added\n", i))

		// Wrap the worker call in a closure that makes sure to tell
		// the WaitGroup that this worker is done. This way the worker
		// itself does not have to be aware of the concurrency primitives
		// involved in its execution.
		go func() {
			defer wg.Done()
			worker(i, mockLogger)
		}()
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they're done.
	wg.Wait()
}

// This is the function we'll run in every goroutine.
func worker(id int, logger *logger.ServiceLogger) {
	logger.Info(fmt.Sprintf("Worker %d starting\n", id))

	// Sleep to simulate an expensive task.
	time.Sleep(2 * time.Second)
	logger.Info(fmt.Sprintf("Worker %d done\n", id))
}

// Container holds a map of counters; since we want to
// update it concurrently from multiple goroutines, we
// add a `Mutex` to synchronize access.
// Note that mutexes must not be copied, so if this
// `struct` is passed around, it should be done by
// pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Lock the mutex before accessing `counters`; unlock
	// it at the end of the function using a [defer](defer)
	// statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func TestMuxLock(t *testing.T) {
	c := Container{
		// Note that the zero value of a mutex is usable as-is, so no
		// initialization is required here.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// This function increments a named counter
	// in a loop.
	doIncrement := func(name string, n int, p int) {
		fmt.Println("thread started", p)
		for i := 0; i < n; i++ {
			c.inc(name)
		}

		time.Sleep(10 * time.Second)
		fmt.Println("thread completed", p)

		wg.Done()
	}

	// Run several goroutines concurrently; note
	// that they all access the same `Container`,
	// and two of them access the same counter.
	wg.Add(3)
	go doIncrement("a", 10000, 1)
	go doIncrement("a", 10000, 2)
	go doIncrement("b", 10000, 3)

	fmt.Println("all threads started")

	// Wait for the goroutines to finish
	wg.Wait()

	fmt.Println("all threads completed, wait done")
	fmt.Println(c.counters)
}

func TestChannel(t *testing.T) {
	jobs := make(chan int, 100)
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
				time.Sleep(10 * time.Second)
				fmt.Println("Processed job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the `jobs`
	// channel, then closes it.
	for j := 1; j <= 100; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	fmt.Println("sent all jobs, continue other process")

	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.
	<-done
}
