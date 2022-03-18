/*
 * @Author: lwnmengjing
 * @Date: 2022/3/11 16:43
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/11 16:43
 */

package mgos

import (
	"go.mongodb.org/mongo-driver/bson"
)

// MakeCondition make condition
func MakeCondition(q interface{}) (bson.M, bson.D) {
	condition := &Public{}
	ResolveSearchQuery(q, condition)
	var filter bson.M
	var andFilter bson.M
	var orFilter bson.M
	if len(condition.And) > 0 {
		andFilter = bson.M{"$and": condition.And}
	}
	if len(condition.Or) > 0 {
		orFilter = bson.M{"$or": condition.Or}
	}
	if len(condition.And) > 0 && len(condition.Or) > 0 {
		filter = bson.M{"$and": []bson.M{andFilter, orFilter}}
	} else if len(condition.And) > 0 {
		filter = andFilter
	} else if len(condition.Or) > 0 {
		filter = orFilter
	}
	return filter, condition.Order
}
