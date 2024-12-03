package domain

type TournamentInfo struct {
	TournamentID string `json:"tournamentid" binding:"required"`
}

// Tournament Resp Body
type ChallongeTournamentInfo struct {
	Tournament Tournament `json:"tournament" binding:"required"`
}

type Tournament struct {
	CreatedAt    string               `json:"created_at"`
	GameID       *int                 `json:"game_id"`
	ID           *int                 `json:"id"`
	Name         string               `json:"name"`
	FullURL      string               `json:"full_challonge_url"`
	Participants []ParticipantWrapper `json:"participants"`
}

type ParticipantWrapper struct {
	Participant Participant `json:"participant"`
}

type Participant struct {
	ID           *int   `json:"id"`
	TournamentID *int   `json:"tournament_id"`
	Name         string `json:"name"`
	Seed         *int   `json:"seed"`
	Active       bool   `json:"active"`
}

// All matches response body
type AllMatches struct {
	Match Match `json:match`
}

type Match struct {
	CreatedAt     string  `json:"created_at"`
	GroupID       *int    `json:"group_id"` // Use *int for nullable values
	ID            int     `json:"id"`
	Identifier    string  `json:"identifier"`
	Location      *string `json:"location"`
	LoserID       *int    `json:"loser_id"`
	Player1ID     int     `json:"player1_id"`
	Player2ID     int     `json:"player2_id"`
	Round         int     `json:"round"`
	ScheduledTime *string `json:"scheduled_time"`
	StartedAt     string  `json:"started_at"`
	State         string  `json:"state"`
	TournamentID  int     `json:"tournament_id"`
	UnderwayAt    *string `json:"underway_at"`
	UpdatedAt     string  `json:"updated_at"`
	WinnerID      *int    `json:"winner_id"`
	ScoresCSV     string  `json:"scores_csv"`
}

// Update match Request
type UpdateMatchRequest struct {
	Score    string `json:"score_csv"`
	WinnerID string `json:"winner_id"`
}
