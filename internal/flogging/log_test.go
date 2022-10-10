/*
 * @Author: huangzhijie
 * @Description:
 * @File: crawler_test
 * @Version: 1.0.0
 * @Date: 2022/10/9 15:06
 */
package flogging

import (
	"testing"
)

func Test_Log(t *testing.T) {

	Log.Info("###### info ######")
	Log.Warnf("###### info ######")
	Log.Error("###### info ######")
	Log.Infow("###### info ######")
	Log.Info("###### info ######")

}
