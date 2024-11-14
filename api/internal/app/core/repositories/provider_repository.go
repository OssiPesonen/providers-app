package repositories

import (
	"errors"
	"log"

	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
	"github.com/ossipesonen/go-traffic-lights/pkg/database"
	"github.com/upper/db/v4"
)

type Identifier struct {
	Id int
}

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
	q := p.db.Handle().SQL().Select("id", "name", "city", "region", "line_of_business").From("providers")
	if err := q.All(&providers); err != nil {
		return nil, err
	}

	return &providers, nil
}

func (p *ProviderRepository) Read(id int) (*models.Provider, error) {
	provider := &models.Provider{}

	q := p.db.Handle().SQL().Select("id", "name", "city", "region", "line_of_business").From("providers")
	if err := q.One(&provider); err != nil {
		return nil, err
	}

	return provider, nil
}

func (p *ProviderRepository) Create(provider *models.Provider) (int, error) {
	success := p.db.Handle().SQL().InsertInto("providers").Values(provider).Returning("id").Iterator().Next(&provider)
	if !success {
		p.logger.Printf("an error occurred when inserting provider and returning id")
		return 0, errors.New("unable to insert entry and save to database")
	}

	return provider.Id, nil
}

func (p *ProviderRepository) Find(name string, city string) (*models.Provider, error) {
	provider := &models.Provider{}

	q := p.db.Handle().SQL().Select("id", "name", "city", "region", "line_of_business").From("providers").Where("name = ?", name).And("city = ?", city)
	if err := q.One(&provider); err != nil && err != db.ErrNoMoreRows {
		return nil, err
	}

	return provider, nil
}
