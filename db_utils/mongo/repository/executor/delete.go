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

func (executor MongoQueryExecutor) DeleteOne(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (mongo.DeleteResult, repoUtils.Status) {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	opts, status := utils.ParseQueryParams[*options.DeleteOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseOpts() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	deleteResp, err := collConnection.DeleteOne(ctx, filter, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("FindOne() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	return *deleteResp, repoUtils.Status{}
}

func (executor MongoQueryExecutor) DeleteMany(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (mongo.DeleteResult, repoUtils.Status) {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	opts, status := utils.ParseQueryParams[*options.DeleteOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseOpts() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	deleteResp, err := collConnection.DeleteMany(ctx, filter, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("FindOne() -> %w", status.Error)
		return mongo.DeleteResult{}, status
	}

	return *deleteResp, repoUtils.Status{}
}
