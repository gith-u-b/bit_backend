package models

type RedeemCode struct {
	ID            int    `gorm:"primary_key" json:"id"`
	Code          string `json:"code"`
	IsUsed        string `json:"is_used"`
	UsedByAddress string `json:"used_by_address"`
	UsedByDomain  string `json:"used_by_domain"`
	Comment       string `json:"comment"`
	TxHash        string `json:"tx_hash"`
}

func GetRedeemCodes(pageNum int, pageSize int, maps interface{}) (redeem_codes []RedeemCode) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&redeem_codes)

	return
}

func GetRedeemCodeTotal(maps interface{}) (count int) {
	db.Model(&RedeemCode{}).Where(maps).Count(&count)

	return
}

func AddRedeemCode(code string, comment string) bool {
	db.Create(&RedeemCode{
		Code:    code,
		Comment: comment,
	})

	return true
}

func DeleteRedeemCode(id int) bool {
	db.Where("id = ?", id).Delete(&RedeemCode{})

	return true
}
