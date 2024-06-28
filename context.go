//CREATING A CONTEXT WITH TIMEOUT

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) //context.Background() returns an empty context, and context.WithTimeout wraps it with a timeout
	defer cancel()                                                          //Ensures that the cancel function is called when the main function exits, which will release resources associated with the context.

	go performTask(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	}
}

func performTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	}
}
