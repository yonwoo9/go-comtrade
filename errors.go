package comtrade

import "github.com/pkg/errors"

var (
	ErrUnknownTypeOfData  = errors.New("unknown type of data")
	ErrInvalidFile        = errors.New("invalid file")
	ErrReadFirstLine      = errors.New("read first line error")
	ErrReadSecondLine     = errors.New("read second line error")
	ErrReadADChannel      = errors.New("read A/D channel num line error")
	ErrReadAnalogChannel  = errors.New("read analog channel line error")
	ErrReadDigitalChannel = errors.New("read digital channel line error")
	ErrReadSample         = errors.New("read sample error")
)
