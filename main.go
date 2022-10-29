package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/config"
	"github.com/th3khan/api-quiniela-world-cup/pkg/seeders"
	"github.com/th3khan/api-quiniela-world-cup/pkg/utils"
	"github.com/th3khan/api-quiniela-world-cup/platform/migrations"
)

func main() {
	config.InitJWT()

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

	port, _ := strconv.Atoi(config.Config("APP_PORT"))

	utils.CreateServer(port)
}
