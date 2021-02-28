package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/digikarya/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

type Pengecekan struct {
	KendaraanID 		string `json:"kendaraan_id"  validate:"required"`
	CheckListID 		string `json:"check_list_id"  validate:"required"`
	KodeSpj 			string `json:"kode_spj"  validate:""`
	Catatan 			string `json:"catatan"  validate:""`
	CheckList           []struct{
		ID 			string `json:"id"  validate:"required"`
		Value 		string `json:"value"  validate:"required"`
		Keterangan 	string `json:"keterangan"  validate:""`
	} `json:"detail"  validate:"required"`
}

type PengecekanAssociate struct{
	PengecekanKendaranID 		uint `gorm:"column:pengecekan_kendaraan_id; PRIMARY_KEY" json:"-"`
	HashID 		string `json:"id"  validate:""`
	KaryawanID 		uint `json:"karyawan_id"  validate:"required"`
	Teknisi 		string `json:"teknisi"  validate:"required"`
	Catatan 		string `json:"catatan"  validate:"required"`
	KodeSPJ 		string `json:"kode_spj"  validate:"required"`
	KendaraanID 		uint `json:"-"  validate:"required"`
	CheckListID 		uint `json:"-"  validate:"required"`
	NoBody 		string `json:"no_body"  validate:"required"`
	NoKendaraan 		string `json:"no_kendaraan"  validate:"required"`
	StatusKondisiFisik 		string `json:"status_kondisi_fisik"  validate:""`
	StatusKondisiSurat 		string `json:"status_kondisi_surat"  validate:""`
	TrayekID uint `json:"-"  validate:""`
	KategoriKendaraan string `json:"kategori_kendaraan"  validate:""`

	Detail 			[]DetailPengecekan `gorm:"foreignKey:pengecekan_id;references:pengecekan_kendaraan_id" json:"detail"  validate:"" `
}


type PengecekanModel struct {
	PengecekanKendaranID 		uint `gorm:"column:pengecekan_kendaraan_id; PRIMARY_KEY" json:"-"`
	HashID 		string `json:"id"  validate:""`
	KaryawanID 		uint `json:"-"  validate:"required"`
	Teknisi 		string `json:"teknisi"  validate:"required"`
	Catatan 		string `json:"catatan"  validate:"required"`
	KodeSPJ 		string `json:"kode_spj"  validate:"required"`
	KendaraanID 		uint `json:"-"  validate:"required"`
	CheckListID 		uint `json:"-"  validate:"required"`
	NoBody 		string `json:"no_body"  validate:"required"`
	NoKendaraan 		string `json:"no_kendaraan"  validate:"required"`
	StatusKondisiFisik 		string `json:"status_kondisi_fisik"  validate:""`
	StatusKondisiSurat 		string `json:"status_kondisi_surat"  validate:""`
	TrayekID uint `json:"-"  validate:""`
	KategoriKendaraan string `json:"kategori_kendaraan"  validate:""`
	KendaraanHashID string `json:"kendaraan_hash_id"  validate:"required"`
	LayoutID string `json:"layout_id"  validate:"required"`
	JumlahSeat string `json:"jumlah_seat"  validate:"required"`
	helper.TimeModel
}

func (PengecekanModel) TableName() string {
	return "pengecekan_kendaraan"
}

func (PengecekanAssociate) TableName() string {
	return "pengecekan_kendaraan"
}



