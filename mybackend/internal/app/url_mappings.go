package app

import "github.com/Shreesha5102/ChallongeTool/mybackend/internal/handlers"

func mapRestUrls() {
	// Handler
	routes := handlers.NewRoutesHandler()

	// Group Route
	resource := router.Group("/challongetool/v1")

	// Tournament Route
	resource.GET("/tournament/:tournamentID", routes.TournamentsView)

	// Matches Routes
	resource.GET("/matches/:tournamentID", routes.GetAllMatchesOfTournament)
	resource.PUT("/matches/:tournamentID/:matchID", routes.UpdateMatch)

}
