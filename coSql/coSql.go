package coSql

import (
	"GoMap/loger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db  = &sql.DB{}
	err error
)

//
//
// Connection / Deconnection

func coPg() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		loger.Crash("Crash à l'ouverture de Postgres", err)
	}
}

func closePg() {
	err = db.Close()
	if err != nil {
		loger.Crash("Crash à la fermeture de Postgres", err)
	}
}

//
//
// Requetes

func GetRefcodeData(refcode3 string) Refcode {

	req := fmt.Sprintf(`
SELECT rc_nro, rc_sro, rc_refcode1, rc_refcode2
FROM data.l_refcode 
WHERE rc_refcode3 = '%s'`,
		refcode3)

	coPg()

	data, _ := db.Query(req)
	rc := Refcode{}

	for data.Next() {

		err = data.Scan(&rc.NRO, &rc.SRO, &rc.RefCode1, &rc.RefCode2)
		if err != nil {
			loger.Crash("Crash lors de la lecture des données", err)
		}
	}

	closePg()
	return rc
}
