// Package main is the sqlite lesson `l1_transactions_and_locking` homework scaffold for Vibe Learn.
//
// Задача: банковский перевод Transfer как атомарная BEGIN IMMEDIATE транзакция с проверкой неотрицательного баланса.
// Реализуй функции ниже — сигнатуры и тестовая поверхность фиксированы;
// CI (.github/workflows/ci.yml) гоняет `go vet` и `go test ./...`.
// Подробности и критерии приёмки — в README.md.
//
// Драйвер: modernc.org/sqlite — ЧИСТЫЙ Go, без CGO. Имя драйвера "sqlite",
// импорт blank-формой `_ "modernc.org/sqlite"`. БД встроена — сервера нет.
package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

// Note — пример доменной строки (notes/CRUD-уроки).
type Note struct {
	ID        int64
	Title     string
	Body      string
	CreatedAt string
}

// RegionRevenue — пример агрегата (join/CTE-уроки).
type RegionRevenue struct {
	Region  string
	Revenue int64
}

// Latencies — собранные перцентили для бенчмарка (WAL/concurrency-уроки).
type Latencies struct{ P50, P99 int64 }

// ----- config -----

// envOr returns the env var for `key` if set, else `fallback`.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// DatabasePath — путь к файлу БД. Дефолт ":memory:" — БД живёт в процессе.
// Для WAL-уроков переопредели на файловый путь (WAL завязан на файл).
func DatabasePath() string {
	return envOr("DATABASE_PATH", ":memory:")
}

// OpenDB открывает sqlite-БД через чистый Go-драйвер modernc и создаёт схему.
func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(Schema); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

// Schema — DDL, исполняемый идемпотентно при открытии БД (CREATE ... IF NOT EXISTS).
const Schema = `CREATE TABLE IF NOT EXISTS accounts (id INTEGER PRIMARY KEY, balance INTEGER NOT NULL)`

// ----- TODO #1: Transfer -----
//
// BEGIN IMMEDIATE, два UPDATE арифметикой в SQL, ROLLBACK с ошибкой при уходе баланса в минус
func Transfer(db *sql.DB, fromID, toID, amount int64) error {
	// TODO: implement
	panic("Transfer: not implemented")
}

// ----- TODO #2: Balance -----
//
// прочитать текущий баланс счёта
func Balance(db *sql.DB, id int64) (int64, error) {
	// TODO: implement
	panic("Balance: not implemented")
}

// ----- TODO #3: TotalBalance -----
//
// SELECT sum(balance) — инвариант не должен меняться после серии переводов
func TotalBalance(db *sql.DB) (int64, error) {
	// TODO: implement
	panic("TotalBalance: not implemented")
}

// _refs keeps the example domain types live while the TODO bodies are stubs.
// Удали эту переменную, когда реализуешь TODO выше.
var _refs = []any{
	Note{},
	RegionRevenue{},
	Latencies{},
}

// ----- main entry -----

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("Vibe Learn — sqlite lesson %s scaffold up", "l1_transactions_and_locking")
	log.Printf("DATABASE_PATH: %s (driver: modernc.org/sqlite, pure-Go)", DatabasePath())
	log.Printf("Реализуй TODO-функции, затем `go test ./...`. README.md содержит задачу.")

	db, err := OpenDB(DatabasePath())
	if err != nil {
		log.Fatalf("OpenDB failed: %v", err)
	}
	defer db.Close()
	log.Printf("схема создана, БД готова")
}
