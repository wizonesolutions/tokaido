package cmd

import (
	"github.com/ironstar-io/tokaido/conf"
	"github.com/ironstar-io/tokaido/initialize"
	"github.com/ironstar-io/tokaido/services/docker"
	"github.com/ironstar-io/tokaido/services/telemetry"
	"github.com/ironstar-io/tokaido/services/tok"
	"github.com/ironstar-io/tokaido/utils"
	"github.com/spf13/cobra"
)

// DestroyCmd - `tok destroy`
var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Stop and destroy all containers",
	Long:  "Destroy your Tokaido environment.",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.TokConfig("destroy")
		utils.CheckCmdHard("docker-compose")
		telemetry.SendCommand("destroy")

		docker.HardCheckTokCompose()

		tok.Destroy(conf.GetConfig().Tokaido.Project.Name)
	},
}
