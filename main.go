package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	requiredVars := []string{"COSMOS_ENDPOINT", "SLACK_WEBHOOK_URL", "CHAIN_ID"}
	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			fmt.Printf("Error: %s is not defined\n", varName)
			os.Exit(1)
		}
	}

	cosmosEndpoint := os.Getenv("COSMOS_ENDPOINT")
	slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	chainID := os.Getenv("CHAIN_ID")

	interval, err := time.ParseDuration(os.Getenv("FETCH_INTERVAL"))
	if err != nil {
		fmt.Printf("Error parsing FETCH_INTERVAL: %s\n", err)
		return
	}

	runsCounter := 0
	lastPostedProposalId := ""
	for {
		latestProposal, err := getLatestProposal(cosmosEndpoint)
		if err != nil {
			fmt.Printf("Error fetching proposal: %s\n", err)
			continue
		}
		// post to slack only if
		// not posted before
		// and this is at least second iteration
		if runsCounter > 0 && latestProposal.ID == lastPostedProposalId {
			err = postToSlack(chainID, *latestProposal, slackWebhookURL)
			if err != nil {
				fmt.Printf("Error posting to slack: %s\n", err)
				continue
			}
			lastPostedProposalId = latestProposal.ID
		}

		time.Sleep(interval)
		runsCounter++
	}
}
