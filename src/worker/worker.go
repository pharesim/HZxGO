package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"hzxconf"
	"hzxdb"
	"hzxdebug"
	"node"

	"hzxdb/peers"
)

func Worker() {
	hzxdb.Init()

	go peersWorker();
}

func peersWorker() {
	var sleepTime time.Duration = 300

	for {
		nodes := peers.SrvAddrBatch(hzxconf.Conf.Nodes)
		saved := peers.Get()
		for i := 1; i < len(saved); i++ {
			nodes = append(nodes,peers.AddressPortString(&saved[i]))
		}
		
		amount := len(nodes)
		for i := 0; i < amount; i++ {
			srv := nodes[i]
			var k int64 = 0
			getpeers := sendRequest("getPeers",srv)
			if getpeers.Status != "ERROR" {
				data := []peers.Peer{}
				json.Unmarshal([]byte(getpeers.Data["peers"]),&data)
				for j := 0; j < len(data); j++ {
					if hzxdb.StringInSlice(peers.AddressPortString(&data[j]),nodes) == false {
						status := sendRequest("getStatus",peers.AddressPortString(&data[j]))
						if status.Status == "OK" {
							data[j].Active = true
						} else {
							data[j].Active = false
						}

						if peers.New(data[j]) {
							k++;
						}
					}
				}
			}

			if k > 0 {
				word := "Peer"
				if k > 1 {
					word = "Peers"
				}

				debug.Info.Println(strconv.FormatInt(k,10)+" "+word+" added")
			}		
		}

		time.Sleep(sleepTime * time.Second)
	}
}

func sendRequest(requestType, srv string) (*node.Reply) {
	resp, err := http.Get(fmt.Sprintf("http://%s/%s",peers.SrvAddr(srv),requestType))
	if err != nil {
		debug.Error.Println(err.Error())
		return &node.Reply{Status: "FAILED", Data: map[string]string{"message":err.Error()}}
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		result := &node.Reply{}
		json.Unmarshal(body, &result)
		return result
	}
}