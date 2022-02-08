package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"htp-platform/app/htpp/interface/internal/conf"
	"htp-platform/app/htpp/interface/internal/data/ent"
	"htp-platform/app/htpp/interface/internal/data/ent/migrate"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "interface-service/data"))

	d := &Data{
		db: entClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			logHelper.Error(err)
		}
	}, nil
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
