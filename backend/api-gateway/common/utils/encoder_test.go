package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeBase62(t *testing.T) {
	// Test case 1: Input 0 should result in "0" in base62
	result := EncodeBase62(0)
	assert.Equal(t, "", result)

	// Test case 2: Input 10 should result in "a" in base62
	result = EncodeBase62(10)
	assert.Equal(t, "a", result)

	// Test case 3: Input 61 should result in "Z" in base62
	result = EncodeBase62(61)
	assert.Equal(t, "Z", result)

	// Test case 4: Input 62 should result in "10" in base62
	result = EncodeBase62(62)
	assert.Equal(t, "10", result)

	// Test case 5: Input 100 should result in "1C" in base62
	result = EncodeBase62(100)
	assert.Equal(t, "1C", result)

	// Add more test cases for different inputs and expected outputs
}
