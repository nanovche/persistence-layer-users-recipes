package logging_functionality

import (
	"fmt"
	"go-training/persistence-layer-for-users-and-recipies/db"
	"go-training/persistence-layer-for-users-and-recipies/entities"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct{
	loggedUser entities.User
}

func (lc *LoginController) Login(username, password string) (user entities.User, err error) {

	row := db.DB.QueryRow("SELECT * FROM users WHERE userName=? ",username)

	if err := row.Scan(&user.Id, &user.UserName, &user.LoginName, &user.Password, &user.Gender, &user.UserRole, &user.PictureURL,
		&user.ShortDescr, &user.Active, &user.Created, &user.Modified); err != nil {
		return user, fmt.Errorf("error scanning record")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, fmt.Errorf("passwords do not match")
	}
	lc.loggedUser = user
	fmt.Printf("\n\nUser %s successfully logged in.\n", username)
	return user,nil
}

func (lc *LoginController) Logout() {
	userName := lc.loggedUser.UserName
	lc.loggedUser = entities.User{}
	fmt.Printf("User %s successfully logged out.\n", userName)
}

func (lc *LoginController) GetLoggedUser() entities.User {
	fmt.Printf("Logged user: %s\n", lc.loggedUser.UserName)
	return lc.loggedUser
}

