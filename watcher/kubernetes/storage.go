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
	CreateApply(data *RegistryRow, status common.DeploymentStatus) (string, error)
	UpdateApply(applyID string, data *RegistryRow, status common.DeploymentStatus) (bool, error)
	GetAppliesByStatus(status common.DeploymentStatus) (map[string]DBSchema, error)
	UpdateAppliesVersionHistory(deploymentName string, hash uint64) bool
	DeleteAppliedVersion(deploymentName string) bool
}

// MySQLStorage ...
type MySQLStorage struct {
	client *state.MySQLManager
	logger *log.Entry
}

// NewMysql create new MyySQL client
func NewMysql(db *state.MySQLManager) *MySQLStorage {

	return &MySQLStorage{
		client: db,
		logger: log.WithField("storage_engine", "mysql"),
	}
}

// CreateApply creating a new apply row
func (my *MySQLStorage) CreateApply(data *RegistryRow, status common.DeploymentStatus) (string, error) {

	my.logger.WithFields(log.Fields{
		"name": data.DBSchema.Application,
	}).Debug("save new apply")

	deploymentDetails, err := json.Marshal(data.DBSchema)
	if err != nil {
		return "", err
	}

	applyID := data.GetApplyID()
	apply := state.TableKubernetes{
		ApplyId:   applyID,
		Name:      data.DBSchema.Application,
		Cluster:   data.DBSchema.Cluster,
		Namespace: data.DBSchema.Namespace,
		Status:    string(status),
		Details:   string(deploymentDetails),
		DeployBy:  data.DBSchema.DeployBy,
		Time:      data.DBSchema.CreationTimestamp,
	}

	if err := my.client.DB.Create(&apply).Error; err != nil {
		my.logger.WithError(err).WithFields(log.Fields{
			"name":     data.DBSchema.Application,
			"apply_id": applyID,
		}).Error("error when trying to create a new apply")
		return "", err
	}

	return applyID, nil

}

// UpdateApply update current deployment
func (my *MySQLStorage) UpdateApply(applyID string, data *RegistryRow, status common.DeploymentStatus) (bool, error) {

	my.logger.WithFields(log.Fields{
		"name":     data.DBSchema.Application,
		"apply_id": applyID,
	}).Debug("update apply")

	applyDetails, err := json.Marshal(data.DBSchema)
	if err != nil {
		return false, err
	}

	if err := my.client.DB.Model(&state.TableKubernetes{}).Where("apply_id = ?", applyID).Updates(state.TableKubernetes{
		Status:  string(status),
		Details: string(applyDetails),
		Time:    data.DBSchema.CreationTimestamp,
	}).Error; err != nil {
		my.logger.WithError(err).WithFields(log.Fields{
			"apply_id": applyID,
		}).Error("error when trying to update apply")
		return false, err
	}

	return true, nil

}

// GetAppliesByStatus return lits of deployment by given status
func (my *MySQLStorage) GetAppliesByStatus(status common.DeploymentStatus) (map[string]DBSchema, error) {

	appRow := &[]state.TableKubernetes{}
	resources := map[string]DBSchema{}

	if err := my.client.DB.Where(map[string]interface{}{"status": status}).Select("apply_id, details").Find(appRow).Error; err != nil {
		my.logger.WithError(err).WithFields(log.Fields{
			"status": status,
		}).Error("error when trying to get applications by status")
		return resources, err
	}

	for _, resource := range *appRow {
		var resourceDetails DBSchema
		err := json.Unmarshal([]byte(resource.Details), &resourceDetails)
		if err != nil {
			my.logger.WithError(err).Error("could not parsing resource results")
			continue
		}
		resources[resource.ApplyId] = resourceDetails

	}

	return resources, nil

}

// UpdateAppliesVersionHistory Checks if we should create/update a new Apply hash
func (my *MySQLStorage) UpdateAppliesVersionHistory(applyName string, hash uint64) bool {

	row := state.TableDeploymentsHash{}

	// Check if the deployment exists in DB
	if err := my.client.DB.Where("deployment = ?", applyName).First(&row).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			my.client.DB.Create(&state.TableDeploymentsHash{
				Deployment: applyName,
				Hash:       hash,
			})
			my.logger.WithFields(log.Fields{
				"apply_name": applyName,
				"hash":       hash,
			}).Debug("apply hash version not found in storage, creating one")
			return true
		}
		return false
	} else if row.Hash == hash {
		my.logger.WithFields(log.Fields{
			"apply_name": applyName,
			"spec_hash":  hash,
		}).Info("apply version already exists, the spec data is equal the the last apply")
		return false
	}

	my.client.DB.Model(&row).Where("deployment = ?", applyName).Update("hash", hash)
	my.logger.WithFields(log.Fields{
		"apply_name": applyName,
		"spec_hash":  hash,
	}).Info("apply version updated")
	return true

}

func (my *MySQLStorage) DeleteAppliedVersion(applyName string) bool {

	my.client.DB.Delete(&state.TableDeploymentsHash{
		Deployment: applyName,
	})
	return true
}
