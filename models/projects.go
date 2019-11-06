package models

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type Project struct {
	gorm.Model
	LeadId uint            `json:"lead_id"`
	Lead   Lead            `gorm:"foreignkey:lead_id"`
	Name   string          `json:"name"`
	Status uint            `json:"status"`
	Sum    decimal.Decimal `json:"sum"`
	Info   string          `json:"info"`
	UserId uint            `json:"user_id"`
	User   Account         `gorm:"foreignkey:user_id"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (project *Project) Validate() (map[string]interface{}, bool) {

	if project.LeadId <= 0 {
		return u.Message(false, "Lead is not recognized"), false
	}
	if project.Name == "" {
		return u.Message(false, "Project name should be on the payload"), false
	}
	if project.Status == 0 {
		return u.Message(false, "Status is not recognized"), false
	}
	if project.UserId <= 0 {
		return u.Message(false, "UserId is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (project *Project) Create() (map[string]interface{}) {

	if resp, ok := project.Validate(); !ok {
		return resp
	}

	GetDB().Create(project)

	resp := u.Message(true, "success")
	resp["project"] = project
	return resp
}

func GetProject(id uint) (*Project) {

	project := &Project{}
	err := GetDB().Table("projects").Where("id = ?", id).First(project).Error
	if err != nil {
		return nil
	}
	return project
}

func GetProjects(user uint) ([]*Project) {

	projects := make([]*Project, 0)
	err := GetDB().Table("projects").Where("user_id = ?", user).Find(&projects).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return projects
}
