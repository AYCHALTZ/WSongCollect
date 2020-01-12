package db_storm

import (
	"time"

	"github.com/asdine/storm/q"
	"github.com/cloudsonic/sonic-server/domain"
)

type _Album struct {
	ID           string    ``
	Name         string    `storm:"index"`
	ArtistID     string    `storm:"index"`
	CoverArtPath string    ``
	CoverArtId   string    ``
	Artist       string    `storm:"index"`
	AlbumArtist  string    ``
	Year         int       `storm:"index"`
	Compilation  bool      ``
	Starred      bool      `storm:"index"`
	PlayCount    int       `storm:"index"`
	PlayDate     time.Time `storm:"index"`
	SongCount    int       ``
	Duration     int       ``
	Rating       int       `storm:"index"`
	Genre        string    ``
	StarredAt    time.Time `storm:"index"`
	CreatedAt    time.Time `storm:"index"`
	UpdatedAt    time.Time ``
}

type albumRepository struct {
	stormRepository
}

func NewAlbumRepository() domain.AlbumRepository {
	r := &albumRepository{}
	r.init(&_Album{})
	return r
}

func (r *albumRepository) Put(a *domain.Album) error {
	ta := _Album(*a)
	return Db().Save(&ta)
}

func (r *albumRepository) Get(id string) (*domain.Album, error) {
	ta := &_Album{}
	err := r.getByID(id, ta)
	if err != nil {
		return nil, err
	}
	a := domain.Album(*ta)
	return &a, err
}

func (r *albumRepository) FindByArtist(artistId string) (domain.Albums, error) {
	var albums []_Album
	err := r.execute(q.Eq("ArtistID", artistId), &albums)
	if err != nil {
		return nil, err
	}
	return r.toAlbums(albums)
}

func (r *albumRepository) GetAll(options ...domain.QueryOptions) (domain.Albums, error) {
	var all []_Album
	err := r.getAll(&all, options...)
	if err != nil {
		return nil, err
	}
	return r.toAlbums(all)
}

func (r *albumRepository) toAlbums(all []_Album) (domain.Albums, error) {
	result := make(domain.Albums, len(all))
	for i, a := range all {
		result[i] = domain.Album(a)
	}
	return result, nil
}

func (r *albumRepository) GetAllIds() ([]string, error) {
	var all []_Album
	err := r.getAll(&all)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(all))
	for i, a := range all {
		result[i] = domain.Album(a).ID
	}
	return result, nil
}

func (r *albumRepository) PurgeInactive(activeList domain.Albums) ([]string, error) {
	return r.purgeInactive(activeList)
}

func (r *albumRepository) GetStarred(options ...domain.QueryOptions) (domain.Albums, error) {
	var starred []_Album
	err := r.execute(q.Eq("Starred", true), &starred, options...)
	if err != nil {
		return nil, err
	}
	return r.toAlbums(starred)
}

var _ domain.AlbumRepository = (*albumRepository)(nil)
var _ = domain.Album(_Album{})
