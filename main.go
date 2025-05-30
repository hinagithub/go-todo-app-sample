package main

import (
	"embed"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TodoRequest struct {
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var f embed.FS

func main() {

	db := connectDb()
	defer db.Close()

	router := gin.New()

	// HTML
	router.LoadHTMLGlob("templates/*.html")

	// Image
	router.Static("assets", "./assets")
	fmt.Println()

	// CROSS ORIGIN
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", " OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 404
	router.NoRoute(func(c *gin.Context) {
		fmt.Println("Page not exists.")
		c.HTML(http.StatusNotFound, "notfound.html", nil)
	})

	// list all TODO.
	router.GET("/todo", func(c *gin.Context) {
		result := models.FindAll(db)
		fmt.Println(result)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Data": result,
		})
	})

	// create new TODO.
	router.POST("/todo", func(c *gin.Context) {
		var todo TodoRequest
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := isValidRequest(todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		models.Add(db, todo.Completed, todo.Body)
		c.JSON(200, gin.H{
			"result": "created.",
		})
	})

	// update TODO.
	router.POST("/todo/:id", func(c *gin.Context) {
		var todo TodoRequest
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := isValidRequest(todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = models.Edit(db, id, todo.Completed, todo.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"result": "updated.",
		})
	})

	// delete TODO.
	router.DELETE("/todo/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = models.Delete(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"result": "deleted.",
		})
	})

	// listen.
	router.Run(":3000")
}

func isValidRequest(todo TodoRequest) error {
	if todo.Body == "" {
		return errors.New("body required.")
	}
	return nil
}

func isValidId(id_str string) error {
	_, err := strconv.Atoi(id_str)
	if err != nil {
		return errors.New("id invalid.")
	}
	return nil
}

func connectDb() *sqlx.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	fmt.Println("ENV: ", os.Getenv("MYSQL_HOST"))
	dbconf := mysql.Config{
		DBName:               os.Getenv("MYSQL_DATABASE"),
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Addr:                 os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	fmt.Println(dbconf)
	db, err := sqlx.Connect("mysql", dbconf.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Println(err)
		db.Close()
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("❌ Failed to connect DB.")
		db.Close()
		panic(err)
	} else {
		fmt.Println("🟢 DB connected.")
	}
	return db
}
