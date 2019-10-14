package retryer

import "time"

type tryChain struct {
	executeFuncError func() error
	executeFuncBool  func() bool
	strategy         func()
}

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

func RetryAndWaitForever(durationAttemptFunc func(int) time.Duration) *tryChain {
	chain := &tryChain{}

	chain.strategy = func() {
		counter := 0

		for true {
			if execute(chain) {
				return
			}

			counter += 1

			time.Sleep(durationAttemptFunc(counter))
		}
	}

	return chain
}

func (chain *tryChain) ExecuteError(executeFunc func() error) {
	chain.executeFuncError = executeFunc
	chain.strategy()
}

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
