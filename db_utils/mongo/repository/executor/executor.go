package executor

type MongoQueryExecutor struct {
	database   string
	collection string
}

func NewMongoQueryExecutor(database, collection string) MongoQueryExecutor {
	return MongoQueryExecutor{
		database:   database,
		collection: collection,
	}
}
