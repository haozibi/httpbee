package cmd

import (
	"fmt"
	"log"

	"github.com/haozibi/httpbee/app"

	"github.com/spf13/cobra"
)

// Execute Execute
func Execute() {
	root := newRootCommand()
	root.AddCommand(newVersionCommand())

	if err := root.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func newRootCommand() *cobra.Command {

	var opt app.Config

	var cmd = &cobra.Command{
		Use:           "httpbee",
		Short:         "Fake HTTP Server",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opt.RespPath == "" {
				return cmd.Help()
			}

			s, err := app.NewFakeServer(opt)
			if err != nil {
				return err
			}
			return s.RunHTTP()
		},
	}

	flag := cmd.Flags()
	flag.IntVarP(&opt.Port, "port", "", 8080, "fake http port")
	flag.StringVarP(&opt.RespPath, "resp-file", "f", "", "response file")

	return cmd
}

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s \ntag: %s\nbuild: %s\nhash: %s\n",
				app.BuildAppName,
				app.BuildVersion,
				app.BuildTime,
				app.CommitHash,
			)
		},
	}
	return cmd
}
