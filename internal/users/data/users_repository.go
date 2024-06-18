package data

import (
	"github.com/Uncensored-Developer/buzz/internal/users/models"
	"github.com/Uncensored-Developer/buzz/pkg/bun_mysql"
	"github.com/Uncensored-Developer/buzz/pkg/repository"
	"github.com/uptrace/bun"
	"time"
)

type IUserRepository interface {
	repository.IRepository[models.User]
}

func NewUserRepository(db bun.IDB) IUserRepository {
	return bun_mysql.NewBunRepository[models.User](db)
}

func UserWithEmail(email string) repository.SelectCriteria {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where("email = ?", email)
	}
}

func UserWithID(id int64) repository.SelectCriteria {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where("id = ?", id)
	}
}

func UserWithEmailAndPassword(email, password string) repository.SelectCriteria {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where("email = ?", email).Where("password = ?", password)
	}
}

func UsersWithinDobRange(min, max time.Time) repository.SelectCriteria {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where("dob >= ? AND dob <= ?", min, max)
	}
}

func UsersWithGender(gender string) repository.SelectCriteria {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where("gender = ?", gender)
	}
}

func UsersExcludingID(id int64) repository.SelectCriteria {
	return func(query *bun.SelectQuery) *bun.SelectQuery {
		return query.Where("id != ?", id)
	}
}
