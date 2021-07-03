package zap

import (
	"fmt"
	"testing"
)

func TestUpdateCore(t *testing.T) {
	logger, err := NewProduction()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	level := NewAtomicLevel()
	logger = logger.WithOptions(UpdateCore(level, NewProductionConfig()))
	level.SetLevel(DebugLevel)
	if !logger.Core().Enabled(DebugLevel) {
		t.Fail()
	}
}
