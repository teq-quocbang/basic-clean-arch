package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/quocbang/learn/repository/auth"
	"github.com/quocbang/learn/repository/interfaces"
)

type DB struct {
	db *gorm.DB
}

func (d *DB) Auth() interfaces.Auth {
	return auth.NewAuthRepository(d.db)
}

func (d *DB) Todo() interfaces.Todo {
	return nil
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func NewDatabaseConnection(d Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		d.Host,
		d.UserName,
		d.Password,
		d.Name,
		d.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewRepository(d Database) (interfaces.Repository, error) {
	db, err := NewDatabaseConnection(d)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}
