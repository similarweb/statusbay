package state

import (
	"fmt"
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
	ApplyId   string `gorm:"unique_index;not null"`
	Name      string `gorm:"not null"`
	Cluster   string `gorm:"not null"`
	Namespace string `gorm:"not null"`
	Status    string `gorm:"not null;type:varchar(12)"`
	Time      int64  `gorm:"not null"`
	DeployBy  string `gorm:"not DeployBy"`
	Details   string `gorm:"not null;type:json"`
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

// open will create a new DB connection
func open(username, password, dns, schema string, port int) (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dns, port, schema))
}

// NewMysqlClient create new MyySQL client
func NewMysqlClient(config *MySQLConfig) *MySQLManager {

	var db *gorm.DB

	c := make(chan int, 1)
	go func() {
		var err error
		for {
			db, err = open(config.Username, config.Password, config.DNS, config.Schema, config.Port)
			if err == nil {
				break
			}
			log.Warn("could not initialize connection to database, retrying for 5 seconds")
			time.Sleep(5 * time.Second)
		}
		c <- 1
	}()

	select {
	case <-c:
	case <-time.After(60 * time.Second):
		log.Fatal("could not connect database, timed out after 1 minute")
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

// MySQLConfig client configuration
type MySQLConfig struct {
	DNS      string `yaml:"dns" env:"STATUSBAY_MYSQL_HOST, overwrite"`
	Port     int    `yaml:"port" env:"STATUSBAY_MYSQL_PORT, overwrite"`
	Username string `yaml:"username" env:"STATUSBAY_MYSQL_USERNAME, overwrite"`
	Password string `yaml:"password" env:"STATUSBAY_MYSQL_PASSWORD, overwrite"`
	Schema   string `yaml:"schema"  env:"STATUSBAY_MYSQL_SCHEMA, overwrite"`
}
