package foo

import (
	"testing"
)

func TestMyUnexportedFunc(t *testing.T) {

	myUnexportedFunc()

}
