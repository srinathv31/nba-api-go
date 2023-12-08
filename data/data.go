package data

type Data map[string]TeamHistory

type TeamHistory map[string]Team

type Team struct {
    Roster   Roster   `json:"Roster"`
    Schedule Schedule `json:"Schedule"`
}

type Roster struct {
    PlayerMap PlayerMap `json:"players"`
    URL       string    `json:"url"`
}

type PlayerMap map[string]Player

type Player struct {
    G     string `json:"G"`
    PER   string `json:"PER"`
    P_G   string `json:"P_G"`
    P_PER string `json:"P_PER"`
    P_TS  string `json:"P_TS%"`
    P_WS  string `json:"P_WS"`
    TS    string `json:"TS%"`
    WS    string `json:"WS"`
}

type Schedule struct {
    GameMap GameMap `json:"games"`
    URL     string  `json:"url"`
}

type GameMap map[string]Game

type Game struct {
    BoxScore string `json:"Box Score"`
    Date     string `json:"Date"`
    L        string `json:"L"`
    Location string `json:"Location"`
    Opponent string `json:"Opponent"`
    PtsFor   string `json:"PtsFor"`
    PtsOpp   string `json:"PtsOpp"`
    Result   string `json:"Result"`
    Start_ET string `json:"Start (ET)"`
    Streak   string `json:"Streak"`
    W        string `json:"W"`
}
