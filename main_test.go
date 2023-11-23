package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/infrastructure/rest"
	"github.com/joao-antonio-gomes/go-gin-rest-api/database"
	init_docker_test "github.com/joao-antonio-gomes/go-gin-rest-api/init-docker-test"
	. "github.com/onsi/ginkgo/v2"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
)

func SetupTestsRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

var _ = BeforeSuite(func() {
	// Init container, open connection, run migrations seed database, init repository
	rand.Seed(time.Now().UnixNano())
	exposedPort := fmt.Sprint(rand.Intn(10000))
	pool, resource = init_docker_test.InitTestDocker(exposedPort)
	gdb := init_docker_test.OpenDatabaseConnection(pool, resource, exposedPort)
	database.RunMigrations(gdb)
	init_docker_test.SeedDatabase(gdb, pool, resource)
})

func TestVerifyStatusCode(t *testing.T) {
	r := SetupTestsRoutes()
	r.GET("api/students", rest.ShowAllStudents)

	request, _ := http.NewRequest("GET", "/api/students", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")
}

var _ = AfterSuite(func() {
	// Purge function destroys container
	if err := pool.Purge(resource); err != nil {
		logrus.Error(err)
	}
})
