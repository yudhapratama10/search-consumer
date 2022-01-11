package usecase

import (
	model "github.com/yudhapratama10/search-consumer/model"
)

func (usecase *footballUsecase) Insert(footballClub model.FootballClub) (model.FootballClub, error) {
	resp, err := usecase.repo.Insert(footballClub)
	if err != nil {
		return model.FootballClub{}, err
	}

	return resp, nil
}

func (usecase *footballUsecase) Update(footballClub model.FootballClub) (model.FootballClub, error) {
	resp, err := usecase.repo.Update(footballClub)
	if err != nil {
		return model.FootballClub{}, err
	}

	return resp, nil
}

func (usecase *footballUsecase) Delete(footballClub model.FootballClub) (model.FootballClub, error) {
	resp, err := usecase.repo.Delete(footballClub)
	if err != nil {
		return model.FootballClub{}, err
	}

	return resp, nil
}
