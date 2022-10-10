package fil

import (
	"bytes"
	"encoding/json"
	"github.com/fil_bs_crawler/config"
	"github.com/fil_bs_crawler/internal/flogging"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/consensus/filcns"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/ipfs/go-cid"
	"testing"
)

var log = flogging.Log

func Test_FilGetTipsetByHeight(t *testing.T) {
	nodeAddress := config.ChainViper.GetString("Network.RpcNodeURL")
	crawler := New(nodeAddress)
	result, err := crawler.FilGetTipsetByHeight(1374715)
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", result.Cids)
	log.Infof("-------Height: %d", result.Height)
	for _, l := range result.Blocks {
		log.Infof("------- %+v", l.ForkSignaling)
		log.Infof("------- %+v", l.BLSAggregate.Type)
		log.Infof("------- %+v", l.BLSAggregate.Data)
		log.Infof("------- %s", string(l.BLSAggregate.Data))
		log.Infof("------- %s", l.Miner.String())
		for _, b := range l.WinPoStProof {
			log.Infof("------- %x", b.PoStProof) // f033716 2329038 f036004 2329039 2329040
			str2 := string(b.ProofBytes[:])
			log.Infof("------- %s", str2)
		}

		//log.Infof("------- %X", l.ElectionProof.VRFProof)
		//log.Infof("------- Messages %+v", l.Messages)

	}
}

func JsonParams(code cid.Cid, method abi.MethodNum, params []byte) (string, error) {
	p, err := stmgr.GetParamType(filcns.NewActorRegistry(), code, method) // todo use api for correct actor registry
	if err != nil {
		return "", err
	}

	if err := p.UnmarshalCBOR(bytes.NewReader(params)); err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(p, "", "  ")
	return string(b), err
}

func parseParams(cid cid.Cid, method abi.MethodNum, params []byte) error {

	pstr, err := JsonParams(cid, abi.MethodNum(method), params)
	if err != nil {
		return err
	}

	log.Infof("precommit info => %s", string(pstr))
	return nil

}

func Test_FilGetMessagesByCID(t *testing.T) {
	nodeAddress := config.ChainViper.GetString("Network.RpcNodeURL")
	crawler := New(nodeAddress)
	result, err := crawler.FilGetTipsetByHeight(1374715)
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", result.Cids)
	log.Infof("-------Height: %d", result.Height)

	for _, l := range result.Blocks {
		msges, err := crawler.FilGetMessagesByCID(l.Cid().String())
		if err != nil {
			log.Errorf(err.Error())

		}
		log.Infof("-------: %+v", l.Messages)
		log.Infof("-------: %+v", msges.SecpkMessages)
		log.Infof("-------: %+v", msges.SecpkMessages)
		if len(msges.BlsMessages) > 0 {
			for _, msg := range msges.BlsMessages {
				log.Infof("-------: %+v", msg.Cid())
				//log.Infof("-------: %s", msg.Params)
				parseParams(msg.Cid(), msg.Method, msg.VMMessage().Params)
			}

		}

		if len(msges.SecpkMessages) > 0 {
			for _, msg := range msges.SecpkMessages {
				parseParams(msg.Cid(), msg.VMMessage().Method, msg.VMMessage().Params)
				//log.Infof("-------: %s", string(msgj))
			}

		}

	}
}
