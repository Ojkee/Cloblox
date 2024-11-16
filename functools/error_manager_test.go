package functools_test

import (
	"errors"
	"testing"

	"Cloblox/functools"
)

func TestErrorManager_1(t *testing.T) {
	em := functools.NewErrorManager(nil)
	em.AppendNew(errors.New("LMAO"))
	if em.ContainsStongError() == true {
		t.Error()
	}
	em.AppendNew(functools.NewStrongError("LMAO1", "LMAO2"))
	if em.ContainsStongError() == false {
		t.Error()
	}
}
