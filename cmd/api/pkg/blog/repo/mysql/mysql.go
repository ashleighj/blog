package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"api/config"
	"api/logger"
	err "api/service_error"
)

// Mysql holds data structures associated
// with a mysql database connection
type Mysql struct {
	db     *sql.DB
	config config.DataSource
}

// New returns an instance of Mysql connection
func New(ctx context.Context,
	conf config.DataSource) (db *Mysql, e error) {

	db = &Mysql{config: conf}

	e = db.Connect(ctx)
	if e != nil {
		logger.Error(ctx, err.DatastoreInit(e))
	}
	return
}

// Connect connects to mysql database
func (m *Mysql) Connect(ctx context.Context) (e error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		m.config.User,
		m.config.Password,
		m.config.Host,
		m.config.Port,
		m.config.Database)

	conn, e := sql.Open("mysql", connStr)
	if e != nil {
		logger.Error(ctx, err.DatastoreInit(e))
		return
	}

	e = conn.Ping()
	if e != nil {
		logger.Error(ctx, err.DatastoreInit(e))
		return
	}

	m.db = conn
	return nil
}

// Close closes current connection to database
func (m *Mysql) Close(ctx context.Context) (e error) {
	e = m.db.Close()
	if e != nil {
		logger.Error(ctx, err.Datastore(e))
	}

	return
}
