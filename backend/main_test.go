package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "John Doe", "john@example.com").
		AddRow(2, "Jane Doe", "jane@example.com")

	mock.ExpectQuery("SELECT \\* FROM users").WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/api/go/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := getUsers(db)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `[{"id":1,"name":"John Doe","email":"john@example.com"},{"id":2,"name":"Jane Doe","email":"jane@example.com"}]`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "John Doe", "john@example.com")
	mock.ExpectQuery("SELECT \\* from users WHERE id = \\$1").WithArgs("1").WillReturnRows(row)

	req, err := http.NewRequest("GET", "/api/go/users/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/go/users/{id}", getUser(db)).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"id":1,"name":"John Doe","email":"john@example.com"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("INSERT INTO users \\(name, email\\) VALUES \\(\\$1, \\$2\\) RETURNING id").
		WithArgs("John Doe", "john@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	user := `{"name":"John Doe","email":"john@example.com"}`
	req, err := http.NewRequest("POST", "/api/go/users", strings.NewReader(user))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := createUser(db)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"id":1,"name":"John Doe","email":"john@example.com"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("UPDATE users SET name = \\$1, email = \\$2 WHERE id = \\$3").
		WithArgs("John Doe Updated", "john_updated@example.com", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT id, name, email FROM users WHERE id = \\$1").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(1, "John Doe Updated", "john_updated@example.com"))

	user := `{"name":"John Doe Updated","email":"john_updated@example.com"}`
	req, err := http.NewRequest("PUT", "/api/go/users/1", strings.NewReader(user))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/go/users/{id}", updateUser(db)).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"id":1,"name":"John Doe Updated","email":"john_updated@example.com"}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(1, "John Doe", "john@example.com"))

	mock.ExpectExec("DELETE FROM users WHERE id = \\$1").WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("DELETE", "/api/go/users/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/go/users/{id}", deleteUser(db)).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `"User deleted"`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestEnableCORS(t *testing.T) {
	handler := enableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, err := http.NewRequest("OPTIONS", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "*", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET, POST, PUT, DELETE, OPTIONS", rr.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type, Authorization", rr.Header().Get("Access-Control-Allow-Headers"))
}

func TestJSONContentTypeMiddleware(t *testing.T) {
	handler := jsonContentTypeMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}
