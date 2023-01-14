package database

import (
	"database/sql"
	"testing"

	"github.com/MuriloAbranches/fc-ms-walletcore/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255),  amount int, created_at date)")
	s.client1, _ = entity.NewClient("John Doe", "j@j.com")
	s.client2, _ = entity.NewClient("John Doe 2", "j@j2.com")

	s.accountFrom = entity.NewAccount(s.client1)
	s.accountFrom.Balance = 1000.0

	s.accountTo = entity.NewAccount(s.client2)
	s.accountTo.Balance = 1000.0

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
