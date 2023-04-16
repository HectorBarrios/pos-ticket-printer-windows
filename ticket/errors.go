package ticket

import (
	"errors"
	"fmt"
)

var (
	ErrFieldNotProvided            = errors.New("field not provided")
	ErrFieldPrinterNameNotProvided = fmt.Errorf("%w: printer name", ErrFieldNotProvided)
)
