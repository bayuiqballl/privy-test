package app

import (
	"fmt"

	"log"
	"os"

	mysql "privy-test/drivers/mysql"
	"privy-test/repositories"

	"privy-test/usecases"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	Repo := repositories.NewRepo(mysql.InitMysql())
	log.Println(Repo)
	app := usecases.NewUsecase(Repo)
	RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
