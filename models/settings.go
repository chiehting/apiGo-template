package models

// Settings is struct of table settings without foreign key
type Settings struct {
	ID          uint   `gorm:"primaryKey;column:id"`
	Class       string `gorm:"column:class"`
	Subject     string `gorm:"column:subject"`
	Value       string `gorm:"column:value"`
	Description string `gorm:"column:description"`
	DataType    string `gorm:"column:data_type"`
	DataEdit    uint8  `gorm:"column:data_edit"`
	MaxValue    uint8  `gorm:"column:max_value"`
}

var _ = Settings{}

// TableName sets the insert table name for this struct type
func (s *Settings) TableName() string {
	return "settings"
}

// GetAll is get all data of settings
func (s *Settings) GetAll() ([]Settings, error) {
	var rows []Settings
	result := _db.Model(&s).Select("*").Scan(&rows)
	return rows, result.Error
}

// Get is get setting by subject
func (s *Settings) Get() error {
	result := _db.Model(&s).Take(&s)
	return result.Error
}

// UpdateValue 更新 item_value 資料
func (s *Settings) UpdateValue() error {
	result := _db.Model(&s).Select("value").Updates(s)
	return result.Error
}

// GetCache is get user info from cache
func (s *Settings) GetCache() map[string]string {
	return _cache.HGetAll(_ctx, s.getCacheKey(0)).Val()
}

// SetCache is set data to cache
func (s *Settings) SetCache(rows []Settings) error {
	var pipeline = _cache.Pipeline()
	for _, v := range rows {
		pipeline.HSet(_ctx, v.getCacheKey(0), map[string]interface{}{v.Subject: v.Value})
	}
	var _, err = pipeline.Exec(_ctx)
	return err
}

// DelCache is to remove user cache
func (s *Settings) DelCache() error {
	var result, _, _ = _cache.Scan(_ctx, 0, s.getCacheKey(0)+"*", 0).Result()
	var pipeline = _cache.Pipeline()
	for i := range result {
		pipeline.Del(_ctx, result[i])
	}
	var _, err = pipeline.Exec(_ctx)
	return err
}

func (s *Settings) getCacheKey(cacheKey int) string {
	return _cachePrefix + [...]string{
		"settings-" + s.Class,
	}[cacheKey]
}
