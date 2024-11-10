package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/mksmstpck/telegram-miniapp/server/internal/config"
	"github.com/mksmstpck/telegram-miniapp/server/internal/db"
	"github.com/mksmstpck/telegram-miniapp/server/internal/handlers"
	"github.com/mksmstpck/telegram-miniapp/server/internal/services"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.Function), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func main() {
	config, err := config.LoadConfig()

	conn, err := pgx.Connect(context.Background(), config.PsqlDns)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	db := db.NewDB(conn)

	services := services.NewServices(*db)

	handlers.NewHandlers(config, services).HandleAll()
}
