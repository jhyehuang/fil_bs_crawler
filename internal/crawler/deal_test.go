/*
 * @Author: huangzhijie
 * @Description:
 * @File: crawler_test
 * @Version: 1.0.0
 * @Date: 2022/10/9 15:06
 */
package crawler

import (
	"context"
	"github.com/fil_bs_crawler/internal/crawler/fil"
	"github.com/ipfs/go-cid"
	"testing"
)

func Test_GetDeal(t *testing.T) {

	var dataChan chan fil.ChainDataResult
	var craw = NewFil(1374670, dataChan)
	pubCids := "bafy2bzacedgbytoxb4w3dcfnzj6kk23j7clbsvxbjmhcnkegezax5fvbtqu36"
	propCids := "bafyreihqxrymnobhiz2cv6hgwa2hu3iskj5jqte3m5akycdregkkqpnypm"
	pubCid, err := cid.Decode(pubCids)
	if err != nil {
		log.Errorf(err.Error())

	}
	propCid, err := cid.Decode(propCids)
	if err != nil {
		log.Errorf(err.Error())

	}
	result, err := craw.GetDealID(context.TODO(), pubCid, propCid)
	if err != nil {
		log.Errorf(err.Error())

	}
	t.Logf("%+v", result)

}
