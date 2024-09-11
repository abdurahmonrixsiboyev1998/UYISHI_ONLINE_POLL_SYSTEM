package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online/internal/models"
	"online/internal/storage"
	"sync"
)

var surveyIDCounter = 1
var surveyMutex sync.Mutex

func CreateSurvey(w http.ResponseWriter, r *http.Request) {
	var survey models.Survey
	if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	surveyMutex.Lock()
	survey.ID = fmt.Sprintf("survey-%d", surveyIDCounter)
	surveyIDCounter++
	survey.Options = make(map[string]int)
	survey.Votes = make(map[string]string)
	surveyMutex.Unlock()

	storage.SurveyMux.Lock()
	storage.Surveys[survey.ID] = &survey
	storage.SurveyMux.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(survey)
}

func Vote(w http.ResponseWriter, r *http.Request) {
	surveyID := r.URL.Query().Get("id")
	option := r.URL.Query().Get("option")
	ip := r.RemoteAddr

	storage.SurveyMux.Lock()
	survey, exists := storage.Surveys[surveyID]
	storage.SurveyMux.Unlock()

	if !exists {
		http.Error(w, "Survey not found", http.StatusNotFound)
		return
	}

	if _, voted := survey.Votes[ip]; voted {
		http.Error(w, "You have already voted", http.StatusForbidden)
		return
	}

	surveyMutex.Lock()
	survey.Options[option]++
	survey.Votes[ip] = option
	surveyMutex.Unlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Voted successfully")
}

func GetSurveyResults(w http.ResponseWriter, r *http.Request) {
	surveyID := r.URL.Query().Get("id")

	storage.SurveyMux.Lock()
	survey, exists := storage.Surveys[surveyID]
	storage.SurveyMux.Unlock()

	if !exists {
		http.Error(w, "Survey not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(survey.Options)
}
