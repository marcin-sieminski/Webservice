package data

import (
	"database/sql"
	"errors"
	"time"
)

type Item struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Version   int32     `json:"-"`
}

type ItemModel struct {
	DB *sql.DB
}

func (itemModel ItemModel) Insert(item *Item) error {
	query := `
		INSERT INTO items (name)
		VALUES ($1)
		RETURNING id, created_at, version`

	args := []interface{}{item.Name}
	return itemModel.DB.QueryRow(query, args...).Scan(&item.ID, &item.CreatedAt, &item.Version)
}

func (itemModel ItemModel) Get(id int64) (*Item, error) {
	if id < 1 {
		return nil, errors.New("record not found")
	}

	query := `
		SELECT id, created_at, name, version
		FROM items
		WHERE id = $1`

	var item Item

	err := itemModel.DB.QueryRow(query, id).Scan(
		&item.ID,
		&item.CreatedAt,
		&item.Name,
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
		SET name = $1, version = version + 1
		WHERE id = $2 AND version = $3
		RETURNING version`

	args := []interface{}{item.Name, item.ID, item.Version}
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
	  SELECT * 
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
