package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/immesys/bw2bind"
)

func tap(uri string) {
	bwc := bw2bind.ConnectOrExit("127.0.0.1:28588")
	_, err := bwc.SetEntity([]byte(ourEntity)[1:])
	if err != nil {
		fmt.Printf("Could not obtain permissions: %v\n", err)
		os.Exit(1)
	}
	sub := bwc.SubscribeOrExit(&bw2bind.SubscribeParams{
		URI:       uri,
		AutoChain: true,
	})
	fmt.Printf("=========================\n")
	//rvc := make(chan map[string]float64, 1000)
	for msg := range sub {
		po := msg.GetOnePODF("2.0.11.2")
		if po == nil {
			fmt.Printf("po mismatch\n")
			continue
		}
		tgt := make(map[string]interface{})
		po.(bw2bind.MsgPackPayloadObject).ValueInto(&tgt)
		startidx := strings.Index(msg.URI, "s.hamilton")
		serial := msg.URI[startidx+11 : startidx+11+4]
		iserial, err := strconv.ParseInt(serial, 16, 32)
		if err != nil {
			panic(err)
		}
		tgt["serial"] = fmt.Sprintf("0x%04x", iserial)
		s, _ := json.MarshalIndent(tgt, "", "  ")
		fmt.Printf("%s\n", string(s))
		//spew.Dump(tgt)
	}
}
