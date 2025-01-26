package valorantapi

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/joho/godotenv"
	govapi "github.com/yldshv/go-valorant-api"
)

func Authorization() *govapi.VAPI {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apikey := os.Getenv("VALOAPIKEY")

	vapi := govapi.New(govapi.WithKey(apikey))
	return vapi
}


func GetAccountPUUID(name string, tag string, vapi *govapi.VAPI) string {
	acc, err := vapi.GetAccountByName(govapi.GetAccountByNameParams{
		Name: name,
		Tag: tag,
	})
	if err != nil {
		log.Fatal(err)
	}

	puuid := acc.Data.Puuid
	fmt.Println(puuid)
	return puuid
}

type Match struct { 
    Id string
    MapName string
    Mode string
    Kills int
    Deaths int
    CharacterName string
}

func GetAccountMatches(puuid string, vapi *govapi.VAPI) *govapi.GetLifetimeMatchesByPUUIDResponse {
	matches, err := vapi.GetLifetimeMatchesByPUUID(govapi.GetLifetimeMatchesByPUUIDParams{
		PUUID: puuid,
		Affinity: "eu",
		Page: "1",
		Size: "12",
		Mode: "competitive",
	})
	if err != nil {
		fmt.Println("Error fetching matches:", err)
	}
	fmt.Printf("Status Code: %v\n", matches.Status)
    
    return matches
}

func GetAccountMMR(puuid string, vapi *govapi.VAPI){
	mmrHistory, err := vapi.GetLifetimeMMRHistoryByPUUID(govapi.GetLifetimeMMRHistoryByPUUIDParams{
		Affinity: "eu",
		Puuid: puuid, 
		Page: "1",
		Size: "12",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, mmrMatch := range mmrHistory.Data {
		fmt.Printf("Match ID: %s\n", mmrMatch.MatchID)
		fmt.Printf("Map Name: %s\n", mmrMatch.Map.Name)
		fmt.Printf("MMR Change: %d\n", mmrMatch.LastMmrChange)
		fmt.Printf("Elo: %d\n", mmrMatch.Elo)
		fmt.Println("---")
	}
}

func FormatMatches(response *govapi.GetLifetimeMatchesByPUUIDResponse) []Match {
    var matches []Match
	for _, match := range response.Data {
        matches = append(matches, Match{
            Id: match.Meta.ID,
            MapName: match.Meta.Map.Name,
            Mode: match.Meta.Mode,
            Kills: match.Stats.Kills,
            Deaths: match.Stats.Deaths,
            CharacterName: match.Stats.Character.Name,
        })
	}
    slices.Reverse(matches)
    return matches
}
