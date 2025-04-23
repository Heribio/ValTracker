package valorantapi

type Username struct {
    Name    string
    Tag     string
}

type TokenData struct {
    ValApiToken string
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
    Rank string
    Headshots int
    Bodyshots int
    Legshots int
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
    Headshots int
    Bodyshots int
    Legshots int
}
