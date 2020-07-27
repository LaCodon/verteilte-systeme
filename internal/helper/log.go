package helper

import (
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"os"
	"path"
	"path/filepath"
)

func GetLogFilePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		lg.Log.Warning(err)
	}
	return path.Join(dir, "../logs", config.Default.Logfile)
}
