package models

type User struct {
	ActiveTag string `json:"activeTag"`
	Backpack  struct {
		Loot map[string]interface{} `json:"loot"`
		Rows int                    `json:"rows"`
	} `json:"backpack"`
	Balance        int64                  `json:"balance"`
	Credits        int64                  `json:"credits"`
	FlightTime     int                    `json:"flightTime"`
	God            bool                   `json:"god"`
	HomeCapacity   int                    `json:"homeCapacity"`
	Homes          map[string]interface{} `json:"homes"`
	MerchantItems  []interface{}          `json:"merchantItems"`
	PlayerWarps    map[string]interface{} `json:"playerWarps"`
	Pvp            bool                   `json:"pvp"`
	StatisticsUser StatisticsUser         `json:"statisticsUser"`
}

type StatisticsUser struct {
	BlocksMined    map[string]interface{} `json:"blocksMined"`
	CropsHarvested map[string]interface{} `json:"cropsHarvested"`
	CropsPlaced    map[string]interface{} `json:"cropsPlaced"`
	Deaths         int                    `json:"deaths"`
	FishCaught     map[string]interface{} `json:"fishCaught"`
	Kills          int                    `json:"kills"`
	Level          int                    `json:"level"`
	MobDeaths      map[string]interface{} `json:"mobDeaths"`
	MobKills       map[string]interface{} `json:"mobKills"`
	OresMined      map[string]interface{} `json:"oresMined"`
	PlayTime       int                    `json:"playTime"`
	RubbishCaught  map[string]interface{} `json:"rubbishCaught"`
	Votes          int                    `json:"votes"`
}
