package retryer

import "time"

type tryChain struct {
	executeFuncError func() error
	executeFuncBool  func() bool
	strategy         func()
}

// Retry is used for running the action until number of retries is reached or the action is runned successfully.
func Retry(numberOfRetry int) *tryChain {
	chain := &tryChain{}

	chain.strategy = func() {
		for i := 0; i < numberOfRetry; i++ {
			if execute(chain) {
				return
			}
		}
	}

	return chain
}

// RetryForever is used for running the action forever until the action is runned successfully.
func RetryForever() *tryChain {
	chain := &tryChain{}

	chain.strategy = func() {
		for true {
			if execute(chain) {
				return
			}
		}
	}

	return chain
}

// RetryAndWait is used for running the action and wait durationsNanoseconds until number of retries is reached or action is runned successfully.
func RetryAndWait(durationsNanoseconds []time.Duration) *tryChain {
	chain := &tryChain{}

	chain.strategy = func() {
		for i := 0; i < len(durationsNanoseconds); i++ {
			if execute(chain) {
				return
			}

			time.Sleep(durationsNanoseconds[i])
		}
	}

	return chain
}

// RetryAndWaitForever is used for running the action and wait until action is runned successfully.
func RetryAndWaitForever(durationAttemptFunc func(int) time.Duration) *tryChain {
	chain := &tryChain{}

	chain.strategy = func() {
		counter := 0

		for true {
			if execute(chain) {
				return
			}

			counter++

			time.Sleep(durationAttemptFunc(counter))
		}
	}

	return chain
}

// ExecuteError attach action which returns error if the action ran unsuccessfully
func (chain *tryChain) ExecuteError(executeFunc func() error) {
	chain.executeFuncError = executeFunc
	chain.strategy()
}

// ExecuteError attach action which returns false if the action ran unsuccessfully
func (chain *tryChain) ExecuteBool(executeFunc func() bool) {
	chain.executeFuncBool = executeFunc
	chain.strategy()
}

func execute(chain *tryChain) bool {
	if chain.executeFuncBool != nil {
		return chain.executeFuncBool()
	}

	return chain.executeFuncError() == nil
}
