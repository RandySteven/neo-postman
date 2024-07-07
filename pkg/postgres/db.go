package postgres

import (
	"context"
	"database/sql"
	"fmt"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	"github.com/RandySteven/neo-postman/pkg/config"
	"github.com/RandySteven/neo-postman/queries"
	"github.com/RandySteven/neo-postman/repositories"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Repositories struct {
	TestDataRepo  repositories_interfaces.TestDataRepository
	JiraIssueRepo repositories_interfaces.JiraIssueRepository
	db            *sql.DB
}

func NewRepositories(config *config.Config) (*Repositories, error) {
	if config.Postgres.Port != "" {
		config.Postgres.Host = fmt.Sprintf("%s:%s", config.Postgres.Host, config.Postgres.Port)
	}
	conn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require",
		config.Postgres.DbUser,
		config.Postgres.DbPass,
		config.Postgres.Host,
		config.Postgres.DbName,
	)
	log.Println(conn)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(8)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(8 * time.Minute)
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return &Repositories{
		TestDataRepo:  repositories.NewTestDataRepository(db),
		JiraIssueRepo: repositories.NewJiraRepository(db),
		db:            db,
	}, nil
}

func initTableMigration() []queries.MigrationQuery {
	return []queries.MigrationQuery{
		queries.CreateTestDataTable,
		queries.CreateJiraIssueTable,
	}
}

func (r *Repositories) Migration(ctx context.Context) {
	migrationQueries := initTableMigration()

	for _, migration := range migrationQueries {
		_, err := r.db.ExecContext(ctx, migration.ToString())
		if err != nil {
			log.Fatalf("Error in migration : %s \n", err.Error())
			return
		}
	}
}
