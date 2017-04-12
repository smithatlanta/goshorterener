package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestShortenedURL
var projectID = ""

func TestShortenedURL(t *testing.T) {
	var key = "test"
	var value = "http://www.mikeandmel.com"

	// create the shortened url
	result := createShortenedURL(key, value, projectID)
	assert.Nil(t, result)

	// get the shortened url
	val, result := getShortenedURL(key, projectID)
	assert.NotNil(t, val)
	assert.Equal(t, "http://www.mikeandmel.com", val)
	assert.Nil(t, result)

	// delete the shortened url
	result = deleteShortenedURL(key, projectID)
	assert.Nil(t, result)

}

func TestShortenedURL100times(t *testing.T) {
	t.SkipNow()

	var value = "http://www.mikeandmel.com"

	for i := 0; i < 100; i++ {
		var key = "Test" + string(i)
		// create the shortened url
		result := createShortenedURL(key, value, projectID)
		assert.Nil(t, result)

		// get the shortened url
		val, result := getShortenedURL(key, projectID)
		assert.NotNil(t, val)
		assert.Equal(t, "http://www.mikeandmel.com", val)
		assert.Nil(t, result)

		// delete the shortened url
		result = deleteShortenedURL(key, projectID)
		assert.Nil(t, result)
	}
}
