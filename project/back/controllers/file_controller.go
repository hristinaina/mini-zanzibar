package controllers

import (
	"back/dtos"
	"back/models"
	"back/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type FileController struct {
	fileService services.FileService
	aclService  services.ACLService
	namespace   string
}

func NewFileController(db *sql.DB) FileController {
	return FileController{
		fileService: services.NewFileService(db),
		aclService:  services.NewACLService(),
		namespace:   "doc",
	}
}

func (fc *FileController) Create(c *gin.Context) {
	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	file.Owner = currentUser.Email
	createdFile, err := fc.fileService.CreateFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}

	// Grant owner access
	relation := dtos.Relation{
		User:     "user:" + currentUser.Email,
		Object:   fc.namespace + ":" + createdFile.ID,
		Relation: "owner",
	}

	resp, err := fc.aclService.AddRelation(relation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set file owner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File created successfully", "file": createdFile})
}

func (fc *FileController) Modify(c *gin.Context) {
	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	// Check if user is owner or editor
	isAuthorized, err := fc.aclService.CheckRelation(dtos.Relation{
		User:     "user:" + currentUser.Email,
		Object:   fc.namespace + ":" + file.ID,
		Relation: "owner",
	})
	if err != nil || !isAuthorized {
		isAuthorized, err = fc.aclService.CheckRelation(dtos.Relation{
			User:     "user:" + currentUser.Email,
			Object:   fc.namespace + ":" + file.ID,
			Relation: "editor",
		})
		if err != nil || !isAuthorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to modify file"})
			return
		}
	}

	modifiedFile, err := fc.fileService.ModifyFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to modify file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File modified successfully", "file": modifiedFile})
}

func (fc *FileController) ShareAccess(c *gin.Context) {
	var relation dtos.Relation
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	// Ensure only owner can share access
	isAuthorized, err := fc.aclService.CheckRelation(dtos.Relation{
		User:     "user:" + currentUser.Email,
		Object:   fc.namespace + ":" + relation.Object,
		Relation: "owner",
	})
	if err != nil || !isAuthorized {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to share access"})
		return
	}

	relation.Object = fc.namespace + ":" + relation.Object
	relation.User = "user:" + relation.User

	if _, err := fc.aclService.AddRelation(relation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share access"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access shared successfully"})
}

func (fc *FileController) GetUserFiles(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	files, err := fc.fileService.GetFilesByUser(currentUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user files"})
		return
	}

	c.JSON(http.StatusOK, files)
}

func (fc *FileController) GetSharedFiles(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	files, err := fc.fileService.GetFilesSharedWithUser(currentUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get shared files"})
		return
	}

	c.JSON(http.StatusOK, files)
}
