package provider

import (
	"log"

	"github.com/burhanwakhid/shopifyx_backend/config"
	"github.com/burhanwakhid/shopifyx_backend/internal"
	"github.com/burhanwakhid/shopifyx_backend/pkg/cache"
	"github.com/burhanwakhid/shopifyx_backend/pkg/orm"
	"github.com/burhanwakhid/shopifyx_backend/pkg/translator"
)

func ProvideErrorTranslator(
	translatorCfg config.TranslatorConfig,
) internal.ErrorTranslator {
	return translator.NewError(translatorCfg.ErrorFilePath)
}

func ProvideDBReplication(
	dbConf *config.DBConfig,

) *orm.Replication {
	return ProvideDBReplicationWithTracer(&dbConf.PSQLConfig)
}

func ProvideDBReplicationWithTracer(
	dbConf *config.PSQLConfig,

) *orm.Replication {
	opt := orm.Option{

		MaxLifeTime:  dbConf.MaxLifeTime,
		MaxOpenConns: dbConf.MaxConnection,
		LogLevel:     dbConf.LogLevel,
	}
	opt.Database = dbConf.Master

	master, err := orm.Open(opt)
	if err != nil {
		log.Panicln("unable to init postgres master: s ", err)
	}

	opt.Database = dbConf.Slave

	slave, err := orm.Open(opt)
	if err != nil {
		log.Panicln("unable to init postgres slave", err)
	}

	return orm.NewReplication(master, slave)
}

func ProvideLocalCache(
	dbConfig *config.DBConfig,
) cache.Local {
	return *cache.NewLocal(
		cache.LocalOption{
			DefaultExpiration: dbConfig.LocalCache.DefaultExpiration,
			CleanupInterval:   dbConfig.LocalCache.CleanupInterval,
		},
	)
}
