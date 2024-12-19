package storage

import (
	"context"
	"testing"

	"github.com/onsi/gomega"
	"github.com/oxelf/pixel-party/internal/canvas"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartRedisContainer() (string, testcontainers.Container, error) {
	ctx := context.Background()

	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "redis:latest",
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor:   wait.ForLog("Ready to accept connections"),
		},
		Started: true,
	})
	if err != nil {
		return "", nil, err
	}

	port, err := redisContainer.MappedPort(ctx, "6379")
	if err != nil {
		return "", nil, err
	}

	return port.Port(), redisContainer, nil
}

func TestCanvas(t *testing.T) {
	g := gomega.NewWithT(t)

	port, _, err := StartRedisContainer()

	g.Expect(err).ToNot(gomega.HaveOccurred())

	var address string
	address = "localhost:" + port

	client := StartRedis(&address)

	testPixel := []canvas.Pixel{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
	}

	err = client.SetCanvas("0", testPixel)
	g.Expect(err).To(gomega.BeNil())

	pixel, err := client.GetCanvasAsPixel("0")
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(pixel).To(gomega.Equal(testPixel))

	testPixel[1] = canvas.Pixel{R: 15, G: 15, B: 10, A: 15}

	err = client.SetPixel("0", testPixel[1], 1)
	g.Expect(err).ToNot(gomega.HaveOccurred())

	pixel, err = client.GetCanvasAsPixel("0")
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(pixel).To(gomega.Equal(testPixel))
}

func TestMain(m *testing.M) {
	m.Run()
}
