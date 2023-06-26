package sql

import (
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/vars"
	"gorm.io/gorm/clause"
)

var connect = vars.Connection

type ResSql struct {
	Status bool        `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors error       `json:"errors,omitempty"`
}

func Insert(coreModel *repositories.Core) *ResSql {
	result := connect.Clauses(clause.OnConflict{
		Columns: []clause.Column{{
			Name: "internal_id",
		}},
		DoUpdates: clause.AssignmentColumns([]string{"email", "first_name", "last_name", "avatar", "access_token"}),
	}).Clauses(clause.Returning{}).Create(coreModel)
	//result := Connection.Create(&coreModel)
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

func Update(coreModel repositories.Core) {
	connect.Model(&coreModel).Where("id", coreModel.ID).Save(&coreModel)
}

func getAll(coreModel repositories.Core) ResSql {
	results := connect.Find(&coreModel)
	if results.Error != nil {
		return ResSql{
			Status: false,
			Errors: results.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   results,
		Errors: nil,
	}
}

func First(coreModel repositories.Core) ResSql {
	result := connect.First(&coreModel)
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

func DeleteById(coreModel repositories.Core) ResSql {
	result := connect.Delete(&coreModel)
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
