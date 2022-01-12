package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudhapratama10/search-consumer/model"
	"github.com/yudhapratama10/search-consumer/repository"
)

func TestInsert(t *testing.T) {

	t.Parallel()

	t.Run("Should be Success Insert", func(t *testing.T) {

		var (
			data = model.FootballClub{
				Id:          20,
				Name:        "Newcastle United",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup"},
				Rating:      2.5,
				Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			}
		)

		repo := new(repository.FootballMock)
		uc := NewFootballClubUsecase(repo)

		//Insert(data model.FootballClub) (model.FootballClub, error)
		repo.On("Insert", data).Return(data, nil)

		resp, err := uc.Insert(data)

		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("Should be Success Insert 2", func(t *testing.T) {

		var (
			data = model.FootballClub{
				Id:          20,
				Name:        "Manchester City",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup", "Uefa Champions League", "FA Community Shield"},
				Rating:      5,
				Description: "Manchester City Football Club (dikenal pula sebagai Man City atau The Citizens) adalah sebuah klub sepak bola profesional dari Inggris yang bermain di Liga Premier Inggris. Klub ini merupakan klub sekota dengan Manchester United dan bermarkas di Stadion Etihad, Manchester.",
			}
		)

		repo := new(repository.FootballMock)
		uc := NewFootballClubUsecase(repo)

		//Insert(data model.FootballClub) (model.FootballClub, error)
		repo.On("Insert", data).Return(data, nil)

		resp, err := uc.Insert(data)

		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})
}

func TestUpdate(t *testing.T) {

	t.Parallel()

	t.Run("Should be Success Update", func(t *testing.T) {

		var (
			data = model.FootballClub{
				Id:          20,
				Name:        "Newcastle United",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup"},
				Rating:      2.5,
				Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			}
		)

		repo := new(repository.FootballMock)
		uc := NewFootballClubUsecase(repo)

		//Insert(data model.FootballClub) (model.FootballClub, error)
		repo.On("Update", data).Return(data, nil)

		resp, err := uc.Update(data)

		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("Should be Success Update 2", func(t *testing.T) {

		var (
			data = model.FootballClub{
				Id:          20,
				Name:        "Manchester City",
				Nation:      "Inggris",
				Tournaments: []string{"English Premier League", "FA Cup", "Uefa Champions League", "FA Community Shield"},
				Rating:      5,
				Description: "Manchester City Football Club (dikenal pula sebagai Man City atau The Citizens) adalah sebuah klub sepak bola profesional dari Inggris yang bermain di Liga Premier Inggris. Klub ini merupakan klub sekota dengan Manchester United dan bermarkas di Stadion Etihad, Manchester.",
			}
		)

		repo := new(repository.FootballMock)
		uc := NewFootballClubUsecase(repo)

		//Insert(data model.FootballClub) (model.FootballClub, error)
		repo.On("Update", data).Return(data, nil)

		resp, err := uc.Update(data)

		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})
}
