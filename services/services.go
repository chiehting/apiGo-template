package services

import (
	"github.com/chiehting/apiGo-template/models"
	"github.com/chiehting/apiGo-template/pkg/config"
	"github.com/chiehting/apiGo-template/pkg/log"
	"github.com/chiehting/apiGo-template/pkg/storage"
	"golang.org/x/sync/singleflight"
)

type service struct{}

// Service is to provide settings configuration
var Service *service

var _sf singleflight.Group

// Init is initialization when the service started
func (service *service) Init() {
	var cfg = config.GetApplication()

	storage.Migration()
	service.setupDatabase(cfg)
	service.setupCache(cfg)
}

func (service *service) setupDatabase(cfg *config.Application) bool {
	var settings models.Settings
	var settingData = []struct {
		id    uint
		value string
	}{
		{id: 1, value: cfg.Name},
		{id: 2, value: cfg.Version},
	}

	for _, setting := range settingData {
		settings.ID = setting.id
		settings.Value = setting.value
		if err := settings.UpdateValue(); err != nil {
			log.Error(err)
			return false
		}
	}

	return true
}

func (service *service) setupCache(cfg *config.Application) bool {
	var settings = models.Settings{}
	settings.DelCache()

	var rows, _ = settings.GetAll()
	if err := settings.SetCache(rows); len(rows) > 0 && err != nil {
		return false
	}
	return true
}

// GetSettingsByClassSubject is get application settings
func (service *service) GetSettingsByClassSubject(class string, subject string) (interface{}, error) {
	var result interface{}
	var err error
	var settings models.Settings
	settings.Class = class
	settings.Subject = subject

	if result := settings.GetCache(); len(result) == 0 {
		return result, nil
	}

	result, err, _ = _sf.Do(subject, func() (i interface{}, err error) {
		settings.Get()
		settings.SetCache([]models.Settings{settings})
		return settings.Value, err
	})

	return result, err
}
