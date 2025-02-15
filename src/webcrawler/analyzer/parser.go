package analyzer

import (
	"net/http"

	base "github.com/Jsoneft/goc2p/src/webcrawler/base"
)

// 被用于解析HTTP响应的函数类型。
type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)
