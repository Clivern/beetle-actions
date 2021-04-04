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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// event file path
var event string

var triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Parse and load event data",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		log.Info("trigger command got called")

		if !util.FileExists(event) {
			log.Error(fmt.Sprintf("Event file %s is missing", event))
			return
		}

		log.Info(fmt.Sprintf("Event file %s exists", event))

		content, err := util.ReadFile(event)

		if err != nil {
			log.Error(fmt.Sprintf("Error while reading event file %s: %s", event, err.Error()))
			return
		}

		log.Info(fmt.Sprintf("File %s content loaded", event))

		eventObj := model.CreateEvent{}

		err = eventObj.LoadFromJSON([]byte(content))

		if err != nil {
			log.Error(fmt.Sprintf("Error while decoding event content %s: %s", event, err.Error()))
			return
		}

		if eventObj.RefType != "tag" {
			log.Info(fmt.Sprintf("Skip event since it is not for a new tag %s", event))
			return
		}

		log.Info(fmt.Sprintf("A new tag got created: %s", eventObj.Ref))

		if os.Getenv("BEETLE_NUMERIC_TAG") == "true" {
			log.Info("Remove any letters from the tag")
			re, _ := regexp.Compile(`[a-zA-Z]`)
			eventObj.Ref = re.ReplaceAllString(eventObj.Ref, "")
			log.Info(fmt.Sprintf("The tag will became %s", eventObj.Ref))
		}

		if os.Getenv("BEETLE_WATCH_DEPLOYMENT") == "" {
			log.Info("Environment variable BEETLE_WATCH_DEPLOYMENT is missing")
		} else {
			log.Info("Environment variable BEETLE_WATCH_DEPLOYMENT is available")
		}

		if os.Getenv("BEETLE_API_URL") == "" {
			log.Info("Environment variable BEETLE_API_URL is missing")
		} else {
			log.Info("Environment variable BEETLE_API_URL is available")
		}

		if os.Getenv("BEETLE_API_KEY") == "" {
			log.Info("Environment variable BEETLE_API_KEY is missing")
		} else {
			log.Info("Environment variable BEETLE_API_KEY is available")
		}

		if os.Getenv("BEETLE_CLUSTER_NAME") == "" {
			log.Info("Environment variable BEETLE_CLUSTER_NAME is missing")
		} else {
			log.Info("Environment variable BEETLE_CLUSTER_NAME is available")
		}

		if os.Getenv("BEETLE_NAMESPACE_NAME") == "" {
			log.Info("Environment variable BEETLE_NAMESPACE_NAME is missing")
		} else {
			log.Info("Environment variable BEETLE_NAMESPACE_NAME is available")
		}

		if os.Getenv("BEETLE_APP_ID") == "" {
			log.Info("Environment variable BEETLE_APP_ID is missing")
		} else {
			log.Info("Environment variable BEETLE_APP_ID is available")
		}

		if os.Getenv("BEETLE_DEPLOYMENT_STRATEGY") == "" {
			log.Info("Environment variable BEETLE_DEPLOYMENT_STRATEGY is missing")
		} else {
			log.Info("Environment variable BEETLE_DEPLOYMENT_STRATEGY is available")
		}

		if os.Getenv("BEETLE_APP_VERSION") == "" {
			log.Info("Environment variable BEETLE_APP_VERSION is missing")
		} else {
			log.Info("Environment variable BEETLE_APP_VERSION is available")
		}

		if os.Getenv("BEETLE_MAX_SURGE") == "" {
			log.Info("Environment variable BEETLE_MAX_SURGE is missing")
		} else {
			log.Info("Environment variable BEETLE_MAX_SURGE is available")
		}

		if os.Getenv("BEETLE_MAX_UNAVAILABLE") == "" {
			log.Info("Environment variable BEETLE_MAX_UNAVAILABLE is missing")
		} else {
			log.Info("Environment variable BEETLE_MAX_UNAVAILABLE is available")
		}
	},
}

func init() {
	triggerCmd.Flags().StringVarP(
		&event,
		"event",
		"e",
		"event.json",
		"Absolute path to event file (required)",
	)
	triggerCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(triggerCmd)
}
