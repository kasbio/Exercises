package dao

import (
	"database/sql"
	"errors"
	"fmt"
)

const (
	disConnected = iota
	connected
)

type EntityService struct {
	Connection string
	Status     int
}

type Row struct {
	Data interface{}
}

var DisConnectedError = errors.New("status : disconnected")

func (s *EntityService) Connect() error {

	if s.Connection == "" {
		s.Status = disConnected
		return errors.New("empty connection string")
	}
	s.Status = connected
	return nil
}

func (s *EntityService) Query(queryStr string) ([]Row, error) {
	if s.Status != connected {
		return nil, DisConnectedError
	}

	if rows, err := s.executeCore(queryStr); err != nil {
		return rows, fmt.Errorf("Query not result: \n\t queryStr:%v \n\t %w", queryStr, err)
	} else {
		return rows, nil
	}

}

func (s *EntityService) executeCore(queryStr string) ([]Row, error) {

	rows := make([]Row, 0)
	//to do ...

	//no result and wrap error
	if len(rows) == 0 {
		return nil, sql.ErrNoRows
	}

	return rows, nil
}
