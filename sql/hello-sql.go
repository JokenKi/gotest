package user

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import  "context"
import  "log"

func main1(){
	
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tbd?charset=utf8")
	
	if err != nil {
	  log.Fatal(err)
	}
	
	ctx,stop := context.WithCancel(context.Background())
	defer stop()
	
	if err := db.PingContext(ctx); err != nil {
	  log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := db.QueryContext(ctx, "SELECT base_area.ID id,base_area.AREA_NAME area_name FROM base_area")
	defer rows.Close()
	
	for rows.Next() {
		var id int64
		var area_name string
		if err := rows.Scan(&id,&area_name); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s id: %d\n", area_name,id)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}


func PrintAreaName(){
	
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tbd?charset=utf8")
	
	if err != nil {
	  log.Fatal(err)
	}
	
	ctx,stop := context.WithCancel(context.Background())
	defer stop()
	
	if err := db.PingContext(ctx); err != nil {
	  log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	rows, err := db.QueryContext(ctx, "SELECT base_area.ID id,base_area.AREA_NAME area_name FROM base_area")
	defer rows.Close()
	
	for rows.Next() {
		var id int64
		var area_name string
		if err := rows.Scan(&id,&area_name); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s id: %d\n", area_name,id)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}