package utils

import (
	"fmt"
	"net/http"
)

func UnwrapURLQuery(request *http.Request, key string) (string, error) {
	value := request.URL.Query().Get(key)
	if value == "" {
		return "", fmt.Errorf("%s is not defined in the query params", key)
	}

	return value, nil
}
