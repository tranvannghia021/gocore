package response

import (
	"encoding/json"
	"net/http"
)

type Res struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func Response(data any, w http.ResponseWriter, err bool, e error) {
	w.Header().Set("Content-Type", "application/json")
	if err {

		json.NewEncoder(w).Encode(Res{
			Status:  false,
			Message: "Something Went wrong,Please try again.",
			Data:    nil,
			Error:   e.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(Res{
		Status:  true,
		Message: "Success",
		Data:    data,
		Error:   "",
	})
	return
}
