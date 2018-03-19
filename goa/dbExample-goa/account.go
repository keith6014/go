package main

import (
	"dbExample-goa/app"
	"dbExample-goa/utils/database"
	"fmt"

	"github.com/goadesign/goa"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

func (c *AccountController) Add(ctx *app.AddAccountContext) error {

	db, err := database.SqlLiteConnect("./data.db")
	checkErr(err)
	smtp, err := db.Prepare("INSERT INTO tbl (one,two) values (?,?)")
	checkErr(err)
	_, err = smtp.Exec(ctx.Left, ctx.Right)
	checkErr(err)
	return nil
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	// AccountController_List: start_implement

	// Put your logic here
	//type AccountCollection []*Account

	var k []*app.Account
	db, err := database.SqlLiteConnect("./data.db")
	checkErr(err)

	rows, err := db.Query("select * from tbl1 ")
	checkErr(err)
	for rows.Next() {
		var (
			x int
			y int
		)
		err = rows.Scan(&x, &y)
		fmt.Println(x, y)
		checkErr(err)
		k = append(k, &app.Account{Cid: y, ID: x})
	}
	res := app.AccountCollection(k)
	return ctx.OK(res)
	// AccountController_List: end_implement
}
