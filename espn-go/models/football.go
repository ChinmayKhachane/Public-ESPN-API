package models

// PlayerHit is the normalized Search v2 player result.
type PlayerHit struct {
	AthleteID   string
	UID         string
	GUID        string
	DisplayName string
	Description string
	League      string
	Raw         JSON
}

// TeamHit is the normalized site-v2 team result.
type TeamHit struct {
	TeamID       string
	Abbreviation string
	DisplayName  string
	Location     string
	Nickname     string
	Raw          JSON
}

// ResolvedGame contains IDs needed for game-specific core endpoint chains.
type ResolvedGame struct {
	EventID       string
	CompetitionID string
	CompetitorIDs []string
	TeamIDs       []string
	Summary       JSON
}

// StatLine is a flattened stat entry from core-v2 splits/categories.
type StatLine struct {
	Category     string
	Name         string
	DisplayName  string
	Value        any
	DisplayValue string
	Raw          JSON
}

// PlayerGameStatsInput describes the player game-stat terminal chain.
type PlayerGameStatsInput struct {
	EventID       string
	AthleteID     string
	PlayerName    string
	TeamID        string
	CompetitionID string
	CompetitorID  string
	Split         string
}

// PlayerGameStats is the normalized terminal response for one player in one game.
type PlayerGameStats struct {
	EventID       string
	CompetitionID string
	CompetitorID  string
	AthleteID     string
	Stats         []StatLine
	Raw           JSON
	Summary       JSON
}

// GameSearch describes the non-eventID game resolution workflow.
type GameSearch struct {
	EventID      string
	TeamID       string
	TeamName     string
	OpponentID   string
	OpponentName string
	Date         string
	Season       string
	SeasonType   string
	Week         string
}

// IDBundle groups IDs resolved by a football workflow.
type IDBundle struct {
	Player *PlayerHit
	Team   *TeamHit
	Game   *ResolvedGame
}
