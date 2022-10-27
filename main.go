package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/th3khan/api-quiniela-world-cup/pkg/seeders"
	"github.com/th3khan/api-quiniela-world-cup/pkg/utils"
	"github.com/th3khan/api-quiniela-world-cup/platform/migrations"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if !fiber.IsChild() {
		migrations.Migrate()
	}

	args := os.Args

	seeder := seeders.NewSeeder()
	if len(args) > 0 {
		otherArgs := args[1:]
		for _, argName := range otherArgs {
			argList := strings.Split(argName, "=")
			if len(argList) == 2 {
				if argList[0] == "--seed" {
					fmt.Println(argList)
					action := argList[1]
					valid := reflect.ValueOf(seeder).MethodByName(action).IsValid()
					if valid {
						reflect.ValueOf(seeder).MethodByName(action).Call([]reflect.Value{})
					} else {
						log.Fatal(fmt.Sprintf("Seeder: %s, not found", action))
					}
				}
			}
		}
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
