package product

import "database/sql"

type ProductRepo struct{
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo{
	return &ProductRepo{db: db}
}

func (u *ProductRepo) Create(product Product) error {
	_, err := u.db.Exec("insert into products(name, descriptions) values($1, $2)",
	product.Name, product.Descriptions)
	
	return err
}

func (u *ProductRepo) Get() ([]Product, error){
	rows, err := u.db.Query("select * from Products")
	if err != nil{
		return nil, err
	}

	products := []Product{}
	for rows.Next(){
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Descriptions)
		if err != nil{
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (u *ProductRepo) GetById(id int) (Product, error){
	product := Product{}
	err := u.db.QueryRow("select * from products where id = $1", id).
	Scan(&product.Id, &product.Name, &product.Descriptions)
	return product, err
}

func (u *ProductRepo) Update(product Product) error {
	_, err := u.db.Exec("update users set name=$1, age=$2, product_id=$3 where id=$4 values($1, $2, $3, $4)",
	product.Name, product.Descriptions, product.Id)
	
	return err
}

func (u *ProductRepo) Delete(id int) error{
	_, err := u.db.Exec("delete from products where id = $1", id)
	return err
}