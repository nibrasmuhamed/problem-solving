package main

import (
	"fmt"
	"sync"
	"time"
)

type SessionCounter struct {
	mu sync.Mutex
}

func (sc *SessionCounter) GetUniqueValue(sessionID string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	// Use the current timestamp to generate a unique value
	timestamp := time.Now().UnixNano()

	// Combine session ID and timestamp to create a unique value
	uniqueValue := int(timestamp) ^ hashString(sessionID)

	return uniqueValue
}

// Simple hash function for string
func hashString(s string) int {
	hash := 0
	for _, c := range s {
		hash = 31*hash + int(c)
	}
	return hash
}

func main() {
	sessionCounter := &SessionCounter{}

	// Simulating timestamp events for session "abc"
	for i := 0; i < 5; i++ {
		sessionID := "abc"
		uniqueValue := sessionCounter.GetUniqueValue(sessionID)

		fmt.Printf("Session ID: %s, Unique Value: %d\n", sessionID, uniqueValue)

		// Sleep for a short duration to simulate different timestamps
		time.Sleep(500 * time.Millisecond)
	}

	// Simulating timestamp events for session "xyz"
	for i := 0; i < 3; i++ {
		sessionID := "xyz"
		uniqueValue := sessionCounter.GetUniqueValue(sessionID)

		fmt.Printf("Session ID: %s, Unique Value: %d\n", sessionID, uniqueValue)

		// Sleep for a short duration to simulate different timestamps
		time.Sleep(500 * time.Millisecond)
	}
}
