package mysql

import (
	"database/sql"
	"log"
	"todo/pkg/models"
)

// Define a TodosModel type which wraps a sql.DB connection pool.
type TodosModel struct {
	DB *sql.DB
}

// This will insert a new todo into the database.
func (m *TodosModel) Insert(Name string) (int, error) {

	stmt := `INSERT INTO todos (Name, created, expires)
	VALUES(?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, Name, 7)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

// This will return a specific todo based on its id.
func (m *TodosModel) GetAll() ([]*models.Todos, error) {
	stmt := "SELECT ID, Name FROM todos"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*models.Todos{}

	for rows.Next() {
		s := &models.Todos{}
		err = rows.Scan(&s.ID, &s.Name)
		if err != nil {
			return nil, err
		}
		todos = append(todos, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

// Delete the task according to the ID given
func (m *TodosModel) Remove(ID int) error {
	log.Println(ID)
	stmt := "DELETE FROM todos WHERE id = ?"
	_, err := m.DB.Exec(stmt, ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *TodosModel) Update(ID int, Name string) (int, error) {
	stmt := "UPDATE todos SET Name=? WHERE ID=?"
	_, err := m.DB.Exec(stmt, Name, ID)
	if err != nil {
		return 0, nil
	}
	return 0, nil

}

// check if the given id is in the db
func (m *TodosModel) CheckIfExists(id int) (bool, error) {
	stmt := `SELECT id FROM todos WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	var todoID int

	err := row.Scan(&todoID)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}
