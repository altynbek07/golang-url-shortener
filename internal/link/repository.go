package link

import (
	"go/adv-demo/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: db,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) GetById(id uint64) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint64) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
