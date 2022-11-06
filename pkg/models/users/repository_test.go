package users

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	commentRepo := ProvideUserRepository(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT 1").
		WillReturnRows(
			sqlmock.NewRows([]string{"email", "password", "name"}).
				AddRow("test@mail.com", "simplepass", "test"))

	commentRepo.GetEmail(1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	rs := sqlmock.NewRows([]string{"email", "password", "name"})

	commentRepo := ProvideUserRepository(gdb)

	mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1").
		WithArgs("test@mail.com").WillReturnRows(rs)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`email`,`password`,`name`) VALUES (?,?,?)").WithArgs("test@mail.com",
		"simplepass", "google").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.GetUser("test@mail.com", "simplepass")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHashedPassword(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	rs := sqlmock.NewRows([]string{"email", "password", "name"})

	commentRepo := ProvideUserRepository(gdb)

	mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` " +
		"LIMIT 1").WithArgs("test@mail.com").WillReturnRows(rs)

	commentRepo.GetHashedPassword("test@mail.com", "simplepass")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	rs := sqlmock.NewRows([]string{"email", "password", "name"})

	commentRepo := ProvideUserRepository(gdb)

	mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1").
		WithArgs("test@mail.com").WillReturnRows(rs)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`email`,`password`,`name`) VALUES (?,?,?)").
		WithArgs("test@mail.com", "simplepass", "test").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.CreateUser("test@mail.com", "simplepass", "test")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
