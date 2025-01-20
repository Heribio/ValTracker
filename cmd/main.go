package main

import(
	"github.com/Heribio/ValTracker/internal/valorantapi"
)

func main() {
	vapi := valorantapi.Authorization()

	puuid := valorantapi.GetAccountPUUID("Name", "Tag", vapi)	
	valorantapi.GetAccountMatches(puuid, vapi)
	valorantapi.GetAccountMMR(puuid, vapi)
}
