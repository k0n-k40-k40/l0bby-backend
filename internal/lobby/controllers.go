package lobby

import (
	_ "database/sql"
	. "l0bby_backend/internal/party"

	_ "github.com/go-sql-driver/mysql"
)

var incomingParties = []Party{}

func matchParties() {
	// copy incomingParties into a constant array
	parties := make([]Party, len(incomingParties))

	const maxTeams = 8
	const teamSize = 5
	const maxPlayers = maxTeams * teamSize

	dp := make([][]int, maxPlayers)
	for i := range dp {
		dp[i] = make([]int, len(parties))
	}

	trace := make([][]int, maxPlayers)
	for i := range trace {
		trace[i] = make([]int, len(parties))
	}

	for i, party := range parties {
		for sum := 0; sum < maxPlayers; sum++ {
			dp[sum][i] = dp[sum][i-1]
			party_size := len(party.Members)

			if sum < party_size {
				continue
			}

			if dp[sum][i] < dp[sum-party_size][i-1]+1 {
				dp[sum][i] = dp[sum-party_size][i-1] + 1
				trace[sum][i] = i
			}
		}
	}

	// trace back to find the parties
	sum := maxPlayers - 1
	for i := len(parties) - 1; i >= 0; i-- {
		if trace[sum][i] == i {
			// party i is in the solution
			sum -= len(parties[i].Members)
		}
	}

	// parties in the solution
	solution := []Party{}
	for i := 0; i < len(parties); i++ {
		if trace[sum][i] == i {
			solution = append(solution, parties[i])
			sum -= len(parties[i].Members)
		}
	}

	// remove the matched parties from incomingParties
	for _, party := range solution {
		for i, p := range incomingParties {
			if p.ID == party.ID {
				incomingParties = append(incomingParties[:i], incomingParties[i+1:]...)
				break
			}
		}
	}

}
