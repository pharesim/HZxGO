package debug

import(
	"io"
	"log"
	"os"
	"runtime"

	"hzxconf"
)

var (
    Info    *log.Logger
    Error   *log.Logger
    Conf    hzxconf.Config
)

func Init() {
	startup(os.Stdout, os.Stderr)
}

func startup(
    infoHandle io.Writer,
    errorHandle io.Writer) {

	Conf = hzxconf.GetConfiguration()

	if Conf.Debug == true {
	    Info = log.New(infoHandle,
	        "INFO: ",
	        log.Ldate|log.Ltime)

	    Error = log.New(errorHandle,
	        "ERROR: ",
	        log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func GetOS() (string) {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		// freebsd, openbsd,
		// plan9, windows...
		return os
	}
}