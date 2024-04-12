package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func fetchProposalsData(cosmosEndpoint string) (*ProposalsData, error) {

	url := cosmosEndpoint
	url += "/cosmos/gov/v1/proposals"
	url += "?pagination.limit=1&pagination.count_total=true&pagination.reverse=true"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp)

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

func getLatestProposal(cosmosEndpoint string) (*Proposal, error) {
	proposalsData, err := fetchProposalsData(cosmosEndpoint)
	if err != nil {
		return nil, err
	}
	return &proposalsData.Proposals[0], nil
}
