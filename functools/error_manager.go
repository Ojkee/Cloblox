package functools

import "fmt"

type ErrorManager struct {
	errs []error
}

func NewErrorManager(errs *[]error) *ErrorManager {
	if errs == nil {
		return &ErrorManager{
			errs: make([]error, 0),
		}
	}
	return &ErrorManager{
		errs: *errs,
	}
}

func (em *ErrorManager) ContainsStongError() bool {
	for _, err := range em.errs {
		if em.IsStrong(err) {
			return true
		}
	}
	return false
}

func (em *ErrorManager) GetConsoleErrors() []string {
	retVal := make([]string, 0)
	for _, err := range em.errs {
		retVal = append(retVal, err.Error())
	}
	return retVal
}

func (em *ErrorManager) GetDebugErrors() []string {
	retVal := make([]string, 0)
	for _, err := range em.errs {
		switch strongErr := err.(type) {
		case *StrongError:
			retVal = append(retVal, strongErr.Debug())
		}
	}
	return retVal
}

func (em *ErrorManager) AppendNew(err error) {
	if err != nil {
		em.errs = append(em.errs, err)
	}
}

func (em *ErrorManager) AppendNewErrors(err []error) {
	if err != nil {
		em.errs = append(em.errs, err...)
	}
}

func (em *ErrorManager) Clear() {
	em.errs = make([]error, 0)
}

func (em *ErrorManager) StrongErrorCount() int {
	retVal := 0
	for _, err := range em.errs {
		if em.IsStrong(err) {
			retVal++
		}
	}
	return retVal
}

func (em *ErrorManager) IsStrong(err error) bool {
	switch err.(type) {
	case *StrongError:
		return true
	}
	return false
}

func (em *ErrorManager) PrintAllErrors() {
	fmt.Println("**************ERRORS**************************************")
	fmt.Printf("ERRORS: %d\t\tSTRONG ERRORS: %d\n", len(em.errs), em.StrongErrorCount())
	for _, err := range em.errs {
		switch strongErr := err.(type) {
		case *StrongError:
			fmt.Println(strongErr.Debug())
			break
		}
		fmt.Println(err.Error())
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}
