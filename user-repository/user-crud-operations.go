package user_repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-training/persistence-layer-for-users-and-recipies/db"
	"go-training/persistence-layer-for-users-and-recipies/entities"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type UserRepository struct{}

func(ur UserRepository) FindAll() ([]entities.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rows, err := db.DB.QueryContext(ctx,"SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("database query failed\n")
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {

		var user entities.User
		err := rows.Scan(&user.Id, &user.UserName, &user.LoginName, &user.Password, &user.Gender, &user.UserRole, &user.PictureURL, &user.ShortDescr,
		&user.Active, &user.Created, &user.Modified)

		if err != nil {
			return nil, fmt.Errorf("error scanning database record\n")
		}
		users = append(users, user)
	}
	return users, nil
}

func(ur UserRepository) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entities.User, error){

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var query string
	offset := (pageNumber - 1) * pageSize
	if ascending {
		 query = fmt.Sprintf("SELECT * FROM users ORDER BY %s LIMIT %d OFFSET %d ", sortingAttribute, pageSize, offset)
	} else {
		query = fmt.Sprintf("SELECT * FROM users ORDER BY %s DESC LIMIT %d OFFSET %d ", sortingAttribute, pageSize, offset)
	}

	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("database query failed\n")
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {

		var user entities.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.LoginName, &user.Password, &user.Gender, &user.UserRole, &user.PictureURL, &user.ShortDescr,
			&user.Active, &user.Created, &user.Modified); err != nil {
			return nil, fmt.Errorf("database query failed\n")
		}
		users = append(users, user)
	}
	return users, nil
}

func(ur UserRepository) FindByID(userId uint) (user entities.User, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = db.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", userId).
		Scan(&user.Id, &user.UserName, &user.LoginName, &user.Password,
		&user.Gender, &user.UserRole, &user.PictureURL, &user.ShortDescr, &user.Active, &user.Created, &user.Modified)
	switch {
		case err == sql.ErrNoRows:
			return user, fmt.Errorf("user with id %d does not exist\n", userId)
		case err != nil:
			return user, fmt.Errorf("error quering the database\n")
	}
	return
}

func(ur UserRepository) FindByUsername(username string) (user entities.User, err error)  {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = db.DB.QueryRowContext(ctx,"SELECT * FROM users WHERE userName = ?", username).
		Scan(&user.Id, &user.UserName, &user.LoginName, &user.Password,
		&user.Gender, &user.UserRole, &user.PictureURL, &user.ShortDescr, &user.Active, &user.Created, &user.Modified)
	switch {
		case err == sql.ErrNoRows:
			return user, fmt.Errorf("user with username %s does not exist\n", username)
		case err != nil:
			return user, fmt.Errorf("database query failed\n")
	}
	return
}

func(ur UserRepository) Create(user *entities.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generating hashed password\n")
	}
	user.Password = string(hashedPassword)
	res, err := db.DB.Exec("INSERT INTO users(userName, loginName, password, gender, userRole, shortDescription, activated, Created, Modified) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.UserName, user.LoginName, user.Password, user.Gender, user.UserRole, user.ShortDescr, user.Active, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("database query failed\n")
	}

	numRows, err :=res.RowsAffected()
	if err != nil || numRows != 1 {
		return fmt.Errorf("error creating user %s\n", user.UserName)
	}
	insId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting id of inserted user: %s\n", user.UserName)
	}
	user.Id = uint(insId)

	return nil
}

func(ur UserRepository) Update(user *entities.User) error {

	rows := db.DB.QueryRow("SELECT id, userName, password FROM users WHERE id = ?", user.Id)

	var dtbsId int
	var dtbsUserName string
	var hashedDtbsPsswd string
	err := rows.Scan(&dtbsId, &dtbsUserName, &hashedDtbsPsswd)
	switch {
	case err == sql.ErrNoRows:
		return fmt.Errorf("no user with id %d exists", user.Id)
	case err != nil:
		return fmt.Errorf("database query failed\n")
	}

	if user.UserName != dtbsUserName {
		return fmt.Errorf("username cannot be changed")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedDtbsPsswd), []byte(user.Password)); err != nil {
		hashedPsswd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing new password\n")
		}
		user.Password = string(hashedPsswd)
	} else {
		user.Password = hashedDtbsPsswd
	}


	res, err := db.DB.Exec("UPDATE users SET loginName=?, password=?, gender=?, userRole=?, shortDescription=?, activated=?, modified=? WHERE id=?",
		user.LoginName, user.Password, user.Gender, user.UserRole, user.ShortDescr, user.Active, time.Now(), user.Id)
	if err != nil {
		return fmt.Errorf("database query failed\n")
	}
	numRows, err := res.RowsAffected()
	if err != nil || numRows != 1 {
		return fmt.Errorf("user updata failed: no rows affected\n")
	}

	return nil
}

func(ur UserRepository) DeleteByID(userID uint) (user entities.User, err error) {

	user, err = ur.FindByID(userID)
	if err != nil {
		return user, fmt.Errorf("error deleting user -> %s", err)
	}

	res, err := db.DB.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return user, fmt.Errorf("database query failed\n")
	}
	numRows, err := res.RowsAffected()
	if err != nil || numRows != 1 {
		return user, fmt.Errorf("error deleting user: zero affected rows\n")
	}

	return
}

func(ur UserRepository) Count() int {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var count int
	rows, err := db.DB.QueryContext(ctx, "SELECT COUNT(*)  FROM users")
	if err != nil {
		log.Println(err)
	}

	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Println(err)
		}
	}

	return count
}


