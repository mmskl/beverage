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
	_ "os"

	"github.com/spf13/cobra"
    "github.com/mmskl/beverage/internal/DB"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List entries",
	Long: `list entries:

beverage list`,
	Run: func(cmd *cobra.Command, args []string) {
        db := DB.Connect()
        entries := db.ListEntries()


    for num, e := range entries {
        num++
        fmt.Printf("%d. %s - %s\n", num, e.Name, e.Url)
    }


	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
