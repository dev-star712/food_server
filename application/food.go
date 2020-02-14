package application

import (
	"food-app/database/rdbms"
	"food-app/domain/entity"
	"food-app/domain/infrastructure"
)

type FoodImpl struct {
}

var FoodApp FoodAppInterface = &FoodImpl{}

type FoodAppInterface interface {
	SaveFood(*entity.Food) (*entity.Food, error)
	GetAllFood() ([]entity.Food, error)
	GetFood(uint64) (*entity.Food, error)
	UpdateFood(*entity.Food) (*entity.Food, error)
	DeleteFood(uint64) error
}

//GetUser returns a user
//func (u *UserImpl) GetFood(id uint64) (*entity.User, error) {
//	//return u.Repository.GetUser(id)
//	return nil, nil
//}

func (u *FoodImpl) SaveFood(food *entity.Food) (*entity.Food, error) {
	db := rdbms.NewDB()
	conn := infrastructure.NewRepositoryFoodCRUD(db)
	//u, err := entity.User{}
	return conn.SaveFood(food)
}

func (u *FoodImpl) GetAllFood() ([]entity.Food, error) {
	db := rdbms.NewDB()
	conn := infrastructure.NewRepositoryFoodCRUD(db)
	return conn.GetAllFood()
}

func (u *FoodImpl) GetFood(foodId uint64) (*entity.Food, error) {
	db := rdbms.NewDB()
	conn := infrastructure.NewRepositoryFoodCRUD(db)
	return conn.GetFood(foodId)
}

func (u *FoodImpl) UpdateFood(food *entity.Food) (*entity.Food, error) {
	db := rdbms.NewDB()
	conn := infrastructure.NewRepositoryFoodCRUD(db)
	return conn.UpdateFood(food)
}


func (u *FoodImpl) DeleteFood(foodId uint64) error {
	db := rdbms.NewDB()
	conn := infrastructure.NewRepositoryFoodCRUD(db)
	return conn.DeleteFood(foodId)
}