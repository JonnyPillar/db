package sqladapter

type Collection interface {
	Name() string
	Exists() bool
	Database() Database
}

type BaseCollection struct {
	database  Database
	tableName string
}

// NewCollection returns a collection with basic methods.
func NewCollection(d Database, tableName string) Collection {
	return &BaseCollection{database: d, tableName: tableName}
}

// Name returns the name of the table.
func (c *BaseCollection) Name() string {
	return c.tableName
}

// Exists returns true if the collection exists.
func (c *BaseCollection) Exists() bool {
	if err := c.Database().TableExists(c.Name()); err != nil {
		return false
	}
	return true
}

// Database returns the database session that backs the collection.
func (c *BaseCollection) Database() Database {
	return c.database
}
