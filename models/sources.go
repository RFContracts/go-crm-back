package models

import (
	u "../utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Source struct {
	gorm.Model
	Name string `json:"name"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (source *Source) Validate() (map[string] interface{}, bool) {

	if source.Name == "" {
		return u.Message(false, "Source name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (source *Source) Create() (map[string] interface{}) {

	if resp, ok := source.Validate(); !ok {
		return resp
	}

	GetDB().Create(source)

	resp := u.Message(true, "success")
	resp["source"] = source
	return resp
}

func GetSource(id uint) (*Source) {

	source := &Source{}
	err := GetDB().Table("sources").Where("id = ?", id).First(source).Error
	if err != nil {
		return nil
	}
	return source
}

func GetSources() ([]*Source) {

	sources := make([]*Source, 0)
	err := GetDB().Table("sources").Find(&sources).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return sources
}