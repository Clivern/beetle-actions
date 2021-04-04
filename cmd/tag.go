// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/clivern/beetle-actions/core/model"
	"github.com/clivern/beetle-actions/core/util"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Get the tag",
	Run: func(cmd *cobra.Command, args []string) {
		if !util.FileExists(event) {
			return
		}

		content, err := util.ReadFile(event)

		if err != nil {
			return
		}

		eventObj := model.CreateEvent{}

		err = eventObj.LoadFromJSON([]byte(content))

		if err != nil {
			return
		}

		if eventObj.RefType != "tag" {
			return
		}

		if os.Getenv("BEETLE_NUMERIC_TAG") == "true" {
			re, _ := regexp.Compile(`[a-zA-Z]`)
			eventObj.Ref = re.ReplaceAllString(eventObj.Ref, "")
			fmt.Printf("%s", eventObj.Ref)
		} else {
			fmt.Printf("%s", eventObj.Ref)
		}
	},
}

func init() {
	tagCmd.Flags().StringVarP(
		&event,
		"event",
		"e",
		"event.json",
		"Absolute path to event file (required)",
	)
	tagCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(tagCmd)
}
