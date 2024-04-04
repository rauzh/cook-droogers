package main

import (
	"context"
	postgres "cookdroogers/internal/repo/pg"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		"rauzh", "cook_droogers", "1337",
		"localhost", "5432")
	fmt.Println(dsnPGConn)

	db, err := sql.Open("pgx", dsnPGConn)

	fmt.Println(err)

	err = db.Ping()

	fmt.Println(err)

	db.SetMaxOpenConns(10)

	artRepo := postgres.NewArtistPgRepo(db)

	//artist := &models.Artist{
	//	UserID:       6,
	//	Nickname:     "tulenik-rocker",
	//	ContractTerm: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
	//	Activity:     true,
	//	ManagerID:    5,
	//}
	//
	//fmt.Println(artRepo.Create(artist))
	//
	//fmt.Println(artist)

	//artist := &models.Artist{
	//	ArtistID:     6,
	//	UserID:       6,
	//	Nickname:     "tulenik-rocker-228",
	//	ContractTerm: time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
	//	Activity:     true,
	//	ManagerID:    5,
	//}
	//fmt.Println(artRepo.Update(artist))
	//fmt.Println(artist)

	a, e := artRepo.Get(context.TODO(), uint64(6))
	fmt.Println(a, e)
}
