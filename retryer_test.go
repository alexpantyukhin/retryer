package retryer

import (
	"errors"
	"testing"
	"time"

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

func TestRetryAndWait_WithBoolean_CallAsRequested(t *testing.T) {
	counter := 0

	t0 := time.Now()
	RetryAndWait([]time.Duration{1000 * time.Millisecond, 1000 * time.Millisecond}).
	ExecuteBool(func() bool {
		counter++
		if counter == 10 {
			return true
		}

		return false
	})
	t1 := time.Now()
	diff := t1.Sub(t0)

	assert.GreaterOrEqual(t, int64(diff), int64(2000*time.Millisecond))
}

func TestRetryAndWait_WithBoolean_CallLessIfTrue(t *testing.T) {
	counter := 0

	t0 := time.Now()
	RetryAndWait([]time.Duration{1000 * time.Millisecond, 1000 * time.Millisecond}).
	ExecuteBool(func() bool {
		counter++
		if counter == 1 {
			return true
		}

		return false
	})
	t1 := time.Now()
	diff := t1.Sub(t0)

	assert.GreaterOrEqual(t, int64(diff), int64(0))
}
