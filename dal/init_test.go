package dal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"testing"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	os.Remove("../test_cases.db")

	testDB = ConnectTestDB()

	// Populate data required to execute test cases
	err := generateDataRequiredToRunTestCases(testDB)
	if err != nil {
		log.Fatal(err)
	}

	// Run test cases
	exitVal := m.Run()
	os.Exit(exitVal)
}

func ConnectTestDB() *sql.DB {
	db, err := sql.Open("sqlite3", "../test_cases.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func generateDataRequiredToRunTestCases(testDB *sql.DB) error {
	_, err := testDB.Exec("CREATE TABLE IF NOT EXISTS deck (id text NOT NULL PRIMARY KEY, is_shuffled BOOLEAN, cards json);")
	if err != nil {
		return err
	}

	_, err = testDB.Exec("INSERT INTO deck (id, is_shuffled, cards) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;", "6818b769-c508-43de-910a-4c0a63ad8dad", false, `[{"value":"KING","suit":"CLUBS","code":"KC"},{"value":"ACE","suit":"CLUBS","code":"AC"},{"value":"5","suit":"HEARTS","code":"5H"},{"value":"KING","suit":"HEARTS","code":"KH"},{"value":"ACE","suit":"DIAMONDS","code":"AD"},{"value":"10","suit":"DIAMONDS","code":"1D"},{"value":"8","suit":"HEARTS","code":"8H"},{"value":"QUEEN","suit":"CLUBS","code":"QC"},{"value":"6","suit":"DIAMONDS","code":"6D"},{"value":"9","suit":"CLUBS","code":"9C"},{"value":"QUEEN","suit":"DIAMONDS","code":"QD"},{"value":"3","suit":"SPADES","code":"3S"},{"value":"3","suit":"CLUBS","code":"3C"},{"value":"3","suit":"DIAMONDS","code":"3D"},{"value":"2","suit":"HEARTS","code":"2H"},{"value":"10","suit":"SPADES","code":"1S"},{"value":"2","suit":"SPADES","code":"2S"},{"value":"4","suit":"CLUBS","code":"4C"},{"value":"3","suit":"HEARTS","code":"3H"},{"value":"7","suit":"HEARTS","code":"7H"},{"value":"ACE","suit":"SPADES","code":"AS"},{"value":"5","suit":"CLUBS","code":"5C"},{"value":"4","suit":"DIAMONDS","code":"4D"},{"value":"QUEEN","suit":"HEARTS","code":"QH"},{"value":"6","suit":"SPADES","code":"6S"},{"value":"7","suit":"DIAMONDS","code":"7D"},{"value":"2","suit":"CLUBS","code":"2C"},{"value":"ACE","suit":"HEARTS","code":"AH"},{"value":"KING","suit":"DIAMONDS","code":"KD"},{"value":"JACK","suit":"CLUBS","code":"JC"},{"value":"QUEEN","suit":"SPADES","code":"QS"},{"value":"7","suit":"SPADES","code":"7S"},{"value":"6","suit":"HEARTS","code":"6H"},{"value":"JACK","suit":"DIAMONDS","code":"JD"},{"value":"10","suit":"CLUBS","code":"1C"},{"value":"JACK","suit":"SPADES","code":"JS"},{"value":"5","suit":"DIAMONDS","code":"5D"},{"value":"4","suit":"HEARTS","code":"4H"},{"value":"2","suit":"DIAMONDS","code":"2D"},{"value":"9","suit":"SPADES","code":"9S"},{"value":"JACK","suit":"HEARTS","code":"JH"},{"value":"10","suit":"HEARTS","code":"1H"},{"value":"KING","suit":"SPADES","code":"KS"},{"value":"5","suit":"SPADES","code":"5S"},{"value":"8","suit":"SPADES","code":"8S"},{"value":"4","suit":"SPADES","code":"4S"},{"value":"7","suit":"CLUBS","code":"7C"},{"value":"8","suit":"DIAMONDS","code":"8D"},{"value":"9","suit":"DIAMONDS","code":"9D"},{"value":"8","suit":"CLUBS","code":"8C"},{"value":"9","suit":"HEARTS","code":"9H"},{"value":"6","suit":"CLUBS","code":"6C"}]`)
	if err != nil {
		return err
	}

	_, err = testDB.Exec("INSERT INTO deck (id, is_shuffled, cards) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;", "2e904357-ae8e-401b-9071-0a0a3f4b5770", false, `[{"value":"ACE","suit":"SPADES","code":"AS"},{"value":"2","suit":"SPADES","code":"2S"},{"value":"3","suit":"SPADES","code":"3S"},{"value":"4","suit":"SPADES","code":"4S"},{"value":"5","suit":"SPADES","code":"5S"},{"value":"6","suit":"SPADES","code":"6S"},{"value":"7","suit":"SPADES","code":"7S"},{"value":"8","suit":"SPADES","code":"8S"},{"value":"9","suit":"SPADES","code":"9S"},{"value":"10","suit":"SPADES","code":"1S"},{"value":"JACK","suit":"SPADES","code":"JS"},{"value":"QUEEN","suit":"SPADES","code":"QS"},{"value":"KING","suit":"SPADES","code":"KS"},{"value":"ACE","suit":"DIAMONDS","code":"AD"},{"value":"2","suit":"DIAMONDS","code":"2D"},{"value":"3","suit":"DIAMONDS","code":"3D"},{"value":"4","suit":"DIAMONDS","code":"4D"},{"value":"5","suit":"DIAMONDS","code":"5D"},{"value":"6","suit":"DIAMONDS","code":"6D"},{"value":"7","suit":"DIAMONDS","code":"7D"},{"value":"8","suit":"DIAMONDS","code":"8D"},{"value":"9","suit":"DIAMONDS","code":"9D"},{"value":"10","suit":"DIAMONDS","code":"1D"},{"value":"JACK","suit":"DIAMONDS","code":"JD"},{"value":"QUEEN","suit":"DIAMONDS","code":"QD"},{"value":"KING","suit":"DIAMONDS","code":"KD"},{"value":"ACE","suit":"CLUBS","code":"AC"},{"value":"2","suit":"CLUBS","code":"2C"},{"value":"3","suit":"CLUBS","code":"3C"},{"value":"4","suit":"CLUBS","code":"4C"},{"value":"5","suit":"CLUBS","code":"5C"},{"value":"6","suit":"CLUBS","code":"6C"},{"value":"7","suit":"CLUBS","code":"7C"},{"value":"8","suit":"CLUBS","code":"8C"},{"value":"9","suit":"CLUBS","code":"9C"},{"value":"10","suit":"CLUBS","code":"1C"},{"value":"JACK","suit":"CLUBS","code":"JC"},{"value":"QUEEN","suit":"CLUBS","code":"QC"},{"value":"KING","suit":"CLUBS","code":"KC"},{"value":"ACE","suit":"HEARTS","code":"AH"},{"value":"2","suit":"HEARTS","code":"2H"},{"value":"3","suit":"HEARTS","code":"3H"},{"value":"4","suit":"HEARTS","code":"4H"},{"value":"5","suit":"HEARTS","code":"5H"},{"value":"6","suit":"HEARTS","code":"6H"},{"value":"7","suit":"HEARTS","code":"7H"},{"value":"8","suit":"HEARTS","code":"8H"},{"value":"9","suit":"HEARTS","code":"9H"},{"value":"10","suit":"HEARTS","code":"1H"},{"value":"JACK","suit":"HEARTS","code":"JH"},{"value":"QUEEN","suit":"HEARTS","code":"QH"},{"value":"KING","suit":"HEARTS","code":"KH"}]`)
	if err != nil {
		return err
	}
	return nil
}
