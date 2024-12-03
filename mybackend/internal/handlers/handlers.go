package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/domain"
	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/rest"
	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/utils/constants"
	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	log = logger.GetLogger()
	// Set Param
	Params = url.Values{
		constants.ApiKey: []string{constants.ChallongeApiToken},
	}
	AllMatches      []domain.AllMatches
	AllParticipants []domain.ParticipantWrapper
)

type RoutesHadlerInterface interface {
	TournamentsView(c *gin.Context)
	GetAllMatchesOfTournament(c *gin.Context)
	UpdateMatch(c *gin.Context)
}

type RoutesHandler struct {
	challongeClient rest.NewRestInterface
}

func NewRoutesHandler() RoutesHadlerInterface {
	// RestClient to interact with Challonge
	challongeClient, restErr := rest.NewRestClient(constants.ChallongeHostName, constants.Timeout)
	if restErr != nil {
		log.Errorf("Failed to create rest client: %v", restErr.Error())
		return &RoutesHandler{}
	}
	// Set BaseURl
	challongeClient.SetBaseURL(constants.HTTPS, constants.ChallongeHostName)
	return &RoutesHandler{
		challongeClient: challongeClient,
	}
}

func (route RoutesHandler) TournamentsView(c *gin.Context) {
	// Recieve tourname slug or tournament id (ex:  fn url is "https://challonge.com/p1l7bugj" tournamet slug/id is p1l7bugj)
	tournamentID := c.Param("tournamentID")
	log.Infof("Tournament ID: %s", tournamentID)

	if route.challongeClient == nil {
		log.Errorf("Rest client creation failed, please check the logs")
		c.JSON(http.StatusInternalServerError, "Something went wrong on our side")
		return
	}

	// Add additonal Params
	Params.Add("include_participants", "1")

	// URL is already "https://api.challonge.com/v1/"
	reqURL := fmt.Sprintf("tournaments/%s.json?%s", tournamentID, Params.Encode())

	// Form Request
	req, err := http.NewRequest("GET", reqURL, http.NoBody)
	if err != nil {
		log.Errorf("Failed to create request: %v", err.Error())
		return
	}

	// Send Request
	resp, err := route.challongeClient.SendRequest(req)
	if err != nil {
		log.Errorf("Request Failed: %v", err.Error())
		return
	}

	log.Info("Status code: %v", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		var respBody domain.ChallongeTournamentInfo
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			log.Error("Error decoding JSON:", err)
		}
		log.Infof("Response Body: %v", respBody)
		c.JSON(http.StatusOK, respBody)
		return
	}
	c.JSON(http.StatusInternalServerError, "Something went wrong")
}

func (route RoutesHandler) GetAllMatchesOfTournament(c *gin.Context) {
	// Recieve tourname slug or tournament id (ex:  fn url is "https://challonge.com/p1l7bugj" tournamet slug/id is p1l7bugj)
	tournamentID := c.Param("tournamentID")
	log.Infof("Tournament ID: %s", tournamentID)

	if route.challongeClient == nil {
		log.Errorf("Rest client creation failed, please check the logs")
		c.JSON(http.StatusInternalServerError, "Something went wrong on our side")
		return
	}

	fetchedAllMatches := GetAllMatches(route.challongeClient, tournamentID)

	AllMatches = fetchedAllMatches

	if len(AllMatches) == 0 {
		c.JSON(http.StatusInternalServerError, "Something went wrong")

	}
	c.JSON(http.StatusOK, AllMatches)
}

func (route RoutesHandler) UpdateMatch(c *gin.Context) {
	// Recieve tournament slug or tournament id (ex:  fn url is "https://challonge.com/p1l7bugj" tournamet slug/id is p1l7bugj)
	tournamentID := c.Param("tournamentID")
	log.Infof("Tournament ID: %s", tournamentID)

	// Match ID
	matchID := c.Param("matchID")
	log.Infof("Match ID: %s", matchID)
	mid, err := strconv.Atoi(matchID)
	if err != nil {
		log.Error("Error converting string to int", err.Error())
		c.JSON(http.StatusInternalServerError, "Something went wrong on our side")
		return
	}

	// Request Body
	var updateMatchRequest domain.UpdateMatchRequest
	err = c.ShouldBindJSON(&updateMatchRequest) // Pass a pointer here
	if err != nil {
		log.Error("Error in update request: ", err.Error())
		c.JSON(http.StatusInternalServerError, "Something went wrong on our side")
		return
	}

	if route.challongeClient == nil {
		log.Errorf("Rest client creation failed, please check the logs")
		c.JSON(http.StatusInternalServerError, "Something went wrong on our side")
		return
	}

	log.Infof("Params: %v", Params)

	allOpenMatches := GetAllOpenMatches(route.challongeClient, tournamentID)

	var status bool
	if len(allOpenMatches) != 0 {
		for _, match := range allOpenMatches {
			if mid == match.Match.ID {
				status = UpdateMatch(route.challongeClient, updateMatchRequest, tournamentID, matchID)
				break
			}
		}
	}

	if status {
		c.JSON(http.StatusOK, "Match updated successfully")
	} else {
		c.JSON(http.StatusInternalServerError, "Failed to update match")
	}
}
