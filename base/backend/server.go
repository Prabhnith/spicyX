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

		fmt.Println("\n\nRequest Received : \n\n ", ven.Name, ven.Owner)

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		var track ID

		// insert into users table
		tx.QueryRow(`
        INSERT INTO vendors (owner, vendorname, email ,mobile ,address  ,imageaddress ,
		                      password ) values ($1, $2, $3, $4, $5, $6, $7) returning vendorid
          `, ven.Owner, ven.Name, ven.Email, ven.Mobile, ven.Address, ven.Image, ven.Password).Scan(&track.id)

		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "ERR")
			return
		}
		fmt.Println("Vendor registered and his ID:", track.id)
		c.JSON(200, track)

	})

	r.POST("/registercustomer", func(c *gin.Context) {
		var cus customer
		c.BindJSON(&cus)

		fmt.Println("\n\nRequest Received : \n\n")

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		var track CSID

		// insert into users table
		tx.QueryRow(`
        INSERT INTO customers (customer_name, emailid ,mobile ,address ,password ) values ($1, $2, $3, $4, $5) returning customer_id
          `, cus.Name, cus.Email, cus.Mobile, cus.Address, cus.Password).Scan(&track.id)

		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "ERR")
			return
		}
		fmt.Println("cutomer registered and his ID:", track.id)
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
	Vendorid    int      `json:"vendorid,omitempty"`
	Owner       string   `json:"owner"`
	Name        string   `json:"vendorname"`
	Email       string   `json:"email"`
	Mobile      []string `json:"mobile"`
	Address     string   `json:"address"`
	Image       string   `json:"imageaddress,omitempty"`
	Description string   `json:"description,omitempty"`
	Password    string   `json:"password"`
}

type ID struct {
	id int64 `json:"vendorid,omitempty"`
}

type customer struct {
	Customerid int      `json:"customer_id,omitempty"`
	Name       string   `json:"customer_name"`
	Email      string   `json:"emailid"`
	Mobile     []string `json:"mobile"`
	Address    string   `json:"address"`
	Password   string   `json:"password"`
}

type CSID struct {
	id int64 `json:"vendorid,omitempty"`
}
