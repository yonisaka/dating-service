package httphandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/gorilla/schema"
)

var jsonContent = regexp.MustCompile(`(?i)application/((\w|\.|-)+\+)?json(-seq)?`)

// Data context for http use case
type Data struct {
	Request *http.Request
}

// Cast casts data based on servcice type
func (d *Data) Cast(target any) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target %T cannot addressable, must pointer target", target)
	}

	return d.httpCast(target)
}

func (d *Data) httpCast(target any) error {
	if d.Request == nil {
		return fmt.Errorf("unable to cast http data, null request")
	}

	// httpCast transform request payload data
	// GET -> params-query-string
	// POST -> json-body
	if err := d.grabMethod(target); err != nil {
		return err
	}

	return nil
}

// Transform query-string into json struct
func (d *Data) transform(target any, src map[string][]string) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.SetAliasTag("url")

	if err := decoder.Decode(target, src); err != nil {
		return fmt.Errorf("unable to decode query string:%s", err.Error())
	}

	return nil
}

// Grab request method
// Take a destination source of struct
func (d *Data) grabMethod(target any) error {
	if d.Request.Method != http.MethodPost && d.Request.Method != http.MethodPut && d.Request.Method != http.MethodPatch {
		return d.transform(target, d.Request.URL.Query())
	}

	cType := d.Request.Header.Get("Content-Type")
	if strings.Contains(cType, "form-urlencoded") {
		if err := d.Request.ParseForm(); err != nil {
			return err
		}

		return d.transform(target, d.Request.PostForm)
	}

	if jsonContent.MatchString(cType) {
		return d.decodeJSON(d.Request.Body, target)
	}

	return fmt.Errorf("unsupported decode content-type=%s", cType)
}

func (d *Data) decodeJSON(body io.ReadCloser, dst any) error {
	if body == nil {
		return nil
	}

	if err := json.NewDecoder(body).Decode(dst); err != nil {
		return fmt.Errorf("unable decode request body, err:%s", err.Error())
	}

	return nil
}
