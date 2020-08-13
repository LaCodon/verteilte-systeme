package lg

import (
	"github.com/op/go-logging"
	"os"
)

var Log *logging.Logger

func init() {
	var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)

	stdOut := logging.NewLogBackend(os.Stdout, "", 0)
	stdOutFormatted := logging.NewBackendFormatter(stdOut, format)

	logging.SetBackend(stdOutFormatted)
	logging.SetLevel(logging.INFO, "")

	Log = logging.MustGetLogger("smkvs")
}
