package main

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

func main() {
	defer slog.MustFlush()
	const myTemplate = "{{datetime}} channel={{channel}} level={{level}} [file={{caller}}] message={{message}} data={{data}}\n"

	// DangerLevels 包含: slog.PanicLevel, slog.ErrorLevel, slog.WarnLevel
	h1 := handler.MustRotateFile(".\\app_error.log", rotatefile.EveryDay,
		handler.WithLogLevels(slog.DangerLevels),
	)

	// NormalLevels 包含: slog.InfoLevel, slog.NoticeLevel, slog.DebugLevel, slog.TraceLevel
	h2 := handler.MustRotateFile(".\\app_info.log", rotatefile.EveryHour,
		handler.WithLogLevels(slog.NormalLevels),
	)
	f := slog.AsTextFormatter(h1.Formatter())
	f.SetTemplate(myTemplate)
	slog.PushHandler(h1)
	slog.PushHandler(h2)

	// add logs
	slog.Info("info message text")
	slog.Error("error message text")
}
