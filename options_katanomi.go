package zap

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

// UpdateCore is an Option for update loogger's core(encoder and LevelEnable)
func UpdateCore(lvl AtomicLevel, cfg Config) Option {
	return optionFunc(func(log *Logger) {
		enc, err := cfg.buildEncoder()
		if err != nil {
			fmt.Fprintf(log.errorOutput, "failed to buildEncoder: %v\n", err)
			return
		}
		sink, _, err := cfg.openSinks()
		if err != nil {
			fmt.Fprintf(log.errorOutput, "failed to sinks: %v\n", err)
			return
		}
		core := zapcore.NewCore(enc, sink, lvl)
		log.core = core
	})
}
