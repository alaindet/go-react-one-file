package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func MustJson(data any) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func SendJson(w http.ResponseWriter, statusCode int, data any) {
	jsonData := MustJson(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func Resp(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}

func CompareLower(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}
