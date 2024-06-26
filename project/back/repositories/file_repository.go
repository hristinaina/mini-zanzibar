package repositories

import (
	"back/models"
	"database/sql"
	"errors"
)

type FileRepository interface {
	CreateFile(file models.File) (*models.File, error)
	UpdateFile(file models.File) (*models.File, error)
	GetFileByID(id string) (*models.File, error)
	GetFilesByUser(owner string) ([]models.File, error)
	GetFilesSharedWithUser(user string) ([]models.File, error)
}

type FileRepositoryImpl struct {
	db *sql.DB
}

var ErrFileNotFound = errors.New("file not found")

func NewFileRepository(db *sql.DB) FileRepository {
	return &FileRepositoryImpl{db: db}
}

func (repo *FileRepositoryImpl) CreateFile(file models.File) (*models.File, error) {
	query := "INSERT INTO files (name, content, owner) VALUES ($1, $2, $3) RETURNING id"
	row := repo.db.QueryRow(query, file.Name, file.Content, file.Owner)

	err := row.Scan(&file.ID)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (repo *FileRepositoryImpl) UpdateFile(file models.File) (*models.File, error) {
	query := "UPDATE files SET name = $1, content = $2 WHERE id = $3 RETURNING id, name, content, owner"
	row := repo.db.QueryRow(query, file.Name, file.Content, file.ID)

	err := row.Scan(&file.ID, &file.Name, &file.Content, &file.Owner)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (repo *FileRepositoryImpl) GetFileByID(id string) (*models.File, error) {
	var file models.File
	query := "SELECT id, name, content, owner FROM files WHERE id = $1"
	row := repo.db.QueryRow(query, id)

	err := row.Scan(&file.ID, &file.Name, &file.Content, &file.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrFileNotFound
		}
		return nil, err
	}
	return &file, nil
}

func (repo *FileRepositoryImpl) GetFilesByUser(owner string) ([]models.File, error) {
	query := "SELECT id, name, content, owner FROM files WHERE owner = $1"
	rows, err := repo.db.Query(query, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		if err := rows.Scan(&file.ID, &file.Name, &file.Content, &file.Owner); err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func (repo *FileRepositoryImpl) GetFilesSharedWithUser(user string) ([]models.File, error) {
	//query := `
	//	SELECT f.id, f.name, f.content, f.owner
	//	FROM files f
	//	JOIN acl a ON f.id = a.object
	//	WHERE a.user = $1 AND a.relation IN ('editor', 'viewer')
	//`
	//rows, err := repo.db.Query(query, user)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//var files []models.File
	//for rows.Next() {
	//	var file models.File
	//	if err := rows.Scan(&file.ID, &file.Name, &file.Content, &file.Owner); err != nil {
	//		return nil, err
	//	}
	//	files = append(files, file)
	//}
	//return files, nil
	return nil, nil
}
