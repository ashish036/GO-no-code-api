package executor

import (
	"context"
	"fmt"

	utils "github.com/ashish036/GO-no-code-api/db_utils/mongo/repository/utils"
	mongoUtils "github.com/ashish036/GO-no-code-api/db_utils/mongo/utils"
	repoUtils "github.com/ashish036/GO-no-code-api/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (executor MongoQueryExecutor) InsertOne(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (interface{}, repoUtils.Status) {
	var insertedId interface{}
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return insertedId, status
	}

	document, status := utils.ParseQueryParams[interface{}](utils.DOCUMENT, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return insertedId, status
	}

	opts, status := utils.ParseQueryParams[*options.InsertOneOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return insertedId, status
	}

	insertResult, err := collConnection.InsertOne(ctx, document, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("InsertOne() -> %w", status.Error)
		return insertedId, status
	}

	return insertResult.InsertedID, repoUtils.Status{}
}

func (executor MongoQueryExecutor) InsertMany(ctx context.Context, queryParams map[utils.QueryParam]interface{}) ([]interface{}, repoUtils.Status) {
	var insertedIds []interface{}
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return insertedIds, status
	}

	documents, status := utils.ParseQueryParams[[]interface{}](utils.DOCUMENTS, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return insertedIds, status
	}

	opts, status := utils.ParseQueryParams[*options.InsertManyOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return insertedIds, status
	}

	insertResults, err := collConnection.InsertMany(ctx, documents, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("InsertMany() -> %w", status.Error)
		return insertedIds, status
	}

	return insertResults.InsertedIDs, repoUtils.Status{}
}
