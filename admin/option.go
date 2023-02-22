package admin

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Key   string `json:"key"`
	Value string `json:"value"`
}

func getOptions() []Option {
	var options []Option
	db.Find(&options)
	return options
}

func saveOptions(option []Option) error {
	for _, opt := range option {
		err := saveOption(&opt)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveOption(option *Option) error {
	var actualOption Option
	result := db.First(&actualOption, "key = ?", option.Key)
	if result.RowsAffected == 0 {
		db.Create(option)
		return nil
	}
	db.Model(&actualOption).Update("Value", option.Value)
	return nil
}

func GetOptionValue(key string) string {
	var option Option
	db.First(&option, "key = ?", key)
	return option.Value
}
