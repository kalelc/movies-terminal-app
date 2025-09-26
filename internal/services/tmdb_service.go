package services

import (
	"log"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/kalelc/movies/internal/domain"
)

type TmdbService struct {
	TmdbClient *tmdb.Client
}

func NewTmdbService() *TmdbService {
	tmdbClient, err := tmdb.Init(os.Getenv("APIKEY"))

	if err != nil {
		log.Fatal(err)
	}

	return &TmdbService{TmdbClient: tmdbClient}
}

func (s *TmdbService) GetMovies() []domain.Movie {
	options := map[string]string{
		"language": "es-ES",
		"page":     "1",
	}
	popularMovies, err := s.TmdbClient.GetMoviePopular(options)
	if err != nil {
		log.Fatal("Error obteniendo películas populares:", err)
	}
	var movies []domain.Movie
	for _, movie := range popularMovies.Results {
		movies = append(
			movies,
			domain.Movie{
				Id:          movie.ID,
				Name:        movie.Title,
				Overview:    movie.Overview,
				ReleaseDate: movie.ReleaseDate,
				Popularity:  movie.Popularity,
				PosterPath:  movie.PosterPath,
				VoteAverage: movie.VoteAverage,
			},
		)
	}

	return movies
}
