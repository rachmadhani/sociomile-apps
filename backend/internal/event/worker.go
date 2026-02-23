package event

import (
	"encoding/json"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Worker struct {
	db *gorm.DB
}

func NewWorker(db *gorm.DB) *Worker {
	return &Worker{db: db}
}

func (w *Worker) Start(eventChan <-chan Event) {
	go func() {
		for e := range eventChan {
			payload, _ := json.Marshal(e.Payload)

			log := model.ActivityLog{
				ID:        uuid.New(),
				TenantID:  uuid.Must(uuid.Parse(e.TenantID)),
				EventType: e.EventType,
				EntityID:  uuid.Must(uuid.Parse(e.EntityID)),
				Payload:   string(payload),
			}

			if err := w.db.Create(&log).Error; err != nil {
				logger, logFile, err := utils.GenerateNewLogger("worker")
				logger.Printf("error create activity log: %v", err)
				logFile.Close()
			}
		}
	}()
}
