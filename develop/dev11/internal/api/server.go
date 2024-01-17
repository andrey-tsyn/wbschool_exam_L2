package api

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"task11/internal/api/handler"
	"task11/internal/api/middleware"
	"task11/internal/configuration"
	"task11/internal/domain/repository"
	"task11/internal/domain/service"
)

var eventService *service.EventService

func StartServer(config configuration.Config, repositories map[string]interface{}) {
	initServices(repositories)

	router := http.NewServeMux()

	// Adding routes
	router.HandleFunc("/create_event", handler.CreateEvent(eventService))
	router.HandleFunc("/update_event", handler.UpdateEvent(eventService))
	router.HandleFunc("/delete_event", handler.DeleteEvent(eventService))
	router.HandleFunc("/events_for_day", handler.GetEventsForDay(eventService))
	router.HandleFunc("/events_for_week", handler.GetEventsForWeek(eventService))
	router.HandleFunc("/events_for_month", handler.GetEventsForMonth(eventService))

	// Wrapping with middlewares
	wrappedRouter := middleware.WrapWithLoggerMiddleware(router)

	log.Infof("Listening for client connections on localhost:%s", config.Port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", config.Port), wrappedRouter)
	if err != nil {
		log.Fatal(err)
	}
}

func initServices(repositories map[string]interface{}) {
	eventService = service.NewEventService(repositories["event"].(repository.EventRepository))
}
