package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
// NullTime handles the time that can be null and
// needs to be persisted with the database/sql package
type NullTime struct {
	sql.NullTime
}

// MarshalJSON converts the struct to a friendly json output
func (n NullTime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Time)
	}

	return []byte(`null`), nil
}

// Connect returns a new instance of Gorm database using the configuration
// passed by argument.
func Connect() *gorm.DB {
	cfg := config.Config()
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.Net = "tcp"
	mysqlConfig.DBName = cfg["MYSQL_DB_SCHEMA"]
	mysqlConfig.User = cfg["MYSQL_DB_USER"]
	mysqlConfig.Passwd = cfg["MYSQL_DB_PASSWORD"]
	mysqlConfig.Addr = cfg["MYSQL_DB_ENDPOINT"]
	mysqlConfig.Params = map[string]string{"charset": "utf8mb4"}
	mysqlConfig.ParseTime = true
	dsn := mysqlConfig.FormatDSN()

	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})

	if err != nil {
		fmt.Errorf("Erro ao conectar com o Banco de Dados.")
	}

	return db
}

var gormDBCacheVar *gorm.DB

// GormDB returns a new instance of Gorm database using the configuration
// passed by argument.
func GormDBCache() *gorm.DB {
	if gormDBCacheVar == nil {
		db := Connect()
		gormDBCacheVar = db
	}
	return gormDBCacheVar
}
*/

const (
	DB_NAME = "instagram_bot"
	DB_HOST = "172.17.0.2"
	DB_USER = "root"
	DB_PASS = "root"
	DB_PORT = "3306"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("erro ao conectar com database")
	}
	return db
}
