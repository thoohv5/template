package dbx

import (
	"errors"
	"fmt"
	
	"github.com/thoohv5/template/internal/ent"
)

func ExecTrans(tx *ent.Tx, ts ...tran) (err error) {
	defer func() {
		// 捕获异常
		if pcm := recover(); nil != pcm {
			err = errors.New(fmt.Sprintf("%+v", pcm))
		}
		if nil != err {
			// 回滚
			if err = tx.Rollback(); nil != err {
				return
			}
			return
		}
		// 提交
		if err = tx.Commit(); nil != err {
			return
		}
	}()
	for _, t := range ts {
		err = t(tx)
		if err != nil {
			return
		}
	}
	return
}

type tran = func(tx *ent.Tx) error
