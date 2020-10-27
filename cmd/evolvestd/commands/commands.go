package commands

import (
	"github.com/EdgarTeng/evolvest/embed/rpc"
	"github.com/EdgarTeng/evolvest/pkg/cmdroot"
	"github.com/EdgarTeng/evolvest/pkg/common/config"
	"github.com/spf13/cobra"
	"log"
)

// Execute adds all child commands to the root command
func Execute() {
	cmdroot.InitCommand(
		"evolvestd",
		`evolvest service`,
		cmdroot.WithReport(), cmdroot.WithMonitor())
	cmdroot.AddCommand(newServer())
	cmdroot.Execute()
}

func newServer() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start an evolvestd ",
		RunE:  startServer,
	}

	return serverCmd
}

func initBeforeStart() {

	// init config
	initConfig()

	initGrpc()

	return
}

func initGrpc() {
	go rpc.StartServer(":" + config.Config().ServerPort)

}

func initConfig() {
	// load config
	if err := config.InitConfig(cmdroot.CmdConfig); err != nil {
		log.Fatalf("init config failed, %v\n", err)
	}
}

func startServer(cmd *cobra.Command, args []string) error {
	initBeforeStart()

	cmd.Println("Server starting ...")

	cmdroot.WaitSignal()
	cmd.Println("Server stopping ...")

	return nil
}
