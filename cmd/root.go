package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// http://www.network-science.de/ascii/ Font: stop
var rootCmd = &cobra.Command{
	Use: "pinecms",
	Long: `
       _                             
      (_)                            
 ____  _ ____   ____ ____ ____   ___ 
|  _ \| |  _ \ / _  ) ___)    \ /___)
| | | | | | | ( (/ ( (___| | | |___ |
| ||_/|_|_| |_|\____)____)_|_|_(___/ 
|_|     		      version: ` + Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
