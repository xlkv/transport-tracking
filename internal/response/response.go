package response

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, r *http.Request, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, code int, message string) {
    WriteJSON(w, nil, code, map[string]interface{}{
        "success": false,
        "error":   message,
    })
}

func Success(w http.ResponseWriter, code int, data any) {
    WriteJSON(w, nil, code, map[string]interface{}{
        "success": true,
        "data":    data,
    })
}