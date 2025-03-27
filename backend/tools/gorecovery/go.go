package gorecovery

import "log/slog"

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic", slog.Any("error", err))
			}
		}()
		f()
	}()
}
