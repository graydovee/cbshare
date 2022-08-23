package cmd

import (
	"github.com/graydovee/clipboardshare/pkg/client"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli := client.NewClipboardClient()
		cli.Start(clientAddr)
	},
}

var (
	clientAddr string
)

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringVarP(&clientAddr, "addr", "a", "127.0.0", "serverAddr")
}
