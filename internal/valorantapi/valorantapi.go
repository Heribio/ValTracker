package valorantapi

import (
	"log"
    "fmt"

    "github.com/Heribio/ValTracker/internal/jsonthings"

	govapi "github.com/Heribio/go-valorant-api"
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

type Player struct {
    PUUID    string
    Username string
    Tag     string
    Kills int
    Deaths int
    Assists int
    CharacterName string
    Score   int 
    Team      string
    Rank    string
    Rounds  int
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

func GetMatch(matchID string) *govapi.GetMatchResponse{
    params := govapi.GetMatchParams{
        MatchId: matchID,
    }
    match, err := vapi.GetMatch(params)
    if err != nil {
        log.Fatal("Problem getting match", err)
    }
    return match
}

func GetPlayers(match *govapi.GetMatchResponse) []Player {
    var players []Player
    for _, player := range match.Data.Players.AllPlayers{
        players = append(players, Player{
            PUUID: player.Puuid,
            Username: player.Name,
            Tag: player.Tag,
            Kills: player.Stats.Kills,
            Deaths: player.Stats.Deaths,
            Assists: player.Stats.Assists,
            Score: player.Stats.Score,
            Rank: player.CurrenttierPatched,
            CharacterName: player.Character,
            Team: player.Team,
            Rounds : match.Data.Metadata.RoundsPlayed,
        })
    }
    return players
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
    params := govapi.GetEsportsScheduleParams{}

    var resp *govapi.GetEsportsScheduleResponse
    resp, err := vapi.GetEsportsSchedule(params)
    if err != nil {
        fmt.Println(resp)
        log.Fatal("GetVersion did not work for CheckToken\n", err)
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
