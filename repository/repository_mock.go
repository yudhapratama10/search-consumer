package repository

import (
	"net/http"

	"github.com/stretchr/testify/mock"
	"github.com/yudhapratama10/search-consumer/model"
)

// Product Mock for Internal Layer (repo)

type TransportMock struct {
	Response    *http.Response
	RoundTripFn func(req *http.Request) (*http.Response, error)
}

func (t *TransportMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFn(req)
}

// ========================================================================================
// Product Mock for External Layer (uc)

type FootballMock struct {
	mock.Mock
}

func (s *FootballMock) Insert(data model.FootballClub) (model.FootballClub, error) {
	args := s.Called(data)

	return args.Get(0).(model.FootballClub), args.Error(1)
}

func (s *FootballMock) Update(data model.FootballClub) (model.FootballClub, error) {
	args := s.Called(data)

	return args.Get(0).(model.FootballClub), args.Error(1)
}

func (s *FootballMock) Delete(data model.FootballClub) (model.FootballClub, error) {
	args := s.Called(data)

	return args.Get(0).(model.FootballClub), args.Error(1)
}
