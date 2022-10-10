package flogging

import (
	logging "github.com/ipfs/go-log/v2"
	"go.uber.org/zap"
)

var GinLogger *zap.Logger
var Logger *zap.SugaredLogger

var appVersion = "v1"

var Log = logging.Logger("crawler").With("app_version", appVersion)

func init() {
	_ = logging.SetLogLevel("*", "info")
}
