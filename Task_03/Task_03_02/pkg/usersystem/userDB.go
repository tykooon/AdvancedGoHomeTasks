package usersystem

import (
	"database/sql"
	"errors"
)

type UserList []*User

type UserSysModel struct {
	DB *sql.DB
}

func (m *UserSysModel) AddRange(list UserList) (n int) {
	for _, user := range list {
		var id int64 = m.Insert(user)
		if id != -1 {
			user.Id = id
			n++
		}
	}
	return
}

func (m *UserSysModel) Insert(user *User) (id int64) {
	query := `
	INSERT INTO Users (FullName, Email, Hash, IsActive)
	VALUES (?, ?, ?, ?)
	RETURNING Id`
	res := m.DB.QueryRow(query,
		user.FullName,
		user.Email,
		user.passHash,
		user.IsActive)
	if res.Scan(&id) != nil {
		return -1
	}
	return
}

func (m *UserSysModel) GetById(id int64) (*User, error) {
	user := User{Id: id}
	query := `
	Select FullName, Email, Hash, IsActive
	FROM Users 
	WHERE Id = ?`
	res := m.DB.QueryRow(query, id)
	err := res.Scan(&user.FullName, &user.Email, &user.passHash, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserSysModel) GetAll() (UserList, error) {
	query := `
	Select Id, FullName, Email, Hash, IsActive
	FROM Users
	ORDER BY Id`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := make(UserList, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Email,
			&user.passHash,
			&user.IsActive)
		if err != nil {
			return nil, err
		}
		list = append(list, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (m *UserSysModel) Update(user *User) bool {
	query := `
	UPDATE Users 
	SET 
		FullName = ?,
		Email = ?,
		IsActive = ?
	WHERE 
		Id = ?`
	res, err := m.DB.Exec(query,
		user.FullName,
		user.Email,
		user.IsActive,
		user.Id)

	if err == nil {
		k, _ := res.RowsAffected()
		if k == 1 {
			return true
		}
	}
	return false
}

func (m *UserSysModel) Delete(id int64) error {
	query := `
	DELETE FROM Users 
	WHERE 
		Id = ?`
	res, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	k, _ := res.RowsAffected()
	if k != 1 {
		return errors.New("User wasn't deleted")
	}
	return nil
}
