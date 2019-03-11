package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/joshsoftware/emp/empServer/db"
)

var (
	createUserQuery   = `INSERT INTO  USERS (age, first_name, last_name,email) VALUES ($1, $2, $3,$4) RETURNING id`
	readUserQuery     = `SELECT * FROM USERS WHERE ID = $1;`
	updateUserQuery   = `UPDATE users SET first_name = $2, last_name = $3 WHERE id = $1;`
	deleteUserQuery   = `DELETE FROM users WHERE id = $1;`
	readAllUsersQuery = `SELECT * FROM USERS;`
)

func createService(data User) error {
	db := db.GetDB()
	id := 0
	err := db.QueryRow(createUserQuery, data.Age, data.FirstName, data.LastName, data.Email).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println("service :user create successful")
	return nil
}

func readService(id string) (ur User, err error) {
	db := db.GetDB()
	row := db.QueryRow(readUserQuery, id)

	err = row.Scan(&ur.ID, &ur.Age, &ur.FirstName, &ur.LastName, &ur.Email)
	switch err {
	case sql.ErrNoRows:
		return User{}, errUserNotExist
	case nil:
		return ur, nil
	default:
		return User{}, err
	}

}
func updateService(id string, data User) error {
	db := db.GetDB()
	res, err := db.Exec(updateUserQuery, id, data.FirstName, data.LastName)
	if err != nil {
		return errors.New("Some query fault")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return errors.New("Some query fault")
	}
	if count == 0 {
		return errors.New("Invalid user Id")
	}

	fmt.Println(count)

	fmt.Println("service: update successful")
	return nil
}
func deleteService(id string) error {
	db := db.GetDB()
	res, err := db.Exec(deleteUserQuery, id)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return errors.New("Internal server Errors")
	}
	if cnt == 0 {
		return errors.New("Invalid user id")
	}
	fmt.Println(" : delete successful")
	return nil
}

func readAllService() ([]User, error) {
	db := db.GetDB()
	var users []User

	rows, err := db.Query(readAllUsersQuery)
	if err != nil {
		return users, errors.New("Err to retrive users")
	}

	defer rows.Close()
	var ur User
	for rows.Next() {
		err = rows.Scan(&ur.ID, &ur.Age, &ur.FirstName, &ur.LastName, &ur.Email)
		if err != nil {
			return users, errors.New("Err to retrive users")
		}
		//  fmt.Println(ur)
		users = append(users, ur)
	}
	fmt.Println("service: all user successful")
	fmt.Println(users)
	return users, nil
}
func Check(id int) bool {
	db := db.GetDB()
	res, err := db.Exec(readUserQuery, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	if count == 0 {

		return false
	}
	return true
}
