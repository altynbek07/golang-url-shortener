package stat

import (
	"go/adv-demo/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	Database *db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Database: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	currentDate := datatypes.Date(time.Now())
	var stat Stat
	repo.Database.DB.First(&stat, "link_id = ? AND date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Database.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		repo.Database.DB.Save(&stat)
	}
}
