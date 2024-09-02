package model

import (
	"fmt"
	"wecat/common/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primaryKey;column:id" json:"id"`
	CreatedBy  string `gorm:"column:created_by" json:"created_by"`
	ModifiedBy string `gorm:"column:modified_by" json:"modified_by"`
	CreatedOn  uint32 `gorm:"column:created_on" json:"created_on"`
	ModifiedOn uint32 `gorm:"column:modified_on" json:"modified_on"`
	DeletedOn  uint32 `gorm:"column:deleted_on" json:"deleted_on"`
	IsDel      uint8  `gorm:"column:is_del" json:"is_del"`
}

// type Model struct {
// 	*gorm.Model
// 	CreatedBy string `json:"created_by"`
// 	UpdatedBy string `json:"updated_by"`
// 	DeletedBy string `json:"deleted_by"`
// 	IsDel     uint8  `json:"is_del"`
// }

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Passworld,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
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
