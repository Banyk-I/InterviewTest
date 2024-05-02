package response

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, data interface{}) {
	// сереалізація data in JSON і відправлення відразу через `w` і StatusOK
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"error": "Error encoding JSON"}`, http.StatusInternalServerError)
		return
	}
}
