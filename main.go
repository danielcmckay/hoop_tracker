package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Connecting to API...")
	
	currentGames := scrape("https://www.ncaa.com/scoreboard/basketball-men/d1")

		fmt.Printf("~~~~~~~~~~~~~~~~~~ There are %d games today ~~~~~~~~~~~~~~~~~~\n\n", len(currentGames))
		for i, g := range currentGames {
		fmt.Printf("%d\t%s: %s vs. %s: %s\t\t %s\n", i, g.team1, g.team1Score, g.team2, g.team2Score, g.timeRemaining)
		fmt.Println("----------------------------------------------------------------")
	}
}

type CBBGame struct {
	team1 string
	team2 string
	team1Rank string
	team2Rank string
	team1Score string
	team2Score string
	timeRemaining string
	isDone bool
}

func scrape(url string) []CBBGame {
	var games []CBBGame
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if resp.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		doc.Find(".gamePod-link").Each(func(i int, s *goquery.Selection) {

				team1 := s.Find(".gamePod-game-team-name").First().Text()
				team2 := s.Find(".gamePod-game-team-name").Last().Text()
				team1Score := s.Find(".gamePod-game-team-score").First().Text()
				team2Score := s.Find(".gamePod-game-team-score").Last().Text()
				team1Rank := s.Find(".gamePod-game-team-rank").First().Text()
				team2Rank := s.Find(".gamePod-game-team-rank").Last().Text()
				timeRemaining := s.Find(".game-clock").Text()
				var isDone bool
				
				if (s.Find(".gamePod-status").Text() == "Live") {
					isDone = false
				} else {
					isDone = true
				}

				game := CBBGame{team1, team2, team1Score, team2Score, team1Rank, team2Rank, timeRemaining, isDone}
				games = append(games, game)
			})
	}

	return games
}
