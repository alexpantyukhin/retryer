package retryer

type tryChain struct {
	executeFuncError func() error
	executeFuncBool  func() bool
	strategy         func()
}

func Retryer() tryChain {
	return tryChain{}
}

func (chain tryChain) Retry(numberOfRetry int) {
	for i := 0; i < numberOfRetry; i++ {
		if execute(chain) {
			return
		}
	}

}

func (chain tryChain) RetryForever() {
	for true {
		if execute(chain) {
			return
		}
	}
}

func (chain tryChain) ExecuteError(executeFunc func() error) tryChain {
	return tryChain{
		executeFuncError: executeFunc,
	}
}

func (chain tryChain) ExecuteBool(executeFunc func() bool) tryChain {
	return tryChain{
		executeFuncBool: executeFunc,
	}
}

func execute(chain tryChain) bool {
	if chain.executeFuncBool != nil {
		return chain.executeFuncBool()
	}

	return chain.executeFuncError() == nil
}
