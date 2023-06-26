package sql

import (
	"github.com/tranvannghia021/gocore/vars"
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
}
type IBaseSubConfig interface {
	SetModel(model interface{}) *SBaseSql
	GetModel() interface{}
}

type Smodel struct {
	*SBaseSql
}

var modelBase interface{}

var baseDB = vars.Connection

type SBaseSql struct {
}
type ResSql struct {
	Status bool        `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors error       `json:"errors,omitempty"`
}

func (s *Smodel) SetModel(model interface{}) *SBaseSql {
	modelBase = model
	return s.SBaseSql
}
func (s *Smodel) GetModel() interface{} {
	return modelBase
}

func (s *Smodel) getConnection() *gorm.DB {
	return baseDB
}

func (s *SBaseSql) GetALL() ResSql {
	result := baseDB.Find(&modelBase)
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   modelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Create() ResSql {
	result := baseDB.Clauses(clause.Returning{}).Create(&modelBase)
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   modelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Update() ResSql {
	result := baseDB.Model(&modelBase).Updates(&modelBase)
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   modelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Delete() ResSql {
	result := baseDB.Delete(&modelBase)
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   modelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) First() ResSql {
	result := baseDB.First(&modelBase)
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   modelBase,
		Errors: nil,
	}
}

func (s *SBaseSql) Last() ResSql {
	result := baseDB.Last(&modelBase)
	if result.Error != nil {
		return ResSql{
			Status: false,
			Data:   nil,
			Errors: result.Error,
		}
	}
	return ResSql{
		Status: true,
		Data:   modelBase,
		Errors: nil,
	}
}
