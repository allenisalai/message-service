package main

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

// Hold a single message row
//
type Message struct {
	ID   uint   `json:"id";gorm:"primary_key"`
	Type string `json:"type";gorm:"type:varchar(10)"`
	//Metadata  postgres.Jsonb
	Message   string `json:"message";gorm:"type:varchar(255)"`
	CreatedAt time.Time
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var users []Message
		db.Find(&users)

		c.JSON(200, users)
	})

	r.POST("/", func(c *gin.Context) {
		b, err := c.GetRawData()

		if err != nil {
			c.JSON(400, err)
			return
		}
		m := Message{}
		err = json.Unmarshal(b, &m)

		db.Create(&m)

		c.JSON(200, m)

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
