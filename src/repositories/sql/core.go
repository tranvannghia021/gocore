package sql

import (
	"github.com/tranvannghia021/gocore/src/repositories"
	"gorm.io/gorm/clause"
)

type SCore struct {
	Smodel
}

func (s *SCore) UpdateOrCreate(coreModel *repositories.Core) *ResSql {
	result := s.getConnection().Clauses(clause.OnConflict{
		Columns: []clause.Column{{
			Name: "internal_id",
		}},
		DoUpdates: clause.AssignmentColumns([]string{"email", "first_name", "last_name", "avatar", "access_token"}),
	}).Clauses(clause.Returning{}).Create(coreModel)
	if result.Error != nil {
		return &ResSql{
			Status: false,
			Errors: result.Error,
		}
	}
	return &ResSql{
		Status: true,
		Data:   result,
		Errors: nil,
	}
}
