package repository

import "github.com/jinzhu/gorm"

// UnitOfWork Contain DB, Readonly Flag
// UOW used for DB Transaction Management
type UnitOfWork struct {
	DB        *gorm.DB
	readOnly  bool
	committed bool
}

// NewUnitOfWork return new object of UnitOfWork
func NewUnitOfWork(db *gorm.DB, readonly bool) *UnitOfWork {
	if readonly {
		return &UnitOfWork{
			DB:        db.New(),
			readOnly:  readonly,
			committed: false,
		}
	}

	return &UnitOfWork{
		DB:        db.Begin(),
		readOnly:  readonly,
		committed: false,
	}
}

// Commit the transaction
func (uow *UnitOfWork) Commit() {
	if !uow.readOnly {
		uow.DB.Commit()
	}
	uow.committed = true
}

// Complete Rollback the transaction in case of failure
func (uow *UnitOfWork) Complete() {
	if !uow.committed && !uow.readOnly {
		uow.DB.Rollback()
	}
}
