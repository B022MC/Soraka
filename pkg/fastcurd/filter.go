package fastcurd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	Filter     = map[string]FilterItem
	FilterCond string
	FilterItem struct {
		Condition FilterCond
		Val       any
	}
)

// FilterCond
const (
	// 筛选条件
	CondUndefined FilterCond = "undefined"
	// 数值
	CondEq           FilterCond = "eq"
	CondLt           FilterCond = "lt"
	CondElt          FilterCond = "elt"
	CondGt           FilterCond = "gt"
	CondEgt          FilterCond = "egt"
	CondNeq          FilterCond = "neq"
	CondBetweenValue FilterCond = "betweenValue"
	// 字符串
	CondEqString  FilterCond = "eqString"
	CondLike      FilterCond = "like"
	CondNotLike   FilterCond = "notLike"
	CondNeqString FilterCond = "neqString"
	// 时间
	CondBefore      FilterCond = "before"
	CondAfter       FilterCond = "after"
	CondBetweenTime FilterCond = "betweenTime"
	// 数组
	CondIn    FilterCond = "in"
	CondNotIn FilterCond = "notIn"
	// 内部使用
	CondRaw FilterCond = "raw"

	// order
	//OrderAsc  = "asc"
	OrderDesc = "desc"
)

var (
	// CondMapDbCond 条件映射 数据库条件
	CondMapDbCond = map[FilterCond]string{
		CondEq:           "=",
		CondLt:           "<",
		CondElt:          "<=",
		CondGt:           ">",
		CondEgt:          ">=",
		CondNeq:          "<>",
		CondBetweenValue: "BETWEEN",
		CondEqString:     "=",
		CondLike:         "LIKE",
		CondNotLike:      "NOT LIKE",
		CondNeqString:    "<>",
		CondBefore:       "<",
		CondAfter:        ">",
		CondBetweenTime:  "BETWEEN",
		CondIn:           "IN",
		CondNotIn:        "NOT IN",
	}
)

func FmtCondVal(cond FilterCond, val any) (any, error) {
	switch cond {
	case CondLike, CondNotLike:
		if val, ok := val.(string); !ok {
			return nil, errors.New("筛选条件为" + string(cond) + "时,val必须为字符串")
		} else {
			return "%" + val + "%", nil
		}
	case CondBetweenValue, CondBetweenTime:
		switch val.(type) {
		case []int:
			return []int{val.([]int)[0], val.([]int)[1]}, nil
		case []string:
			return []string{val.([]string)[0], val.([]string)[1]}, nil
		case []time.Time:
			location := time.Local
			return []string{val.([]time.Time)[0].In(location).Format(time.DateTime),
				val.([]time.Time)[1].In(location).Format(time.DateTime)}, nil
		case []any:
			return val, nil
		default:
			return nil, errors.New("筛选条件为" + string(cond) + "时,val必须为数组")
		}
	default:
		return val, nil
	}
}
func FmtValPlaceholder(cond FilterCond) interface{} {
	switch cond {
	//case CondIn, CondNotIn:
	//	return "(?)"
	case CondBetweenTime, CondBetweenValue:
		return "? and ?"
	default:
		return "?"
	}
}
func BuildFilterCond(filterMap map[string]string, db *gorm.DB, filter Filter) (*gorm.DB, error) {
	for filterKey, filterItem := range filter {
		if dbField, ok := filterMap[filterKey]; (ok || filterItem.Condition ==
			CondRaw) && filterItem.Condition != CondUndefined && filterItem.Val != nil {
			switch filterItem.Condition {
			case CondLike, CondNotLike:
				dbFieldList := strings.Split(dbField, "|")
				sql := ""
				actValArr := make([]any, 0, 1)
				for _, field := range dbFieldList {
					if !IsValidQueryField(field) {
						continue
					}
					actCondition := CondMapDbCond[filterItem.Condition]
					actVal, err := FmtCondVal(filterItem.Condition, filterItem.Val)
					if err != nil {
						return nil, err
					}
					valPlaceholder := FmtValPlaceholder(filterItem.Condition)
					if arrVal, ok := actVal.([]string); ok {
						sql += fmt.Sprintf("%s %s %s", field, actCondition, valPlaceholder)
						actValArr = append(actValArr, arrVal[0], arrVal[1])
					} else {
						sql += fmt.Sprintf("%s %s %s", field, actCondition, valPlaceholder)
						actValArr = append(actValArr, actVal)
					}
					sql += " or "
				}
				sql = sql[:len(sql)-4]
				db = db.Where(sql, actValArr...)
				// todo maybe sql inspect wait review
			case CondRaw:
				//rawSQLData := filterItem.Val.([]any)
				//db = db.Where(rawSQLData[0].(string), rawSQLData[1].([]any)...)
			default:
				if !IsValidQueryField(dbField) {
					continue
				}
				actCondition := CondMapDbCond[filterItem.Condition]
				actVal, err := FmtCondVal(filterItem.Condition, filterItem.Val)
				if err != nil {
					return nil, err
				}
				valPlaceholder := FmtValPlaceholder(filterItem.Condition)
				switch filterItem.Condition {
				case CondBetweenTime, CondBetweenValue:
					if arrVal, ok := actVal.([]any); ok && len(arrVal) == 2 {
						db = db.Where(fmt.Sprintf("%s %s %s", dbField, actCondition, valPlaceholder),
							arrVal[0], arrVal[1])
					}
				default:
					db = db.Where(fmt.Sprintf("%s %s %s", dbField, actCondition, valPlaceholder), actVal)
				}
			}
		}
	}
	return db, nil
}
func BuildOrderCond(orderKeyMap map[string]string, q *gorm.DB, order map[string]string) *gorm.DB {
	for orderKey, orderVal := range order {
		if actKey, ok := orderKeyMap[orderKey]; ok {
			if orderVal == OrderDesc {
				q = q.Order(actKey + " desc")
			} else {
				q = q.Order(actKey + " asc")
			}
		}
	}
	return q
}
