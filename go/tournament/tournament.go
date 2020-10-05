// Package tournament keeps score in tournaments
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
}

func (t *team) points() int {
	return t.wins*3 + t.draws
}

// Tally calculates the the tally for a series of games.
func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)

	teams := make(map[string]team)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Split(line, ";")

		if len(fields) != 3 {
			return errors.New("wrong number of fields")
		}

		homeTeamName := fields[0]
		awayTeamName := fields[1]
		result := fields[2]

		homeTeam := teams[homeTeamName]
		awayTeam := teams[awayTeamName]
		homeTeam.name = homeTeamName
		awayTeam.name = awayTeamName

		homeTeam.matchesPlayed++
		awayTeam.matchesPlayed++

		switch result {
		case "win":
			homeTeam.wins++
			awayTeam.losses++
		case "loss":
			awayTeam.wins++
			homeTeam.losses++
		case "draw":
			homeTeam.draws++
			awayTeam.draws++
		default:
			return errors.New("unexpected result")
		}

		teams[homeTeamName] = homeTeam
		teams[awayTeamName] = awayTeam
	}

	scores := make([]team, 0)

	for _, team := range teams {
		scores = append(scores, team)
	}

	sort.Slice(scores, func(i, j int) bool {
		iPoints := scores[i].points()
		jPoints := scores[j].points()
		if iPoints == jPoints {
			return scores[i].name < scores[j].name
		}

		return iPoints > jPoints
	})

	// first line is headers
	fmt.Fprintf(writer, "%-31v| %2v | %2v | %2v | %2v | %2v\n", "Team", "MP", "W", "D", "L", "P")

	for _, s := range scores {
		// find all teams with this s
		// print it
		fmt.Fprintf(writer, "%-31v| %2v | %2v | %2v | %2v | %2v\n", s.name, s.matchesPlayed, s.wins, s.draws, s.losses, s.points())
	}

	return nil
}
