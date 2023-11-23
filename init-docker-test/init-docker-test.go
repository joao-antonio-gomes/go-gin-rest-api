package init_docker_test

import (
	"fmt"
	"github.com/joao-antonio-gomes/go-gin-rest-api/database/seeds"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitTestDocker(exposedPort string) (*dockertest.Pool, *dockertest.Resource) {
	var passwordEnv = "POSTGRES_PASSWORD=postgres"
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"listen_addresses = '*'",
			fmt.Sprint(passwordEnv),
		},
		ExposedPorts: []string{exposedPort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {
				{HostIP: "0.0.0.0", HostPort: exposedPort},
			},
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"} // Important option when container crash, and you want to debug container
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := resource.Expire(30); err != nil { // Tell docker to hard kill the container in 30 seconds
		logrus.Error(err)
	}

	// retry if container is not ready
	pool.MaxWait = 30 * time.Second
	if err = pool.Retry(func() error {
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return pool, resource
}

func OpenDatabaseConnection(pool *dockertest.Pool, resource *dockertest.Resource, exposedPort string) *gorm.DB {
	user := "postgres"
	password := "postgres"
	db := "postgres"
	port := "5432"
	dns := "host=%s port=%s user=%s sslmode=disable password=%s dbname=%s"

	retries := 5
	host := resource.GetBoundIP(fmt.Sprintf("%s/tcp", port))
	gdb, err := gorm.Open(postgres.Open(fmt.Sprintf(dns, host, exposedPort, user, password, db)), &gorm.Config{})

	// Sometimes it happens that after first time container is not ready.
	// It's always better to create retry if necessary and be sure that tests run without problems
	for err != nil {
		if retries > 1 {
			retries--
			time.Sleep(1 * time.Second)
			gdb, err = gorm.Open(postgres.Open(fmt.Sprintf(dns, host, exposedPort, user, password, db)), &gorm.Config{})
			continue
		}

		if err := pool.Purge(resource); err != nil {
			logrus.Error(err)
		}

		log.Panic("Fatal error in connection: ", err, resource.GetBoundIP("5432/tcp"))
	}

	return gdb
}

func SeedDatabase(gdb *gorm.DB, pool *dockertest.Pool, resource *dockertest.Resource) {
	for _, seed := range seeds.All() {
		if err := seed.Run(gdb); err != nil {
			if err := pool.Purge(resource); err != nil {
				logrus.Error(err)
			}

			log.Fatalf("Running seeds '%s', failed with error: %s", seed.Name, err)
		}
	}
}
