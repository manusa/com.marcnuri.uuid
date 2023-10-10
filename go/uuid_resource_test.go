package main

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"regexp"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = setupRouter()
	m.Run()
}

func TestNoQuantityParam(t *testing.T) {
	t.Run("Returns 200 status code", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, httptest.NewRequest("GET", "/", nil))
		if recorder.Code != 200 {
			t.Error("Expected 200, got ", recorder.Code)
		}
	})
	t.Run("Returns a single UUID", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, httptest.NewRequest("GET", "/", nil))
		body := recorder.Body.String()
		isUuid, _ := regexp.MatchString("^\\[\\n\\s+\"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\"\\n]$", body)
		if !isUuid {
			t.Error("Returned body doesn't contain a UUID: ", body)
		}
	})
}

func TestInvalidQuantityParam(t *testing.T) {
	t.Run("Returns 400 status code", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, httptest.NewRequest("GET", "/?quantity=not-a-number", nil))
		if recorder.Code != 400 {
			t.Error("Expected 400, got ", recorder.Code)
		}
	})
	t.Run("Returns body with message", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, httptest.NewRequest("GET", "/?quantity=not-a-number", nil))
		body := recorder.Body.String()
		if body != "Invalid quantity not-a-number" {
			t.Error("Returned body doesn't contain expected message: ", body)
		}
	})
}

func TestOutOfBoundsQuantityParam(t *testing.T) {
	t.Run("Returns 400 status code", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, httptest.NewRequest("GET", "/?quantity=99999999", nil))
		if recorder.Code != 400 {
			t.Error("Expected 400, got ", recorder.Code)
		}
	})
	t.Run("Returns body with message", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, httptest.NewRequest("GET", "/?quantity=-100", nil))
		body := recorder.Body.String()
		if body != "Quantity must be between 1 and 20000" {
			t.Error("Returned body doesn't contain expected message: ", body)
		}
	})
}
