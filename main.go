package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	promremote "github.com/castai/promwrite"
)

var (
	prometheusRemoteHost     string
	prometheusRemoteUsername string
	prometheusRemotePassword string
)

func main() {
	prometheusRemoteHost = os.Getenv("INPUT_PROMETHEUS_REMOTE_WRITE_URL")
	prometheusRemoteUsername = os.Getenv("INPUT_PROMETHEUS_USERNAME")
	prometheusRemotePassword = os.Getenv("INPUT_PROMETHEUS_PASSWORD")

	fmt.Println("url: %s", prometheusRemoteHost)
	fmt.Println("username: %s", prometheusRemoteUsername)
	fmt.Println("pwd: %s", prometheusRemotePassword)

	if prometheusRemoteHost == "" || prometheusRemotePassword == "" || prometheusRemoteUsername == "" {
		log.Fatal("Invalid options")
	}

	ctx := context.Background()
	cli := promremote.NewClient(prometheusRemoteHost)
	headers := make(map[string]string)

	userPass := prometheusRemoteUsername + ":" + prometheusRemotePassword
	encodedPass := base64.StdEncoding.EncodeToString([]byte(userPass))
	headers["Authorization"] = fmt.Sprintf("Basic %s", encodedPass)

	authorizationHeader := promremote.WriteHeaders(headers)

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64() * 100
	fmt.Println("random: %d", randomNumber)

	metric := promremote.WriteRequest{
		TimeSeries: []promremote.TimeSeries{
			{
				Labels: []promremote.Label{
					{
						Name:  "__name__",
						Value: "ghaction_metric",
					},
				},
				Sample: promremote.Sample{
					Time:  time.Now(),
					Value: randomNumber,
				},
			},
		},
	}

	_, err := cli.Write(ctx, &metric, authorizationHeader)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

}
