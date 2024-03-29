package database

import (
	"database/sql"
	"fmt"
	cnfg "github.com/daria/Portfolio/backend/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Picture struct {
	ID          int             `db:"id"`
	Name        string          `db:"name"`
	Path        sql.NullString  `db:"path"`
	Price       sql.NullFloat64 `db:"price"`
	Date        sql.NullString  `db:"date"`
	Material    sql.NullString  `db:"material"`
	Size        sql.NullString  `db:"size"`
	Description sql.NullString  `db:"description"`
	SeriesId    int             `db:"series_id"`
	ClientId    sql.NullInt32   `db:"client_id"`
}
type Series struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}
type Client struct {
	ID      int           `db:"id"`
	Name    string        `db:"name"`
	Contact sql.NullInt32 `db:"contact_id"`
	Type    int           `db:"type_id"`
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
	}
	return series[0].ID, nil
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
func (r *Repo) AddSeries(item Series) (int, error) {
	sqlStatement := `INSERT INTO series (name, description) VALUES ($1, $2) RETURNING id`
	id := -1
	err := r.db.QueryRow(sqlStatement, item.Name, item.Description.String).Scan(&id)

	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return -1, err
	}
	return id, nil
}
func (r *Repo) UpdateSeries(item Series) error {
	_, err := r.db.Queryx(fmt.Sprintf("UPDATE series SET name = '%s', description = '%s' WHERE id = '%d'", item.Name, item.Description.String, item.ID))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}
func (r *Repo) GetPicturePathBySeriesID(id string) (string, error) {
	pictures := []Picture{}
	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM pictures WHERE series_id='%s' AND (path IS NOT NULL OR path <> ' ')", id))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return "", err
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
		return "", nil
	}
	return pictures[0].Path.String, nil
}
func (r *Repo) GetClients(num int) ([]Client, error) {
	clients := []Client{}
	rows, err := r.db.Queryx(fmt.Sprintf("SELECT * FROM clients WHERE type_id <> 2 LIMIT '%d'", num))
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
	}
	return pictures[0], nil
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
	_, err := r.db.Queryx("INSERT INTO pictures (name,path,price,date,material,size, description,series_id)" +
		fmt.Sprintf(" VALUES ('%s','%s', '%f', '%s', '%s','%s','%s','%d')", item.Name, item.Path.String, item.Price.Float64, item.Date.String, item.Material.String, item.Size.String, item.Description.String, item.SeriesId))
	if err != nil {
		fmt.Errorf("failed to execute the query: %v", err.Error())
		return err
	}
	return nil
}
func (r *Repo) UpdatePicture(item Picture) error {
	_, err := r.db.Queryx(fmt.Sprintf("UPDATE pictures SET name = '%s',path = '%s',price = '%f',", item.Name, item.Path.String, item.Price.Float64) +
		fmt.Sprintf("date = '%s',material='%s',size='%s', description='%s',series_id = '%d' WHERE id = '%d'", item.Date.String, item.Material.String, item.Size.String, item.Description.String, item.SeriesId, item.ID))
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
