package state

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"

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

// NewMysqlClient create new MyySQL client
func NewMysqlClient(config *MySQLConfig) *MySQLManager {

	var err error
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.DNS, config.Schema))
	if strings.ToLower(fmt.Sprintf("%s", log.GetLevel())) == "debug" {
		db.LogMode(true)
	}

	if err != nil {
		log.Panic(err)
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
	DNS      string `yaml:"dns"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}
