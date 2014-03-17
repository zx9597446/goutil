//db util functions
//drop dbutil.go to source directory
//modify package name
//this package can define sql statement by:
//var stmt = newStmt("select * from table")
//then you can prepare all statements by:
//prepareAllStmts(db)
package util

import "database/sql"

type dbStmt struct {
	sql  string
	stmt *sql.Stmt
}

var allStmts = make([]*dbStmt, 0)

func newStmt(sql string) *dbStmt {
	r := dbStmt{sql, nil}
	allStmts = append(allStmts, &r)
	return &r
}

func (s *dbStmt) prepare(db *sql.DB) {
	var err error
	s.stmt, err = db.Prepare(s.sql)
	panicfIf(err, "prepare sql statement [%s] failed: %s\n", s.sql, err)
}

func prepareAllStmts(db *sql.DB) {
	for _, v := range allStmts {
		v.prepare(db)
	}
}
