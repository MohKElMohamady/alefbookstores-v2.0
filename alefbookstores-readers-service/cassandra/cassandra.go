package cassandra

import (
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

type ReaderRepository interface {
	getUserByEmail(email string) User
	getAllUsers() []User
	registerUser(user User) User
	updateUser(user User)
	deleteUser(user User)
}

type ReaderCassandraRepository struct {
}

func (r *ReaderCassandraRepository) getUserByEmail(email string) User {
	//TODO implement me
	panic("implement me")
}

func (r *ReaderCassandraRepository) getAllUsers() []User {
	//TODO implement me
	panic("implement me")
}

func (r *ReaderCassandraRepository) registerUser(user User) User {
	//TODO implement me
	panic("implement me")
}

func (r *ReaderCassandraRepository) updateUser(user User) {
	//TODO implement me
	panic("implement me")
}

func (r *ReaderCassandraRepository) deleteUser(user User) {
	//TODO implement me
	panic("implement me")
}
