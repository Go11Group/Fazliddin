package user

import "gorm.io/gorm"

type UserRepo struct{
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func(u *UserRepo) CreateTable(user User) error{
	err := u.Db.AutoMigrate(user)
	if err != nil{
		return err
	}
	return nil
}

func(u *UserRepo) Create(user User){
	u.Db.Create(&user)
}

func (u *UserRepo) Delete(email string){
	u.Db.Where("email = ?", email).Delete(&User{})
}

func (u *UserRepo) Update(user User){
	u.Db.Model(User{}).Where("email = ?", user.Email).Updates(user)
}

func (u *UserRepo) GetById(email string) User{
	user := User{}
	u.Db.Where("email = ?", email).First(&user)
	return user
}

func (u *UserRepo) GetUsers() []User{
	var users []User
    u.Db.Find(&users)
	return users
}

func (c *UserRepo) Filter(chose, filter_by string) []User {
	users := []User{}
	
	switch chose{
	case "first_name":
		c.Db.Where("firs_name = ?", filter_by).Find(&users)
	case "last_name":
		c.Db.Where("last_name = ?", filter_by).Find(&users)
	case "email":
		c.Db.Where("email = ?", filter_by).Find(&users)
	case "age":
		c.Db.Where("age = ?", filter_by).Find(&users)
	case "field":
		c.Db.Where("field = ?", filter_by).Find(&users)
	case "gender":
		c.Db.Where("gender = ?", filter_by).Find(&users)
	case "is_employe":
		c.Db.Where("is_employe = ?", filter_by).Find(&users)
	}
	return users
}