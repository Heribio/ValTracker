package valorantapi

import(
	"os"
	"fmt"
	"log"

	govapi "github.com/yldshv/go-valorant-api"
	"github.com/joho/godotenv"
)

func Authorization() *govapi.VAPI {
	err := godotenv.Load(".env")
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

func GetAccountMatches(puuid string, vapi *govapi.VAPI){
	matches, err := vapi.GetLifetimeMatchesByPUUID(govapi.GetLifetimeMatchesByPUUIDParams{
		PUUID: puuid,
		Affinity: "eu",
		Page: "1",
		Size: "12",
		Mode: "competitive",
	})
	if err != nil {
		fmt.Println("Error fetching matches:", err)
		return
	}

	fmt.Printf("Status Code: %v\n", matches.Status)

	for _, match := range matches.Data {
		fmt.Printf("Match ID: %s\n", match.Meta.ID)
		fmt.Printf("Map Name: %s\n", match.Meta.Map.Name)
		fmt.Printf("Game Mode: %s\n", match.Meta.Mode)
		fmt.Printf("Kills: %d\n", match.Stats.Kills)
		fmt.Printf("Deaths: %d\n", match.Stats.Deaths)
		fmt.Printf("Character: %s\n", match.Stats.Character.Name)
		fmt.Println("---")
	}
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
