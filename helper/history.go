package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"technical-test-backend/model/domain"

	"github.com/gertd/go-pluralize"
	"gorm.io/gorm"
)

const (
	HistoryDelete = "DELETE"
	HistoryUpdate = "UPDATE"
	HistoryCreate = "CREATE"
)

func CreateHistory(db *gorm.DB, model interface{}, action string, userId uint) error {
	// Getting the table name of the model.
	var tableName string
	pluralize := pluralize.NewClient()
	typeModel := reflect.TypeOf(model)
	val := reflect.ValueOf(model).Elem()
	if typeModel.Kind() == reflect.Ptr {
		tableName = typeModel.Elem().Name()
	} else {
		tableName = typeModel.Name()
	}
	tableName = pluralize.Plural(ToSnakeCase(tableName))

	// Creating a map of the model's fields and values.
	data := map[string]interface{}{}
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag
		name := ToSnakeCase(val.Type().Field(i).Name)
		value := val.Field(i).Interface()
		if name == "model" {

			modelCustom, ok := val.FieldByName("Model").Interface().(domain.Model)
			if ok {
				data["id"] = modelCustom.ID
				data["created_at"] = modelCustom.CreatedAt
				data["updated_at"] = modelCustom.UpdatedAt
			}

			modelGorm, ok := val.FieldByName("Model").Interface().(gorm.Model)
			if ok {
				data["id"] = modelGorm.ID
				data["created_at"] = modelGorm.CreatedAt
				data["updated_at"] = modelGorm.UpdatedAt
			}

		}
		if strings.Contains(string(tag), "gorm") && !strings.Contains(string(tag), "foreignKey") {
			if reflect.ValueOf(value).Kind() == reflect.Ptr {
				if action != "UPDATE" && name != "deleted_by_id" {
					if value == nil {
						data[name] = reflect.ValueOf(value).Elem().Interface()
					} else {
						data[name] = reflect.ValueOf(value)
					}
				}
			} else {
				data[name] = value
			}
		}
	}
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Creating a new history record in the database.
	err = db.Create(&domain.History{
		TableID:     fmt.Sprintf("%v", data["id"]),
		CreatedByID: userId,
		TableName:   tableName,
		Data:        string(json),
		Type:        action,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
