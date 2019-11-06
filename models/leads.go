package models

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Info      string  `json:"info"`
	SourceId  uint    `json:"source_id"`
	Profit    bool    `json:"profit"`
	UserId    uint    `json:"user_id"` //The user that this lead belongs to
	User      Account `gorm:"foreignkey:user_id"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (lead *Lead) Validate() (map[string]interface{}, bool) {

	if lead.FirstName == "" {
		return u.Message(false, "Lead first name should be on the payload"), false
	}
	if lead.LastName == "" {
		return u.Message(false, "Lead last name should be on the payload"), false
	}
	if lead.Info == "" {
		return u.Message(false, "Lead info should be on the payload"), false
	}
	if lead.SourceId <= 0 {
		return u.Message(false, "Source is not recognized"), false
	}
	if lead.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (lead *Lead) Create() (map[string]interface{}) {

	if resp, ok := lead.Validate(); !ok {
		return resp
	}

	GetDB().Create(lead)

	resp := u.Message(true, "success")
	resp["lead"] = lead
	return resp
}

func GetLead(id int) (*Lead) {

	lead := &Lead{}
	err := GetDB().Table("leads").Where("id = ?", id).First(lead).Error
	if err != nil {
		return nil
	}
	return lead
}

func GetLeads(user uint) ([]*Lead) {

	leads := make([]*Lead, 0)
	err := GetDB().Table("leads").Where("user_id = ?", user).Find(&leads).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return leads
}
