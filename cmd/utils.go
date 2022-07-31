package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/go-resty/resty/v2"
	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/spf13/viper"
)

type WebexError struct {
	Message string `json:"message,omitempty"` // Error message from Webex
}

func NewWebexTeamsClient() *webexteams.Client {
	wc := webexteams.NewClient()
	token := viper.GetString("token")
	wc.SetAuthToken(token)
	return wc
}

func StringSliceContains(s []string, v string, insensitive bool) bool {
	for _, a := range s {
		if a == v || (insensitive && strings.EqualFold(a, v)) {
			return true
		}
	}
	return false
}

func printStructSliceAsTable(s []interface{}, fields []string) {
	if len(s) <= 0 {
		fmt.Println("Nothing to display")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

	headers := make([]string, 0)
	values := make([][]string, len(s))
	v := reflect.ValueOf(s[0])
	typ := v.Type()
	for i, n := 0, typ.NumField(); i < n; i++ {
		if len(fields) == 0 || StringSliceContains(fields, typ.Field(i).Name, true) {
			headers = append(headers, typ.Field(i).Name)
		}
	}
	for n, i := range s {
		v := reflect.ValueOf(i)
		for _, header := range headers {
			values[n] = append(values[n], fmt.Sprint(v.FieldByName(header).Interface()))
		}
	}
	fmt.Fprint(w, strings.Join(headers, "\t")+"\n")
	for _, v := range values {
		fmt.Fprint(w, strings.Join(v, "\t")+"\n")
	}

	w.Flush()
}

func checkWebexError(resp resty.Response) error {
	if resp.IsSuccess() {
		return nil
	}
	var err error
	if resp.StatusCode() == 401 {
		return errors.New("webex returned a 401, please update your token using saveToken command")
	}
	body := resp.Body()
	we := &WebexError{}
	err = json.Unmarshal(body, &we)
	if err != nil {
		return fmt.Errorf("unable to unmarsharl the webex error: %v", err)
	}
	return fmt.Errorf("webex error: %v", we.Message)
}
