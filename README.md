# cosmos-notifier
Grab notifications from Cosmos SDK 

## Development

Example governance proposal 
```json
{
  "id": "34",
  "messages": [
    {
      "@type": "/cosmos.staking.v1beta1.MsgUpdateParams",
      "authority": "centauri10d07y265gmmuvt4z0w9aw880jnsr700j7g7ejq",
      "params": {
        "unbonding_time": "300s",
        "max_validators": 15,
        "max_entries": 7,
        "historical_entries": 10000,
        "bond_denom": "ppica",
        "min_commission_rate": "0.000000000000000000"
      }
    }
  ],
  "status": "PROPOSAL_STATUS_PASSED",
  "final_tally_result": {
    "yes_count": "821306325657092432991588",
    "abstain_count": "0",
    "no_count": "0",
    "no_with_veto_count": "0"
  },
  "submit_time": "2024-04-11T16:14:34.599148021Z",
  "deposit_end_time": "2024-04-11T16:19:34.599148021Z",
  "total_deposit": [
    {
      "denom": "ppica",
      "amount": "10000000"
    }
  ],
  "voting_start_time": "2024-04-11T16:14:34.599148021Z",
  "voting_end_time": "2024-04-11T16:31:14.599148021Z",
  "metadata": "",
  "title": "Update max_validators to 15",
  "summary": "Update max_validators to 15 for testing",
  "proposer": "centauri1r2zlh2xn85v8ljmwymnfrnsmdzjl7k6w6lytan"
}

```
- reverse order
```bash
curl -X GET "https://osmosis-rest.publicnode.com/cosmos/gov/v1/proposals?pagination.reverse=true" -H "accept: application/json"
```

- reverse order and get last one
```
curl -X GET "https://osmosis-rest.publicnode.com/cosmos/gov/v1beta1/proposals?pagination.limit=1&pagination.count_total=true&pagination.reverse=true" -H "accept: application/json" -s | jq .
```


