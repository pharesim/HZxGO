package hzxdb

import(
	"github.com/jinzhu/gorm"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"

	"hzxconf"
	"hzxdebug"
)

type checkForPing interface {
	Ping()
}

type Ping struct {
}

type badType struct {
}

var Db gorm.DB

func connected(v interface{}) (bool) {
	if _, ok := v.(checkForPing); ok {
		return true
	}

	return false
}

func Connect(db *gorm.DB) {
	if !connected(db) {
		var err error
		*db, err = gorm.Open("sqlite3",hzxconf.Conf.Database)
		if err != nil {
			debug.Error.Println(err.Error())
		}

		err = db.DB().Ping()
		if err != nil {
			panic(err)
		}
	}
}

func Init() (gorm.DB) {
	Db, err := gorm.Open("sqlite3",hzxconf.Conf.Database)
	if err != nil {
		debug.Error.Println(err.Error())
	}

	err = Db.DB().Ping()
	if err != nil {
		panic(err)
	}

	debug.Info.Println("Connected to database")

	return Db
}

func StringInSlice(a string, list []string) (bool) {
    for _, b := range list {
        if b == a {
            return true
        }
    }

    return false
}