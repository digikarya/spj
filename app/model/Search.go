package model

import (
	"errors"
	"github.com/digikarya/helper"
	"gorm.io/gorm"
	"net/http"
)

type SearchRequest struct {
	Condition []struct{
		Column string `json:"column"  validate:"required,alpha"`
		Value string `json:"value"  validate:"required"`
	} `json:"condition"  validate:"required"`
}

type tmpKendaraan struct {
	KendaraanID string `json:"kendaraan_id"  validate:"required"`
	NoKendaraan string `json:"no_kendaraan"  validate:"required"`
	NoBody string `json:"no_body"  validate:"required"`
	Kategori string `json:"kategori"  validate:"required"`
	JumlahSeat string `json:"jumlah_seat"  validate:"required"`
}

func (tmpKendaraan) TableName() string {
	return "pengecekan_kendaraan"
}


func (payload *SearchRequest) KendaraanAvail(db *gorm.DB,r *http.Request)  (interface{},error) {
	err := payload.setPayload(r)
	if err != nil {
		return nil,errors.New("data tidak sesuai")
	}
	tmpData := []tmpKendaraan{}
	trayekID,err := helper.DecodeHash(payload.Condition[0].Value)
	//trans := db.Limit(limit).Find(&result)
	sql := `SELECT  DISTINCT no_body,
				kendaraan_hash_id 'kendaraan_id',no_kendaraan,kategori_kendaraan 'kategori' , jumlah_seat
			FROM pengecekan_kendaraan 
			WHERE status_kondisi_fisik = 'BAIK' AND status_kondisi_fisik = 'BAIK' AND used='0' AND trayek_id = ?`
	exec := db.Raw(sql, trayekID).Scan(&tmpData)
	if exec.Error != nil {
		return tmpData,exec.Error
	}
	return tmpData,nil
}

func (payload *SearchRequest) setPayload(r *http.Request)  (err error)  {
	if err := helper.DecodeJson(r,&payload);err != nil {
		return errors.New("invalid payload")
	}
	if err := helper.ValidateData(payload);err != nil {
		return err
	}
	if len(payload.Condition) < 1 {
		return errors.New("invalid payload")
	}
	return nil
}