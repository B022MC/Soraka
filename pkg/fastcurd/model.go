package fastcurd

import (
	"context"
	"errors"
	"time"

	"Soraka/pkg/constraints"
	"Soraka/pkg/json"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	errCreateFailed = errors.New("create failed")
)
var (
	defaultFilterKeyMapDbField = map[string]string{
		PrimaryField:    PrimaryField,
		CreateTimeField: CreateTimeField,
		UpdateTimeField: UpdateTimeField,
	}
	defaultOrderKeyMapDbField = map[string]string{
		PrimaryField:    PrimaryField,
		CreateTimeField: CreateTimeField,
		UpdateTimeField: UpdateTimeField,
	}
)

type (
	BaseModelBase[M any] interface {
		schema.Tabler
		GetDB() *gorm.DB
		GetCtx() context.Context
		GetFmtDetail(sceneParam ...string) any
		GetFilterKeyMapDBField() map[string]string
		GetOrderKeyMapDBField() map[string]string
	}
	BaseModel[P any] interface {
		constraints.Ptr[P]
		BaseModelBase[*P]
	}
	Base struct {
		ID                 uint64          `json:"id" redis:"id" gorm:"type:bigint unsigned auto_increment;primaryKey;"`
		Ctime              *time.Time      `json:"ctime,omitempty" gorm:"type:datetime;default:CURRENT_TIMESTAMP;not null;"`
		Utime              *time.Time      `json:"utime,omitempty" gorm:"type:datetime ON UPDATE CURRENT_TIMESTAMP;default:CURRENT_TIMESTAMP;not null;"`
		RelationAffectRows int             `json:"-" gorm:"-"` // 更新时用来保存其他关联数据的更新数
		Ctx                context.Context `json:"-" gorm:"-"` // ctx
		DB                 *gorm.DB        `json:"-" gorm:"-"` // db
	}
)

func (m *Base) GetDB() *gorm.DB {
	return m.DB
}
func (m *Base) GetCtx() context.Context {
	return m.Ctx
}
func (m *Base) GetFilterKeyMapDBField() map[string]string {
	return defaultFilterKeyMapDbField
}
func (m *Base) GetOrderKeyMapDBField() map[string]string {
	return defaultOrderKeyMapDbField
}
func (m *Base) GetFmtDetail(scenes ...string) any {
	var scene string
	if len(scenes) == 1 {
		scene = scenes[0]
	}
	var model any
	switch scene {
	default:
		model = NewDefaultSceneBase(m)
	}
	return model
}
func (m *Base) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
func (m *Base) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func NewDefaultSceneBase(m *Base) map[string]any {
	return map[string]any{
		"id":    m.ID,
		"ctime": m.Ctime,
		"utime": m.Utime,
	}
}

func GetGormQuery[P BaseModel[M], M any](m P) *gorm.DB {
	db := m.GetDB()
	if m.GetCtx() != nil {
		db = db.WithContext(m.GetCtx())
	}
	return db.Model(m)
}
func GetTxGormQuery[P BaseModel[M], M any](m P, tx *gorm.DB) *gorm.DB {
	db := tx
	if m.GetCtx() != nil {
		db = db.WithContext(m.GetCtx())
	}
	return db.Model(m)
}

