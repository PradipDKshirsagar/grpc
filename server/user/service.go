package user

import (
	"grpc/server/db"
	"database/sql"
	"errors"
	"fmt"
)

var (
	createUserQuery           = `INSERT INTO users (age, first_name, last_name) VALUES ($1, $2, $3) RETURNING id`
	readUserQuery             = `SELECT * FROM USERS WHERE ID = $1;`
	updateUserQuery           = `UPDATE users SET age = $2, first_name = $3, last_name = $4 WHERE id = $1;`
	deleteUsersInterestsQuery = `DELETE FROM users_interests WHERE user_id = $1;`
	deleteUserQuery           = `DELETE FROM users WHERE id = $1;`
	readAllUsersQuery         = `SELECT * FROM USERS;`
)

func createService(data User) error {
	db := db.GetDB()
	id := 0
	err := db.QueryRow(createUserQuery, data.Age, data.FirstName, data.LastName).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func ReadService(id string) (ur User, err error) {
	db := db.GetDB()
	row := db.QueryRow(readUserQuery, id)

	err = row.Scan(&ur.ID, &ur.Age, &ur.FirstName, &ur.LastName)
	switch err {
	case sql.ErrNoRows:
		return User{}, errUserNotExist
	case nil:
		return ur, nil
	default:
		return User{}, err
	}
}
func readService(id string) (ur User, err error) {
	db := db.GetDB()
	row := db.QueryRow(readUserQuery, id)

	err = row.Scan(&ur.ID, &ur.Age, &ur.FirstName, &ur.LastName)
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
	res, err := db.Exec(updateUserQuery, id, data.Age, data.FirstName, data.LastName)
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
	return nil
}

func deleteService(id string) error {
	db := db.GetDB()
	_, err := db.Exec(deleteUsersInterestsQuery, id)
	if err != nil {
		return err
	}

	res, err := db.Exec(deleteUserQuery, id)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return errors.New("Internal Errors")
	}
	if cnt == 0 {
		return errors.New("Invalid user id")
	}
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
		err = rows.Scan(&ur.ID, &ur.Age, &ur.FirstName, &ur.LastName)
		if err != nil {
			return users, errors.New("Err to retrive users")
		}
		//  fmt.Println(ur)
		users = append(users, ur)
	}
	fmt.Println(users)
	return users, nil
}

// Check function to check user present or not which having parameter id.
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
