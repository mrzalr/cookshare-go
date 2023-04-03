package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	) + "?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
