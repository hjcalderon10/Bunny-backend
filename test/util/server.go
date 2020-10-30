package util

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	validatorV10 "github.com/go-playground/validator/v10"
	"github.com/hjcalderon10/bunny-backend/util/validator"
	"github.com/labstack/echo"
)

func SetupServerTest(method string, endPoint string, payload io.Reader) (*echo.Echo, *http.Request, *httptest.ResponseRecorder, echo.Context) {
	sw := echo.New()
	sw.Validator = &validator.ApiValidator{Validator: validatorV10.New()}

	req := httptest.NewRequest(
		method,
		endPoint,
		payload,
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := sw.NewContext(req, rec)

	return sw, req, rec, c
}

func StringReader(payload string) io.Reader {
	return strings.NewReader(payload)
}
