package services

import (
	"back/models"
	"back/repositories"
	"database/sql"
)

type FileService struct {
	fileRepo repositories.FileRepository
}

func NewFileService(db *sql.DB) FileService {
	return FileService{
		fileRepo: repositories.NewFileRepository(db),
	}
}

func (fs *FileService) CreateFile(file models.File) (*models.File, error) {
	return fs.fileRepo.CreateFile(file)
}

func (fs *FileService) ModifyFile(file models.File) (*models.File, error) {
	return fs.fileRepo.UpdateFile(file)
}

func (fs *FileService) GetFilesByUser(userID string) ([]models.File, error) {
	return fs.fileRepo.GetFilesByUser(userID)
}

func (fs *FileService) GetFilesSharedWithUser(userID string) ([]models.File, error) {
	return fs.fileRepo.GetFilesSharedWithUser(userID)
}
