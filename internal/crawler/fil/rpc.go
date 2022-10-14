package fil

import (
	"bytes"
	"encoding/json"
	"github.com/fil_bs_crawler/internal/crawler/net"
	"github.com/fil_bs_crawler/internal/flogging"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/lotus/api"
	chaintypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"net/http"
)

type FilRPC struct {
	httprpc *net.HTTPRpc
}

func New(url string, options ...func(rpc *FilRPC)) *FilRPC {

	rpc := &FilRPC{
		httprpc: &net.HTTPRpc{
			url,
			http.DefaultClient,
			flogging.Log,
		},
	}

	for _, option := range options {
		option(rpc)
	}

	return rpc
}

func (e *FilRPC) getTipset(method string, params ...interface{}) (*chaintypes.ExpTipSet, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset chaintypes.ExpTipSet
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) getOriTipset(method string, params ...interface{}) (*chaintypes.TipSet, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset chaintypes.TipSet
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) FilGetOriTipsetByHeight(number int) (*chaintypes.TipSet, error) {
	return e.getOriTipset("Filecoin.ChainGetTipSetByHeight", number, nil)
}

func (e *FilRPC) getMessages(method string, params ...interface{}) (Messages, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return Messages{}, err
	}
	if bytes.Equal(result, []byte("null")) {
		return Messages{}, nil
	}

	var msgs Messages
	err = json.Unmarshal(result, &msgs)
	if err != nil {
		return Messages{}, err
	}

	return msgs, nil
}

func (e *FilRPC) FilGetTipsetByHeight(number int) (*chaintypes.ExpTipSet, error) {
	return e.getTipset("Filecoin.ChainGetTipSetByHeight", number, nil)
}

func (e *FilRPC) FilGetMessagesByCID(blockcid string) (Messages, error) {

	cid := make(map[string]string)
	cid["/"] = blockcid

	msgs, err := e.getMessages("Filecoin.ChainGetBlockMessages", cid)
	msgs.Blockcid = blockcid
	return msgs, err
}

func (e *FilRPC) getActor(method string, params ...interface{}) (*chaintypes.Actor, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset chaintypes.Actor
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) FilGetActor(to address.Address, ts chaintypes.TipSetKey) (*chaintypes.Actor, error) {

	msgs, err := e.getActor("Filecoin.StateGetActor", to, chaintypes.EmptyTSK)

	return msgs, err
}

func (e *FilRPC) getStateSearchMsg(method string, params ...interface{}) (*api.MsgLookup, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset api.MsgLookup
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) FilGetStateSearchMsg(number cid.Cid, ts chaintypes.TipSetKey) (*api.MsgLookup, error) {
	//return e.getStateSearchMsg("Filecoin.StateSearchMsg", ts, number, 1000, false)
	return e.getStateSearchMsg("Filecoin.StateSearchMsg", number)
}

func (e *FilRPC) getOriMessages(method string, params ...interface{}) (*chaintypes.Message, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return &chaintypes.Message{}, err
	}
	if bytes.Equal(result, []byte("null")) {
		return &chaintypes.Message{}, nil
	}

	var msgs chaintypes.Message
	err = json.Unmarshal(result, &msgs)
	if err != nil {
		return &chaintypes.Message{}, err
	}

	return &msgs, nil
}

func (e *FilRPC) getNetworkVersion(method string, params ...interface{}) (uint, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return 0, err
	}
	if bytes.Equal(result, []byte("null")) {
		return 0, nil
	}

	var msgs network.Version
	err = json.Unmarshal(result, &msgs)
	if err != nil {
		return 0, err
	}

	return uint(msgs), nil
}

func (e *FilRPC) FilGetMessages(blockcid cid.Cid) (*chaintypes.Message, error) {

	msgs, err := e.getOriMessages("Filecoin.ChainGetMessage", blockcid)
	return msgs, err
}

func (e *FilRPC) FilGetStateNetworkVersion() (uint, error) {

	msgs, err := e.getNetworkVersion("Filecoin.StateNetworkVersion", chaintypes.EmptyTSK)
	return msgs, err
}
