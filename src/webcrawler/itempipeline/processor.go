package itemproc

import (
	base "github.com/Jsoneft/goc2p/src/webcrawler/base"
)

// 被用来处理条目的函数类型。
type ProcessItem func(item base.Item) (result base.Item, err error)
