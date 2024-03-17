package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const (
	WorkloadWrite = "write"
	WorkloadRead  = "read"
)

var ErrNoRowsAffected = errors.New("no rows affected")

type Replication struct {
	master *DB
	slave  *DB
}

func NewReplication(master, slave *DB) *Replication {
	return &Replication{
		master: master,
		slave:  slave,
	}
}

func (r *Replication) GetInstance(workload string) *DB {
	if workload == WorkloadWrite {
		return r.GetMasterInstance()
	}
	return r.GetSlaveInstance()
}

func (r *Replication) GetMasterInstance() *DB {
	return r.master
}

func (r *Replication) GetSlaveInstance() *DB {
	return r.slave
}

func (r *Replication) Pagination(page, limit int) func(query string, args ...interface{}) (*sql.Rows, error) {
	return func(query string, args ...interface{}) (*sql.Rows, error) {
		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		offset := (page - 1) * limit
		return r.master.QueryContext(context.Background(), fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset), args...)
	}
}

// func (r *Replication) HandleUpdateResult(tx *sql.Tx, result *pq.Result, rowsAffected int64) error {
// 	if result.Err() != nil || result.RowsAffected() != rowsAffected {
// 		if tx != nil {
// 			_ = tx.Rollback() // handle potential rollback error
// 		}
// 		err := result.Err()
// 		if result.RowsAffected() != rowsAffected {
// 			return ErrNoRowsAffected
// 		}
// 		return err
// 	}
// 	return nil
// }

func (r *Replication) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := r.master.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	return tx, nil
}
