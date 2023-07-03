package sql

import (
	"github.com/tranvannghia021/gocore/vars"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBaseSql interface {
	GetALL() ResSql
	Create() ResSql
	Update() ResSql
	Delete() ResSql
	First() ResSql
	Last() ResSql
	GetModel() interface{}
	GetConnection() *gorm.DB
	UpdateOrCreate() ResSql
}

type SBaseSql struct {
	ModelBase interface{}
	baseDB    *gorm.DB
}

type ResSql struct {
	Status bool        `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors error       `json:"errors,omitempty"`
}

func (s *SBaseSql) GetModel() interface{} {
	return s.ModelBase
}

func (s *SBaseSql) GetConnection() *gorm.DB {
	s.baseDB = vars.Connection
	return s.baseDB
}

func (s *SBaseSql) GetALL() ResSql {
	result := s.baseDB.Find(s.GetModel())
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   s.ModelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Create() ResSql {
	result := s.GetConnection().Clauses(clause.Returning{}).Create(s.GetModel())
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   s.ModelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Update() ResSql {
	result := s.GetConnection().Model(s.GetModel()).Updates(s.GetModel())
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   s.ModelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Delete() ResSql {
	result := s.GetConnection().Delete(s.GetModel())
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   s.ModelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) First() ResSql {
	result := s.GetConnection().First(s.GetModel())
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   s.ModelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Last() ResSql {
	result := s.GetConnection().Last(s.GetModel())
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   s.ModelBase,
		Errors: nil,
	}
}
func (s *SBaseSql) UpdateOrCreate() ResSql {
	result := s.GetConnection().Clauses(clause.OnConflict{
		Columns: []clause.Column{{
			Name: "internal_id",
		}},
		DoUpdates: clause.AssignmentColumns([]string{}),
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
