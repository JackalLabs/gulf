package cmd

import (
	"fmt"
	"github.com/JackalLabs/gulf/core"
	"github.com/JackalLabs/gulf/jackal/uploader"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/spf13/cobra"
	"os"
)

func RootCMD(q *uploader.Queue, w *wallet.Wallet) *cobra.Command {
	c := &cobra.Command{
		Use:   "gulf",
		Short: "Gulf is a command line application for posting streamable videos to Jackal and hosting them on IPFS.",
		Long:  `Gulf is a command line application for posting streamable videos to Jackal and hosting them on IPFS. These files are organized in folders as they appear on disk using IPFS folder nodes saved as virtual files on Jackal.`,
	}

	c.AddCommand(LaunchCMD(q, w))

	return c

}

func LaunchCMD(q *uploader.Queue, w *wallet.Wallet) *cobra.Command {
	c := &cobra.Command{
		Use:   "post [file]",
		Short: "Starts the upload process.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			file := args[0]

			o, err := cmd.Flags().GetString("out")
			if err != nil {
				return err
			}

			err = core.Convert(file, o)
			if err != nil {
				return err
			}

			cid, _, err := core.PostDir(o, q, w)
			if err != nil {
				return err
			}

			_ = os.RemoveAll(o)

			fmt.Printf("Lift Off! 🚀\n\nYou can now view your files at ipfs://%s\n", cid)

			return nil
		},
	}

	c.Flags().StringP("out", "o", "out", "sets the output directory")

	return c
}
