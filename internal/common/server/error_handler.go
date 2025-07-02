package server

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func defaultHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	log.Println("Got error processing a request: ", err)

	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}

		var message string
		switch m := he.Message.(type) {
		case string:
			message = m
		case error:
			message = m.Error()
		default:
			message = "internal error"
		}

		err = fmt.Errorf("%s", message)
	}

	err = ReturnError(c, he.Code, err)
	if err != nil {
		log.Println("Got error processing an error:", err)
	}
}
