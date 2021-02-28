package model

import (
	"errors"
	"github.com/digikarya/helper"
	"gorm.io/gorm"
	"net/http"
)

type SpjsRequest struct {
	JadwalID string `json:"jadwal_id"  validate:""`
	KendaraanID string `json:"kendaraan_id"  validate:""`
	SupirID string `json:"supir_id"  validate:""`
	KondekturID string `json:"kondektur_id"  validate:""`
	TglKeberangkatan string `json:"tgl_keberangkatan"  validate:""`
}
type SpjModel struct {
	SpjID 	uint `gorm:"column:spj_id; PRIMARY_KEY" json:"-"`
	HashID string `json:"id"  validate:""`
	KodeSPJ string `json:"kode_spj"  validate:""`
	Petugas uint `json:"petugas"  validate:""`
	NamaPetugas string `json:"nama_petugas"  validate:""`
	Sopir  string `json:"sopir"  validate:""`
	Kondektur  string `json:"kondektur"  validate:""`
	JadwalID  uint `json:"jadwal_id"  validate:""`
	WaktuKeberangkatan string `json:"waktu_keberangkatan"  validate:""`
	WaktuKedatangan string `json:"waktu_kedatangan"  validate:""`
	Asal string `json:"asal"  validate:""`
	Tujuan string `json:"tujuan"  validate:""`
	NoBody string `json:"no_body"  validate:""`
	KendaraanID string `json:"-"  validate:""`
	KendaraanHashID  string `json:"kendaraan_id"  validate:""`
	NoKendraan  string `json:"no_kendaraan"  validate:""`
	KategoriKendaraan string `json:"kategori_kendaraan"  validate:""`
	Kapasitas string `json:"kapasitas"  validate:""`
	TotalPenumpang string `json:"total_penumpang"  validate:""`
	NamaSupir string `json:"nama_supir"  validate:""`
	NamaKondektur string `json:"nama_kondektur"  validate:""`
	TglKeberangkatan string `json:"tanggal_keberangkatan"  validate:""`
	helper.TimeModel
}

func (data *SpjModel) All(db *gorm.DB,r *http.Request,string ...string) (interface{}, error) {
	 result := []SpjModel{}
		if err := helper.DecodeJson(r,&data);err != nil {
		return result,errors.New("invalid payload")
	}
	//getJWT := helper.TokenJWT{}
	//decodeJWT,_,err :=helper.DecodeJWT(&getJWT,r)
	//if err != nil {
	//	return nil,err
	//}
	//KaryawanID,err := helper.DecodeHash(fmt.Sprintf("%v", decodeJWT["user"]))
	trans := db.Where("tgl_keberangkatan > ?",1).Find(&result)
	if trans.Error != nil {
		return result,trans.Error
	}
	return result,nil
}
