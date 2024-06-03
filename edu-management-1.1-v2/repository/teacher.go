package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type TeacherRepository interface {
	FetchAll() ([]model.Teacher, error)
	FetchByID(id int) (*model.Teacher, error)
	Store(g *model.Teacher) error
	Update(id int, g *model.Teacher) error
}

type teacherRepoImpl struct {
	db *sql.DB
}

func NewTeacherRepo(db *sql.DB) *teacherRepoImpl {
	return &teacherRepoImpl{db}
}

func (g *teacherRepoImpl) FetchAll() ([]model.Teacher, error) {
	// TODO: replace this
	rows, err := g.db.Query("SELECT id, name, address, subject FROM teachers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Teachers []model.Teacher
	for rows.Next() {
		var Teacher model.Teacher
		err := rows.Scan(&Teacher.ID, &Teacher.Name, &Teacher.Address, &Teacher.Subject)
		if err != nil {
			return nil, err
		}
		Teachers = append(Teachers, Teacher)
	}
	return Teachers, nil
}

func (g *teacherRepoImpl) FetchByID(id int) (*model.Teacher, error) {
	row := g.db.QueryRow("SELECT id, name, address, subject FROM teachers WHERE id = $1", id)

	var Teacher model.Teacher
	err := row.Scan(&Teacher.ID, &Teacher.Name, &Teacher.Address, &Teacher.Subject)
	if err != nil {
		return nil, err
	}

	return &Teacher, nil
}

func (g *teacherRepoImpl) Store(teacher *model.Teacher) error {
	// TODO: replace this
	_, err := g.db.Exec("INSERT INTO teachers (name, address, subject) VALUES ($1, $2, $3)", teacher.Name, teacher.Address, teacher.Subject)
	return err
}

func (g *teacherRepoImpl) Update(id int, teacher *model.Teacher) error {
	// TODO: replace this
	_, err := g.db.Exec("UPDATE teachers SET name = $1, address = $2, subject = $3 WHERE id = $4", teacher.Name, teacher.Address, teacher.Subject, id)
	return err
}
