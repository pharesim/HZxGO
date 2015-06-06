package hzxconf

import (
	"os"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Debug      bool
	Database   string
    UIListen   string
    UIPort     int64
    NodeListen string
    NodePort   int64
    NodePublic bool
    Nodes      []string
}

var Conf Config

var	Version string = "0.1"
var	DefaultPort int64 = 8260

func fileExists(name string) (bool) {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
            }
    }
    return true
}

func GetConfiguration() (Config) {
	filename := "config/hzx.conf.toml"

	if fileExists(filename) == false {
	    panic("Config file not found")
	}

	if _, err := toml.DecodeFile(filename, &Conf); err != nil {
	    panic("Config file parse error")
	}

	return Conf
}