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
	logService  *services.LogService
}

func NewFileController(db *sql.DB, logService *services.LogService) FileController {
	return FileController{
		fileService: services.NewFileService(db),
		aclService:  services.NewACLService(),
		namespace:   "doc",
		logService:  logService,
	}
}

func (fc *FileController) Create(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	fc.logService.Info("Processing Create File request. USER: " + currentUser.Email)
	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		fc.logService.Error("Bad input data. USER: " + currentUser.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	file.Owner = currentUser.Email
	createdFile, err := fc.fileService.CreateFile(file)
	if err != nil {
		fc.logService.Error("Failed to create file. USER: " + currentUser.Email)
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
		fc.logService.Error("Failed to send request to zanzibar. USER: " + currentUser.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Zanzibar"})
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fc.logService.Error("Failed to set file owner. USER: " + currentUser.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set file owner"})
		return
	}

	fc.logService.Info("File created successfully. USER: " + currentUser.Email + " . File: " + createdFile.ID)
	c.JSON(http.StatusOK, gin.H{"message": "File created successfully", "file": createdFile})
}

func (fc *FileController) Modify(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	fc.logService.Info("Processing Modify File request. USER: " + currentUser.Email)
	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		fc.logService.Error("Bad input data. USER: " + currentUser.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

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
			fc.logService.Error("Not authorized to modify file. USER: " + currentUser.Email)
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to modify file"})
			return
		}
	}

	modifiedFile, err := fc.fileService.ModifyFile(file)
	if err != nil {
		fc.logService.Error("Failed to modify file. USER: " + currentUser.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to modify file"})
		return
	}

	fc.logService.Info("File modified successfully. USER: " + currentUser.Email + " . File: " + modifiedFile.ID)
	c.JSON(http.StatusOK, gin.H{"message": "File modified successfully", "file": modifiedFile})
}

func (fc *FileController) ShareAccess(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	fc.logService.Info("Processing Share Access request. USER: " + currentUser.Email)
	var relation dtos.Relation
	if err := c.ShouldBindJSON(&relation); err != nil {
		fc.logService.Error("Bad input data. USER: " + currentUser.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad input data"})
		return
	}

	// Ensure only owner can share access
	isAuthorized, err := fc.aclService.CheckRelation(dtos.Relation{
		User:     "user:" + currentUser.Email,
		Object:   fc.namespace + ":" + relation.Object,
		Relation: "owner",
	})
	if err != nil || !isAuthorized {
		fc.logService.Error("Not authorized to share access. USER: " + currentUser.Email)
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to share access"})
		return
	}

	relation.Object = fc.namespace + ":" + relation.Object
	relation.User = "user:" + relation.User

	if _, err := fc.aclService.AddRelation(relation); err != nil {
		fc.logService.Error("Failed to share access. USER: " + currentUser.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share access"})
		return
	}

	fc.logService.Info("Access shared successfully. USER: " + currentUser.Email + " . Object: " + relation.Object + " . User: " + relation.User)
	c.JSON(http.StatusOK, gin.H{"message": "Access shared successfully"})
}

func (fc *FileController) GetUserFiles(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	fc.logService.Info("Processing Get User Files request. USER: " + currentUser.Email)
	files, err := fc.fileService.GetFilesByUser(currentUser.Email)
	if err != nil {
		fc.logService.Error("Failed to get user files. USER: " + currentUser.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user files"})
		return
	}

	fc.logService.Info("User files retrieved successfully. USER: " + currentUser.Email)
	c.JSON(http.StatusOK, files)
}

func (fc *FileController) GetSharedFiles(c *gin.Context) {
	currentUserFromCookie, _ := c.Get("user")
	currentUser := currentUserFromCookie.(*models.User)

	fc.logService.Info("Processing Get Shared Files request. USER: " + currentUser.Email)
	files, err := fc.fileService.GetFilesSharedWithUser(currentUser.Email)
	if err != nil {
		fc.logService.Error("Failed to get shared files. USER: " + currentUser.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get shared files"})
		return
	}

	fc.logService.Info("Shared files retrieved successfully. USER: " + currentUser.Email)
	c.JSON(http.StatusOK, files)
}
