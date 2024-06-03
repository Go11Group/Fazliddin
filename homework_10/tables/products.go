package tables

import "database/sql"

type Product struct{
	Id int
	Name string
	Description string
	Price float64
	Stock_quantity int
}

type PRepo struct{
	DB *sql.DB
}

func NewPRepo(db *sql.DB) *PRepo{
	return &PRepo{db}
}

func (p *PRepo) Create(product Product) error{
	a, _:= p.DB.Begin()
	defer a.Commit()
	_, err :=p.DB.Exec("insert into products(name, description, price, stock_quantity) values($1, $2, $3, $4)", 
	product.Name, product.Description, product.Price, product.Stock_quantity)
	if err != nil{
		return err
	}
	return nil
}

func (p *PRepo) GetAllProduct() ([]Product, error){
	a, _:= p.DB.Begin()
	defer a.Commit()
	products := []Product{}
	data, err := p.DB.Query("select * from products")
	if err != nil{
		return nil, err
	}
	for data.Next(){
		product := Product{}
		err := data.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
		if err != nil{
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

func (p *PRepo) GetById(id int) (Product, error){
	a, _:= p.DB.Begin()
	defer a.Commit()
	product := Product{}
	err :=p.DB.QueryRow("select * from products where id = $1", id).
	Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
	return product, err
}

func (p *PRepo) Update(product Product) error{
	a, _:= p.DB.Begin()
	defer a.Commit()
	_, err := p.DB.Exec("update table users set name = $1, description = $2, price = $3, Srock_quantity = $4 where id = $5",
	product.Name, product.Description, product.Price, product.Stock_quantity, product.Id)
	if err != nil{
		return err
	}
	return nil
}

func (p *PRepo) Delete(id int) error{
	_, err := p.DB.Exec("delete from products where id = $1", id)
	return err
}
