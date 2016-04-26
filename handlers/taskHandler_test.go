package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	getHandler := GetTasks(db)
	mock.ExpectQuery("select taskId,task,priority from tasks").WillReturnRows(sqlmock.NewRows([]string{"som"}))
	req, _ := http.NewRequest("GET", "/getAllTasks", nil)
	w := httptest.NewRecorder()
	getHandler.ServeHTTP(w, req)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}

}
