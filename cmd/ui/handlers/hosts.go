package handlers

import (
	"encoding/json"
	"net/http"
	"seems.cloud/badwolf/server/api/hosts"
)

func HostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		jsonBytes, err := json.Marshal(hosts.GetHosts())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	default:
		http.Error(w, "Unsupported request method", http.StatusBadRequest)
	}
}
