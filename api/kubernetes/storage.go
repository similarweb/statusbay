package kubernetes

import (
	"fmt"
	"statusbay/state"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Storage interface {
	Applications(queryFillter FilterApplications) (*[]state.TableKubernetes, error)
	ApplicationsCount(queryFillter FilterApplications) (int64, error)
	GetDeployment(applyID string) (state.TableKubernetes, error)
	GetUniqueFieldValues(tableName, columnName string) ([]string, error)
}

type MySQLStorage struct {
	client *state.MySQLManager
	logger *log.Entry
}

//NewMysql create new MyySQL client
func NewMysql(db *state.MySQLManager) *MySQLStorage {

	return &MySQLStorage{
		client: db,
		logger: log.WithField("storage_engine", "mysql"),
	}
}

func (my *MySQLStorage) ApplicationsCount(queryFillter FilterApplications) (int64, error) {

	var dummy *state.TableKubernetes
	queryBuilder := my.builderApplications(queryFillter)
	var count int64

	if err := queryBuilder.Table(dummy.TableName()).Select("count(apply_id)").Count(&count).Error; err != nil {
		my.logger.WithError(err).Error("could not fetch application list count")
		return 0, err
	}
	return count, nil

}

func (my *MySQLStorage) Applications(queryFillter FilterApplications) (*[]state.TableKubernetes, error) {

	table := &[]state.TableKubernetes{}
	queryBuilder := my.builderApplications(queryFillter)

	if err := queryBuilder.Find(table).Error; err != nil {
		my.logger.WithError(err).Error("could not fetch application list")
		return nil, err
	}
	return table, nil

}

//GetUniqueFieldValues return list of unique values by given table name and column name
func (my *MySQLStorage) GetUniqueFieldValues(tableName, columnName string) ([]string, error) {

	var values []string

	rows, err := my.client.DB.Select(fmt.Sprintf("%s as val, COUNT(*) as count", columnName)).Table(tableName).Group(columnName).Order("count DESC").Rows()
	defer rows.Close()

	if err != nil {
		my.logger.WithError(err).WithFields(log.Fields{
			"table_name":  tableName,
			"column_name": columnName,
		}).Error("could not fetch unique field values")
		return values, err
	}

	for rows.Next() {
		var val string
		var count int16
		rows.Scan(&val, &count)
		values = append(values, val)
	}

	return values, nil
}

func (my *MySQLStorage) GetDeployment(applyID string) (state.TableKubernetes, error) {

	var empty state.TableKubernetes
	deploymentRow := &state.TableKubernetes{}

	if err := my.client.DB.Where(&state.TableKubernetes{ApplyId: applyID}).First(deploymentRow).Error; err != nil {
		my.logger.WithError(err).WithFields(log.Fields{
			"apply_id": applyID,
		}).Error("could not get deployment")
		return empty, err
	}

	return *deploymentRow, nil

}

func (my *MySQLStorage) builderApplications(queryFillter FilterApplications) *gorm.DB {

	queryBuilder := my.client.DB.Offset(queryFillter.Offset).Limit(queryFillter.Limit)

	if len(queryFillter.Clusters) > 0 {
		for i, cluster := range queryFillter.Clusters {
			if i == 0 {
				queryBuilder = queryBuilder.Where(&state.TableKubernetes{Cluster: cluster})
				continue
			}
			queryBuilder = queryBuilder.Or(&state.TableKubernetes{Cluster: cluster})
		}

	}
	if len(queryFillter.Namespaces) > 0 {

		for i, namespace := range queryFillter.Namespaces {
			if i == 0 {
				queryBuilder = queryBuilder.Where(&state.TableKubernetes{Namespace: namespace})
				continue
			}
			queryBuilder = queryBuilder.Or(&state.TableKubernetes{Namespace: namespace})
		}

	}

	if queryFillter.Name != "" {
		queryBuilder = queryBuilder.Where("name LIKE ?", fmt.Sprintf("%%%s%%", queryFillter.Name))

	}

	if queryFillter.ExactName != "" {
		queryBuilder = queryBuilder.Where("name = ?", fmt.Sprintf("%s", queryFillter.ExactName))
	}

	if queryFillter.DeployBy != "" {
		queryBuilder = queryBuilder.Where("deploy_by LIKE ?", fmt.Sprintf("%%%s%%", queryFillter.DeployBy))
	}

	// By Default SortBy will sort by desc direction
	if queryFillter.SortBy != "" {
		queryBuilder = queryBuilder.Order(fmt.Sprintf("%s %s", queryFillter.SortBy, queryFillter.SortDirection))
	}

	if queryFillter.Distinct {
		queryBuilder = queryBuilder.Where("time IN (?)", my.client.DB.Select("MAX(time)").Model(&state.TableKubernetes{}).Group("name").QueryExpr())
	}

	// We only support a case where we have both filter of From and To
	if queryFillter.From != 0 && queryFillter.To != 0 {
		queryBuilder = queryBuilder.Where("time >= ? and time <= ?", queryFillter.From, queryFillter.To)
	}

	if len(queryFillter.Statuses) > 0 {

		for i, status := range queryFillter.Statuses {
			if i == 0 {
				queryBuilder = queryBuilder.Where(&state.TableKubernetes{Status: status})
				continue
			}
			queryBuilder = queryBuilder.Or(&state.TableKubernetes{Status: status})
		}
	}

	return queryBuilder
}
