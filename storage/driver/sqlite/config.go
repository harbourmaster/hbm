package sqlite

// AddConfig function
func (c *Config) AddConfig(name string) {
	c.DB.Create(&AppConfig{Key: name})
}

// RemoveConfig function
func (c *Config) RemoveConfig(name string) error {
	c.DB.Where("key = ?", name).Delete(AppConfig{})

	return nil
}

// ListConfigs function
func (c *Config) ListConfigs() []string {
	var configs []AppConfig

	result := []string{}

	c.DB.Find(&configs)

	for _, config := range configs {
		result = append(result, config.Key)
	}

	return result
}

// FindConfig function
func (c *Config) FindConfig(name string) bool {
	var count int64

	c.DB.Model(&AppConfig{}).Where("key = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}
