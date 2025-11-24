package utils

import (
	"fmt"
	"net/http"
)

func PureseJSON(r *http.Request , payload any) error {
	if r.Body ==nil{
		return fmt.Errorf("missing request body")
	}
	
	return json.NewDecoder(r.Body).Decode(payload)

}