package executor

import (
	"context"
	"fmt"

	utils "github.com/ashish036/GO-no-code-api/db_utils/mongo/repository/utils"
	mongoUtils "github.com/ashish036/GO-no-code-api/db_utils/mongo/utils"
	repoUtils "github.com/ashish036/GO-no-code-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// [TODO] Use mongo transaction to perform whole operations
func (executor MongoQueryExecutor) BulkWrite(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (mongo.BulkWriteResult, repoUtils.Status) {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return mongo.BulkWriteResult{}, status
	}

	models, status := utils.ParseQueryParams[[]mongo.WriteModel](utils.WRITE_MODELS, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return mongo.BulkWriteResult{}, status
	}

	opts, status := utils.ParseQueryParams[*options.BulkWriteOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseOpts() -> %w", status.Error)
		return mongo.BulkWriteResult{}, status
	}

	bulkWriteResult, err := collConnection.BulkWrite(ctx, models, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("BulkWrite() -> %w", status.Error)
		return mongo.BulkWriteResult{}, status
	}

	return *bulkWriteResult, repoUtils.Status{}
}

// [TODO] Use mongo transaction to perform whole operations
func (executor MongoQueryExecutor) Aggregate(ctx context.Context, result interface{}, queryParams map[utils.QueryParam]interface{}) repoUtils.Status {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return status
	}

	pipeline, status := utils.ParseQueryParams[[]interface{}](utils.PIPELINE, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return status
	}

	opts, status := utils.ParseQueryParams[*options.AggregateOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseOpts() -> %w", status.Error)
		return status
	}

	cursor, err := collConnection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("Aggregate() -> %w", status.Error)
		return status
	}

	err = cursor.All(context.TODO(), result)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("All() -> %w", status.Error)
		return status
	}

	return repoUtils.Status{}
}
