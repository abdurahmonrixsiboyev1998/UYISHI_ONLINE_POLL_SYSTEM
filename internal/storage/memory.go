package storage

import (
	"online/internal/models"
	"sync"
)

var (
	Surveys   = make(map[string]*models.Survey)
	Users     = make(map[string]string)         
	SurveyMux sync.Mutex                        
)
