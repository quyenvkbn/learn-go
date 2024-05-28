package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CursorToMap(cursor *mongo.Cursor, ctx context.Context) map[string]any {
	result := make(map[string]any)
	for cursor.Next(ctx) {
		var rowData bson.M
		cursor.Decode(&rowData)
		id := ToString(rowData["_id"])
		result[id] = rowData
	}
	defer cursor.Close(ctx)

	return result
}
