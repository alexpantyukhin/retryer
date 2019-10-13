package retryer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetry_WithBoolean_CallAsRequested(t *testing.T) {
	counter := 0

	Retry(10).
	ExecuteBool(func() bool {
		counter++
		if counter == 100 {
			return true
		}

		return false
	})

	assert.Equal(t, counter, 10)
}

func TestRetry_WithBoolean_CallLessIfTrue(t *testing.T) {
	counter := 0

	Retry(10).
	ExecuteBool(func() bool {
		counter++
		if counter == 3 {
			return true
		}

		return false
	})

	assert.Equal(t, counter, 3)
}

func TestRetry_WithError_CallAsRequested(t *testing.T) {
	counter := 0

	Retry(10).
	ExecuteError(func() error {
		counter++
		if counter == 100 {
			return nil
		}

		return errors.New("Error")
	})

	assert.Equal(t, counter, 10)
}

func TestRetry_WithError_CallLessIfTrue(t *testing.T) {
	counter := 0

	Retry(10).
	ExecuteError(func() error {
		counter++
		if counter == 3 {
			return nil
		}

		return errors.New("Error")
	})

	assert.Equal(t, counter, 3)
}


func TestRetryForever_WithBoolean_CallAsRequested(t *testing.T) {
	counter := 0

	RetryForever().
	ExecuteBool(func() bool {
		counter++
		if counter == 100 {
			return true
		}

		return false
	})

	assert.Equal(t, counter, 100)
}

func TestRetryForever_WithBoolean_CallLessIfTrue(t *testing.T) {
	counter := 0

	RetryForever().
	ExecuteBool(func() bool {
		counter++
		if counter == 3 {
			return true
		}

		return false
	})

	assert.Equal(t, counter, 3)
}