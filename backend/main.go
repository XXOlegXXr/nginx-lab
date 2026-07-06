package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type StatusResponse struct {
	ServerName string `json:"server_name"`
	Time       string `json:"time"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "Go-API-Default"
	}

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] API запит на бекенді %s\n", time.Now().Format("15:04:05"), serverName)

		w.Header().Set("Content-Type", "application/json")
		
		response := StatusResponse{
			ServerName: serverName,
			Time:       time.Now().Format("15:04:05.000"),
		}

		json.NewEncoder(w).Encode(response)
	})

	fmt.Printf("API-сервер %s стартував на порті %s...\n", serverName, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Помилка:", err)
	}
}
