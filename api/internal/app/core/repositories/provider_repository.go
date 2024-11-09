package repositories

import (
	"log"

	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
	"github.com/ossipesonen/go-traffic-lights/pkg/database"
)

type ProviderRepository struct {
	db     database.Database
	logger *log.Logger
	dbName string
}

func NewProviderRepository(db database.Database, logger *log.Logger) *ProviderRepository {
	return &ProviderRepository{
		db:     db,
		logger: logger,
		dbName: "providers",
	}
}

func (p *ProviderRepository) List() (*[]models.Provider, error) {
	providers := []models.Provider{}
	rows, err := p.db.Handle().Query("select id, name from providers")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)

		if err != nil {
			return nil, err
		}

		p := models.Provider{Id: id, Name: name}
		providers = append(providers, p)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return &providers, nil
}

func (p *ProviderRepository) Read(id int) (*models.Provider, error) {
	var providerId int
	var providerName string
	err := p.db.Handle().QueryRow("select id, name from providers where id = $1", id).Scan(&providerId, &providerName)

	if err != nil {
		return nil, err
	}

	return &models.Provider{
		Id:   providerId,
		Name: providerName,
	}, nil
}
