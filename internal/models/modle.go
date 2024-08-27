package models

import (
	"fmt"
	"wecat/common/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(s,
			databaseSetting.UserName,
			databaseSetting.Passworld,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// if global.ServerSetting.RunMode == "debug" {
	// 	db.LogMode(true)
	// }
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// func updateTimeStampForCreateCallback(db *gorm.DB) {
// 	if db.Error == nil {
// 		nowTime := time.Now().Unix()
// 		if createTimeField, ok := db.Statement.Schema.FieldsByName["CreateOn"]; ok {
// 			createTimeField.Set(nowTime)
// 		}
// 	}
// }

// func updateTimeStampForUpdateCallback(db *gorm.DB) {

// }

// func deleteCallback(db *gorm.DB) {

// }
