CREATE TYPE transaction_type AS enum('creedit', 'debit');

create table users(
    user_id uuid default gen_random_uuid() primary key,
    name varchar,
    age int,
    phone varchar(13),
    deleted_at bigint default 0
);

create table card(
    card_id uuid default gen_random_uuid() primary key,
    number varchar(13),
    user_id uuid references users(user_id),
    deleted_at bigint default 0
);

CREATE TABLE transaction (
    transaction_id uuid default gen_random_uuid() PRIMARY KEY,
    card_id uuid NOT NULL REFERENCES card(card_id),
    amount DECIMAL(10, 2) NOT NULL,
    terminal_id UUID,
    transaction_type transaction_type NOT NULL,
    deleted_at bigint default 0
);