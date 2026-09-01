package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
)

func RunMigrations() error {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "internal")
	sslmode := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error opening gorm connection for migrations: %w", err)
	}

	gormDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	gormDB.Exec("CREATE EXTENSION IF NOT EXISTS \"pgcrypto\"")

	err = gormDB.AutoMigrate(
		&entity.UsersEntity{},
		&entity.RefreshTokenEntity{},
		&entity.ContactEntity{},
		&entity.GroupEntity{},
		&entity.GroupMemberEntity{},
		&entity.CostEntity{},
	)
	if err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	gormDB.Exec(`CREATE INDEX IF NOT EXISTS idx_cost_splits_gin ON cost_entities USING GIN (splits)`)

	// Recreate all FKs with the correct ON DELETE behavior.
	// AutoMigrate does not alter existing constraints, so we drop and recreate them.
	type fkFix struct {
		table      string
		constraint string
		column     string
		references string
		onDelete   string
	}
	fixes := []fkFix{
		{"contact_entities", "fk_users_entities_contacts", "owner_id", "users_entities(id)", "CASCADE"},
		{"cost_entities", "fk_users_entities_costs", "user_id", "users_entities(id)", "CASCADE"},
		{"group_member_entities", "fk_group_entities_members", "group_id", "group_entities(id)", "CASCADE"},
		{"group_member_entities", "fk_contact_entities_group_members", "contact_id", "contact_entities(id)", "CASCADE"},
	}
	for _, f := range fixes {
		gormDB.Exec(fmt.Sprintf(`ALTER TABLE %s DROP CONSTRAINT IF EXISTS %s`, f.table, f.constraint))
		gormDB.Exec(fmt.Sprintf(
			`ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s ON DELETE %s`,
			f.table, f.constraint, f.column, f.references, f.onDelete,
		))
	}

	sqlDB, _ := gormDB.DB()
	sqlDB.Close()

	return nil
}
