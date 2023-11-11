package unidb

import (
	"fmt"
)

type IUniDB interface {
	// Find is a method what can be used to find any data from database
	Find(dest interface{}, query string, args ...interface{}) error
	// Create is a method what can be used to create any data to database
	Create()
}

type UniDB struct {
	repo Manager
}

func New(repo Manager) IUniDB {
	return &UniDB{
		repo: repo,
	}
}

func (uni *UniDB) Find(dest interface{}, query string, args ...interface{}) error {
	err := uni.repo.Find(dest, query, args...)
	if err != nil {
		return fmt.Errorf("unidb.Find: %w", err)
	}
	return nil
}

func (uni *UniDB) Create() {
	fmt.Println("Create")
}
