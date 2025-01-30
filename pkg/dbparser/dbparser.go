package dbparser

import (
	"github.com/Valentin-Kaiser/go-dbase/dbase"
	"go.uber.org/zap"
	"os"
	"uber_eats/config"
)

type DbParser struct {
	dbfFile string
	dbase   dbase.Table
	logger  *zap.SugaredLogger
	table   *dbase.File
}

// NewDbParser creates a new DbParser
func NewDbParser(dbConfig config.DBConfig, logger zap.SugaredLogger) DbParser {
	dbase.Debug(dbConfig.EnableDebug, os.Stdout)

	// Open the example database table.
	table, err := dbase.OpenTable(&dbase.Config{
		Filename:   dbConfig.DataDir + "/" + dbConfig.StorageFile,
		TrimSpaces: true,
		Untested:   dbConfig.IgnoreVersionTest,
	})
	if err != nil {
		panic(err)
	}
	//defer table.Close()
	return DbParser{
		dbfFile: dbConfig.DataDir + "/" + dbConfig.StorageFile,
		logger:  logger.Desugar().Named("DbParser").Sugar(),
		table:   table,
	}
}

// Parse parses the dbf file
func (dbp *DbParser) Parse() {
	dbp.logger.Debug("Parsing dbf file")

}
