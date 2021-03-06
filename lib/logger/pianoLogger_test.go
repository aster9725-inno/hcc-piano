package logger

import (
	"hcc/piano/lib/errors"
	"testing"
)

func Test_CreateDirIfNotExist(t *testing.T) {
	err := CreateDirIfNotExist("/var/log/" + LogName)
	if err != nil {
		t.Fatal("Failed to create dir!")
	}
}

func Test_Logger_Prepare(t *testing.T) {
	err := Init()
	if err != nil {
		errors.SetErrLogger(Logger)
		errors.NewHccError(errors.PianoInternalInitFail, "logger.Init(): "+err.Error()).Fatal()
	}
	errors.SetErrLogger(Logger)
	defer func() {
		_ = FpLog.Close()
	}()
}
