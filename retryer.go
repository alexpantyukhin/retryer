package retryer

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
