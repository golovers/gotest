package user

import (
	"encoding/csv"
	"io"
	"strings"
)

type (
	Users []User
)

func (users *Users) ReadFrom(r io.Reader) (int64, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	c := 0
	if err != nil {
		return int64(c), err
	}
	for _, row := range records {
		*users = append(*users, User{
			ID:   row[0],
			Name: row[1],
		})
		c += len(strings.Join(row, ","))
	}
	return int64(c), nil
}
