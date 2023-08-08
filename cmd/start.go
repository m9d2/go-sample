package cmd

import (
	"sample/config"
	"sample/database"
	"sample/database/migrate"
	"sample/router"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	port = 8080

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  "Start the running web server gracefully.",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func start() {
	config.InitConfig()
	database.InitDataBase()
	migrate.AutoMigrate()
	r := router.InitRouter()
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		return
	}
}

func init() {
	RootCmd.AddCommand(startCmd)
}
