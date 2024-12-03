package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/domain"
	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/rest"
)

func GetAllMatches(restClient rest.NewRestInterface, tournamentID string) []domain.AllMatches {
	// URL is already "https://api.challonge.com/v1/"
	reqURL := fmt.Sprintf("tournaments/%s/matches.json?%s", tournamentID, Params.Encode())

	// Form Request
	req, err := http.NewRequest("GET", reqURL, http.NoBody)
	if err != nil {
		log.Errorf("Failed to create request: %v", err.Error())
		return []domain.AllMatches{}
	}

	// Send Request
	resp, err := restClient.SendRequest(req)
	if err != nil {
		log.Errorf("Request Failed: %v", err.Error())
		return []domain.AllMatches{}
	}

	log.Info("Status code: %v", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		var respBody []domain.AllMatches
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			log.Error("Error decoding JSON:", err)
		}
		log.Infof("Response Body: %v", respBody)
		return respBody
	}
	return []domain.AllMatches{}
}

func GetAllOpenMatches(restClient rest.NewRestInterface, tournamentID string) []domain.AllMatches {
	allMatches := GetAllMatches(restClient, tournamentID)

	var allOpenMatches []domain.AllMatches
	for _, match := range allMatches {
		if match.Match.State == "open" {
			allOpenMatches = append(allOpenMatches, match)
		}
	}

	return allOpenMatches
}

func UpdateMatch(restClient rest.NewRestInterface, updateRequest domain.UpdateMatchRequest, tournamentId, matchID string) bool {
	score := updateRequest.Score
	winner := updateRequest.WinnerID

	// Add params
	Params.Add("match[scores_csv]", score)
	if winner != "" {
		Params.Add("match[winner_id]", winner)
	}

	// URL is already "https://api.challonge.com/v1/"
	reqURL := fmt.Sprintf("tournaments/%s/matches/%s.json?%s", tournamentId, matchID, Params.Encode())

	// Form Request
	req, err := http.NewRequest("PUT", reqURL, http.NoBody)
	if err != nil {
		log.Errorf("Failed to create request: %v", err.Error())
		return false
	}

	// Send Request
	resp, err := restClient.SendRequest(req)
	if err != nil {
		log.Errorf("Request Failed: %v", err.Error())
		return false
	}

	log.Infof("Status code: %v", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		var respBody domain.AllMatches
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			log.Error("Error decoding JSON:", err)
		}
		log.Infof("Response Body: %v", respBody)
		return true
	}
	return false
}
