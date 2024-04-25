package main

import (
	"github.com/ashwanthkumar/slack-go-webhook"
)

const proposalPrefixTestnetUrl = "https://explorer.stavr.tech/Composable-Testnet/gov/"
const proposalPrefixMainnetUrl = "https://explorer.stavr.tech/Composable-Mainnet/gov/"

func postToSlack(chainId string, proposal Proposal, slackWebookUrl string) error {

	attachment := slack.Attachment{}
	attachment.AddField(slack.Field{Title: "Chain:", Value: chainId})
	attachment.AddField(slack.Field{Title: "ID", Value: proposal.ID})
	attachment.AddField(slack.Field{Title: "Title", Value: proposal.Title})
	attachment.AddField(slack.Field{Title: "Summary", Value: proposal.Summary})
	attachment.AddField(slack.Field{Title: "Proposer", Value: proposal.Proposer})
	attachment.AddField(slack.Field{Title: "Status", Value: proposal.Status})
	attachment.AddField(slack.Field{Title: "VotingStartTime", Value: proposal.VotingStartTime})
	attachment.AddField(slack.Field{Title: "VotingEndTime", Value: proposal.VotingEndTime})

	switch chainId {
	case "centauri-1":
		attachment.AddAction(slack.Action{Type: "button", Text: "Open", Url: proposalPrefixMainnetUrl + proposal.ID, Style: "primary"})
	case "banksy-testnet-5":
		attachment.AddAction(slack.Action{Type: "button", Text: "Open", Url: proposalPrefixTestnetUrl + proposal.ID, Style: "primary"})
	}

	payload := slack.Payload{
		Text:        "Found new governance proposal",
		Username:    "robot",
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(slackWebookUrl, "", payload)
	if len(err) > 0 {
		return err[0]
	}
	return nil
}
