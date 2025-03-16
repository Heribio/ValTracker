package valorantapi

import (
	"log"
    "fmt"

    "github.com/Heribio/ValTracker/internal/jsonthings"

	govapi "github.com/yldshv/go-valorant-api"
)

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
    Rank string
}

var vapi = Authorization()

func Authorization() *govapi.VAPI {
    tokenData := jsonthings.GetTokenData()
    apikey := tokenData.ValApiToken

	vapi := govapi.New(govapi.WithKey(apikey))
	return vapi
}

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

func AppendMatchList(list []Match, page string, affinity string, mode string) []Match {
    puuid := GetAccountPUUID(jsonthings.GetFileData("data.json").Name, jsonthings.GetFileData("data.json").Tag)
    moreMatches := GetAccountMatches(puuid, page, affinity, mode) 

    list = append(list, moreMatches...)
    return list
}

func GetAccountMatches(puuid string, page string, affinity string, mode string) []Match {
    size := "10"
	apiresp, err := vapi.GetLifetimeMatchesByPUUID(
        govapi.GetLifetimeMatchesByPUUIDParams{
            PUUID: puuid,
            Affinity: affinity, //eu
            Page: page,
            Size: size,
            Mode: mode, 
        })
	if err != nil {
		fmt.Println("Error fetching matches:", err)
	}
    mmrApiResp, err := vapi.GetLifetimeMMRHistoryByPUUID(
        govapi.GetLifetimeMMRHistoryByPUUIDParams{
            Puuid: puuid,
            Affinity: affinity,
            Size: size,
            Page: page,
        })
    matches := FormatMatches(apiresp, mmrApiResp)
    return matches
}

func GetAccountMMR(puuid string, affinity string) *govapi.GetMMRByPUUIDv2Response{
	mmrHistory, err := vapi.GetMMRByPUUIDv2(govapi.GetMMRByPUUIDv2Params{
        Affinity: affinity,
		Puuid: puuid, 
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

func FormatMatches(response *govapi.GetLifetimeMatchesByPUUIDResponse, mmrResponse *govapi.GetLifetimeMMRHistoryByPUUIDResponse) []Match {
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

func getMMR() string {
    puuid := GetAccountPUUID(jsonthings.GetFileData("data.json").Name, jsonthings.GetFileData("data.json").Tag)
    mmrList := GetAccountMMR(puuid, "eu") 
    return mmrList.Data.CurrentData.CurrenttierPatched
}
