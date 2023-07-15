package repositories

import (
	"context"
	"log"
	"pi/internal/core/domains"
	"pi/internal/core/ports"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	if err := db.AutoMigrate(&domains.User{}); err != nil {
		log.Fatalf("error migrate database: %v", err)
	}
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, req *domains.CreateUserRequest) (*domains.User, error) {
	user := domains.User{
		Username: req.Username,
		Email:    req.Email,
	}
	err := r.db.Model(&domains.User{}).Create(&user).Error
	return &user, err
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*domains.User, error) {
	var user *domains.User
	err := r.db.Model(&domains.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) Update(ctx context.Context, id uint, req *domains.UpdateUserRequest) error {
	user := domains.User{
		Username: req.Username,
		Email:    req.Email,
	}
	return r.db.Model(&domains.User{}).Where("id = ?", id).Updates(&user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.Model(&domains.User{}).Where("id = ?", id).Delete(&domains.User{}).Error
}
