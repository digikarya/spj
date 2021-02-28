package handler

import (
	"github.com/digikarya/helper"
	"github.com/digikarya/spj/app/model"
	"gorm.io/gorm"
	"net/http"
)

func PengecekanCreate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	serv := model.Pengecekan{}
	data,err := serv.Create(db,r)
	if err != nil {
		helper.RespondJSONError(w, http.StatusBadRequest, err)
		return
	}
	helper.RespondJSON(w, "Success",http.StatusOK, data)
	return
}


func PengecekanAll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	serv := model.PengecekanModel{}
	data,err := serv.All(db,r)
	if err != nil {
		helper.RespondJSONError(w, http.StatusBadRequest, err)
		return
	}
	helper.RespondJSON(w, "Success",http.StatusOK, data)
	return
}