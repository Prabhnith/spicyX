package main

import (
	"fmt"
	"io/ioutil"

	"strings"

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

	//*************************Hosting client.html page
	r.GET("/main", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/foodies/spicyX/base/main.html")
		c.Data(200, "text/html", res)
	})

	r.GET("/menuItems", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/foodies/spicyX/base/dashboard/menuItems.html")
		c.Data(200, "text/html", res)
	})

	r.GET("/dashboard", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/foodies/spicyX/base/dashboard/dashboard.html")
		c.Data(200, "text/html", res)
	})
	r.GET("/dash", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/foodies/spicyX/base/dashboard/dash.html")
		c.Data(200, "text/html", res)
	})
	r.GET("/table", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/foodies/spicyX/base/dashboard/table.html")
		c.Data(200, "text/html", res)
	})
	r.GET("/user", func(c *gin.Context) {
		res, _ := ioutil.ReadFile("/home/anil/foodies/spicyX/base/dashboard/user.html")
		c.Data(200, "text/html", res)
	})

	//**********************fetching Javascript files file
	r.GET("/js/:js_file", func(c *gin.Context) {
		//to ser
		jsFile := c.Param("js_file")

		res, err := ioutil.ReadFile("/home/anil/foodies/spicyX/base/js/" + jsFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, "error while fetching file")
		}
		c.Data(200, "application/javascript", res)

		// c.Data(200, path.Join("applications", "javascript"), res)
	})

	//********************fetching CSS files
	r.GET("/css/:css_file", func(c *gin.Context) {
		//to ser
		cssFile := c.Param("css_file")

		res, err := ioutil.ReadFile("/home/anil/foodies/spicyX/base/css/" + cssFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, "error while fetching file")
		}
		c.Data(200, "text/css", res)

		// c.Data(200, path.Join("applications", "javascript"), res)
	})

	//********************fetching Images
	r.GET("/img/:img_file", func(c *gin.Context) {
		//to ser
		imgFile := c.Param("img_file")
		extension := strings.ToLower(strings.Split(imgFile, ".")[1])

		res, err := ioutil.ReadFile("/home/anil/foodies/spicyX/base/img/" + imgFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, "error while fetching Image")
		}

		if extension == "jpg" {
			c.Data(200, "image/jpg", res)
		} else if extension == "png" {
			c.Data(200, "image/png", res)
		} else if extension == "jpeg" {
			c.Data(200, "image/png", res)
		}

		// c.Data(200, path.Join("applications", "javascript"), res)
	})

	//********************Registering vendors
	r.POST("/registervendor", func(c *gin.Context) {
		var ven vendor

		c.BindJSON(&ven)

		fmt.Println("\n\nRequest Received  for vendor registration: \n\n ")

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function
		fmt.Println(ven.Owner, ven.Name, ven.Email, ven.Mobile, ven.Address, ven.Image, ven.Description, ven.Offer, ven.Password)
		// var track ID
		var num int64
		// insert into users table
		err := tx.QueryRow(`
        INSERT INTO vendors (owner, vendorname, email ,mobile ,address  ,imageaddress ,description,offer, password ) values ($1, $2, $3, $4, $5, $6, $7,$8,$9) returning vendorid
          `, ven.Owner, ven.Name, ven.Email, ven.Mobile, ven.Address, ven.Image, ven.Description, ven.Offer, ven.Password).Scan(&num)
		fmt.Println(err)
		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "ERR")
			return
		}
		fmt.Println("Vendor registered and his ID:", num)
		c.JSON(200, num)

	})

	//I**************************tem menu updation
	r.POST("/additems", func(c *gin.Context) {
		var menu MENU

		c.BindJSON(&menu)

		fmt.Println("\n\nRequest Received for menu updation: \n\n ")
		// fmt.Printf("%#v", menu)
		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		for _, item := range menu.ITEMS {
			// fmt.Println(item.Vendorid, item.Itemno, item.Name, item.IType, item.Nature, item.Price, item.Description,
			// 	item.Offer, item.Image, item.Discount)
			_, err := tx.Exec(`
				INSERT INTO itemmenu (vendor_id ,item_no ,item_name ,item_type ,item_nature ,price , item_description ,imageaddress ,discount) 
				values ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`, item.Vendorid, item.Itemno, item.Name, item.IType, item.Nature, item.Price, item.Description, item.Image, item.Discount)

			if err != nil {
				// c.JSON(500, "error")
				fmt.Println("error", err)
			}
		}

		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "ERR")
			return
		}
		// fmt.Println("Menu  updated")
		c.JSON(200, 1)

	})

	//*************************customer registration
	r.GET("/registercustomer", func(c *gin.Context) {
		var cus customer
		c.BindJSON(&cus)

		fmt.Println("\n\nRequest Received : \n\n")

		tx, _ := db.Begin() // tx => transaction , _ => error and execute

		defer tx.Rollback() // it will be executed after the completion of local function

		var track CSID

		// insert into users table
		tx.QueryRow(`
        INSERT INTO customers (customer_name, emailid ,mobile ,address ,password ) values ($1, $2, $3, $4, $5) returning customer_id
          `, cus.Name, cus.Email, cus.Mobile, cus.Address, cus.Password).Scan(&track.Customerid)

		commit_err := tx.Commit()

		if commit_err != nil {
			tx.Rollback()
			c.JSON(500, "ERR")
			return
		}
		fmt.Println("cutomer registered and his ID:", track.Customerid)
		c.JSON(200, track)

	})

	//*****************************Serving vendors and their id's
	r.GET("/getvendors", func(c *gin.Context) {
		// c.BindJSON(&cus)

		fmt.Println("\n\nRequest Received : \n\n")

		rows, err := db.Query(` SELECT  vendorid, vendorname from vendors `)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, "error while retreiving vendors data")
		}

		defer rows.Close()

		// var vendors = make(map[string]int)
		ven := make([]VendorsToSend, 0)

		for rows.Next() {
			var t VendorsToSend
			err := rows.Scan(&t.Vendorid, &t.Vendorname)
			ven = append(ven, t)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error while retreiving vendors data")
			}
		}
		c.JSON(200, ven)
		fmt.Println("Vendors names are sent")
	})

	//****************** method to serve request for MENU of particular vendor
	r.GET("/getvendorsmenu", func(c *gin.Context) {
		var id VID
		c.BindJSON(&id)

		fmt.Println("\n\nRequest for retreiving vendors menu Received : \n\n")

		rows, err := db.Query(` SELECT  item_no, item_name, item_type, item_nature, price, item_description, imageaddress, discount
		                        from itemmenu where vendor_id = $1 `, id.Vendorid)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, "error while retreiving vendors menu")
		}

		defer rows.Close()

		// var vendors = make(map[string]int)
		var items MENU

		for rows.Next() {
			var t item
			err := rows.Scan(&t.Itemno, &t.Name, &t.IType, &t.Nature, &t.Price, &t.Description, &t.Image, &t.Discount)
			items.ITEMS = append(items.ITEMS, t)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, "error while retreiving vendors menu")
			}
		}
		c.JSON(200, items)
		fmt.Println("Vendors Menu  sent")
	})

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
	Offer       string   `json:"offer,omitempty"`
	Password    string   `json:"password"`
}

type VID struct {
	Vendorid int `json:"vendorid"`
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
	Customerid int `json:"customerid,omitempty"`
}

type MENU struct {
	ITEMS []item `json:"items"`
}

//Menu updation
type item struct {
	Vendorid    int     `json:"vendor_id"`
	Itemno      int     `json:"item_no"`
	Name        string  `json:"item_name"`
	IType       string  `json:"item_type"`
	Nature      bool    `json:"item-nature"`
	Description string  `json:"item_description"`
	Price       string  `json:"price"`
	Image       string  `json:"imageaddress,omitempty"`
	Discount    float64 `json:"discount,omitempty"`
}

type VendorsToSend struct {
	Vendorid   int    `json:"vendor_id"`
	Vendorname string `json:"vendorname"`
}
