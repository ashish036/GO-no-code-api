package executor

import (
	"context"
	"fmt"

	utils "github.com/ashish036/GO-no-code-api/db_utils/mongo/repository/utils"
	mongoUtils "github.com/ashish036/GO-no-code-api/db_utils/mongo/utils"
	repoUtils "github.com/ashish036/GO-no-code-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (executor MongoQueryExecutor) UpdateOne(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (mongo.UpdateResult, repoUtils.Status) {
	var resp mongo.UpdateResult
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return resp, status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}
	update, status := utils.ParseQueryParams[interface{}](utils.UPDATE, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	opts, status := utils.ParseQueryParams[*options.UpdateOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	updateResp, err := collConnection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("UpdateOne() -> %w", status.Error)
		return resp, status
	}

	return *updateResp, repoUtils.Status{}
}

func (executor MongoQueryExecutor) UpdateMany(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (mongo.UpdateResult, repoUtils.Status) {
	var resp mongo.UpdateResult
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return resp, status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}
	update, status := utils.ParseQueryParams[interface{}](utils.UPDATE, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	opts, status := utils.ParseQueryParams[*options.UpdateOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	updateResp, err := collConnection.UpdateMany(ctx, filter, update, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("UpdateMany() -> %w", status.Error)
		return resp, status
	}

	return *updateResp, repoUtils.Status{}
}

func (executor MongoQueryExecutor) UpsertOne(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (mongo.UpdateResult, repoUtils.Status) {
	var resp mongo.UpdateResult
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return resp, status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}
	upsert, status := utils.ParseQueryParams[interface{}](utils.UPSERT, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	opts, status := utils.ParseQueryParams[*options.UpdateOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	updateResp, err := collConnection.UpdateOne(ctx, filter, upsert, opts.SetUpsert(true))
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("UpdateOne() -> %w", status.Error)
		return resp, status
	}

	return *updateResp, repoUtils.Status{}
}

func (executor MongoQueryExecutor) ReplaceOne(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (mongo.UpdateResult, repoUtils.Status) {
	var resp mongo.UpdateResult
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return resp, status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}
	replacement, status := utils.ParseQueryParams[interface{}](utils.REPLACEMENT, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	opts, status := utils.ParseQueryParams[*options.ReplaceOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return resp, status
	}

	updateResp, err := collConnection.ReplaceOne(ctx, filter, replacement, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("ReplaceOne() -> %w", status.Error)
		return resp, status
	}

	return *updateResp, repoUtils.Status{}

}
