package configs

import (
	"log"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dialector := selectDialector(Env.Driver, Env.URL)

    db, err := gorm.Open(dialector, &gorm.Config{
        Logger: logger.Default.LogMode(logger.Warn),
    })
    if err != nil {
        log.Fatal("❌ Failed to connect to database: ", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("❌ Failed to get sql.DB from GORM: ", err)
    }

    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    DB = db
}

func selectDialector(driver, dsn string) gorm.Dialector {

	switch driver {

		case "mysql", "sqlx-mysql":
			return mysql.Open(dsn)

		case "postgres":
			return postgres.Open(dsn)

		default:
			log.Fatalf("❌ Unsupported database driver: %s", driver)
			return nil

    }

}
