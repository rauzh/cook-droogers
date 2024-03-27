package service

type IReportService interface {
	GetReportForManager(mngID uint64) (map[string][]byte, error)
	GetReportForArtist(artistID uint64) (map[string][]byte, error)
}
