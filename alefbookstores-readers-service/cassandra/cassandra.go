package cassandra

import (
	"context"
	"crypto/tls"
	"github.com/joho/godotenv"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	datastax "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"sync"
)

var datastaxConnectionPool = sync.Pool{
	New: func() interface{} {

		datastaxRemoteUri := os.Getenv("DATASTAX_REMOTE_URI")
		_ = os.Getenv("DATASTAX_REMOTE_CLIENT")
		_ = os.Getenv("DATASTAX_REMOTE_CLIENT_SECRET")
		datastaxBearerToken := os.Getenv("DATASTAX_BEARER_TOKEN")

		config := &tls.Config{InsecureSkipVerify: false}
		conn, err := grpc.Dial(
			datastaxRemoteUri,
			grpc.WithTransportCredentials(credentials.NewTLS(config)),
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithPerRPCCredentials(auth.NewStaticTokenProvider(datastaxBearerToken)),
		)
		if err != nil {
			log.Fatalf("failed to connect to remote cassandra instance by datastax reason %s\n", err)
		}

		stargateClient, err := client.NewStargateClientWithConn(conn)
		if err != nil {
			log.Fatalf("failed to create stargate client %s\n", err)
		}

		return stargateClient

	},
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to parse environment variable %s", err)
	}

	stargateClient1 := datastaxConnectionPool.Get().(*client.StargateClient)
	_, err = stargateClient1.ExecuteQuery(&datastax.Query{
		Cql: `CREATE TABLE IF NOT EXISTS readers.users 
              (email text, user_id text, password text ,first_name text, last_name text, PRIMARY KEY ((email)));`,
	})
	if err != nil {
		log.Fatalf("failed to execute cql table creation statement %s\n", err)
	}
}

type ReadersRepository interface {
	GetUserByEmail(context context.Context, email string) (User, error)
	getAllUsers(context.Context) ([]User, error)
	registerUser(context context.Context, user User) (User, error)
	updateUser(context context.Context, user User) error
	deleteUser(context context.Context, user User) error
}

type ReadersCassandraRepository struct {
}

func (r *ReadersCassandraRepository) GetUserByEmail(context context.Context, email string) (User, error) {
	c := datastaxConnectionPool.Get().(*client.StargateClient)

	res, err := c.ExecuteQuery(&datastax.Query{
		Cql: `SELECT * FROM readers.users WHERE email = ?`,
		Values: &datastax.Values{
			Values: []*datastax.Value{
				{
					Inner: &datastax.Value_String_{String_: email},
				},
			},
		},
		Parameters: nil,
	})
	if err != nil {
		log.Printf("failed to execute query and fetch the user given the email %s\n", err)
		return User{}, err
	}

	resultSet := res.GetResultSet()
	var foundUser User = User{}
	for _, user := range resultSet.Rows {
		foundUser = User{
			Email:     user.Values[0].GetString_(),
			FirstName: user.Values[1].GetString_(),
			LastName:  user.Values[2].GetString_(),
			Password:  user.Values[3].GetString_(),
			UserId:    user.Values[4].GetString_(),
		}
	}
	return foundUser, nil
}

func (r *ReadersCassandraRepository) getAllUsers(ctx context.Context) ([]User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReadersCassandraRepository) registerUser(context context.Context, user User) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReadersCassandraRepository) updateUser(context context.Context, user User) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (r *ReadersCassandraRepository) deleteUser(context context.Context, user User) error {
	//TODO implement me
	panic("implement me")
	return nil
}
