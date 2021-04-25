package dao

import (
	"database/sql"
	"errors"
	"testing"
)

func TestQueryNoRowErr(t *testing.T) {
	s := &EntityService{Connection: "123456"}
	s.Connect()
	if _, err := s.Query("select * from table"); err != nil {
		t.Logf("wrap error:%v \n\t", err)

		if errors.Is(err, sql.ErrNoRows) {
			t.Logf("root error:%v \n\t", errors.Unwrap(err))
			return
		}
		t.Error("no return errnorows")
	}

}
