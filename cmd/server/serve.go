package server

import (
	"github.com/spf13/cobra"
	"github.com/xiusin/pine/di"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动pinecms服务器",
}

func init()  {
	// 标识非serve模式
	di.Set("pinecms.serve.mode", func(builder di.AbstractBuilder) (interface{}, error) {
		return true, nil
	}, false)
}
