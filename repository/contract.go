package repository

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/yudhapratama10/search-consumer/model"
)

type footballRepository struct {
	client *elasticsearch.Client
}

type FootballRepositoryContract interface {
	Insert(data model.FootballClub) (model.FootballClub, error)
	Update(data model.FootballClub) (model.FootballClub, error)
	Delete(data model.FootballClub) (model.FootballClub, error)
}

func NewFootballRepository(client *elasticsearch.Client) FootballRepositoryContract {
	return &footballRepository{
		client: client,
	}
}
