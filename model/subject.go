package model

const SubjectDQL = "SELECT id, name, type_id  FROM subjects"

type Subject struct {
	ID     int
	TypeID int
	Name   string
}

