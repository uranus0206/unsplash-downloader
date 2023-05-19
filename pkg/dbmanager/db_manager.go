package dbmanager

import (
	// "gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动

	"time"
	log "unsplash-downloader/pkg/qlogger"

	"github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/gorm"
)

type DbManager struct {
	gormDB *gorm.DB
}

var Dbm *DbManager

const DbTopicKey = "DbTopicKey"
const DbEditorialKey = "DbEditorialKey"

type DbTopics struct {
	PRKey      string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	TopicsIndx uint64
	TopicsId   string
	PageOffset uint64
}

type DbEditorial struct {
	PRKey      string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	PageOffset uint64
}

func InitWithPath(path string) *DbManager {
	gormd, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	Dbm = &DbManager{
		gormDB: gormd,
	}

	return Dbm
}

func (db *DbManager) CreateTable() {
	db.gormDB.AutoMigrate(&DbTopics{}, &DbEditorial{})
}

func (db *DbManager) AddTopicRecord(record *DbTopics) {
	// Update if already exist in db, or create a new one
	// var found DbTopics
	var old DbTopics
	res := db.gormDB.First(&old)
	log.Println("db search: ", res.Error)
	if res.Error == gorm.ErrRecordNotFound ||
		res.Error == gorm.ErrInvalidValue {
		res := db.gormDB.Create(record)
		log.Println("Add new: ", res.Error)
	} else if res.Error == nil {
		// Update
		res := db.gormDB.Updates(record)
		log.Println("Update exist: ", record, ", err: ", res.Error)
	} else {
		log.Errorln("DB search error: ", res.Error)
	}
}

func (db *DbManager) AddEditorialRecord(record *DbEditorial) {
	// Update if already exist in db, or create a new one
	var found DbEditorial
	res := db.gormDB.First(&found).Where("pr_key = ?", record.PRKey)
	log.Println("db search: ", res.Error)
	if res.Error == gorm.ErrRecordNotFound ||
		res.Error == gorm.ErrInvalidValue {
		res := db.gormDB.Create(record)
		log.Println("Add new: ", res.Error)
	} else if res.Error == nil {
		// Update
		res := db.gormDB.Model(&DbEditorial{}).
			First(&found).Where("pr_key = ?", record.PRKey).
			Updates(record)
		log.Println("Update exist: ", res.Error)
	} else {
		log.Errorln("DB search error: ", res.Error)
	}
}

func (db *DbManager) SearchTopicsByKey() (DbTopics, error) {
	var found DbTopics

	res := db.gormDB.First(&found).Where("pr_key = ?", DbTopicKey)

	log.Println("SearchTopicsByKey: ", res.Error, found)

	return found, res.Error
}

func (db *DbManager) SearchEditorialByKey() (DbEditorial, error) {
	var found DbEditorial

	res := db.gormDB.First(&found).Where("pr_key = ?", DbEditorialKey)

	log.Println("SearchEditorialByKey: ", res.Error, found)

	return found, res.Error
}
