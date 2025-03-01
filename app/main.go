// Package main ...
//
// @title My PDF API
// @version 1.0
// @description This is a sample server for PDF management
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @BasePath /
// @schemes http
package main

import (
	// "database/sql"
	// "fmt"
	"log"
	// "net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	// "github.com/PHAWAPHON/go-clean-arch/article"
	// mysqlRepo "github.com/PHAWAPHON/go-clean-arch/internal/repository/mysql"
	_ "github.com/PHAWAPHON/go-clean-arch/app/docs" // เปลี่ยนเป็น module path ของโปรเจกต์คุณ
	pdfRepo "github.com/PHAWAPHON/go-clean-arch/internal/repository/pdf_repo"
	"github.com/PHAWAPHON/go-clean-arch/internal/rest"
	"github.com/PHAWAPHON/go-clean-arch/internal/rest/middleware"
	pdfSrv "github.com/PHAWAPHON/go-clean-arch/pdf"
	"github.com/joho/godotenv"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//prepare database
	// dbHost := os.Getenv("DATABASE_HOST")
	// dbPort := os.Getenv("DATABASE_PORT")
	// dbUser := os.Getenv("DATABASE_USER")
	// dbPass := os.Getenv("DATABASE_PASS")
	// dbName := os.Getenv("DATABASE_NAME")
	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	// dbConn, err := sql.Open(`mysql`, dsn)
	// if err != nil {
	// 	log.Fatal("failed to open connection to database", err)
	// }
	// err = dbConn.Ping()
	// if err != nil {
	// 	log.Fatal("failed to ping database ", err)
	// }

	// defer func() {
	// 	err := dbConn.Close()
	// 	if err != nil {
	// 		log.Fatal("got error when closing the DB connection", err)
	// 	}
	// }()
	// // prepare echo

	e := echo.New()
	e.Use(middleware.CORS)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// // Prepare Repository
	// authorRepo := mysqlRepo.NewAuthorRepository(dbConn)
	// articleRepo := mysqlRepo.NewArticleRepository(dbConn)

	// // Build service Layer
	// svc := article.NewService(articleRepo, authorRepo)
	// rest.NewArticleHandler(e, svc)

	PDFrepo := pdfRepo.NewPDFRepository()

	PDFsvc := pdfSrv.NewService(PDFrepo)

	rest.NewPDFHandler(e, PDFsvc)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}

	// rest.NewPDFHandler(e, pdfSvc)
	log.Fatal(e.Start(address)) //nolint
}
