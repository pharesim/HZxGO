package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"

	"hzxconf"
	// "hzxdb"
	
	"hzxdb/peers"
)

type Reply struct {
	Request string
	Status  string
	Data    map[string]string
}

var (
	reply   Reply
)

func nodeHandler(w http.ResponseWriter, r *http.Request) {
	request := r.URL.Path[1:]
	reply := requests(request,r)
	result, _ := json.Marshal(reply)
	fmt.Fprintf(w, string(result))
}

func ServeNode() {
	conf := &hzxconf.Conf
	serverMuxNODE := http.NewServeMux()
	serverMuxNODE.HandleFunc("/", nodeHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d",conf.NodeListen,conf.NodePort), serverMuxNODE)
}

func requests(request string, r *http.Request) (*Reply) {
	empty := &Reply{Request: request, Status: "ERROR", Data: map[string]string{"message":"Unknown request"}}
	switch request {
	case "getStatus":
		return &Reply{Request: "getStatus", Status: "OK", Data: map[string]string{"version":hzxconf.Version}}
	case "getPeers":
		res, _ := json.Marshal(peers.Get())
		return &Reply{Request: "getPeers", Status: "OK", Data: map[string]string{"peers":string(res)}}
	default:
		return empty
	}
}