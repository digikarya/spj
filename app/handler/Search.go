package handler

import (
	"github.com/digikarya/helper"
	"github.com/digikarya/spj/app/model"
	"gorm.io/gorm"
	"net/http"
)

func AvailKendaraan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	serv := model.SearchRequest{}
	data,err := serv.KendaraanAvail(db,r)
	if err != nil {
		helper.RespondJSONError(w, http.StatusBadRequest, err)
		return
	}
	helper.RespondJSON(w, "Found",http.StatusOK, data)
	return
}
