package loadgen

import (
	"github.com/Jsoneft/goc2p/src/logging"
)

var logger logging.Logger

func init() {
	logger = logging.NewSimpleLogger()
}
