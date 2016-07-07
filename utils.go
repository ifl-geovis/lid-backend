package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func fileExists(fp string) bool {
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}
	return info.IsDir() == false
}

func dirExists(fp string) bool {
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func printError(w http.ResponseWriter, err interface{}) {
	printJson(w, err)
}

func printJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	json.NewEncoder(w).Encode(obj)
}
