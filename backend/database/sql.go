package database

import (
	"database/sql"
	"fmt"
	"time"

	cnfg "github.com/daria/Portfolio/backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const layoutISO = "2006-01-02"

type Picture struct {
	ID          int             `db:"id"`
	Name        string          `db:"name"`
	Path        sql.NullString  `db:"path"`
	Price       sql.NullFloat64 `db:"price"`
	Date        sql.NullString  `db:"date"`
	Material    sql.NullString  `db:"material"`
	Size        sql.NullString  `db:"size"`
	Description sql.NullString  `db:"description"`
	SeriesId    sql.NullInt32   `db:"series_id"`
	ClientId    sql.NullInt32   `db:"client_id"`
}
type Series struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}
type Client struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Contact int    `db:"contact_id""`
	Type    int    `db:"type_id"`
}

type Contact struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	Number    int    `db:"number""`
	AddNumber int    `db:"add_number"`
}
type ClientType struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetSeries() ([]Series, error) {
	series := []Series{}
	rows, err := r.db.Queryx("SELECT * FROM series ORDER BY id")
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return nil, err
	}
	for rows.Next() {
		var p Series
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		series = append(series, p)
	}
	return series, nil
}
func (r *Repo) GetSeriesById(id string) (Series, error) {
	series := []Series{}
	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM series WHERE id=%s", id))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return Series{}, err
	}
	for rows.Next() {
		var p Series
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		series = append(series, p)
	}
	if len(series) == 0 {
		return Series{ID: -1}, nil
	} else {
		return series[0], nil
	}

}
func (r *Repo) GetSeriesIDByName(name string) (int, error) {
	series := []Series{}
	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM series WHERE name='%s'", name))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return 0, err
	}
	for rows.Next() {
		var p Series
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		series = append(series, p)
	}
	if len(series) == 0 {
		return 0, nil
	} else {
		return series[0].ID, nil
	}

}
func (r *Repo) DeleteSeries(id string) error {
	_, err := r.db.Queryx(fmt.Sprintf("DELETE FROM pictures WHERE series_id=%s", id))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	_, err = r.db.Queryx(fmt.Sprintf("DELETE FROM series WHERE id=%s", id))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}
func (r *Repo) AddSeries(item Series) error {
	_, err := r.db.Queryx(fmt.Sprintf("INSERT INTO series (name, description) VALUES ('%s','%s')", item.Name, item.Description.String))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}
func (r *Repo) UpdateSeries(item Series) error {
	_, err := r.db.Queryx(fmt.Sprintf("UPDATE series SET name = '%s', description = '%s' WHERE id = '%d'", item.Name, item.Description.String, item.ID))

	//_, err := r.db.NamedExec(`UPDATE series SET name=:name, description=:description.String WHERE id =:id`, item)
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}

func (r *Repo) GetClients(num int) ([]Client, error) {
	clients := []Client{}
	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM clients WHERE type_id <> 2 LIMIT %d", num))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return nil, err
	}
	for rows.Next() {
		var p Client
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		clients = append(clients, p)
	}
	return clients, nil
}

func (r *Repo) GetPictures() ([]Picture, error) {
	pictures := []Picture{}
	rows, err := r.db.Queryx("SELECT * FROM pictures ORDER BY id")
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return nil, err
	}
	for rows.Next() {
		var p Picture
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		pictures = append(pictures, p)
	}
	return pictures, nil
}
func (r *Repo) GetPictureById(id string) (Picture, error) {
	pictures := []Picture{}
	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM pictures WHERE id=%s", id))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return Picture{}, err
	}
	for rows.Next() {
		var p Picture
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		pictures = append(pictures, p)
	}
	if len(pictures) == 0 {
		return Picture{ID: -1}, nil
	} else {
		return pictures[0], nil
	}
}
func (r *Repo) GetPictureBySeries(id string) ([]Picture, error) {
	pictures := []Picture{}
	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM pictures WHERE series_id=%s ORDER BY id", id))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return nil, err
	}
	for rows.Next() {
		var p Picture
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		pictures = append(pictures, p)
	}
	return pictures, nil
}
func (r *Repo) DeletePictures(id string) error {
	_, err := r.db.Queryx(fmt.Sprintf("DELETE FROM pictures WHERE id=%s", id))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}
func (r *Repo) AddPicture(item Picture) error {
	date, err := time.Parse(layoutISO, item.Date.String)

	_, err = r.db.Queryx("INSERT INTO pictures (name,path,price,date,material,size, description,series_id,client_id)" +
		fmt.Sprintf(" VALUES ('%s','%s', '%f', '%s', '%s','%s','%s','%d','%d')", item.Name, item.Path.String, item.Price.Float64, date, item.Material.String, item.Size.String, item.Description.String, item.SeriesId.Int32, item.ClientId.Int32))

	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}
func (r *Repo) UpdatePicture(item Picture) error {
	_, err := r.db.NamedExec(`UPDATE series SET name=:name, description=:description WHERE id =:id`, item)
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}

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
