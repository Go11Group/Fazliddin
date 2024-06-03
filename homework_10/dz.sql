1. Doimgidek.Yaxshilab yozing.
2. Leetcodedan 2ta masala.
3. Bizda users va products tablelari bor. userni productslarida create, update, delete qilinganda, transactiondan foydalanilsin. Siklda aylantirgan holda. 
-- Users jadvali
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

-- Products jadvali
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);
Crudlar:
Create (Yaratish):
CreateUser: Foydalanuvchi yaratish.
CreateProduct: Mahsulot yaratish.
Read (O'qish):
GetUser: Foydalanuvchini o'qish.
GetProduct: Mahsulotni o'qish.
Update (Yangilash):
UpdateUser: Foydalanuvchi ma'lumotlarini yangilash.
UpdateProduct: Mahsulot ma'lumotlarini yangilash.
Delete (O'chirish):
DeleteUser: Foydalanuvchini o'chirish.
DeleteProduct: Mahsulotni o'chirish.
user_products:(tableni o'zingiz yozing)
Bu tableda UserProduct ma'lumotlari saqlanadi. 
Aynan Userni productlarini create, read update qilishda transactionlardan foydalaning.


SAvollar bo'lsa so'rangiz



select u.id, u.username, u.email, array_agg(p.id), array_agg(p.name), u.password
from users u
left join users_products up on up.user_id = u.id
left join products p on p.id = up.product_id
group by u.id