package types

import (
	"database/sql"
)

// MySQLDBHandlerInterface contains the implementable methods for the MySQL DB handler
type MySQLDBHandlerInterface interface {
	Connect(params ConnectionParams) error
	ConnectViaSSH(paramsSSH SSHConnectionParams, params ConnectionParams) error
	Execute(stmt string, model interface{}) (sql.Result, error)
	Query(qstmt string, model interface{}, bindModel interface{}) error
	QueryRow(qstmt string, model interface{}, bindModel interface{}) error
}
