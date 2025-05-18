package sqlc_gen

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/looksaw/exchange_go/config"
	"github.com/looksaw/exchange_go/util"
	"github.com/stretchr/testify/require"
)

const (
	Uri = "postgresql://postgres@localhost:5432/exchange_go?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config := config.InitConfig()
	config.ExchangeConfig.DbConfig.Uri = Uri
	conn, err := sql.Open(config.ExchangeConfig.DbConfig.Driver,
		config.ExchangeConfig.DbConfig.Uri)
	if err != nil {
		log.Fatal("Can't connect database ........")
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		UserName:     util.GenRandomName(26),
		UserPassword: util.GenRandomNumber(50),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Name, arg.UserName)
	require.Equal(t, user.Password, arg.UserPassword)
}

func TestDeleteUserByName(t *testing.T) {

}

func TestDeleteUserByID(t *testing.T) {

}

func TestGetUserByID(t *testing.T) {

}

func TestGetUserByName(t *testing.T) {
	user, err := testQueries.GetUserByName(context.Background(), "looksaw")
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Name, "looksaw")
	require.Equal(t, user.Password, "123456")
}

func TestGetUserList(t *testing.T) {

}

func GetUserWithRange(t *testing.T) {

}

func UpdateUserById(t *testing.T) {

}

func UpdateUserByName(t *testing.T) {

}
