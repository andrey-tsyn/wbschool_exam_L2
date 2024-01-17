package handler

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"task11/internal/domain/model"
	"task11/internal/domain/repository"
	"task11/internal/domain/service"
	"task11/internal/domain/validation"
	"task11/internal/infrastructure/parsing"
	"time"
)

func CreateEvent(eventService *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			log.Error(err)
			sendErrorResponse(w, errorResponse{
				StatusCode: http.StatusServiceUnavailable,
				Msg:        err.Error(),
			})
			return
		}

		event, err := parsing.ParseEventFromMap(formToStringMap(r.Form))
		if err != nil {
			sendErrorResponse(w, errorResponse{
				StatusCode: http.StatusBadRequest,
				Msg:        err.Error(),
			})
			return
		}

		id, err := eventService.Create(*event)
		if err != nil {
			switch {
			case errors.As(err, &validation.Error{}):
				fallthrough
			case errors.As(err, &repository.EventAlreadyExist{}):
				sendErrorResponse(w, errorResponse{
					StatusCode: http.StatusBadRequest,
					Msg:        err.Error(),
				})
			default:
				sendErrorResponse(w, errorResponse{
					StatusCode: http.StatusServiceUnavailable,
					Msg:        err.Error(),
				})
			}
			return
		}

		event.Id = id

		sendSuccessResponse(w, successResponse{
			ObjectToSerialize: event,
			Msg:               "",
		})
	}
}

func UpdateEvent(eventService *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			log.Error(err)
			sendErrorResponse(w, errorResponse{
				StatusCode: http.StatusServiceUnavailable,
				Msg:        err.Error(),
			})
			return
		}

		event, err := parsing.ParseEventFromMap(formToStringMap(r.Form))
		if err != nil {
			sendErrorResponse(w, errorResponse{
				StatusCode: http.StatusBadRequest,
				Msg:        err.Error(),
			})
			return
		}

		err = eventService.Update(*event)
		if err != nil {
			switch {
			case errors.As(err, &validation.Error{}):
				fallthrough
			case errors.As(err, &repository.EventNotFoundError{}):
				sendErrorResponse(w, errorResponse{
					StatusCode: http.StatusBadRequest,
					Msg:        err.Error(),
				})
			default:
				sendErrorResponse(w, errorResponse{
					StatusCode: http.StatusServiceUnavailable,
					Msg:        err.Error(),
				})
			}
			return
		}

		sendSuccessResponse(w, successResponse{
			ObjectToSerialize: nil,
			Msg:               "",
		})
	}
}

func DeleteEvent(eventService *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			sendErrorResponse(w, errorResponse{
				StatusCode: http.StatusBadRequest,
				Msg:        "id must be provided and be an number",
			})
			return
		}

		err = eventService.Delete(id)
		if err != nil {
			switch {
			case errors.As(err, &validation.Error{}):
				fallthrough
			case errors.As(err, &repository.EventNotFoundError{}):
				sendErrorResponse(w, errorResponse{
					StatusCode: http.StatusBadRequest,
					Msg:        err.Error(),
				})
			default:
				sendErrorResponse(w, errorResponse{
					StatusCode: http.StatusServiceUnavailable,
					Msg:        err.Error(),
				})
			}
			return
		}

		sendSuccessResponse(w, successResponse{
			ObjectToSerialize: nil,
			Msg:               "",
		})
	}
}

func getEventsFromTodayTo(w http.ResponseWriter, r *http.Request, eventService *service.EventService, daysAfter time.Duration) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	strUserId := r.URL.Query().Get("user_id")

	userId, err := strconv.Atoi(strUserId)

	if err != nil {
		sendErrorResponse(w, errorResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        fmt.Sprintf("user_id must be provided and be an number, %s provided", strUserId),
		})
		return
	}

	date, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))

	events, err := eventService.GetByUserId(userId, date, date.Add(daysAfter*24*time.Hour))
	if err != nil {
		sendErrorResponse(w, errorResponse{
			StatusCode: http.StatusServiceUnavailable,
			Msg:        err.Error(),
		})
		return
	}

	result := struct {
		UserId int           `json:"user_id"`
		Events []model.Event `json:"events"`
	}{userId, events}

	sendSuccessResponse(w, successResponse{
		ObjectToSerialize: result,
		Msg:               "",
	})
}

func GetEventsForDay(eventService *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		getEventsFromTodayTo(w, r, eventService, time.Duration(0))
	}
}

func GetEventsForWeek(eventService *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		getEventsFromTodayTo(w, r, eventService, time.Duration(7))
	}
}

func GetEventsForMonth(eventService *service.EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		getEventsFromTodayTo(w, r, eventService, time.Duration(31))
	}
}
