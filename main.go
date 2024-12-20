package main

import (
	"api_banklist/docs"
	"api_banklist/model"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	// _ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// swagger embed files
)

var db *sql.DB
var driver string

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("Environment variable %s not set, using default value: %s", key, defaultValue)
	return defaultValue
}

func initDB() {
	var err error
	username := getEnv("DB_USERNAME", "root")
	password := getEnv("DB_PASSWORD", "")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	database := getEnv("DB_DATABASE", "bank-indonesia")
	driver = getEnv("DB_DRIVER", "mysql")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	if driver == "postgres" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require&binary_parameters=yes", username, password, host, port, database)
	}

	log.Printf("Attempting to connect to %s database at %s:%s...", driver, host, port)
	db, err = sql.Open(driver, dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Successfully connected to the database")
}

func checkTableExist() {
	log.Println("Checking if required tables exist...")
	tables := []string{"banks"}
	for _, table := range tables {
		log.Printf("Checking table: %s", table)
		if _, err := db.Query(fmt.Sprintf("SELECT 1 FROM %s LIMIT 1", table)); err != nil {
			log.Fatalf("Table %s does not exist: %v", table, err)
		}
		log.Printf("Table %s exists", table)
	}
	log.Println("All required tables exist")
}

func info(c *gin.Context) {
	log.Println("Handling request for info endpoint")
	counts := make(map[string]int)
	tables := []string{"banks"}

	for _, table := range tables {
		var count int
		log.Printf("Counting records in %s table", table)
		if err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&count); err != nil {
			log.Printf("Error counting %s: %v", table, err)
			c.JSON(500, gin.H{"error": fmt.Sprintf("Error counting %s", table)})
			return
		}
		counts[table] = count
		log.Printf("%s count: %d", table, count)
	}

	c.JSON(200, gin.H{
		"banks_count": counts["banks"],
	})
	log.Println("Info request handled successfully")
}

func getItems(c *gin.Context, tableName, columnName string) {
	log.Printf("Handling request for %s items", tableName)
	searchQuery := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset := (page - 1) * limit

	log.Printf("Search query: %s, Page: %d, Limit: %d", searchQuery, page, limit)

	query := fmt.Sprintf("SELECT id, %s,city, swift_code FROM %s", columnName, tableName)
	args := make([]interface{}, 0)

	if searchQuery != "" {
		query += fmt.Sprintf(" WHERE %s LIKE ?", columnName)
		args = append(args, "%"+searchQuery+"%")
	}

	query += " ORDER BY id ASC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	if driver == "postgres" {
		query = convertToPostgres(query)
	}

	log.Printf("Executing query: %s", query)
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var itemsBank []model.Bank
	// Data to append

	for rows.Next() {
		var id int
		var name string
		var city string
		var swiftCode string
		if err := rows.Scan(&id, &name, &city, &swiftCode); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		if tableName == "banks" {
			itemsBank = append(itemsBank, model.Bank{
				ID:        id,
				Bank:      name,
				City:      city,
				SwiftCode: swiftCode,
			})
		} else {

		}

	}
	// Create a Province struct and append it to items

	if tableName == "banks" {
		response := model.BanksDataResponse{
			TotalCount: len(itemsBank), // Assuming no provinces are returned, hence count is 0
			Banks:      itemsBank,
		}
		c.JSON(200, response)
	} else {
		// if we have another table response
	}

}

// @Tags banks
// banks godoc
// @Summary banks
// @Description banks
// @Accept json
// @Produce json
// @Router /banks [GET]
// @Accepts json
// @Produce json
// @Success 200 {object} model.BanksDataResponse 	"List of Banks with total count"
//
// @Failure 206 {object} model.FailResponse             "Error lain"
// @Failure 207 {object} model.FailResponse             "Gagal Login"
// @Failure 209 {object} model.FailResponse             "Invalid JSON input"
// @Failure 210 {object} model.FailResponse             "Error Middleware Internal"
// @Failure 214 {object} model.FailResponse             "Data Pengguna Tidak Ditemukan"
// @Failure 500 {string} Internal Server Fatal Error
func getBanks(c *gin.Context) {
	getItems(c, "banks", "bank")
}

// @Tags banks
// banks godoc
// @Summary bank
// @Description bank
// @Param BankData body model.BankDataRequest true "Bank data"
// @Accept json
// @Produce json
// @Router /bank [POST]
// @Accepts json
// @Produce json
// @Success 200 {object} model.AddBanksDataResponse 	"List of Banks with total count"
//
// @Failure 206 {object} model.FailResponse             "Error lain"
// @Failure 207 {object} model.FailResponse             "Gagal Login"
// @Failure 209 {object} model.FailResponse             "Invalid JSON input"
// @Failure 210 {object} model.FailResponse             "Error Middleware Internal"
// @Failure 214 {object} model.FailResponse             "Data Pengguna Tidak Ditemukan"
// @Failure 500 {string} Internal Server Fatal Error
func addBank(c *gin.Context) {
	log.Println("post add bank")
	// var bank model.Bank
	//init

	// bankName := c.PostForm("bankName")
	// bankCity := c.PostForm("bankCity")
	// bankBranch := c.PostForm("bankBranch")
	// bankSwiftCode := c.PostForm("bankSwiftCode")
	var bank model.Bank
	err := c.BindJSON(&bank)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("coba", bank)
	// log.Println("coba", bankName, bankCity, bankBranch, bankSwiftCode)
	username := getEnv("DB_USERNAME", "root")
	database := getEnv("DB_DATABASE", "bank-indonesia")
	driver = getEnv("DB_DRIVER", "mysql")
	db, err := sql.Open("mysql", username+"@tcp(127.0.0.1:3306)/"+database)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO `banks` (`bank`, `city`, `branch`, `swift_code`) VALUES (?, ?, ?, ?)", bank.Bank, bank.City, bank.Branch, bank.SwiftCode)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	bank.ID = int(lastID)

	c.JSON(200, gin.H{
		"success": true,
		"message": "success create post",
		"data":    bank,
	})

}

func convertToPostgres(query string) string {
	log.Println("Converting MySQL query to PostgreSQL format")
	paramCount := 1
	for i := 0; i < len(query); i++ {
		if query[i] == '?' {
			query = query[:i] + fmt.Sprintf("$%d", paramCount) + query[i+1:]
			paramCount++
		}
	}
	log.Printf("Converted query: %s", query)
	return query
}

// @title API for the get of Data Bank SwiftCode in Indonesian
// @version 1.0
// @description This is a sample API documentation.
// @contact.name API Support
// @contact.url http://www.terbitdigitalagency.com
// @contact.email hegy.febrianto@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1

// @tag.name banks
// @tag.description API yang digunakan untuk mendapatkan data banks
//

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using default environment variables")
		log.Println("masuk if")
	} else {
		log.Println("Loaded environment variables from .env file")
		log.Println("masuk else")
	}

	port := getEnv("PORT", "8080")
	log.Printf("Using port: %s", port)

	initDB()
	checkTableExist()
	defer db.Close()

	log.Println("Setting up Gin router...")
	router := gin.Default()

	router.GET("/", info)
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// router.GET("/swagger", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/banks", getBanks)
	router.POST("/addbank", addBank)

	log.Printf("Starting server on 0.0.0.0:%s", port)
	router.Run("0.0.0.0:" + port)

}
