package state

import (
	"fmt"
	"statusbay/config"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQLManager create new MySQL client
type MySQLManager struct {
	DB *gorm.DB
}

// TableKubernetes define deployment table schema
type TableKubernetes struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Cluster   string `gorm:"not null"`
	Namespace string `gorm:"not null"`
	Status    string `gorm:"not null;type:varchar(12)"`
	Time      int64  `gorm:"not null"`
	// CreatedAt time.Time // TODO: need to change Time field to this field + creating a generic  tiny
	DeployBy string `gorm:"not DeployBy"`
	Details  string `gorm:"not null;type:json"`
}

// TableName set deployment name table
func (u *TableKubernetes) TableName() string {
	return "kubernetes"
}

// TableDeploymentsHash define deployment hash
type TableDeploymentsHash struct {
	Deployment string `gorm:"not null;primary_key:yes"`
	Hash       uint64 `gorm:"not null"`
}

// TableName set deployment hash name
func (u *TableDeploymentsHash) TableName() string {
	return "last_deployment_version"
}

// open will creates a new DB connection
func open(username, password, dns, schema string) (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dns, schema))
}

// NewMysqlClient create new MyySQL client
func NewMysqlClient(config *config.MySQLConfig) *MySQLManager {

<<<<<<< HEAD
	var err error
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.DNS, config.Port, config.Schema))
	if strings.ToLower(fmt.Sprintf("%s", log.GetLevel())) == "debug" {
		db.LogMode(true)
=======
	var db *gorm.DB

	c := make(chan int, 1)
	go func() {
		var err error
		for {
			db, err = open(config.Username, config.Password, config.DNS, config.Schema)
			if err == nil {
				break
			}
			log.Warn("Failed to open DB connection, sleeping for 5 second.")
			time.Sleep(5 * time.Second)
		}
		c <- 1
	}()

	select {
	case <-c:
	case <-time.After(60 * time.Second):
		log.Fatal("Failed to connect DB after 1 minute, time out.")
>>>>>>> ENH: Create DB connection retries
	}

	if strings.ToLower(fmt.Sprintf("%s", log.GetLevel())) == "debug" {
		db.LogMode(true)
	}

	return &MySQLManager{
		DB: db,
	}
}

// Migration create tables migration
func (my *MySQLManager) Migration() {
	my.DB.AutoMigrate(&TableKubernetes{})
	my.DB.AutoMigrate(&TableDeploymentsHash{})
}
