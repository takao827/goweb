package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	fileName *string
	export   *bool
	cmd      = cobra.Command{
		Use: "exportenv",
		Example: `  exportenv
  eval $(exportenv -f .env)
  eval $(exportenv -f .env -e)`,
		Run: func(cmd *cobra.Command, args []string) {
			var env map[string]string
			env, err := godotenv.Read(*fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				os.Exit(1)
			}

			keys := []string{}
			for key := range env {
				keys = append(keys, key)
			}
			slices.Sort(keys)

			for _, key := range keys {
				if *export {
					fmt.Printf("export %s=\"%s\"\n", key, env[key])
				} else {
					fmt.Printf("%s=\"%s\"\n", key, env[key])
				}
			}
		},
	}
)

func init() {
	fileName = cmd.Flags().StringP("file", "f", ".env", "path to dotenv file")
	export = cmd.Flags().BoolP("export", "e", false, "print as export")
}

func main() {
	cmd.Execute()
}