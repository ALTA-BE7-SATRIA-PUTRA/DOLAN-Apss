package user

import (
	_entities "group-project/dolan-planner/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) CreatUser(user _entities.User) (_entities.User, error) {
	tx := ur.database.Save(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (ur *UserRepository) GetUser(idToken int) (_entities.User, int, error) {
	var user _entities.User
	tx := ur.database.Where("ID = ?", idToken).Find(&user)
	if tx.Error != nil {
		return user, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, 0, nil
	}
	return user, int(tx.RowsAffected), nil
}

func (ur *UserRepository) UpdateUser(userUpdate _entities.User) (_entities.User, int, error) {
	tx := ur.database.Save(&userUpdate)
	if tx.Error != nil {
		return userUpdate, 0, tx.Error
	}
	return userUpdate, int(tx.RowsAffected), nil
}

func (ur *UserRepository) DeleteUser(id int) (int, error) {
	var user _entities.User
	tx := ur.database.Delete(&user, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
