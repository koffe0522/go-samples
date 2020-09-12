package database

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// DB struct
type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

// Book struct
type book struct {
	gorm.Model
	Title string `json:"title"`
}

// Db is instance
var Db *DB

// Init is initializes databese
func Init() {
	d := &DB{
		Host:     "go_rest_mysql:3306",
		Username: "docker",
		Password: "docker",
		DBName:   "test_database",
	}
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

	db.SingularTable(true)

	// Migrate the schema
	db.AutoMigrate(&book{})
	d.Connection = db
	Db = d
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

// Connect begins a connection
func (db *DB) Connect() *gorm.DB {
	return db.Connection
}

// Close close database
func (db *DB) Close() {
	db.Close()
}
