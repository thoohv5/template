#!/bin/bash

camelCase() {
    PARA=$1
    arr=(`echo $PARA | tr '_' ' '`)
    result=''
    for var in ${arr[@]}
    do
      firstLetter=`echo ${var:0:1} | awk '{print toupper($0)}'`
      otherLetter=${var:1}
      result=$result$firstLetter$otherLetter
    done

    firstResult=$(echo ${result:0:1} | tr '[A-Z]' '[a-z]')
    result=$firstResult${result:1}
    echo $result
}

firstLetterUpper() {
  str=$1
  firstLetter=`echo ${str:0:1} | awk '{print toupper($0)}'`
  otherLetter=${str:1}
  result=$firstLetter$otherLetter
  echo $result
}

readDir="./model"

for file_a in ${readDir}/*
do
    temp_file=`basename $file_a`

    dao_name=`basename $file_a .go`
    daoName=`camelCase "$dao_name"`
    DaoName=`firstLetterUpper "$daoName"`

    touch $temp_file
    cat > $temp_file << EOF
package dao

import (
	"context"

	"github.com/garyburd/redigo/redis"
	"github.com/go-kratos/kratos/pkg/log"
	"gorm.io/gorm"

	"med-common/app/service/dtx-prescription-service/api/sdk/common"
	"med-common/app/service/dtx-prescription-service/internal/dao/model"
	ecode "med-common/library/code"
	mmysql "med-common/library/module/mysql"
)

type (
	I${DaoName} interface {
		// One 查询一条
		One(ctx context.Context, condition *${DaoName}Condition, opts ...Option) (info *model.${DaoName}, err error)
		// List 查询多条
		List(ctx context.Context, condition *${DaoName}Condition, opts ...Option) (list []*model.${DaoName}, err error)
		// Add 添加
		Add(ctx context.Context, data *model.${DaoName}, opts ...Option) (err error)
		// Update 更新
		Update(ctx context.Context, condition *${DaoName}Condition, updateData map[string]interface{}, opts ...Option) (err error)
	}
	${daoName} struct {
		dao
	}
	${DaoName}Condition struct {
		*common.BaseRequest

		// 查询参数
		Id  int32
		Ids []int32
	}
)

// new${DaoName} 创建
func new${DaoName}(db mmysql.IDb, redis *redis.Pool) I${DaoName} {
	m := new(${daoName})
	m.model = new(model.${DaoName})
	m.db = db
	m.redis = redis
	return m
}

// 查询
func (m *${daoName}) findIds(ctx context.Context, condition *${DaoName}Condition, opts ...Option) (ids []int32, db *gorm.DB, err error) {
	ids = make([]int32, 0)
	conn, db := m.buildFilterQuery(ctx, condition, opts...)
	if err = conn.Pluck("id", &ids).Error; nil != err {
		err = m.toErr(ctx, err)
	}
	if len(ids) == 0 {
		err = ecode.MedErrDBRecordNotFound
		log.Debugc(ctx, "conn Pluck empty, condition:%+v", condition)
	}
	return
}

// One 查询一条
func (m *${daoName}) One(ctx context.Context, condition *${DaoName}Condition, opts ...Option) (info *model.${DaoName}, err error) {
	info = new(model.${DaoName})

	// 查询
	ids, db, err := m.findIds(ctx, condition, append([]Option{WithDb(m.db.GetRDb()), WithLimit(1)}, opts...)...)
	if nil != err {
		return
	}

	id := ids[0]

	// 缓存
	if err = m.Cache(ctx, id, info, func(ctx context.Context, id int32, info interface{}) error {
		db, _ = m.buildFilterQuery(ctx, &${DaoName}Condition{
			BaseRequest: &common.BaseRequest{},
			Id: id,
		}, WithDb(db))
		if err = db.Limit(1).Order("id ASC").Find(info).Error
			nil != err {
			return m.toErr(ctx, err)
		}
		return nil
	}); nil != err {
		return
	}
	return
}

// List 查询多条
func (m *${daoName}) List(ctx context.Context, condition *${DaoName}Condition, opts ...Option) (list []*model.${DaoName}, err error) {
	ids, initDb, err := m.findIds(ctx, condition, append([]Option{WithDb(m.db.GetRDb())}, opts...)...)
	if nil != err {
		return
	}

	// 查询
	list = make([]*model.${DaoName}, len(ids), cap(ids))

	// 缓存
	if err = m.CacheList(ctx, ids, &list, func(ctx context.Context, ids []int32, list interface{}) error {
		db, _ := m.buildFilterQuery(ctx, &${DaoName}Condition{
			BaseRequest: &common.BaseRequest{},
			Ids: ids,
		}, WithDb(initDb))
		if err = db.Find(list).Error; nil != err {
			return m.toErr(ctx, err)
		}
		return nil
	}); nil != err {
		return
	}
	return
}

// Add 添加
func (m *${daoName}) Add(ctx context.Context, data *model.${DaoName}, opts ...Option) (err error) {
	conn, _ := m.buildFilterQuery(ctx, nil, append([]Option{WithDb(m.db.GetWDb())}, opts...)...)
	if err = conn.Create(data).Error; nil != err {
		err = m.toErr(ctx, err)
	}
	return
}

// Update 更新
func (m *${daoName}) Update(ctx context.Context, condition *${DaoName}Condition, updateData map[string]interface{}, opts ...Option) (err error) {

	// 查询
	ids, db, err := m.findIds(ctx, condition, append([]Option{WithDb(m.db.GetWDb())}, opts...)...)
	if nil != err {
		return
	}

	// 删除
	db, _ = m.buildFilterQuery(ctx, &${DaoName}Condition{
		BaseRequest: &common.BaseRequest{},
			Ids: ids,
	}, WithDb(db))
	if err = db.Updates(updateData).Error; nil != err {
		err = m.toErr(ctx, err)
	}

	// 缓存
	if err = m.CacheRemove(ctx, ids); nil != err {
		return
	}

	return
}

// 参数构建
func (m *${daoName}) buildFilterQuery(ctx context.Context, condition *${DaoName}Condition, opts ...Option) (db *gorm.DB, initDb *gorm.DB) {
	// 处理公共参数
	if condition != nil && condition.BaseRequest != nil {
		opts = append(opts, WithBaseRequest(condition.BaseRequest))
	}
	// 预处理
	opts = append(opts, WithModel(m.model))
	db, initDb = m.dao.buildFilterQuery(ctx, opts...)

	if condition == nil {
		return
	}

	// 查询判断
	ids := condition.Ids
	if id := condition.Id; id > 0 {
		ids = append(ids, id)
	}
	if len(ids) > 0 {
		if len(ids) == 1 {
			db = db.Where("id = ?", ids[0])
		} else {
			db = db.Where("id IN(?)", ids)
		}
	}

	return
}


EOF

cat >> ./dao.go << EOF
func (d *dao) Get${DaoName}() I${DaoName} {
	return new${DaoName}(d.db, d.redis)
}
EOF
done
