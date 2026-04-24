package src

import (
	"fmt"
	"os"
	"time"
)

func LoadingAction(label string, task func() error) {
	chars := []string{"|", "/", "-", "\\"}
	done := make(chan error)
	go func() { done <- task() }()

	for i := 0; ; i++ {
		select {
		case err := <-done:
			fmt.Printf("\r%s... Done.\n", label)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			return
		default:
			fmt.Printf("\r%s %s", label, chars[i%len(chars)])
			time.Sleep(100 * time.Millisecond)
		}
	}
}