func GetFmtList[P BaseModel[M], M any](arr []P, sceneParam ...string) any {
	scene := ""
	if len(sceneParam) > 0 {
		scene = sceneParam[0]
	}
	fmtList := make([]any, 0, len(arr))
	actList := arr
	for _, item := range actList {
		fmtList = append(fmtList, item.GetFmtDetail(scene))
	}
	return fmtList
}
func GetDetailByID[P BaseModel[M], M any](m P, id uint64) (P, error) {
	record := new(M)
	err := GetGormQuery(m).Where("id = ?", id).First(record).Error
	if err != nil {
		record = nil
	}
	return record, err
}
func ListByIDArr[P BaseModel[M], M any](m P, idArr []uint64) ([]P, error) {
	if len(idArr) == 0 {
		return nil, nil
	}
	list := make([]P, 0, len(idArr))
	err := GetGormQuery(m).Where("id in ?", idArr).Find(&list).Error
	return list, err
}
func dbEditByID(db *gorm.DB, id uint64, values map[string]any) (int64, error) {
	res := db.Where("id = ?", id).Updates(values)
	return res.RowsAffected, res.Error
}
func EditByID[P BaseModel[M], M any](m P, id uint64, values map[string]any) (int64, error) {
	return dbEditByID(GetGormQuery(m), id, values)
}
func TxEditByID[P BaseModel[M], M any](m P, tx *gorm.DB, id uint64, values map[string]any) (int64, error) {
	return dbEditByID(GetTxGormQuery(m, tx), id, values)
}
func dbEditByIDArr(db *gorm.DB, idArr []uint64, values map[string]any) (int64, error) {
	res := db.Where("id in ?", idArr).Updates(values)
	return res.RowsAffected, res.Error
}
func EditByIDArr[P BaseModel[M], M any](m P, idArr []uint64, values map[string]any) (int64, error) {
	return dbEditByIDArr(GetGormQuery(m), idArr, values)
}
func TxEditByIDArr[P BaseModel[M], M any](m P, tx *gorm.DB, idArr []uint64, values map[string]any) (int64, error) {
	return dbEditByIDArr(GetTxGormQuery(m, tx), idArr, values)
}
func dbDelByIDArr[P BaseModel[M], M any](m P, db *gorm.DB, idArr []uint64) (int64, error) {
	if len(idArr) == 0 {
		return 0, nil
	}
	res := db.Where("id in ?", idArr).Delete(m)
	return res.RowsAffected, res.Error
}
func DelByIDArr[P BaseModel[M], M any](m P, idArr []uint64) (int64, error) {
	return dbDelByIDArr(m, GetGormQuery(m), idArr)
}
func TxDelByIDArr[P BaseModel[M], M any](m P, tx *gorm.DB, idArr []uint64) (int64, error) {
	return dbDelByIDArr(m, GetTxGormQuery(m, tx), idArr)
}
func dbCreateRecord[P BaseModel[M], M any](m P, db *gorm.DB, record *P) (*P, error) {
	res := db.Create(record)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, errCreateFailed
	}
	return record, nil
}
func CreateRecord[P BaseModel[M], M any](m P, record *P) (*P, error) {
	return dbCreateRecord(m, GetGormQuery(m), record)
}
func TxCreateRecord[P BaseModel[M], M any](m P, tx *gorm.DB, record *P) (*P, error) {
	return dbCreateRecord(m, GetTxGormQuery(m, tx), record)
}
func dbCreateList[P BaseModel[M], M any](db *gorm.DB, list []P) ([]P, error) {
	res := db.Create(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}
func CreateList[P BaseModel[M], M any](m P, list []P) ([]P, error) {
	return dbCreateList(GetGormQuery(m), list)
}
func TxCreateList[P BaseModel[M], M any](m P, tx *gorm.DB, list []P) ([]P, error) {
	return dbCreateList(GetTxGormQuery(m, tx), list)
}
func ListRecord[P BaseModel[M], M any](m P, page, limit int, filter Filter, order map[string]string) ([]P, int64, error) {
	var count int64
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	list := make([]P, 0, limit)
	db := GetGormQuery(m)
	query, err := BuildFilterCond(m.GetFilterKeyMapDBField(), db, filter)
	if err != nil {
		return nil, count, err
	}
	dataQuery := query.WithContext(m.GetCtx())
	g := errgroup.Group{}
	g.Go(func() error {
		return query.Count(&count).Error
	})
	g.Go(func() error {
		dataQuery = BuildOrderCond(m.GetOrderKeyMapDBField(), dataQuery, order)
		return dataQuery.Offset(offset).Limit(limit).Find(&list).Error
	})
	return list, count, g.Wait()
}
