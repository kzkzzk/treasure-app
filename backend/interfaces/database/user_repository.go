package database

import "treasure-app/backend/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (name, email, password, image) VALUES (?,?,?,?)", u.Name, u.Email, u.Password, u.Image,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, name, email, password, image FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	var email string
	var password string
	var image string
	row.Next()
	if err = row.Scan(&id, &name, &email, &password, &image); err != nil {
		return
	}
	user.ID = id
	user.Name = name
	user.Email = email
	user.Image = image
	user.Password = password
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, name, email, password, image FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		var email string
		var password string
		var image string
		if err := rows.Scan(&id, &name, &email, &password, &image); err != nil {
			continue
		}
		user := domain.User{
			ID:       id,
			Name:     name,
			Email:    email,
			Password: password,
			Image:    image,
		}
		users = append(users, user)
	}
	return
}

func (repo *UserRepository) Update(u domain.User) (err error) {
	_, err = repo.Execute(
		"UPDATE users SET name = ?, email = ?, image = ? where id = ?", u.Name, u.Email, u.Image, u.ID,
	)
	return
}

func (repo *UserRepository) FindByFirebaseId(identifier string) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, name, email, password, image FROM users WHERE firebase_uid = ? LIMIT 1", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	var email string
	var password string
	var image string
	row.Next()
	if err = row.Scan(&id, &name, &email, &password, &image); err != nil {
		return
	}
	user.ID = id
	user.Name = name
	user.Email = email
	user.Image = image
	user.Password = password
	return
}
