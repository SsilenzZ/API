package comments

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var comment = []Comments{{Post: 1, ID: 1, Name: "test", Email: "test@mail.com", Body: "some body"},
	{Post: 1, ID: 2, Name: "test 2", Email: "test@mail.com", Body: "some body"},
	{Post: 1, ID: 3, Name: "test 3", Email: "test@mail.com", Body: "some another body"}}

func TestGetAllComments(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	commentRepo := ProvideCommentRepository(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `comments`").
		WillReturnRows(
			sqlmock.NewRows([]string{"post", "id", "name", "email", "body"}).
				AddRow(comment[0].Post, comment[0].ID, comment[0].Name, comment[0].Email, comment[0].Body).
				AddRow(comment[1].Post, comment[1].ID, comment[1].Name, comment[1].Email, comment[1].Body).
				AddRow(comment[2].Post, comment[2].ID, comment[2].Name, comment[2].Email, comment[2].Body))

	res := commentRepo.GetAllComments()
	require.Equal(t, comment, res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	commentRepo := ProvideCommentRepository(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `comments` WHERE `comments`.`id` = ? ORDER BY `comments`.`id` LIMIT 1").
		WithArgs("3").
		WillReturnRows(
			sqlmock.NewRows([]string{"post", "id", "name", "email", "body"}).
				AddRow(comment[2].Post, comment[2].ID, comment[2].Name, comment[2].Email, comment[2].Body))

	res, err := commentRepo.GetComment("3")

	require.NoError(t, err)
	require.Equal(t, comment[2], res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	comment := map[string]interface{}{"Post": 1, "id": 3, "name": "test 3", "email": "test@mail.com", "body": "some another body"}

	commentRepo := ProvideCommentRepository(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `comments` (`post`,`body`,`email`,`id`,`name`) "+
		"VALUES (?,?,?,?,?)").WithArgs(comment["Post"], comment["body"], comment["email"], comment["id"],
		comment["name"]).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.CreateComment(comment)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	comment := map[string]interface{}{"Post": 1, "id": 3, "name": "test 3", "email": "test@mail.com", "body": "some another body"}

	commentRepo := ProvideCommentRepository(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `comments` SET `post`=?,`body`=?,`email`=?,`name`=? "+
		"WHERE `id` = ?").WithArgs(comment["Post"], comment["body"], comment["email"], comment["name"],
		comment["id"]).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.UpdateComment(comment)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	commentRepo := ProvideCommentRepository(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `comments` WHERE `comments`.`id` = ?").
		WithArgs("3").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.DeleteComment("3")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
