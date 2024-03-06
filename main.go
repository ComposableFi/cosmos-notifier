package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Define struct to match JSON structure
type ProposalsData struct {
	Proposals  []Proposal `json:"proposals"`
	Pagination Pagination `json:"pagination"`
}

type Proposal struct {
	ID               string      `json:"id"`
	Messages         interface{} `json:"messages"`
	Status           string      `json:"status"`
	FinalTallyResult interface{} `json:"final_tally_result"`
	SubmitTime       string      `json:"submit_time"`
	DepositEndTime   string      `json:"deposit_end_time"`
	TotalDeposit     interface{} `json:"total_deposit"`
	VotingStartTime  string      `json:"voting_start_time"`
	VotingEndTime    string      `json:"voting_end_time"`
	Metadata         string      `json:"metadata"`
	Title            string      `json:"title"`
	Summary          string      `json:"summary"`
	Proposer         string      `json:"proposer"`
}

type Pagination struct {
	NextKey *json.RawMessage `json:"next_key"`
	Total   string           `json:"total"`
}

func fetchProposals(cosmosEndpoint string) (*ProposalsData, error) {
	resp, err := http.Get(cosmosEndpoint + "/cosmos/gov/v1/proposals")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var proposals ProposalsData
	err = json.Unmarshal(body, &proposals)
	if err != nil {
		return nil, err
	}

	return &proposals, nil
}

func postToSlack(slackWebhookURL, message string) error {
	// payload := fmt.Sprintf(`{"text":"%s"}`, message)

	// req, err := http.NewRequest(http.MethodPost, slackWebhookURL, bytes.NewBuffer([]byte(payload)))
	// if err != nil {
	// 	return err
	// }
	// req.Header.Add("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	return nil
}

func main() {

	requiredVars := []string{"COSMOS_ENDPOINT", "SLACK_WEBHOOK_URL"}

	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			fmt.Printf("Error: %s is not defined\n", varName)
			os.Exit(1) // Exit with a status code of 1
		}
	}

	cosmosEndpoint := os.Getenv("COSMOS_ENDPOINT")
	slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	interval, err := time.ParseDuration(os.Getenv("FETCH_INTERVAL"))
	if err != nil {
		fmt.Printf("Error parsing FETCH_INTERVAL: %s\n", err)
		return
	}

	seenProposals := make(map[string]bool)

	runsCounter := 0
	for {
		proposals, err := fetchProposals(cosmosEndpoint)
		if err != nil {
			fmt.Printf("Error fetching proposals: %s\n", err)
			continue
		}

		for _, proposal := range proposals.Proposals {
			if _, seen := seenProposals[proposal.ID]; !seen {
				message := fmt.Sprintf("New Governance Proposal: ID %s, Title: %s", proposal.ID, proposal.Title)

				// do not post when looping for 1st time
				var err error
				if runsCounter > 0 {
					err = postToSlack(slackWebhookURL, message)
				}
				if err != nil {
					fmt.Printf("Error posting to Slack: %s\n", err)
				} else {
					seenProposals[proposal.ID] = true
				}

			}
		}

		time.Sleep(interval)
		runsCounter++
	}
}
