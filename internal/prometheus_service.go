package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	promremote "github.com/castai/promwrite"
)

func WriteToPrometheus(prometheusRemoteHost string, prometheusRemoteUsername string, prometheusRemotePassword string) error {
	ctx := context.Background()
	cli := promremote.NewClient(prometheusRemoteHost)
	headers := make(map[string]string)

	userPass := prometheusRemoteUsername + ":" + prometheusRemotePassword
	encodedPass := base64.StdEncoding.EncodeToString([]byte(userPass))
	headers["Authorization"] = fmt.Sprintf("Basic %s", encodedPass)

	authorizationHeader := promremote.WriteHeaders(headers)

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64() * 100

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
	return err
}
