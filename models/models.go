package models

import (
	"context"

	"github.com/chiehting/apiGo-template/pkg/config"
	"github.com/chiehting/apiGo-template/pkg/storage"
)

var _db = storage.MySQL
var (
	_cache = storage.Redis
	_ctx   = context.Background()
)
var _cfg = config.GetApplication()
var _cachePrefix = _cfg.Name + "-"
