package sql

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm/clause"
)

type SCore struct {
	SBaseSql
}

func (s *SCore) UpdateOrCreate() ResSql {
	result := s.GetConnection().Clauses(clause.OnConflict{
		Columns: []clause.Column{{
			Name: "internal_id",
		}},
		DoUpdates: clause.AssignmentColumns([]string{"email", "email_verify_at", "password", "first_name", "last_name", "avatar", "gender", "status", "phone", "birth_day", "address", "refresh_token", "access_token", "expire_token", "is_disconnect", "domain", "raw_domain"}),
	}).Clauses(clause.Returning{}).Create(s.GetModel())
	if result.Error != nil {
		return ResSql{
			Status: false,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   result,
		Errors: nil,
	}
}
