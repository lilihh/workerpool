package workerpool

import (
	"fmt"
)

type errorCode int

const (
	IllegalArgument      errorCode = 1001
	NormalBufferOverflow errorCode = 1002
	UrgentBufferOverflow errorCode = 1003
)

var errorMessage = map[errorCode]string{
	IllegalArgument:      "the argument(s) is illegal",
	NormalBufferOverflow: "the buffer of normal tasks is over flow",
	UrgentBufferOverflow: "the buffer of urgent tasks is over flow",
}

func newError(errCode errorCode) error {
	return fmt.Errorf("workerpool: %d %s", errCode, errorMessage[errCode])
}
