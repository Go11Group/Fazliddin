package problems

import (
	"database/sql"
	"fmt"
)

type ProblemRepo struct{
	DB *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo{
	return &ProblemRepo{DB: db}
}

func (p *ProblemRepo) Create(problem Problem) error {
	_, err := p.DB.Exec("insert into problems(description, level) values($1, $2)",
	problem.Description, problem.Level)
	if err != nil{
		fmt.Println(err)
	}
	return err
}

func (p *ProblemRepo) Get() ([]Problem, error){
	rows, err := p.DB.Query("select * from problems")
	if err != nil{
		return nil, err
	}

	problems := []Problem{}
	for rows.Next(){
		problem := Problem{}
		err = rows.Scan(&problem.Id, &problem.Description, &problem.Level)
		if err != nil{
			return nil, err
		}
		problems = append(problems, problem)
	}
	return problems, nil
}

func (p *ProblemRepo) GetById(id int) (Problem, error){
	problem := Problem{}
	err := p.DB.QueryRow("select * from problems where id = $1", id).
	Scan(&problem.Id, &problem.Description, &problem.Level)
	return problem, err
}

func (p *ProblemRepo) Update(problem Problem) error {
	_, err := p.DB.Exec("update problems set description=$1, level=$2, where id=$3",
	problem.Description, problem.Level, problem.Id)
	
	return err
}

func (p *ProblemRepo) Delete(id int) error{
	_, err := p.DB.Exec("delete from problems where id = $1", id)
	return err
}