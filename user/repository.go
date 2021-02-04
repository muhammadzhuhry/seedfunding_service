package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	// membuat objek baru dr struk repository yg diatas
	// lalu karena struct repository memiliki field db maka perlu diisi nilainya menggunakan parameter yg ada di NewRepository
	return &repository{db}
}

// proses simpan ke database
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
