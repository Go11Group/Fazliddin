1-DB_URL = postgres://postgres:123321@localhost:5432/users?sslmode=disable
2-DB_URL = postgres://postgres:123321@localhost:5432/metro?sslmode=disable

1-migrate_create:
	migrate create -ext sql -dir migrate_user -seq h_user

1_up:
	migrate -path migrate_user -database ${1-DB_URL} -verbose up

1_down:
	migrate -path migrate_user -database ${1-DB_URL} -verbose down

1_force:
	migrate -path migrate_user -database ${1-DB_URL} -verbose force 1


2-migrate_create:
	migrate create -ext sql -dir migrate_metro -seq h_motro

2_up:
	migrate -path migrate_metro -database ${2-DB_URL} -verbose up

2_down:
	migrate -path migrate_metro -database ${2-DB_URL} -verbose down

2_force:
	migrate -path migrate_metro -database ${2-DB_URL} -verbose force 1