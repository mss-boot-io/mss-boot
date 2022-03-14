/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 9:32
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 9:32
 */

package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Tabler interface {
	TableName() string
	Make()
	C() *mongo.Collection
}
