package techUI

import (
	"cookdroogers/app"
	"cookdroogers/internal/requests/publish"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
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
	Выберите пункт меню: `

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
				menu.log.Error("Can't apply request: ", slog.Any("error", err))
			}
		case 2:
			err := lookupReqs(menu.a, menu.user)
			if err != nil {
				menu.log.Error("Can't look up requests: ", slog.Any("error", err))
			}
		case 3:
			err := menu.uploadRelease()
			if err != nil {
				menu.log.Error("Can't upload release: ", slog.Any("error", err))
			}
		case 4:
			menu.stats()
		case 5:
			printInfo()
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

	release := &models.Release{
		Title:    title,
		ArtistID: menu.artist.ArtistID,
	}

	fmt.Printf("%s", "Введите количество треков: ")
	var tracksNum int
	_, _ = fmt.Scanf("%d", &tracksNum)

	tracks := make([]models.Track, tracksNum)

	for i := 0; i < tracksNum; i++ {

		fmt.Printf("%s", "Введите информацию о треке\n(название длительность жанр тип): ")
		var trackTitle, genre, trackType string
		var duration uint64
		_, _ = fmt.Scanf("%s %d %s %s", &trackTitle, &duration, &genre, &trackType)

		track := models.Track{
			Title:    trackTitle,
			Duration: duration,
			Genre:    genre,
			Type:     trackType,
			Artists:  []uint64{menu.artist.ArtistID},
		}
		tracks = append(tracks, track)
	}

	return menu.a.Services.ReleaseService.Create(release, tracks)
}

func (menu *artistMenu) stats() {

	report, err := menu.a.Services.ReportService.GetReportForArtist(menu.artist.ArtistID)
	if err != nil {
		menu.log.Error("Can't get stats: ", slog.Any("error", err))
	}

	fmt.Println(report)
}
