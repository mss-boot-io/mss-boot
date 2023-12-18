package gorms

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// FromQueryTag tag标记
	FromQueryTag = "search"
	// Mysql 数据库标识
	Mysql = "mysql"
	// Postgres 数据库标识
	Postgres = "postgres"
	// Dm 数据库标识
	Dm = "dm"
)

// ResolveSearchQuery 解析
/**
 * 	exact / iexact 等于
 * 	contains / icontains 包含
 *	gt / gte 大于 / 大于等于
 *	lt / lte 小于 / 小于等于
 *	startswith / istartswith 以…起始
 *	endswith / iendswith 以…结束
 *	in
 *	isnull
 *  order 排序		e.g. order[key]=desc     order[key]=asc
 */
func ResolveSearchQuery(driver string, q interface{}, condition Condition) {
	qType := reflect.TypeOf(q)
	qValue := reflect.ValueOf(q)
	if qType.Kind() == reflect.Ptr {
		qType = qType.Elem()
	}
	if qValue.Kind() == reflect.Ptr {
		qValue = qValue.Elem()
	}
	var t *resolveSearchTag

	//var sep = "`"
	//if driver == Postgres {
	//	sep = "\""
	//}

	for i := 0; i < qType.NumField(); i++ {
		tag, ok := qType.Field(i).Tag.Lookup(FromQueryTag)
		if !ok {
			//递归调用
			ResolveSearchQuery(driver, qValue.Field(i).Interface(), condition)
			continue
		}
		switch tag {
		case "-":
			continue
		}
		t = makeTag(tag)
		if qValue.Field(i).IsZero() {
			continue
		}

		parseSQL(driver, t, condition, qValue, i)
	}
}

func parseSQL(driver string, searchTag *resolveSearchTag, condition Condition, qValue reflect.Value, i int) {
	if driver == Dm {
		searchTag.Table = strings.ToUpper(searchTag.Table)
		searchTag.Column = strings.ToUpper(searchTag.Column)
	}
	iStr := ""
	if driver == Postgres {
		iStr = "i"
	}
	if searchTag.Table != "" {
		searchTag.Table = fmt.Sprintf("`%s`.", searchTag.Table)
	}
	switch searchTag.Type {
	case "left":
		//左关联
		join := condition.SetJoinOn(searchTag.Type, fmt.Sprintf(
			"left join `%s` on `%s`.`%s` = %s.`%s`",
			searchTag.Join,
			searchTag.Join,
			searchTag.On[0],
			searchTag.Table,
			searchTag.On[1],
		))
		ResolveSearchQuery(driver, qValue.Field(i).Interface(), join)
	case "exact", "iexact":
		condition.SetWhere(fmt.Sprintf("%s`%s` = ?", searchTag.Table, searchTag.Column), []interface{}{qValue.Field(i).Interface()})
	case "contains":
		condition.SetWhere(fmt.Sprintf("%s`%s` like ?", searchTag.Table, searchTag.Column), []interface{}{"%" + qValue.Field(i).String() + "%"})
	case "icontains":
		condition.SetWhere(fmt.Sprintf("%s`%s` %slike ?", searchTag.Table, searchTag.Column, iStr), []interface{}{"%" + qValue.Field(i).String() + "%"})
	case "gt":
		condition.SetWhere(fmt.Sprintf("%s`%s` > ?", searchTag.Table, searchTag.Column), []interface{}{qValue.Field(i).Interface()})
	case "gte":
		condition.SetWhere(fmt.Sprintf("%s`%s` >= ?", searchTag.Table, searchTag.Column), []interface{}{qValue.Field(i).Interface()})
	case "lt":
		condition.SetWhere(fmt.Sprintf("%s`%s` < ?", searchTag.Table, searchTag.Column), []interface{}{qValue.Field(i).Interface()})
	case "lte":
		condition.SetWhere(fmt.Sprintf("%s`%s` <= ?", searchTag.Table, searchTag.Column), []interface{}{qValue.Field(i).Interface()})
	case "startswith":
		condition.SetWhere(fmt.Sprintf("%s`%s` like ?", searchTag.Table, searchTag.Column), []interface{}{qValue.Field(i).String() + "%"})
	case "istartswith":
		condition.SetWhere(fmt.Sprintf("%s`%s` %slike ?", searchTag.Table, searchTag.Column, iStr), []interface{}{qValue.Field(i).String() + "%"})
	case "endswith":
		condition.SetWhere(fmt.Sprintf("%s`%s` like ?", searchTag.Table, searchTag.Column), []interface{}{"%" + qValue.Field(i).String()})
	case "iendswith":
		condition.SetWhere(fmt.Sprintf("%s`%s` %slike ?", searchTag.Table, searchTag.Column, iStr), []interface{}{"%" + qValue.Field(i).String()})
	case "in":
		condition.SetWhere(fmt.Sprintf("%s`%s` in (?)", searchTag.Table, searchTag.Column), []interface{}{qValue.Field(i).Interface()})
	case "isnull":
		if !(qValue.Field(i).IsZero() && qValue.Field(i).IsNil()) {
			condition.SetWhere(fmt.Sprintf("%s`%s` isnull", searchTag.Table, searchTag.Column), make([]interface{}, 0))
		}
	case "order":
		switch strings.ToLower(qValue.Field(i).String()) {
		case "desc", "asc":
			condition.SetOrder(fmt.Sprintf("%s`%s` %s", searchTag.Table, searchTag.Column, qValue.Field(i).String()))
		}
	}
}
