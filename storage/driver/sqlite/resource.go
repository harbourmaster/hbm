package sqlite

import (
	"fmt"

	"github.com/kassisol/hbm/storage/driver"
)

// AddResource function
func (c *Config) AddResource(name, rtype, value, options string) {
	c.DB.Create(&Resource{
		Name:   name,
		Type:   rtype,
		Value:  value,
		Option: options,
	})
}

// RemoveResource function
func (c *Config) RemoveResource(name string) error {
	if c.memberOfCollection(name) {
		return fmt.Errorf("resource \"%s\" cannot be removed. It is being used by a collection", name)
	}

	c.DB.Where("name = ?", name).Delete(Resource{})

	return nil
}

// ListResources function
func (c *Config) ListResources(filter map[string]string) map[driver.ResourceResult][]string {
	result := make(map[driver.ResourceResult][]string)

	sql := c.DB.Table("resources").Select("resources.name, resources.type, resources.value, resources.option, collections.name").Joins("LEFT JOIN collection_resources ON collection_resources.resource_id = resources.id").Joins("LEFT JOIN collections ON collections.id = collection_resources.collection_id")

	if v, ok := filter["name"]; ok {
		sql = sql.Where("resources.name = ?", v)
	}

	if v, ok := filter["type"]; ok {
		sql = sql.Where("resources.type = ?", v)
	}

	if v, ok := filter["value"]; ok {
		sql = sql.Where("resources.value = ?", v)
	}

	if v, ok := filter["elem"]; ok {
		sql = sql.Where("collections.name = ?", v)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	for rows.Next() {
		var res_name string
		var res_type string
		var res_value string
		var res_option string
		var collection string

		rows.Scan(&res_name, &res_type, &res_value, &res_option, &collection)

		rr := driver.ResourceResult{Name: res_name, Type: res_type, Value: res_value, Option: res_option}

		result[rr] = append(result[rr], collection)
	}

	return result
}

// FindResource function
func (c *Config) FindResource(name string) bool {
	var count int64

	c.DB.Model(&Resource{}).Where("name = ?", name).Count(&count)

	if count == 1 {
		return true
	}

	return false
}

// CountResource function
func (c *Config) CountResource(rtype string) int {
	var count int64

	if rtype == "all" {
		c.DB.Model(&Resource{}).Count(&count)
	} else {
		c.DB.Model(&Resource{}).Where("type = ?", rtype).Count(&count)
	}

	return int(count)
}

// AddResourceToCollection function
func (c *Config) AddResourceToCollection(collection, resource string) {
	col := Collection{}
	res := Resource{}

	c.DB.Where("name = ?", resource).Find(&res)
	c.DB.Where("name = ?", collection).Find(&col)

	c.DB.Model(&col).Association("Resources").Append(&res)
}

// RemoveResourceFromCollection function
func (c *Config) RemoveResourceFromCollection(collection, resource string) {
	col := Collection{}
	res := Resource{}

	c.DB.Where("name = ?", resource).Find(&res)
	c.DB.Where("name = ?", collection).Find(&col)

	c.DB.Model(&col).Association("Resources").Delete(&res)
}

func (c *Config) memberOfCollection(name string) bool {
	var count int64

	c.DB.Table("collections").Joins("JOIN collection_resources ON collection_resources.collection_id = collections.id").Joins("JOIN resources ON resources.id = collection_resources.resource_id").Where("resources.name = ?", name).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
