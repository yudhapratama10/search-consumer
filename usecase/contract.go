package usecase

import (
	"github.com/yudhapratama10/search-consumer/model"
	"github.com/yudhapratama10/search-consumer/repository"
)

type footballUsecase struct {
	repo repository.FootballRepositoryContract
}

type FootballUsecaseContract interface {
	Insert(footballClub model.FootballClub) (model.FootballClub, error)
	Update(footballClub model.FootballClub) (model.FootballClub, error)
	Delete(footballClub model.FootballClub) (model.FootballClub, error)
}

func NewFootballClubUsecase(repo repository.FootballRepositoryContract) FootballUsecaseContract {
	return &footballUsecase{
		repo: repo,
	}
}
