//go:generate go run main.go src
package main

import (
	"github.com/bytedance/sonic"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/cmd"
)

func init() {
	cache.SetTranscoderFunc(sonic.Marshal, sonic.Unmarshal)
}

func main() {
	cmd.Execute()
}
