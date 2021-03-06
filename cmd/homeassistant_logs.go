package cmd

import (
	"fmt"

	helper "github.com/home-assistant/hassio-cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var homeassistantLogsCmd = &cobra.Command{
	Use:     "logs",
	Aliases: []string{"log", "lg"},
	Short:   "View the log output of Home Assistant",
	Long: `
Allowing you to look at the log output generated by the Home Assistant running
on your Hass.io system.`,
	Example: `
  hassio homeassistant logs`,
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("homeassistant logs")

		section := "homeassistant"
		command := "logs"
		base := viper.GetString("endpoint")

		url, err := helper.URLHelper(base, section, command)
		if err != nil {
			fmt.Printf("Error: %v", err)
			ExitWithError = true
			return
		}

		request := helper.GetRequest()
		resp, err := request.SetHeader("Accept", "text/plain").Get(url)

		if err != nil {
			fmt.Println(err)
			ExitWithError = true
		} else {
			fmt.Println(resp.String())
		}
		return
	},
}

func init() {
	homeassistantCmd.AddCommand(homeassistantLogsCmd)
}
