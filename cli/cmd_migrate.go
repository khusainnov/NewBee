package cli

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// TODO: handle migration by command
var cmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log := initLog()

		log.Info("init migrations")
		//migrate.New("")

	},
}

func initLog() *zap.Logger {
	log, _ := zap.NewDevelopment()

	return log
}
