/*
Copyright © 2020 Marcin Moskal

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
    "fmt"
	"github.com/spf13/cobra"
    "github.com/mmskl/beverage/internal/DB"
)

// upgradeSchemaCmd represents the upgradeSchema command
var upgradeSchemaCmd = &cobra.Command{
	Use:   "upgrade-schema",
    Aliases: []string{"upgrade-schema", "upgradeSchema"},
	Short: "Upgrade schema DB according to models",
	Long: `Due to application updates structure of models may change.
    To keep it in sync with database run this command.

beverage upgrade-schema - to show what changes will be made
beverage upgrade-schema --force - to perform upgrade
`,
	Run: func(cmd *cobra.Command, args []string) {
        DB.DoUpgrade()
        fmt.Printf("Schema upgraded")
	},
}


func init() {
	rootCmd.AddCommand(upgradeSchemaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeSchemaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeSchemaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
