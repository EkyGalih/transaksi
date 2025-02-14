package modeltransaksi

import (
	"net/http"

	"github.com/EkyGalih/transaksi/db"
	"github.com/EkyGalih/transaksi/entities"
	"github.com/google/uuid"
)

func FetchAll() (entities.Response, error) {
	var transaksi []entities.Transaction
	var res entities.Response

	con := db.CreateCon()

	rows := con.Find(&transaksi)
	if rows.Error != nil {
		return res, rows.Error
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = transaksi

	return res, nil
}

func Detail(id string) (entities.Response, error) {
	var transaksi entities.Transaction
	var res entities.Response

	con := db.CreateCon()

	row := con.Where("id = ?", id).First(&transaksi)
	if row.Error != nil {
		return res, row.Error
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = transaksi

	return res, nil
}

func Store(transaksi entities.Transaction) (entities.Response, error) {
	var res entities.Response

	con := db.CreateCon()

	if transaksi.ID == "" {
		transaksi.ID = uuid.New().String()
	}

	result := con.Create(&transaksi)
	if result.Error != nil {
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "Data Transaksi berhasil disimpan"

	return res, nil
}

func Update(id string, newTransaksi entities.Transaction) (entities.Response, error) {
	var res entities.Response

	con := db.CreateCon()

	// if err := con.First(&bidang, id).Error; err != nil {
	// 	return res, err
	// }

	// bidang.NamaBidang = newBidang.NamaBidang
	// bidang.KodeBidang = newBidang.KodeBidang
	// bidang.AliasBidang = newBidang.AliasBidang

	// if err := con.Save(&bidang).Error; err != nil {
	// 	return res, err
	// }

	if err := con.Model(&entities.Transaction{}).Where("id = ?", id).Updates(newTransaksi).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Data Transaksi berhasil diupdate"
	res.Data = newTransaksi

	return res, nil
}

func Delete(id string) (entities.Response, error) {
	var res entities.Response

	con := db.CreateCon()

	if err := con.Where("id = ?", id).Delete(&entities.Transaction{}).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Delete Success"

	return res, nil
}
