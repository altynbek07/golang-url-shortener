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

func (repo *StatRepository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse

	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks) as clicks"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks) as clicks"
	}

	repo.Database.
		Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period desc").
		Scan(&stats)

	return stats
}
