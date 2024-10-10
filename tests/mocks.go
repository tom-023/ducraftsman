package tests

import "database/sql"

// MockDBManager モック用のDBManager実装
type MockDBManager struct {
	ConnectFunc    func(string, string, string, string) (*sql.DB, error)
	CreateUserFunc func(*sql.DB, string, string, string) error
}

func (m *MockDBManager) Connect(rootUser, rootPassword, host, dbName string) (*sql.DB, error) {
	return m.ConnectFunc(rootUser, rootPassword, host, dbName)
}

func (m *MockDBManager) CreateUser(db *sql.DB, username, password, privileges string) error {
	return m.CreateUserFunc(db, username, password, privileges)
}
