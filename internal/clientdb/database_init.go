package clientdb

import (
	"elichika/internal/generic"
	"log"

	"elichika/internal/config"
	"elichika/internal/utils"

	"bufio"
	"fmt"
	"os"

	"xorm.io/xorm"
)

type Migration struct {
	Name     string
	Path     string
	Database string
}

func discoverMigrations(path string, migrations *generic.List[Migration]) {
	files, err := os.ReadDir(path)
	utils.CheckErr(err)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		migrations.Append(Migration{
			Name:     name,
			Path:     path + name,
			Database: name[4 : len(name)-4],
		})
	}
}

// note that this is subject to change, do not depend on it too much
func initLocale(locale string) {
	dbDir := fmt.Sprint("db/", locale, "/")

	var migrations generic.List[Migration]
	var err error

	// Get locale specific migrations in the format sql/db/<locale>/<order>.filename.sql
	discoverMigrations(fmt.Sprint(config.AssetPath, "sql/", locale, "/"), &migrations)

	// Get locale agnostic migrations in the format sql/db/<order>.filename.sql
	discoverMigrations(fmt.Sprint(config.AssetPath, "sql/"), &migrations)

	needUpdate := map[string]bool{}
	engines := map[string]*xorm.Engine{}
	sessions := map[string]*xorm.Session{}

	for _, migration := range migrations.Slice {
		dbName := migration.Database
		need, exists := needUpdate[dbName]
		if !exists {
			needUpdate[dbName] = isNotChanged(dbDir + dbName)
			need = needUpdate[dbName]
		}
		if !need {
			continue
		}
		session, exists := sessions[dbName]
		if !exists {
			engines[dbName], err = xorm.NewEngine("sqlite", config.AssetPath+dbDir+dbName)
			utils.CheckErr(err)
			engines[dbName].SetMaxOpenConns(50)
			engines[dbName].SetMaxIdleConns(10)
			sessions[dbName] = engines[dbName].NewSession()
			session = sessions[dbName]
			session.Begin()
		}
		log.Println("Running SQL migration: ", migration.Name)

		f, err := os.Open(migration.Path)
		utils.CheckErr(err)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			_, err = session.Exec(scanner.Text())
			utils.CheckErr(err)
		}
		utils.CheckErr(scanner.Err())
	}

	// Close all database sessions
	for _, session := range sessions {
		err := session.Commit()
		utils.CheckErr(err)
		session.Close()
	}
}

func databaseInit() {
	initLocale("gl")
	initLocale("jp")
}
