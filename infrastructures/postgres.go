package infrastructures

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) *gorm.DB {
	// Connect to database
	fmt.Printf("Connecting to database: %s\n", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	// if you have replica, you can configure it with:
	// dbresolver will automatically select replicas (replicaDsn) for read operation
	// and select master (dsn) for write operation

	// err = db.Use(dbresolver.Register(dbresolver.Config{
	// 	Replicas: []gorm.Dialector{
	// 		postgres.Open(replicaDsn),
	// 	},
	// 	TraceResolverMode: true,
	// }))
	// if err != nil {
	// 	panic(err)
	// }

	return db
}
