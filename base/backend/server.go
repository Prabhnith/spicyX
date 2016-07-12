package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// Database connectivity variables
var db *pgx.ConnPool
var db_err error

//Initialise connection to the database
func init() {
	db, db_err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Database: "foodies",
			User:     "anil",
			Password: "205474",
			Port:     5432,
		},
		MaxConnections: 10,
	})

	if db_err != nil {
		fmt.Println("Can't connect to database")
	}
}

func main() {
	r := gin.Default()

	r.POST("/registervendor", func(c *gin.Context) {
		var ven vendor
		c.BindJSON(&ven)

		fmt.Println("\n\nRequest Received : \n\n")

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		var track VenID

		// insert into users table
		tx.QueryRow(`
        INSERT INTO vendors (owner, vendorname, email ,mobile ,address ,offer ,imageaddress ,
		                      password ) values ($1, $2, $3, $4, $5, $6, $7, $8) returning vendorid
          `, ven.Owner, ven.Name, ven.Email, ven.Mobile, ven.Address, ven.Offer, ven.Image, ven.Password).Scan(&track.id)

		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "ERR")
			return
		}
		fmt.Println("Vendor registored and his ID:", track.id)
		c.JSON(200, track)

	})

	// r.GET("/api/verifyemail", func(c *gin.Context) {
	// 	// receive userid and map it with the table users and get email
	// 	var userid UserIDResp
	// 	c.BindJSON(&userid)

	// 	if userid.Userid <= 0 {
	// 		response_map := make(map[string]string)
	// 		response_map["error"] = "Invalid Userid"
	// 		c.JSON(404, response_map)
	// 		return
	// 	}

	// 	var email string
	// 	db.QueryRow(`
	// 		SELECT email
	// 		FROM users
	// 		WHERE userid = $1
	// 	`, userid.Userid).Scan(&email)

	// 	if email == "" {
	// 		response_email := make(map[string]string)
	// 		response_email["error"] = "Userid Not found"
	// 		c.JSON(403, response_email)
	// 		return
	// 	}

	// 	if email != EmailBefore {
	// 		responsefail := make(map[string]string)
	// 		responsefail["error"] = "can't generate user"
	// 		c.JSON(405, responsefail)
	// 		return
	// 	}

	// 	fmt.Printf("\n\nUserid allocated to corresponding Email\n\n")

	// 	emailMap := make(map[string]string)
	// 	emailMap["email"] = email
	// 	c.JSON(200, emailMap)
	// })
	fmt.Println("\n\n\t #####     Foodies server live on :7070     #####")
	r.Run(":7070")
}

// vendor holds the incoming requests for a vendor registration.
type vendor struct {
	Vendorid int      `json:"vendorid,omitempty"`
	Owner    string   `json: owner`
	Name     string   `json: vendorname`
	Email    string   `json: email`
	Mobile   []string `json: mobile`
	Address  string   `json: address`
	Offer    string   `json: offer`
	Image    string   `json:"imageaddress,omitempty"`
	Password string   `json: password`
}

type VenID struct {
	id int64 `json:"vendorid,omitempty"`
}

// Request holds the incoming request for a push notification.
type Request struct {
	Route        string   `json:"route,omitempty"`
	RecentlyUsed device   `json:"recently_used,omitempty"`
	DeviceList   []device `json:"device_list,omitempty"`
	Content      string   `json:"content,omitempty"`
	Timeout      int64    `json:"timeout_after,omitempty"`
	Priority     []string `json:"priority,omitempty"`
}

type device struct {
	Platform string `json:"platform,omitempty"`
	DeviceID string `json:"device_id,omitempty"`
}

// Response is the struct that holds the notification ID and the UNIX timestamp of when it was sent
type Response struct {
	NotificationID int64 `json:"notification_id,omitempty"`
	Timestamp      int64 `json:"timestamp,omitempty"`
}

// Request for creating userid for received email

type UserIDCreate struct {
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	UserName string `json:"username,omitempty"`
}

type UserIDResp struct {
	Userid int64 `json:"userid,omitempty"`
}

type DumpUsersdescription struct {
	UserID   int64  `json:"userid,omitempty"`
	DeviceID string `json:"deviceid,omitempty"`
	Platform string `json:"platform,omitempty"`
}

type ToPNS struct {
	DeviceID string `json:"deviceid,omitempty"`
	Platform string `json:"platform,omitempty"`
}
