package kubernetes

import (
	"fmt"
	"statusbay/state"

	"github.com/apex/log"
	"github.com/jinzhu/gorm"
)

type Storage interface {
	Applications(queryFillter FilterApplications) (*[]state.TableKubernetes, error)
	ApplicationsCount(queryFillter FilterApplications) (int64, error)
	GetDeployment(name string, time int64) (state.TableKubernetes, error)
	GetUniqueFieldValues(tableName, columnName string) ([]string, error)
}

type MySQLStorage struct {
	client *state.MySQLManager
}

//NewMysql create new MyySQL client
func NewMysql(db *state.MySQLManager) *MySQLStorage {

	return &MySQLStorage{
		client: db,
	}
}

func (my *MySQLStorage) ApplicationsCount(queryFillter FilterApplications) (int64, error) {

	var dummy *state.TableKubernetes
	queryBuilder := my.builderApplications(queryFillter)
	var count int64

	if err := queryBuilder.Table(dummy.TableName()).Select("count(id)").Count(&count).Error; err != nil {
		log.WithError(err).Error("MYSQL: could not fetch application list for count")
		return 0, err
	}
	return count, nil

}

func (my *MySQLStorage) Applications(queryFillter FilterApplications) (*[]state.TableKubernetes, error) {

	table := &[]state.TableKubernetes{}
	queryBuilder := my.builderApplications(queryFillter)

	if err := queryBuilder.Find(table).Error; err != nil {
		log.WithError(err).Error("MYSQL: could not fetch application list")
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
		log.WithError(err).WithFields(log.Fields{
			"table_name":  tableName,
			"column_name": columnName,
		}).Error("MYSQL: could not fetch uniq fields value")
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

func (my *MySQLStorage) GetDeployment(name string, time int64) (state.TableKubernetes, error) {

	var empty state.TableKubernetes
	deploymentRow := &state.TableKubernetes{}

	if err := my.client.DB.Where(&state.TableKubernetes{Name: name}).Where(&state.TableKubernetes{Time: time}).First(deploymentRow).Error; err != nil {
		log.WithError(err).WithFields(log.Fields{
			"application": name,
			"time":        time,
		}).Error("MySQL: Error when trying to get deployment")
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
