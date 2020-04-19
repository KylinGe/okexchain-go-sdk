package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)


func TestQueryTx(t *testing.T) {
	cli := NewClient(rpcUrl)
	// get tx hash bytes
	txHash, err := hex.DecodeString("12CF714D13D9B86EDCCBE41BF55845BF96613977AFF8E503C5A5349A50841F9A")
	assertNotEqual(t, err, nil)
	resp, err := cli.QueryTx(txHash, true)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryProposals(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryProposals()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryProposalByID(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryProposalByID(1)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}