func (data Pengecekan) Create(db *gorm.DB,r *http.Request) (interface{},error){
	err := data.setPayload(r)
	if err != nil {
		return nil, err
	}
	endpoint := helper.GetEndpoint().Kendaraan
	code,body,err := helper.Curl("GET",endpoint+"/kendaraan/"+data.KendaraanID,r,nil)
	if err != nil{
		return nil,err
	}
	if code != http.StatusOK {
		return nil,errors.New("kendaraan tidak ditemukan ")
	}
	var getKendaraan map[string]map[string]interface{}
	err = json.Unmarshal(body, &getKendaraan)
	if err != nil{
		return nil,err
	}
	code,body,err = helper.Curl("GET",endpoint+"/check_list/"+data.CheckListID,r,nil)
	if err != nil{
		return nil,err
	}
	if code != http.StatusOK {
		return nil,errors.New("check list tidak ditemukan ")
	}
	var getChecklist map[string]map[string]interface{}
	err = json.Unmarshal(body, &getChecklist)
	if err != nil{
		return nil,err
	}
	var getCheckListDetail []map[string]interface{}
	detail,err := json.Marshal(getChecklist["Data"]["detail"])
	if err != nil{
		return nil,err
	}
	err = json.Unmarshal(detail, &getCheckListDetail)
	if err != nil{
		return nil,err
	}
	list := make(map[string]map[string]string)
	for _,item := range getCheckListDetail{
		lista := make(map[string]string)
		lista["nama"] = fmt.Sprintf("%v", item["nama"])
		lista["tipe"] = fmt.Sprintf("%v", item["tipe"])
		list[fmt.Sprintf("%v", item["id"])] = lista
	}
	savePengecekan := PengecekanModel{}
	savePengecekan.KendaraanID,err = helper.DecodeHash(fmt.Sprintf("%v", data.CheckListID))
	if err != nil {
		return nil, err
	}

	savePengecekan.TrayekID,err = helper.DecodeHash(fmt.Sprintf("%v",getKendaraan["Data"]["trayek_id"] ))
	if err != nil {
		return nil, err
	}

	savePengecekan.CheckListID,err = helper.DecodeHash(fmt.Sprintf("%v", data.CheckListID))
	if err != nil {
		return nil, err
	}
	savePengecekan.StatusKondisiSurat = "BAIK"
	count,err := strconv.Atoi(fmt.Sprintf("%v", getKendaraan["Data"]["jumlah_surat_kadaluarsa"]))
	if count > 0 {
		savePengecekan.StatusKondisiSurat = "TIDAK BAIK"
	}
	savePengecekan.KategoriKendaraan = fmt.Sprintf("%v", getKendaraan["Data"]["kategori"])
	savePengecekan.KendaraanHashID = data.KendaraanID
	savePengecekan.StatusKondisiFisik = "BAIK"
	savePengecekan.NoBody = fmt.Sprintf("%v", getKendaraan["Data"]["no_body"])
	savePengecekan.NoKendaraan = fmt.Sprintf("%v", getKendaraan["Data"]["no_kendaraan"])
	savePengecekan.Catatan = data.Catatan
	savePengecekan.KodeSPJ = data.KodeSpj
	savePengecekan.JumlahSeat = fmt.Sprintf("%v", getKendaraan["Data"]["jumlah_seat"])
	savePengecekan.LayoutID = fmt.Sprintf("%v", getKendaraan["Data"]["layout_id"])
	getJWT := helper.TokenJWT{}
	decodeJWT,_,err :=helper.DecodeJWT(&getJWT,r)
	if err != nil {
		return nil,err
	}
	savePengecekan.Teknisi = fmt.Sprintf("%v", decodeJWT["name"])
	savePengecekan.KaryawanID,err = helper.DecodeHash(fmt.Sprintf("%v", decodeJWT["user"]))
	if err != nil {
		return nil, err
	}
	savePengecekan.GenerateTime(false)
	//nonlist := make(map[string]string)
	trx := db.Begin()
	result := trx.Select( "status_kondisi_surat","kendaraan_hash_id","layout_id","jumlah_seat","status_kondisi_fisik", "trayek_id","kategori_kendaraan","karyawan_id", "catatan", "used", "kode_spj", "kendaraan_id", "check_list_id", "no_body", "no_kendaraan", "teknisi", "created_at", "updated_at").Create(&savePengecekan)
	if result.Error != nil {
		trx.Rollback()
		return nil,result.Error
	}
	if result.RowsAffected < 1 {
		trx.Rollback()
		return nil,errors.New("failed to create data")
	}

	if  err = savePengecekan.updateHashId(trx,int(savePengecekan.PengecekanKendaranID)); err != nil{
		trx.Rollback()
		return nil,errors.New("failed to create data")
	}
	bad := false
	for _,tmpCheckList := range data.CheckList {
		if _, ok := list[tmpCheckList.ID]; ok {
			tmpCreateDetail := DetailPengecekan{}
			tmpCreateDetail.Nama = list[tmpCheckList.ID]["nama"]
			tmpCreateDetail.Tipe = list[tmpCheckList.ID]["tipe"]
			tmpCreateDetail.PengecekanID = savePengecekan.PengecekanKendaranID
			tmpCreateDetail.DetailCheckListID,err = helper.DecodeHash(tmpCheckList.ID)
			if err != nil {
				trx.Rollback()
				return nil, err
			}
			if tmpCheckList.Value == "0" {
				tmpCreateDetail.Status = "BAIK"
			}else {
				tmpCreateDetail.Status = "TIDAK BAIK"
				bad = true
			}
			if _,err = tmpCreateDetail.Create(trx); err != nil {
				trx.Rollback()
				return nil, err
			}

		}
	}
	if bad {
		result = trx.Select("used").Updates(&savePengecekan)
	}
	trx.Commit()
	tmp := PengecekanAssociate{}
	result = db.Preload(clause.Associations).Where("pengecekan_kendaraan_id",savePengecekan.PengecekanKendaranID).Find(&tmp)
	return tmp,nil
}


func (data *PengecekanModel) All(db *gorm.DB,r *http.Request,string ...string) (interface{}, error) {
	var result []PengecekanModel
	getJWT := helper.TokenJWT{}
	decodeJWT,_,err :=helper.DecodeJWT(&getJWT,r)
	if err != nil {
		return nil,err
	}
	KaryawanID,err := helper.DecodeHash(fmt.Sprintf("%v", decodeJWT["user"]))
	trans := db.Where("karyawan_id = ?",KaryawanID).Find(&result)
	if trans.Error != nil {
		return result,trans.Error
	}
	return result,nil
}



func (data *Pengecekan) setPayload(r *http.Request)  error  {
	if err := helper.DecodeJson(r,&data);err != nil {
		return errors.New("invalid payload")
	}
	if err := helper.ValidateData(data);err != nil {
		return err
	}
	return nil
}

func (data *PengecekanModel) updateHashId(db *gorm.DB, id int)  error {
	hashID,err := helper.EncodeHash(id)
	if err != nil {
		return err
	}
	//log.Print(tmp.DaerahID)
	response := db.Model(&data).Where("pengecekan_kendaraan_id",id).Update("hash_id", hashID)
	if response.Error != nil{
		return response.Error
	}
	if response.RowsAffected < 1 {
		return errors.New("gagal rubah id")
	}
	return nil
}


