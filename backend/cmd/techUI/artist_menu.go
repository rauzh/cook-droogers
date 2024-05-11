package techUI

import (
	"cookdroogers/app"
	"cookdroogers/internal/requests/publish"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"errors"
	"fmt"
	"log/slog"
)

type artistMenu struct {
	a      *app.App
	user   *models.User
	artist *models.Artist
	log    *slog.Logger
}

func initArtistMenu(a *app.App, user *models.User, log *slog.Logger) (*artistMenu, error) {

	artist, err := a.Services.ArtistService.GetByUserID(user.UserID)
	if err != nil {
		log.Error("Can't init artist menu: can't get artist: ", slog.Any("error", err))
		return nil, err
	}

	return &artistMenu{
		a:      a,
		user:   user,
		artist: artist,
		log:    log,
	}, nil
}

func (menu *artistMenu) Loop() error {

	startPosition :=
		`
		0 -- выйти
		1 -- подать заявку на публикацию релиза
		2 -- посмотреть список заявок
		3 -- загрузить релиз
		4 -- получить статистику своих релизов
		5 -- посмотреть информацию о лейбле
		6 -- посмотреть свои релизы
	Выберите пункт меню: `

	fmt.Println(menu.artist.ArtistID, menu.artist.ManagerID, menu.artist.ContractTerm, menu.artist.Nickname, menu.artist.UserID)

	for {
		fmt.Printf("%s", startPosition)

		var action int
		_, _ = fmt.Scanf("%d", &action)

		switch action {
		case 0:
			return ErrEXIT
		case 1:
			err := menu.applyPublishRequest()
			if err != nil {
				fmt.Println("Не удается создать заявку по причине ", err)
				menu.log.Error("Can't apply request: ", slog.Any("error", err))
			} else {
				fmt.Println("Заявка на публикацию релиза успешно подана.")
			}
		case 2:
			err := lookupReqs(menu.a, menu.user)
			if err != nil {
				fmt.Println("Не удается просмотреть заявки по причине ", err)
				menu.log.Error("Can't look up requests: ", slog.Any("error", err))
			}
		case 3:
			err := menu.uploadRelease()
			if err != nil {
				fmt.Println("Не удается загрузить релиз по причине ", err)
				menu.log.Error("Can't upload release: ", slog.Any("error", err))
			} else {
				fmt.Println("Релиз успешно загружен.")
			}
		case 4:
			menu.stats()
		case 5:
			printInfo()
		case 6:
			menu.releases()
		default:
			fmt.Printf("Неверный пункт меню")
		}
	}
}

func (menu *artistMenu) applyPublishRequest() error {

	fmt.Printf("%s", "Введите id релиза, который вы хотите опубликовать: ")
	var releaseID uint64
	_, _ = fmt.Scanf("%d", &releaseID)

	fmt.Printf("%s", "Введите желаемую дату публикации (год месяц день): ")
	var year, day, month int
	_, _ = fmt.Scanf("%d %d %d", &year, &month, &day)

	pubReq := publish.NewPublishRequest(menu.user.UserID, releaseID, cdtime.Date(year, month, day))

	return menu.a.UseCases.PublishReqUC.Apply(pubReq)
}

func (menu *artistMenu) uploadRelease() error {

	fmt.Printf("%s", "Введите название релиза, который вы хотите загрузить: ")
	var title string
	_, _ = fmt.Scanf("%s", &title)

	fmt.Printf("%s", "Введите дату создания релиза, который вы хотите загрузить (год месяц день): ")
	var year, day, month int
	_, _ = fmt.Scanf("%d %d %d", &year, &month, &day)

	if year < 1000 || month > 12 || month < 1 || day < 0 || day > 31 {
		return errors.New("invalid date")
	}

	release := &models.Release{
		Title:        title,
		ArtistID:     menu.artist.ArtistID,
		DateCreation: cdtime.Date(year, month, day),
	}

	fmt.Printf("%s", "Введите количество треков: ")
	var tracksNum int
	_, _ = fmt.Scanf("%d", &tracksNum)

	if tracksNum < 1 {
		return errors.New("tracks num must be positive")
	}

	tracks := make([]*models.Track, tracksNum)

	for i := 0; i < tracksNum; i++ {

		fmt.Printf("%s", "Введите информацию о треке\n(название длительность жанр тип): ")
		var trackTitle, genre, trackType string
		var duration uint64
		_, _ = fmt.Scanf("%s %d %s %s", &trackTitle, &duration, &genre, &trackType)

		track := &models.Track{
			Title:    trackTitle,
			Duration: duration,
			Genre:    genre,
			Type:     trackType,
			Artists:  []uint64{menu.artist.ArtistID},
		}
		tracks[i] = track
	}

	return menu.a.Services.ReleaseService.Create(release, tracks)
}

func (menu *artistMenu) stats() {

	report, err := menu.a.Services.ReportService.GetReportForArtist(menu.artist.ArtistID)
	if err != nil {
		fmt.Println("Не удается посмотреть статистику по причине ", err)
		menu.log.Error("Can't get stats: ", slog.Any("error", err))
	}

	for release, releaseStats := range report {
		fmt.Printf("%s:\n%s\n\n", release, string(releaseStats[:]))
	}
}

func (menu *artistMenu) releases() {

	releases, err := menu.a.Services.ReleaseService.GetAllByArtist(menu.artist.ArtistID)
	if err != nil {
		menu.log.Error("Can't get releases: ", slog.Any("error", err))
	}

	for _, release := range releases {
		fmt.Printf("\n\t release id: %d\n\t creation date: %s\n\t status: %s\n\t title: %s",
			release.ReleaseID, release.DateCreation, release.Status, release.Title)
		for _, trackID := range release.Tracks {
			fmt.Printf("\n\t track: %d", trackID)
		}
	}
}
