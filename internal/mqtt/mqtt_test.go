package mqtt

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	// Capture the standard output for later assertion
	expectedOutput := captureStandardOutput(func() {
		// Call the Publish method with a test message
		Publish("Test Message")
	})

	// Assert that the expected print statements are made
	assert.Contains(t, expectedOutput, "Mqtt client is Connected Successfully")
}

// captureStandardOutput captures the output of a function that writes to the standard output
func captureStandardOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return (buf.String())

}

func TestCreateClient(t *testing.T) {
	
	// Assume a reset client
	client = nil

	// Call the createClient function
	result := createClient()

	// Check that the client is not nil
	assert.NotNil(t, result)

	// Call the createClient function once again
	// This time, check that the client is not nil and is the same as the previous one
	result2 := createClient()
	assert.NotNil(t, result2)
	assert.Equal(t, result, result2)
}
