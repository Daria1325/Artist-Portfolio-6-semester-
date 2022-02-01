package database

import (
	"fmt"
	cnfg "github.com/daria/Portfolio/backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Picture struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Owner       string  `db:"owner"`
	Size        string  `db:"size"`
	Materials   string  `db:"materials"`
	Price       float32 `db:"price"`
	Description string  `db:"description"`
	Path        string  `db:"path"`
}
type Series struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
type Client struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Type  string `db:"type"`
	Email string `db:"email"`
	Phone string `db:"phone"`
}

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

//func (r *Repo) GetSeries() ([]Series, error) {
//	series := []Series{}
//	rows, err := r.db.Queryx("SELECT * FROM ports")
//	if err != nil {
//		fmt.Errorf("failed to execute the query: %v", err.Error())
//		return nil, err
//	}
//	for rows.Next() {
//		var p Series
//		err = rows.StructScan(&p)
//		if err != nil {
//			fmt.Errorf("%s", err.Error())
//			continue
//		}
//		series = append(series, p)
//	}
//	return ports, nil
//}
//func (r *Repo) GetClients() ([]Client, error) {
//	clients := []Client{}
//	rows, err := r.db.Queryx("SELECT * FROM ports")
//	if err != nil {
//		fmt.Errorf("failed to execute the query: %v", err.Error())
//		return nil, err
//	}
//	for rows.Next() {
//		var p Client
//		err = rows.StructScan(&p)
//		if err != nil {
//			fmt.Errorf("%s", err.Error())
//			continue
//		}
//		clients = append(clients, p)
//	}
//	return clients, nil
//}
//func (r *Repo) GetPictures() ([]Picture, error) {
//	pictures := []Picture{}
//	rows, err := r.db.Queryx("SELECT * FROM ports")
//	if err != nil {
//		fmt.Errorf("failed to execute the query: %v", err.Error())
//		return nil, err
//	}
//	for rows.Next() {
//		var p Picture
//		err = rows.StructScan(&p)
//		if err != nil {
//			fmt.Errorf("%s", err.Error())
//			continue
//		}
//		pictures = append(pictures, p)
//	}
//	return pictures, nil
//}
//
//func (r *Repo) AddSeries(item Series) error {
//	_, err := r.db.NamedExec(`INSERT INTO ports (name, description)
//        VALUES (:name, :description)`, item)
//	if err != nil {
//		fmt.Errorf("failed to execute the query: %v", err.Error())
//		return err
//	}
//	return nil
//}
//func (r *Repo) UpdateSeries(item Series) error {
//	_, err := r.db.NamedExec(`UPDATE ports SET name=:name, description=:description WHERE id =:id`, item)
//	if err != nil {
//		fmt.Errorf("failed to execute the query: %v", err.Error())
//		return err
//	}
//	return nil
//}

func (r *Repo) Close() error {
	err := r.db.Close()
	return err
}

func Init(config *cnfg.Config) *Repo {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host =%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName))
	if err != nil {
		return nil
	}
	repo := New(db)

	return repo
}
