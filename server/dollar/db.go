package dollar

import (
	"context"
	"database/sql"
	"github.com/kenesparta/fullcycle-client-server-api/server/constants"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DatabaseConnection struct {
	db  *sql.DB
	ctx context.Context
}

func (dc *DatabaseConnection) connect() {
	db, connErr := sql.Open("sqlite3", constants.DbFileName)
	if connErr != nil {
		log.Printf("failed to connect: %v", connErr)
		return
	}

	dc.db = db
}

func (dc *DatabaseConnection) close() {
	err := dc.db.Close()
	if err != nil {
		log.Printf("failed to close database: %v", err)
	}
}

func (dc *DatabaseConnection) crateTable() {
	_, createErr := dc.db.Exec(`
CREATE TABLE cotacao (
    id CHARACTER(36) PRIMARY KEY, 
    code TEXT,
    codein TEXT,
    name TEXT,
    high DOUBLE,
    low DOUBLE,
    varBid DOUBLE,
    pctChange DOUBLE,
    bid DOUBLE,
    ask DOUBLE,
    timestamp TEXT,
    create_date TEXT
)`)
	if createErr != nil {
		log.Printf("Failed to create table: %v", createErr)
	}
}

func (dc *DatabaseConnection) insert(conv Cotacao) {
	ctx, cancel := context.WithTimeout(dc.ctx, constants.MaxDbTimeout)
	defer cancel()

	_, err := dc.db.ExecContext(ctx, `
INSERT INTO cotacao (id,code,codein,name,high,low,varBid,pctChange,bid,ask,timestamp,create_date) 
VALUES (?,?,?,?,?,?,?,?,?,?,?,?)
`,
		conv.Id,
		conv.Code,
		conv.Codein,
		conv.Name,
		conv.High,
		conv.Low,
		conv.VarBid,
		conv.PctChange,
		conv.Bid,
		conv.Ask,
		conv.Timestamp,
		conv.CreateDate)
	if err != nil {
		log.Printf("error in the database: %v", err)
	}
}

func (dc *DatabaseConnection) read() ([]Cotacao, error) {
	ctx, cancel := context.WithTimeout(dc.ctx, constants.MaxDbTimeout)
	defer cancel()

	var cotList []Cotacao
	rows, err := dc.db.QueryContext(ctx, `SELECT * FROM cotacao`)
	for rows.Next() {
		var cotItem Cotacao
		rows.Scan(
			&cotItem.Id,
			&cotItem.Code,
			&cotItem.Codein,
			&cotItem.Name,
			&cotItem.High,
			&cotItem.Low,
			&cotItem.VarBid,
			&cotItem.PctChange,
			&cotItem.Bid,
			&cotItem.Ask,
			&cotItem.Timestamp,
			&cotItem.CreateDate,
		)
		cotList = append(cotList, cotItem)
	}
	if err != nil {
		log.Printf("Failed to query data: %v", err)
		return nil, err
	}

	return cotList, nil
}

func CreateTables(ctx context.Context) {
	dbConn := DatabaseConnection{
		db:  nil,
		ctx: ctx,
	}
	dbConn.connect()
	defer dbConn.close()
	dbConn.crateTable()
}

func Save(ctx context.Context, cot Cotacao) error {
	dbConn := DatabaseConnection{
		db:  nil,
		ctx: ctx,
	}
	dbConn.connect()
	defer dbConn.close()

	dbConn.insert(cot)
	return nil
}

func Read(ctx context.Context) ([]Cotacao, error) {
	dbConn := DatabaseConnection{
		db:  nil,
		ctx: ctx,
	}
	dbConn.connect()
	defer dbConn.close()
	return dbConn.read()
}
