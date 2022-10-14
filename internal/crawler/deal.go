/*
 * @Author: huangzhijie
 * @Description:
 * @File: deal
 * @Version: 1.0.0
 * @Date: 2022/10/12 15:00
 */
package crawler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fil_bs_crawler/internal/flogging"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

func (c *CrawlerFil) GetDealID(ctx context.Context, pubcid cid.Cid, PropCid cid.Cid) (abi.DealID, error) {
	mlookup, err := c.filRpc.FilGetStateSearchMsg(pubcid, types.EmptyTSK)
	if err != nil {
		return 0, xerrors.Errorf("could not find published deal on chain: %w", err)
	}

	if mlookup.Message != pubcid {
		// TODO: can probably deal with this by checking the message contents?
		return 0, xerrors.Errorf("publish deal message was replaced on chain")
	}

	msg, err := c.filRpc.FilGetMessages(mlookup.Message)
	if err != nil {
		return 0, err
	}

	//ret, err := c.filRpc.FilGetActor(msg.To, types.EmptyTSK)
	//if err != nil {
	//	fmt.Errorf("%s", err.Error())
	//	return 0, err
	//}

	//rets, err := analyzer.ParseParamsV2(ret.Code, msg.Method, msg.Params)
	//if err != nil {
	//
	//	fmt.Errorf("%s", err.Error())
	//	return 0, err
	//}

	var params market.PublishStorageDealsParams
	if err := params.UnmarshalCBOR(bytes.NewReader(msg.Params)); err != nil {
		return 0, err
	}
	flogging.Log.Warnf("%+v", params)

	dealix := -1
	for i, pd := range params.Deals {
		nd, err := cborutil.AsIpld(&pd)
		if err != nil {
			return 0, xerrors.Errorf("failed to compute deal proposal ipld node: %w", err)
		}

		flogging.Log.Warnf("nd %s PropCid %s", nd.Cid().String(), PropCid.String())
		if nd.Cid() == PropCid {
			dealix = i
			break
		}
	}

	if dealix == -1 {
		return 0, fmt.Errorf("our deal was not in this publish message")
	}

	if mlookup.Receipt.ExitCode != 0 {
		return 0, xerrors.Errorf("miners deal publish failed (exit: %d)", mlookup.Receipt.ExitCode)
	}

	var retval market.PublishStorageDealsReturn
	if err := retval.UnmarshalCBOR(bytes.NewReader(mlookup.Receipt.Return)); err != nil {
		return 0, xerrors.Errorf("publish deal return was improperly formatted: %w", err)
	}

	if len(retval.IDs) != len(params.Deals) {
		return 0, fmt.Errorf("return value from publish deals did not match length of params")
	}
	fmt.Println(retval)
	return retval.IDs[dealix], nil
}
