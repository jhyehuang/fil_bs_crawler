package fil

import (
	"github.com/fil_bs_crawler/config"
	"github.com/fil_bs_crawler/internal/analyzer"
	"github.com/fil_bs_crawler/internal/flogging"
	chaintypes "github.com/filecoin-project/lotus/chain/types"
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

func Test_FilGetMessagesByCID(t *testing.T) {
	nodeAddress := config.ChainViper.GetString("Network.RpcNodeURL")
	crawler := New(nodeAddress)
	result, err := crawler.FilGetTipsetByHeight(1374715)
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", result.Cids)
	log.Infof("-------Height: %d", result.Height)

	for _, tl := range result.Blocks {
		l := tl
		msges, err := crawler.FilGetMessagesByCID(l.Cid().String())
		if err != nil {
			log.Errorf(err.Error())

		}
		log.Infof("-------block : %+v", l.Cid())
		log.Infof("-------BlsMessages : %+v", len(msges.BlsMessages))
		log.Infof("-------SecpkMessages: %+v", len(msges.SecpkMessages))
		if len(msges.BlsMessages) > 0 {
			for _, tmsg := range msges.BlsMessages {
				msg := tmsg
				log.Infof("-------msg.Cid: %+v", msg.Cid())
				log.Infof("-------msg.Method: %+v", msg.Method)
				ret, err := crawler.FilGetActor(msg.To, chaintypes.EmptyTSK)
				if err != nil {
					log.Errorf(err.Error())
				}

				log.Infof("-------: %+v", ret.Code)

				as, err := analyzer.ParseParamsV1(ret.Code, int(msg.Method), msg.Params)
				if err != nil {
					t.Fatal(err)
				}
				log.Infof("-------: %+v", as)
			}

		}

		if len(msges.SecpkMessages) > 0 {
			for _, tmsg := range msges.SecpkMessages {
				msg := tmsg
				log.Infof("-------msg.Cid: %+v", msg.Cid())
				log.Infof("-------msg.Method: %+v", msg.Message.Method)
				ret, err := crawler.FilGetActor(msg.Message.To, chaintypes.EmptyTSK)
				if err != nil {
					log.Errorf(err.Error())
				}

				log.Infof("-------: %+v", ret.Code)

				as, err := analyzer.ParseParamsV1(ret.Code, int(msg.Message.Method), msg.Message.Params)
				if err != nil {
					t.Fatal(err)
				}
				log.Infof("-------: %+v", as)

			}

		}

	}
}

func Test_FilGetMessagesByCID2(t *testing.T) {
	nodeAddress := config.ChainViper.GetString("Network.RpcNodeURL")
	crawler := New(nodeAddress)
	result, err := crawler.FilGetOriTipsetByHeight(1374715)

	if err != nil {
		log.Errorf(err.Error())

	}

	log.Infof("------- %+v", result.Cids)
	log.Infof("-------Height: %d", result.Height)

	for _, tl := range result.Blocks() {
		l := tl
		msges, err := crawler.FilGetMessagesByCID(l.Cid().String())
		if err != nil {
			log.Errorf(err.Error())

		}
		log.Infof("-------block : %+v", l.Cid())
		log.Infof("-------BlsMessages : %+v", len(msges.BlsMessages))
		log.Infof("-------SecpkMessages: %+v", len(msges.SecpkMessages))
		if len(msges.BlsMessages) > 0 {
			for _, tmsg := range msges.BlsMessages {
				msg := tmsg
				log.Infof("-------msg.Cid: %+v", msg.Cid())
				log.Infof("-------msg.Method: %+v", msg.Method)
				ret, err := crawler.FilGetActor(msg.To, chaintypes.EmptyTSK)
				if err != nil {
					log.Errorf(err.Error())
				}

				log.Infof("-------: %+v", ret.Code)

				as, err := analyzer.ParseParamsV1(ret.Code, int(msg.Method), msg.Params)
				if err != nil {
					t.Fatal(err)
				}
				log.Infof("-------: %+v", as)
			}

		}

		if len(msges.SecpkMessages) > 0 {
			for _, tmsg := range msges.SecpkMessages {
				msg := tmsg
				log.Infof("-------msg.Cid: %+v", msg.Cid())
				log.Infof("-------msg.Method: %+v", msg.Message.Method)
				ret, err := crawler.FilGetActor(msg.Message.To, chaintypes.EmptyTSK)
				if err != nil {
					log.Errorf(err.Error())
				}

				log.Infof("-------: %+v", ret.Code)

				as, err := analyzer.ParseParamsV1(ret.Code, int(msg.Message.Method), msg.Message.Params)
				if err != nil {
					t.Fatal(err)
				}
				log.Infof("-------: %+v", as)

			}

		}

	}
}

func Test_FilGetSend(t *testing.T) {
	nodeAddress := config.ChainViper.GetString("Network.RpcNodeURL")
	crawler := New(nodeAddress)
	result, err := crawler.FilGetOriTipsetByHeight(1381788) //1380526
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", result.Cids)
	log.Infof("-------Height: %d", result.Height)

	for _, tl := range result.Blocks() {
		l := tl
		msges, err := crawler.FilGetMessagesByCID(l.Cid().String())
		if err != nil {
			log.Errorf(err.Error())

		}
		log.Infof("-------block : %+v", l.Cid())
		log.Infof("-------BlsMessages : %+v", len(msges.BlsMessages))
		log.Infof("-------SecpkMessages: %+v", len(msges.SecpkMessages))
		if len(msges.BlsMessages) > 0 {
			for _, tmsg := range msges.BlsMessages {
				msg := tmsg
				log.Infof("-------msg.Cid: %+v", msg.Cid())
				log.Infof("-------msg.Method: %+v", msg.Method)
				log.Infof("-------msg: %+v", msg)
				ret, err := crawler.FilGetActor(msg.To, chaintypes.EmptyTSK)
				if err != nil {
					log.Errorf(err.Error())
				}

				log.Infof("-------: %+v", ret.Code)

				as, err := analyzer.ParseParamsV1(ret.Code, int(msg.Method), msg.Params)
				if err != nil {
					t.Fatal(err)
				}
				log.Infof("-------: %+v", as)
			}

		}

		if len(msges.SecpkMessages) > 0 {
			for _, tmsg := range msges.SecpkMessages {
				msg := tmsg
				log.Infof("-------msg.Cid: %+v", msg.Cid())
				log.Infof("-------msg.Method: %+v", msg.Message.Method)
				log.Infof("-------msg: %+v", msg.Message)

				ret, err := crawler.FilGetActor(msg.Message.To, chaintypes.EmptyTSK)
				if err != nil {
					log.Errorf(err.Error())
				}

				log.Infof("-------: %+v", ret.Code)

				as, err := analyzer.ParseParamsV1(ret.Code, int(msg.Message.Method), msg.Message.Params)
				if err != nil {
					log.Error(err)
					continue
				}
				log.Infof("-------: %+v", as)

			}

		}

	}
}

func Test_FilGetActor(t *testing.T) {
	nodeAddress := config.ChainViper.GetString("Network.RpcNodeURL")
	crawler := New(nodeAddress)
	result, err := crawler.FilGetOriTipsetByHeight(1381788) //1380526
	if err != nil {
		log.Errorf(err.Error())

	}
	for _, tl := range result.Blocks() {
		l := tl
		msges, err := crawler.FilGetMessagesByCID(l.Cid().String())
		if err != nil {
			log.Errorf(err.Error())

		}

		log.Infof("-------block : %+v", l.Cid())
		log.Infof("-------BlsMessages : %+v", len(msges.BlsMessages))
		log.Infof("-------SecpkMessages: %+v", len(msges.SecpkMessages))

		if len(msges.BlsMessages) > 0 {
			for _, tmsg := range msges.BlsMessages {
				msg := tmsg
				log.Infof("-------msg.Cid: %+v", msg.Cid())
				log.Infof("-------msg.Method: %+v", msg.Method)

				ret, err := crawler.FilGetActor(msg.To, chaintypes.EmptyTSK)
				if err != nil {
					log.Errorf(err.Error())
				}

				log.Infof("-------: %+v", ret.Code)

				rets, err := analyzer.ParseParamsV2(ret.Code, msg.Method, msg.Params)
				if err != nil {
					log.Errorf(err.Error())
				}
				log.Infof("-------params: %+v", rets)
				s, _ := msg.MarshalJSON()
				log.Infof("-------params: %+v", string(s))

			}

		}
	}
	//for key, value := range filcns.NewActorRegistry().Methods {
	//	fmt.Println(key, "------>", value)
	//}

}

func Test_FilGetStateNetworkVersion(t *testing.T) {
	nodeAddress := config.ChainViper.GetString("Network.RpcNodeURL")
	crawler := New(nodeAddress)
	result, err := crawler.FilGetStateNetworkVersion() //1380526
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("-------nodeAddress: %+v", nodeAddress)
	log.Infof("-------result: %+v", result)
}
