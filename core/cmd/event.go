// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Parse and load event data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`BAM`)
	},
}

func init() {
	rootCmd.AddCommand(triggerCmd)
}
