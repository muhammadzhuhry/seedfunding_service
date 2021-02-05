package user

import "gorm.io/gorm"

type Repository interface {
	// interface disini berlaku sebagai kontrak awal
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
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
	// gorm save -> menyimpan data baru
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	// gorm save -> untuk merubah data yg sudah ada sebelumnya di db
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
