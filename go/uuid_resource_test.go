package main

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"regexp"
	"testing"
)

type testContext struct {
	router   *gin.Engine
	recorder *httptest.ResponseRecorder
}

func (c *testContext) beforeEach() {
	c.router = setupRouter()
	c.recorder = httptest.NewRecorder()
}

func (c *testContext) afterEach() {
}

func testCase(test func(t *testing.T, c *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &testContext{}
		context.beforeEach()
		defer context.afterEach()
		test(t, context)
	}
}

func TestNoQuantityParam(t *testing.T) {
	t.Run("Returns 200 status code", testCase(func(t *testing.T, c *testContext) {
		c.router.ServeHTTP(c.recorder, httptest.NewRequest("GET", "/", nil))
		if c.recorder.Code != 200 {
			t.Error("Expected 200, got ", c.recorder.Code)
		}
	}))
	t.Run("Returns a single UUID", testCase(func(t *testing.T, c *testContext) {
		c.router.ServeHTTP(c.recorder, httptest.NewRequest("GET", "/", nil))
		body := c.recorder.Body.String()
		isUuid, _ := regexp.MatchString("^\\[\\n\\s+\"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\"\\n]$", body)
		if !isUuid {
			t.Error("Returned body doesn't contain a UUID: ", body)
		}
	}))
}

func TestInvalidQuantityParam(t *testing.T) {
	t.Run("Returns 400 status code", testCase(func(t *testing.T, c *testContext) {
		c.router.ServeHTTP(c.recorder, httptest.NewRequest("GET", "/?quantity=not-a-number", nil))
		if c.recorder.Code != 400 {
			t.Error("Expected 400, got ", c.recorder.Code)
		}
	}))
	t.Run("Returns body with message", testCase(func(t *testing.T, c *testContext) {
		c.router.ServeHTTP(c.recorder, httptest.NewRequest("GET", "/?quantity=not-a-number", nil))
		body := c.recorder.Body.String()
		if body != "Invalid quantity not-a-number" {
			t.Error("Returned body doesn't contain expected message: ", body)
		}
	}))
}

func TestOutOfBoundsQuantityParam(t *testing.T) {
	t.Run("Returns 400 status code", testCase(func(t *testing.T, c *testContext) {
		c.router.ServeHTTP(c.recorder, httptest.NewRequest("GET", "/?quantity=99999999", nil))
		if c.recorder.Code != 400 {
			t.Error("Expected 400, got ", c.recorder.Code)
		}
	}))
	t.Run("Returns body with message", testCase(func(t *testing.T, c *testContext) {
		c.router.ServeHTTP(c.recorder, httptest.NewRequest("GET", "/?quantity=-100", nil))
		body := c.recorder.Body.String()
		if body != "Quantity must be between 1 and 20000" {
			t.Error("Returned body doesn't contain expected message: ", body)
		}
	}))
}

func TestCors(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", "https://example.com")
	t.Run("No CORS applied when request has no origin", testCase(func(t *testing.T, c *testContext) {
		c.router.ServeHTTP(c.recorder, httptest.NewRequest("GET", "/", nil))
		if c.recorder.Code != 200 {
			t.Error("Expected 200, got ", c.recorder.Code)
		}
	}))
	t.Run("Returns Access-Control-Allow-Origin header when CORS request", testCase(func(t *testing.T, c *testContext) {
		request := httptest.NewRequest("GET", "/", nil)
		request.Host = "server1.example.com"
		request.Header.Add("Origin", "https://example.com")
		c.router.ServeHTTP(c.recorder, request)
		if c.recorder.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
			t.Error("Expected Access-Control-Allow-Origin header, got", c.recorder.Header().Get("Access-Control-Allow-Origin"))
		}
	}))
	t.Run("Returns 403 Forbidden status code when request has origin", testCase(func(t *testing.T, c *testContext) {
		request := httptest.NewRequest("GET", "/", nil)
		request.Header.Add("Origin", "https://different.com")
		c.router.ServeHTTP(c.recorder, request)
		if c.recorder.Code != 403 {
			t.Error("Expected 403, got ", c.recorder.Code)
		}
	}))
}
