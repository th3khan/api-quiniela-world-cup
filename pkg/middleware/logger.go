package middleware

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) *os.File {
	basepath, _ := filepath.Abs("./")

	currentDate := time.Now()

	currentDateStr := fmt.Sprintf(
		"%d-%d-%d",
		currentDate.Year(),
		currentDate.Month(),
		currentDate.Day(),
	)

	file, err := os.OpenFile(basepath+"/logs/"+currentDateStr+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Caracas",
	}))

	return nil
}
