package kuberneteswatcher

import (
	"encoding/json"
	"statusbay/state"
	"statusbay/watcher/kubernetes/common"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Storage interface
type Storage interface {
	CreateApply(data *RegistryRow, status common.DeploymentStatus) (uint, error)
	UpdateApply(id uint, data *RegistryRow, status common.DeploymentStatus) (bool, error)
	GetAppliesByStatus(status common.DeploymentStatus) (map[uint]DBSchema, error)
	UpdateAppliesVersionHistory(deploymentName string, hash uint64) bool
	DeleteAppliedVersion(deploymentName string) bool
}

// MySQLStorage ...
type MySQLStorage struct {
	client *state.MySQLManager
}

// NewMysql create new MyySQL client
func NewMysql(db *state.MySQLManager) *MySQLStorage {

	return &MySQLStorage{
		client: db,
	}
}

// CreateDeployment creating a new deployment
func (my *MySQLStorage) CreateApply(data *RegistryRow, status common.DeploymentStatus) (uint, error) {

	log.WithFields(log.Fields{
		"name": data.DBSchema.Application,
	}).Debug("Save new apply")

	deploymentDetails, err := json.Marshal(data.DBSchema)
	if err != nil {
		return 0, err
	}

	apply := state.TableKubernetes{
		Name:      data.DBSchema.Application,
		Cluster:   data.DBSchema.Cluster,
		Namespace: data.DBSchema.Namespace,
		Status:    string(status),
		Details:   string(deploymentDetails),
		DeployBy:  data.DBSchema.DeployBy,
		Time:      data.DBSchema.CreationTimestamp,
	}

	if err := my.client.DB.Create(&apply).Error; err != nil {
		log.WithError(err).WithFields(log.Fields{
			"name":        data.DBSchema.Application,
			"deploy_time": data.DBSchema.CreationTimestamp,
		}).Error("MySQL: Error when trying to create a new apply")
		return 0, err
	}

	return apply.ID, nil

}

// UpdateApply update current deployment
func (my *MySQLStorage) UpdateApply(id uint, data *RegistryRow, status common.DeploymentStatus) (bool, error) {

	log.WithFields(log.Fields{
		"name": data.DBSchema.Application,
		"id":   id,
	}).Debug("Update apply")

	applyDetails, err := json.Marshal(data.DBSchema)
	if err != nil {
		return false, err
	}

	if err := my.client.DB.Model(&state.TableKubernetes{}).Where("id = ?", id).Updates(state.TableKubernetes{
		Status:  string(status),
		Details: string(applyDetails),
		Time:    data.DBSchema.CreationTimestamp,
	}).Error; err != nil {
		log.WithError(err).WithFields(log.Fields{
			"id": id,
		}).Error("MySQL: Error when trying to update apply")
		return false, err
	}

	return true, nil

}

// GetAppliesByStatus return lits of deployment by given status
func (my *MySQLStorage) GetAppliesByStatus(status common.DeploymentStatus) (map[uint]DBSchema, error) {

	appRow := &[]state.TableKubernetes{}
	resources := map[uint]DBSchema{}

	if err := my.client.DB.Where(map[string]interface{}{"status": status}).Select("id, details").Find(appRow).Error; err != nil {
		log.WithError(err).WithFields(log.Fields{
			"status": status,
		}).Error("MySQL: Error when trying to get applications by status")
		return resources, err
	}

	for _, resource := range *appRow {
		var resourceDetails DBSchema
		err := json.Unmarshal([]byte(resource.Details), &resourceDetails)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{}).Error("MySQL: Could not parsing resource results")
			continue
		}
		resources[resource.ID] = resourceDetails

	}

	return resources, nil

}

// UpdateAppliesVersionHistory Checks if we should create/update a new Apply hash
func (my *MySQLStorage) UpdateAppliesVersionHistory(applyName string, hash uint64) bool {

	row := state.TableDeploymentsHash{}

	if err := my.client.DB.Where("deployment = ?", applyName).First(&row).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			my.client.DB.Create(&state.TableDeploymentsHash{
				Deployment: applyName,
				Hash:       hash,
			})
			log.WithFields(log.Fields{
				"apply_name": applyName,
				"hash":       hash,
			}).Info("deployment version created")
			return true
		}
		return false
	} else if row.Hash == hash {
		log.WithFields(log.Fields{
			"apply_name": applyName,
			"hash":       hash,
		}).Info("deployment version already exists")
		return false
	}

	my.client.DB.Model(&row).Where("deployment = ?", applyName).Update("hash", hash)
	log.WithFields(log.Fields{
		"apply_name": applyName,
		"hash":       hash,
	}).Info("Deployment version updated")
	return true

}

func (my *MySQLStorage) DeleteAppliedVersion(applyName string) bool {

	my.client.DB.Delete(&state.TableDeploymentsHash{
		Deployment: applyName,
	})
	return true
}
