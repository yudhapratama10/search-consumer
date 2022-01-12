package repository

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
	"github.com/yudhapratama10/search-consumer/model"
)

var (
	fixtures = make(map[string]io.ReadCloser)
)

func init() {
	fixtureFiles, err := filepath.Glob("mock_data/*.json")
	if err != nil {
		panic(fmt.Sprintf("Cannot glob fixture files: %s", err))
	}

	for _, fpath := range fixtureFiles {
		f, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(fmt.Sprintf("Cannot read fixture file: %s", err))
		}
		fixtures[filepath.Base(fpath)] = ioutil.NopCloser(bytes.NewReader(f))
	}
}

func fixture(fname string) io.ReadCloser {
	out := new(bytes.Buffer)
	b1 := bytes.NewBuffer([]byte{})
	b2 := bytes.NewBuffer([]byte{})
	tr := io.TeeReader(fixtures[fname], b1)

	defer func() { fixtures[fname] = ioutil.NopCloser(b1) }()
	io.Copy(b2, tr)
	out.ReadFrom(b2)

	return ioutil.NopCloser(out)
}

func TestIndex(t *testing.T) {
	t.Parallel()

	// Create Mock for Elasticsearch Server
	mock := TransportMock{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
			Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
		},
	}
	mock.RoundTripFn = func(req *http.Request) (*http.Response, error) { return mock.Response, nil }

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Transport: &mock,
	})
	if err != nil {
		t.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	repo := NewFootballRepository(client)

	data := model.FootballClub{
		Id:          1,
		Name:        "FC Barcelona",
		Tournaments: []string{"Laliga", "Uefa Champions League", "Copa Del Rey", "Spanish Supercup"},
		Nation:      "Spanyol",
		HasStadium:  true,
		Description: `Fútbol Club Barcelona, juga dikenal sebagai Barcelona atau Barça, adalah klub sepak bola profesional yang berbasis di Barcelona, Catalunya, Spanyol. Didirikan pada tahun 1899 oleh sekelompok pemain Swiss, Inggris, Jerman dan Katalan yang dipimpin oleh Joan Gamper, klub telah menjadi simbol budaya Catalan dan Catalanisme, yang mempunyai motto "Més que un club" (Lebih dari sebuah klub).`,
		Rating:      5.0,
	}

	t.Run("Success Insert Index", func(t *testing.T) {
		mock.Response = &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixture("insert_doc.json"),
		}

		_, err = repo.Insert(data)
		assert.NoError(t, err)
	})

	t.Run("Success Update Index", func(t *testing.T) {
		mock.Response = &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixture("update_doc.json"),
		}

		_, err = repo.Update(data)
		assert.NoError(t, err)
	})
}
