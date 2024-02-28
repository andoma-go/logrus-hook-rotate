package rotatehook

import (
	"sync"

	"github.com/andoma-go/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
	Formatter  logrus.Formatter
	Level      logrus.Level
	Enabled    bool
}

type RotateHook struct {
	logger *lumberjack.Logger
	mu     sync.RWMutex
	cfg    *Config
}

func New(cfg *Config) logrus.Hook {
	r := &RotateHook{
		cfg: cfg,
		logger: &lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,
			MaxAge:     cfg.MaxAge,
			MaxBackups: cfg.MaxBackups,
			LocalTime:  cfg.LocalTime,
			Compress:   cfg.Compress,
		},
	}

	return r
}

func (r *RotateHook) Levels() []logrus.Level {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return logrus.AllLevels[:r.cfg.Level+1]
}

func (r *RotateHook) Fire(entry *logrus.Entry) (err error) {
	r.mu.RLock()
	enabled := r.cfg.Enabled
	r.mu.RUnlock()

	b, err := r.cfg.Formatter.Format(entry)
	if err != nil {
		return err
	}
	if enabled {
		r.logger.Write(b)
	}

	return nil
}

func (r *RotateHook) IsEnabled() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.cfg.Enabled
}

func (r *RotateHook) SetEnabled(enabled bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cfg.Enabled = enabled
}
