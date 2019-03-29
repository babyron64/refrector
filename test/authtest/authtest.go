package authtest

import (
	"github.com/kanatatsu64/refrector/mongodb"
	"github.com/kanatatsu64/refrector/server"
)

func main() {
	mongodb.InitMongoDB()
	AuthTest()
}

// AuthTest is an entry point of the authorization test.
func AuthTest() {
	InitAuthTestServer(server.GetPort())
}
