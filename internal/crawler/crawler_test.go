/*
 * @Author: huangzhijie
 * @Description:
 * @File: crawler_test
 * @Version: 1.0.0
 * @Date: 2022/10/9 15:06
 */
package crawler

import (
	"github.com/fil_bs_crawler/internal/crawler/fil"
	"github.com/fil_bs_crawler/internal/flogging"
	"testing"
)

var log = flogging.Log

func Test_GetTipset(t *testing.T) {

	var dataChan chan fil.ChainDataResult
	var craw = NewFil(1374670, dataChan)
	result, err := craw.getChainData(1374715)
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", err)
	log.Infof("------- %+v", result.Height)
	for _, l := range result.MessagesList {
		log.Infof("------- %+v", l)
		log.Infof("------- Cids %+v", l.Cids)
		log.Infof("------- Blockcid %+v", l.Blockcid)
		log.Infof("------- BlsMessages %+v", l.BlsMessages)
		log.Infof("------- SecpkMessages %+v", l.SecpkMessages)
	}

}
