/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/15 03:54:31
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/15 03:54:31
 */

package actions

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type testModel struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	AuthorID         primitive.ObjectID `json:"authorID" bson:"author_id"`
	Author           author             `json:"author" bson:"author"`
}

type author struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func Test_getLinkTag(t *testing.T) {
	type args struct {
		model any
	}

	tests := []struct {
		name               string
		args               args
		wantCollectionName string
		wantLocalField     string
		wantForeignField   string
	}{
		{
			name: "test",
			args: args{
				model: &testModel{},
			},
			wantCollectionName: "authors",
			wantLocalField:     "authorID",
			wantForeignField:   "_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCollectionName, gotLocalField, gotForeignField := getLinkTag(tt.args.model)
			if gotCollectionName != tt.wantCollectionName {
				t.Errorf("getLinkTag() gotCollectionName = %v, want %v", gotCollectionName, tt.wantCollectionName)
			}
			if gotLocalField != tt.wantLocalField {
				t.Errorf("getLinkTag() gotLocalField = %v, want %v", gotLocalField, tt.wantLocalField)
			}
			if gotForeignField != tt.wantForeignField {
				t.Errorf("getLinkTag() gotForeignField = %v, want %v", gotForeignField, tt.wantForeignField)
			}
		})
	}
}
