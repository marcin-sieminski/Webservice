package data

import (
	"database/sql"
	"errors"
	"time"
)

type Item struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Version     int32     `json:"-"`
	Description string    `json:"description"`
}

type ItemModel struct {
	DB *sql.DB
}

func (itemModel ItemModel) Insert(item *Item) error {
	query := `
		INSERT INTO items (name, description)
		VALUES ($1, $2)
		RETURNING id, created_at, version`

	args := []interface{}{item.Name, item.Description}
	return itemModel.DB.QueryRow(query, args...).Scan(&item.ID, &item.CreatedAt, &item.Version)
}

func (itemModel ItemModel) Get(id int64) (*Item, error) {
	if id < 1 {
		return nil, errors.New("record not found")
	}

	query := `
		SELECT id, created_at, name, description, version
		FROM items
		WHERE id = $1`

	var item Item

	err := itemModel.DB.QueryRow(query, id).Scan(
		&item.ID,
		&item.CreatedAt,
		&item.Name,
		&item.Description,
		&item.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}

	return &item, nil
}

func (itemModel ItemModel) Update(item *Item) error {
	query := `
		UPDATE items
		SET name = $1, description = $2, version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version`

	args := []interface{}{item.Name, item.ID, item.Version, item.Description}
	return itemModel.DB.QueryRow(query, args...).Scan(&item.Version)
}

func (itemModel ItemModel) Delete(id int64) error {
	if id < 1 {
		return errors.New("record not found")
	}

	query := `
		DELETE FROM items
		WHERE id = $1`

	results, err := itemModel.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (itemModel ItemModel) GetAll() ([]*Item, error) {
	query := `
	  SELECT id, created_at, name, version, description 
	  FROM items
	  ORDER BY id`

	rows, err := itemModel.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []*Item{}

	for rows.Next() {
		var item Item

		err := rows.Scan(
			&item.ID,
			&item.CreatedAt,
			&item.Name,
			&item.Version,
			&item.Description,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
