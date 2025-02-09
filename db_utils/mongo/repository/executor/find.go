package executor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	utils "github.com/ashish036/GO-no-code-api/db_utils/mongo/repository/utils"
	mongoUtils "github.com/ashish036/GO-no-code-api/db_utils/mongo/utils"
	repoUtils "github.com/ashish036/GO-no-code-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOne executes a find operation to return a single document from MongoDB.
//
// Parameters:
//   - ctx: Context for the database operation
//   - result: Pointer to the variable where the result will be decoded
//   - queryParams: Map containing query parameters
//
// Returns:
//   - repoUtils.Status containing operation status and any error information
func (executor MongoQueryExecutor) FindOne(ctx context.Context, result interface{}, queryParams map[utils.QueryParam]interface{}) repoUtils.Status {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return status
	}

	opts, status := utils.ParseQueryParams[*options.FindOneOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseOpts() -> %w", status.Error)
		return status
	}

	err := collConnection.FindOne(ctx, filter, opts).Decode(result)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("FindOne() -> %w", status.Error)
		return status
	}

	return repoUtils.Status{}
}

func (executor MongoQueryExecutor) FindAll(ctx context.Context, result interface{}, queryParams map[utils.QueryParam]interface{}) repoUtils.Status {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return status
	}

	opts, status := utils.ParseQueryParams[*options.FindOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseOpts() -> %w", status.Error)
		return status
	}

	cursor, err := collConnection.Find(ctx, filter, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("Find() -> %w", status.Error)
		return status
	}

	err = cursor.All(ctx, result)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("All() -> %w", status.Error)
		return status
	}

	return repoUtils.Status{}
}

func (executor MongoQueryExecutor) Distinct(ctx context.Context, result interface{}, queryParams map[utils.QueryParam]interface{}) repoUtils.Status {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return status
	}

	opts, status := utils.ParseQueryParams[*options.DistinctOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseOpts() -> %w", status.Error)
		return status
	}

	field, status := utils.ParseQueryParams[string](utils.FIELD, queryParams, true)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return status
	}

	distinct, err := collConnection.Distinct(ctx, field, filter, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("Distinct() -> %w", status.Error)
		return status
	}

	byteData, err := json.Marshal(distinct)
	if err != nil {
		return repoUtils.Status{
			Code:    http.StatusInternalServerError,
			Message: "Failed to marshal distinct results",
			Error:   fmt.Errorf("Marshal() -> %w", err),
		}
	}

	err = json.Unmarshal(byteData, result)
	if err != nil {
		return repoUtils.Status{
			Code:    http.StatusInternalServerError,
			Message: "Failed to unmarshal distinct results",
			Error:   fmt.Errorf("Unmarshal() -> %w", err),
		}
	}

	return repoUtils.Status{}
}

func (executor MongoQueryExecutor) CountDocuments(ctx context.Context, queryParams map[utils.QueryParam]interface{}) (int64, repoUtils.Status) {
	collConnection, status := mongoUtils.GetDBCollectionConnection(executor.database, executor.collection)
	if status.IsError() {
		status.Error = fmt.Errorf("GetDBCollectionConnection() -> %w", status.Error)
		return -1, status
	}

	filter, status := utils.ParseQueryParams[bson.M](utils.FILTER, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return -1, status
	}

	opts, status := utils.ParseQueryParams[*options.CountOptions](utils.OPTS, queryParams, false)
	if status.IsError() {
		status.Error = fmt.Errorf("ParseQueryParams() -> %w", status.Error)
		return -1, status
	}

	count, err := collConnection.CountDocuments(ctx, filter, opts)
	if err != nil {
		status = mongoUtils.ParseMongoError(err)
		status.Error = fmt.Errorf("CountDocuments() -> %w", status.Error)
		return -1, status
	}

	return count, repoUtils.Status{}
}
