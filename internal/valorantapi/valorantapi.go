package valorantapi

import (
	"fmt"
	"log"
	// "os"

    "github.com/Heribio/ValTracker/internal/jsonthings"

	// "github.com/joho/godotenv"
	govapi "github.com/yldshv/go-valorant-api"
)

func Authorization() *govapi.VAPI {
	// err := godotenv.Load("cmd/.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	//
	// apikey := os.Getenv("VALOAPIKEY")
    tokenData := jsonthings.GetTokenData()
    apikey := tokenData.ValApiToken

	vapi := govapi.New(govapi.WithKey(apikey))
	return vapi
}


var vapi = Authorization()

func GetAccountPUUID(name string, tag string) string {
    vapi := Authorization()
	acc, err := vapi.GetAccountByName(govapi.GetAccountByNameParams{
		Name: name,
		Tag: tag,
	})
	if err != nil {
		log.Fatal(err)
	}

	puuid := acc.Data.Puuid
	return puuid
}

type Match struct { 
    Id string
    MapName string
    Mode string
    Kills int
    Deaths int
    Assists int
    CharacterName string
    StartedAt string 
    Score   int 
    Team      string
    RedTeamScore  int 
    BlueTeamScore int
}

func AppendMatchList(list []Match, page string, affinity string, mode string) []Match {
    puuid := GetAccountPUUID(jsonthings.GetFileData("data.json").Name, jsonthings.GetFileData("data.json").Tag)
    moreMatches := GetAccountMatches(puuid, page, affinity, mode) 

    list = append(list, moreMatches...)
    return list
}

func GetAccountMatches(puuid string, page string, affinity string, mode string) []Match {
	apiresp, err := vapi.GetLifetimeMatchesByPUUID(
        govapi.GetLifetimeMatchesByPUUIDParams{
            PUUID: puuid,
            Affinity: affinity, //eu
            Page: page,
            Size: "12",
            Mode: mode, //competitive
        })
	if err != nil {
		fmt.Println("Error fetching matches:", err)
	}
    matches := FormatMatches(apiresp)
    return matches
}

func GetAccountMMR(puuid string, affinity string, page string) *govapi.GetLifetimeMMRHistoryByPUUIDResponse {
	mmrHistory, err := vapi.GetLifetimeMMRHistoryByPUUID(govapi.GetLifetimeMMRHistoryByPUUIDParams{
		Affinity: affinity,
		Puuid: puuid, 
		Page: page,
		Size: "12",
	})
	if err != nil {
		log.Fatal(err)
	}
    return mmrHistory
}

func CheckToken() bool {
    token := jsonthings.GetTokenData().ValApiToken
    vapi := govapi.New(govapi.WithKey(token))
    params := govapi.GetStatusParams{
        Affinity: "eu",
    }

    var resp *govapi.GetStatusResponse
    resp, err := vapi.GetStatus(params)
    if err != nil {
        log.Fatal("GetStatus did not work for CheckToken")
    }
    if resp.Errors != nil {
        return false
    } else {
        return true
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
            Assists: match.Stats.Assists,
            CharacterName: match.Stats.Character.Name,
            StartedAt: match.Meta.StartedAt,
            Score: match.Stats.Score,
            RedTeamScore: match.Teams.Red,
            BlueTeamScore: match.Teams.Blue,
            Team: match.Stats.Team,
        })
	}
    return matches
}
