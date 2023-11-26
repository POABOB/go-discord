package svc

import (
	"fmt"
	"os"

	"github.com/POABOB/go-discord/apps/profile/rpc/internal/config"
	"github.com/POABOB/go-discord/prisma/db"
)

type ServiceContext struct {
	Config       config.Config
	PrismaClient *db.PrismaClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 Prisma client
	os.Setenv("DATABASE_URL", c.DatabaseUrl)
	client := db.NewClient(db.WithDatasourceURL(c.DatabaseUrl))
	if err := client.Prisma.Connect(); err != nil {
		fmt.Printf("%v\n", err)
	}
	return &ServiceContext{
		Config:       c,
		PrismaClient: client,
	}
}
