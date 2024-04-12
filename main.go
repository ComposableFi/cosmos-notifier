package main

import (
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const defaultInterval time.Duration = time.Minute

func setLogLevel() {
	level, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		log.Info("LOG_LEVEL not set, defaulting to warning")
		log.SetLevel(log.WarnLevel)
		return
	}

	parsedLevel, err := log.ParseLevel(strings.ToLower(level))
	if err != nil {
		log.Warnf("Invalid LOG_LEVEL '%s', defaulting to warning", level)
		log.SetLevel(log.WarnLevel)
		return
	}

	log.SetLevel(parsedLevel)
	log.Infof("Logging level set to %s", parsedLevel)
}

func init() {

	log.SetOutput(os.Stdout)
	setLogLevel()
}

func main() {

	requiredVars := []string{"COSMOS_ENDPOINT", "SLACK_WEBHOOK_URL", "CHAIN_ID"}
	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			log.Panicf("Error: %s is not defined\n", varName)
		}
	}

	cosmosEndpoint := os.Getenv("COSMOS_ENDPOINT")
	slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	chainID := os.Getenv("CHAIN_ID")

	interval, err := time.ParseDuration(os.Getenv("FETCH_INTERVAL"))
	if err != nil {
		log.Errorf("Error parsing FETCH_INTERVAL: %s, defaulting to 1m\n", err)
		interval = defaultInterval
	}

	var lastProposalId string
	for {
		latestProposal, err := getLatestProposal(cosmosEndpoint)
		if err != nil {
			log.Errorf("Error fetching proposal: %s\n", err)
			continue
		}

		// post to slack only if didn't posted before
		if lastProposalId != latestProposal.ID {
			err = postToSlack(chainID, *latestProposal, slackWebhookURL)
			if err != nil {
				log.Errorf("Error posting to slack: %s\n", err)
				continue
			}
			lastProposalId = latestProposal.ID
		}

		time.Sleep(interval)
	}
}
