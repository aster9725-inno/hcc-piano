package logger

import (
	"hcc/piano/lib/syscheck"
	"testing"
)

func Test_CreateDirIfNotExist(t *testing.T) {
	err := CreateDirIfNotExist("/var/log/" + LogName)
	if err != nil {
		t.Fatal("Failed to create dir!")
	}
}

func Test_Logger_Prepare(t *testing.T) {
	if !syscheck.CheckRoot() {
		t.Fatal("Failed to get root permission!")
	}

	if !Prepare() {
		t.Fatal("Failed to prepare logger!")
	}
	defer FpLog.Close()
}
