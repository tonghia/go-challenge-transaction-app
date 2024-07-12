package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sunary/sqlize"
	"github.com/tonghia/go-challenge-transaction-app/internal/model"
)

func AllModels() []interface{} {
	return []interface{}{
		model.User{},
		model.Account{},
		model.AccountTransaction{},
	}
}

func main() {
	migrationFolder := "migrations/"
	sqlLatest := sqlize.NewSqlize(sqlize.WithSqlTag("sql"),
		sqlize.WithMigrationFolder(migrationFolder),
		sqlize.WithCommentGenerate())

	var ms []interface{}
	ms = append(ms, AllModels()...)
	err := sqlLatest.FromObjects(ms...)
	if err != nil {
		log.Fatal("sqlize FromObjects", err)
	}
	sqlVersion := sqlLatest.HashValue()

	sqlMigrated := sqlize.NewSqlize(sqlize.WithMigrationFolder(migrationFolder))
	err = sqlMigrated.FromMigrationFolder()
	if err != nil {
		log.Fatal("sqlize FromMigrationFolder", err)
	}

	sqlLatest.Diff(*sqlMigrated)

	fmt.Println("sql version", sqlVersion)

	fmt.Println("\n\n### migration up")
	migrationUp := sqlLatest.StringUp()
	fmt.Println(migrationUp)

	fmt.Println("\n\n### migration down")
	fmt.Println(sqlLatest.StringDown())

	initVersion := os.Args[1] == "$init"
	if initVersion {
		log.Println("write to init version")
		err = sqlLatest.WriteFilesVersion("init version", 0, false)
		if err != nil {
			log.Fatal("sqlize WriteFilesVersion", err)
		}
	}

	if len(os.Args) > 1 {
		log.Println("write to file", os.Args[1])
		// err = sqlLatest.WriteFilesWithVersion(os.Args[1], sqlVersion, false)
		err = sqlLatest.WriteFiles(os.Args[1])
		if err != nil {
			log.Fatal("sqlize WriteFiles", err)
		}
	}
}
