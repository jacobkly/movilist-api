package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
)

type Movie struct {
	MovieID             int              `db:"movie_id" json:"movie_id"`
	Adult               bool             `db:"adult" json:"adult"`
	BackdropPath        *string          `db:"backdrop_path" json:"backdrop_path"`
	BelongsToCollection *json.RawMessage `db:"belongs_to_collection" json:"belongs_to_collection"`
	Budget              int              `db:"budget" json:"budget"`
	Genres              *json.RawMessage `db:"genres" json:"genres"`
	Homepage            *string          `db:"homepage" json:"homepage"`
	IMDBID              *string          `db:"imdb_id" json:"imdb_id"`
	OriginalLanguage    *string          `db:"original_language" json:"original_language"`
	OriginalTitle       *string          `db:"original_title" json:"original_title"`
	Overview            *string          `db:"overview" json:"overview"`
	Popularity          float64          `db:"popularity" json:"popularity"`
	PosterPath          *string          `db:"poster_path" json:"poster_path"`
	ProductionCompanies *json.RawMessage `db:"production_companies" json:"production_companies"`
	ProductionCountries *json.RawMessage `db:"production_countries" json:"production_countries"`
	ReleaseDate         *time.Time       `db:"release_date" json:"release_date"`
	Revenue             int64            `db:"revenue" json:"revenue"`
	Runtime             int              `db:"runtime" json:"runtime"`
	SpokenLanguages     *json.RawMessage `db:"spoken_languages" json:"spoken_languages"`
	Status              *string          `db:"status" json:"status"`
	Tagline             *string          `db:"tagline" json:"tagline"`
	Title               *string          `db:"title" json:"title"`
	Video               bool             `db:"video" json:"video"`
	VoteAverage         float64          `db:"vote_average" json:"vote_average"`
	VoteCount           int              `db:"vote_count" json:"vote_count"`
	CreatedAt           time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time        `db:"updated_at" json:"updated_at"`
}

type MovieCollection struct {
	CollectionID int       `db:"collection_id" json:"collection_id"`
	MovieID      int       `db:"movie_id" json:"movie_id"`
	Name         string    `db:"name" json:"name"`
	PosterPath   *string   `db:"poster_path" json:"poster_path"`
	VoteAverage  *float64  `db:"vote_average" json:"vote_average"`
	Position     *int      `db:"position" json:"position"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type MovieRepository struct {
	db *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetByTMDBID(ctx context.Context, tmdbID int) (*Movie, error) {
	var movie Movie
	err := r.db.GetContext(ctx, &movie,
		`select * from movies where movie_id = $1`, tmdbID)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &movie, err
}

func (r *MovieRepository) GetMovieIDByMediaID(ctx context.Context, mediaID int) (int, error) {

	var movieID int

	err := r.db.GetContext(ctx, &movieID, `
		select id
		from media_index
		where media_id = $1 and media_type = 'movie'
	`, mediaID)

	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return movieID, nil
}

func (r *MovieRepository) InsertMovie(ctx context.Context, movie *Movie) error {
	_, err := r.db.NamedExecContext(ctx, `
		insert into movies (
			movie_id, adult, backdrop_path, belongs_to_collection,
			budget, genres, homepage, imdb_id,
			original_language, original_title, overview,
			popularity, poster_path, production_companies,
			production_countries, release_date, revenue,
			runtime, spoken_languages, status,
			tagline, title, video, vote_average, vote_count
		) values (
			:movie_id, :adult, :backdrop_path, :belongs_to_collection,
			:budget, :genres, :homepage, :imdb_id,
			:original_language, :original_title, :overview,
			:popularity, :poster_path, :production_companies,
			:production_countries, :release_date, :revenue,
			:runtime, :spoken_languages, :status,
			:tagline, :title, :video, :vote_average, :vote_count
		)
	`, movie)

	return err
}

func (r *MovieRepository) EnsureMediaIndex(ctx context.Context, movieID int) error {
	_, err := r.db.ExecContext(ctx, `
		insert into media_index (id, media_type)
		values ($1, 'movie')
		on conflict (id, season_number, media_type) do nothing
	`, movieID)

	return err
}

func (r *MovieRepository) GetCollectionIDByMovieID(ctx context.Context, movieID int) (int, error) {

	var collectionID int

	err := r.db.GetContext(ctx, &collectionID, `
		select collection_id
		from movie_collections
		where movie_id = $1
		limit 1
	`, movieID)

	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return collectionID, nil
}

func (r *MovieRepository) GetCollectionByCollectionID(
	ctx context.Context,
	collectionID int,
) ([]MovieCollection, error) {

	var movies []MovieCollection

	err := r.db.SelectContext(ctx, &movies, `
		select *
		from movie_collections
		where collection_id = $1
		order by position asc nulls last
	`, collectionID)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return movies, err
}

func (r *MovieRepository) InsertMovieCollectionBatch(
	ctx context.Context,
	entries []MovieCollection,
) error {

	if len(entries) == 0 {
		return nil
	}

	_, err := r.db.NamedExecContext(ctx, `
		insert into movie_collections (
			collection_id,
			movie_id,
			name,
			poster_path,
			vote_average,
			position
		)
		values (
			:collection_id,
			:movie_id,
			:name,
			:poster_path,
			:vote_average,
			:position
		)
		on conflict do nothing
	`, entries)

	return err
}
