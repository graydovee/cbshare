package cmd

import (
	"fmt"
	"github.com/graydovee/clipboardshare/pkg/server"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewClipboardServer()
		server.Start(fmt.Sprintf("%s:%d", serverAddr, serverPort))
	},
}

var (
	serverPort uint16
	serverAddr string
)

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Uint16VarP(&serverPort, "port", "p", 29933, "server listen port")
	serverCmd.Flags().StringVarP(&serverAddr, "addr", "a", "", "server listen address")
}
