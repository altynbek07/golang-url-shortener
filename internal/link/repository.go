package link

import "go/adv-demo/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: db,
	}
}

func (rep *LinkRepository) Create(link Link) {

}
