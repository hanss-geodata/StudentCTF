package oppgave2

import (
	_ "embed"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/fingann/swat/pkg/flags"

	"github.com/fingann/swat/public"
	"github.com/gin-gonic/gin"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
)

func RegisterRoutes(r *gin.Engine) error {
	setupCfolder()

	r.GET("/oppgave2", func(c *gin.Context) {
		// Refresh the whole page if the header is not set
		if c.GetHeader("HX-Request") == "" {
			c.HTML(http.StatusOK, "", public.Base(Page()))
			return
		}

		c.HTML(http.StatusOK, "", Page())
	})

	r.GET("/oppgave2/download", func(c *gin.Context) {
		fileName := c.Query("file")
		finalPath := path.Join(uploadFolder, fileName)
		file, err := GetFile(finalPath)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to download file: %w", err))
			return
		}
		info, err := cFolder.Stat(file.Name())
		if err != nil {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to get file info: %w", err))
			return
		}
		filename := path.Base(finalPath)
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		c.DataFromReader(http.StatusOK, info.Size(), "application/octet-stream", file, nil)
	})

	r.POST("/oppgave2/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to get file: %w", err))
			return
		}
		_, err = cFolder.Create(path.Join(uploadFolder, file.Filename))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create file: %w", err))
			return
		}
		//sseChan <- "upload"
		c.Redirect(http.StatusSeeOther, "/oppgave2/list/items")
	})

	r.GET("/oppgave2/list/items", func(c *gin.Context) {
		files, err := GetUploadedFiles()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get items: %w", err))
			return
		}
		fmt.Println("Sending items", len(files))
		c.HTML(http.StatusOK, "", ListItems(files))
	})

	return nil
}

// GetFile returns a file from the cFolder, helps to act like a file system
func GetFile(filePath string) (billy.File, error) {
	file, err := cFolder.Open(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "chroot boundary crossed") {
			return GetFile(strings.TrimPrefix(filePath, "../"))
		}
		if strings.HasPrefix(filePath, "C:/") {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}

		return GetFile(path.Join("C:", filePath))
	}

	return file, nil
}

var cFolder billy.Filesystem
var uploadFolder = path.Join("C:", "Program Files", "swat", "files")

//go:embed instruction.txt
var instructionFile []byte

func setupCfolder() error {
	cFolder = memfs.New()
	// create root folder
	cFolder.MkdirAll(uploadFolder, 0755)
	cFolder.Chroot("C:")

	// create InstructionFile
	file, err := cFolder.Create(path.Join(uploadFolder, "instructions.txt"))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	_, err = file.Write(instructionFile)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// Create Flag
	cFolder.MkdirAll("C:/windows", 0755)
	file, err = cFolder.Create("C:/windows/flag.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	_, err = file.Write([]byte(flags.Oppgave2Flag))
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func GetUploadedFiles() ([]UploadedFile, error) {
	dir, err := cFolder.ReadDir(uploadFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	files := make([]UploadedFile, 0, len(dir))
	for _, f := range dir {
		files = append(files, UploadedFile{Name: f.Name(), Size: f.Size()})
	}

	return files, nil
}
