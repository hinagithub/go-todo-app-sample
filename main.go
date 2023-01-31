package main

import (
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
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {

	db := connectDb()
	defer db.Close()

	router := gin.New()

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
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// list all TODO.
	router.GET("/todo", func(c *gin.Context) {
		result := models.FindAll(db)
		fmt.Println(result)
		c.JSON(200, gin.H{
			"result": result,
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
		models.Add(db, todo.Title, todo.Body)
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
		err = models.Edit(db, id, todo.Title, todo.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"result": "updated.",
		})
	})

	// listen.
	router.Run(":3000")
}

func isValidRequest(todo TodoRequest) error {
	if todo.Title == "" {
		return errors.New("title required.")
	}
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
