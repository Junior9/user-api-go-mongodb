package userservice

import (
	m "api/models"

	user_repository "api/repositories/user.repository"
)

func Add(user m.User) (bool, error) {
	resp, err := user_repository.Add(user)

	if err != nil {
		return false, err
	}

	if !resp {
		return false, nil
	}

	return true, nil
}

func Get() ([]m.User, error) {
	users, err := user_repository.Get()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetById(id string) (m.User, error) {
	user, err := user_repository.GetById(id)

	if err != nil {
		return m.User{}, err
	}
	return user, nil
}

func Update(user m.User) (bool, error) {
	resp, err := user_repository.Update(user)

	if err != nil {
		return false, err
	}

	if !resp {
		return false, nil
	}

	return true, nil
}

func Delete(id string) (bool, error) {
	_, err := user_repository.Delete(id)

	if err != nil {
		return false, err
	}
	return true, nil
}
