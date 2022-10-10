package fil

import (
	"bytes"
	"encoding/json"
	"github.com/fil_bs_crawler/internal/crawler/net"
	"github.com/fil_bs_crawler/internal/flogging"
	chaintypes "github.com/filecoin-project/lotus/chain/types"
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
