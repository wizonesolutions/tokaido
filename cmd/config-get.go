package cmd

import (
	"reflect"

	"github.com/ironstar-io/tokaido/conf"

	"fmt"
	"log"
	"regexp"

	"github.com/sanity-io/litter"
	"github.com/spf13/cobra"
)

// ConfigGetCmd - `tok config-get`
var ConfigGetCmd = &cobra.Command{
	Use:   "config-get",
	Short: "Get a config property value",
	Long:  "Get a config property value. Eg. `tok config-get drupal path`. See https://tokaido.io/docs/config for a full list of available options",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := conf.GetConfigValueByArgs(args)
		if err != nil {
			log.Fatal(err)
		}

		if c.Kind() == reflect.Struct {
			d := litter.Sdump(c.Interface())
			s := regexp.MustCompile(`(?s)struct {.*?}{`).ReplaceAllString(d, "{")

			fmt.Println(s)
			return
		}

		fmt.Printf("%+v\n", c)
	},
}
