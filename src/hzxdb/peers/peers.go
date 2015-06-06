package peers

import(
	"github.com/jinzhu/gorm"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"strconv"
	"strings"

	"hzxdb"
	"hzxconf"
	// "hzxdebug"
)

type Peer struct {
	gorm.Model
    Address string
    Port int64
    Parent uint
    Version string
	Active bool
}

var db gorm.DB

func New(peer Peer) (bool) {
	hzxdb.Connect(&db)

	db.Where("address = ?",peer.Address).First(&peer)
	if db.NewRecord(peer) {
		db.Create(&peer)
		if db.NewRecord(peer) == false {
			return true
		}
	}

	return false
}

func Get() ([]Peer) {
	hzxdb.Connect(&db)

	peers := []Peer{}
	db.Find(&peers)

	if hzxconf.Conf.NodePublic == true {
		ownpeer := Peer{
			Address: hzxconf.Conf.NodeListen,
			Port:    hzxconf.Conf.NodePort,
			Version: hzxconf.Version,
		}

		peers = append(peers,ownpeer)
	}

	return peers
}

func AddressPortString(peer *Peer) (string) {
	return peer.Address+":"+strconv.FormatInt(peer.Port,10)
}

func SrvAddr(srv string) (string) {
	if !strings.Contains(srv,":") {
		srv = srv+":"+strconv.FormatInt(hzxconf.DefaultPort,10)
	}

	return srv
}

func SrvAddrBatch(srv []string) ([]string) {
	for i := 0; i < len(srv); i++ {
		srv[i] = SrvAddr(srv[i])
	}

	return srv
}