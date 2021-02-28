package model

import (
	"errors"
	"gorm.io/gorm"
)

type DetailPengecekan struct {
	DetailPengecekanID		 		uint `gorm:"column:detail_pengecekan_id; PRIMARY_KEY" json:"-"`
	Nama		 		string `json:"nama"  validate:"required"`
	Tipe		 		string `json:"tipe"  validate:"required"`
	Catatan 			string `json:"catatan"  validate:""`
	Status	 			string `json:"status"  validate:""`
	DetailCheckListID	uint `json:"-"  validate:""`
	PengecekanID			uint `json:"-"  validate:""`
}

func (DetailPengecekan) TableName() string {
	return "detail_pengecekan"
}



func (data *DetailPengecekan) Create(db *gorm.DB) (interface{},error){
	result := db.Select("nama", "tipe", "catatan", "status", "detail_check_list_id", "pengecekan_id").Create(&data)
	if result.Error != nil {
		return nil,result.Error
	}
	if result.RowsAffected < 1 {
		return nil,errors.New("failed to create data")
	}
	return data,nil
}

