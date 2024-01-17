package parsing

import (
	"strconv"
	"task11/internal/domain/model"
	"time"
)

func ParseEventFromMap(m map[string]string) (*model.Event, error) {
	event := &model.Event{}

	if strId, ok := m["id"]; ok {
		id, err := strconv.Atoi(strId)
		if err != nil {
			return nil, err
		}
		event.Id = id
	}

	if strUserId, ok := m["user_id"]; ok {
		userId, err := strconv.Atoi(strUserId)
		if err != nil {
			return nil, err
		}
		event.UserId = userId
	}

	if name, ok := m["name"]; ok {
		event.Name = name
	}

	if strDate, ok := m["date"]; ok {
		date, err := time.Parse(time.DateOnly, strDate)
		if err != nil {
			return nil, err
		}
		event.Date = date
	}

	return event, nil
}
