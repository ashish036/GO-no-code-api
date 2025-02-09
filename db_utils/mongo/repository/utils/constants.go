package utils

type QueryParam string

const (
	FILTER       QueryParam = "filter"
	OPTS         QueryParam = "opts"
	DOCUMENT     QueryParam = "document"
	DOCUMENTS    QueryParam = "documents"
	UPDATE       QueryParam = "update"
	UPSERT       QueryParam = "upsert"
	REPLACEMENT  QueryParam = "replacement"
	PIPELINE     QueryParam = "pipeline"
	FIELD        QueryParam = "field"
	WRITE_MODELS QueryParam = "writeModels"
)
