package core

type AppSetting struct {
	ID                         string `gorm:"primaryKey;column:id;size:450"`
	Value                      string `gorm:"column:value"`
	Module                     string `gorm:"column:module"`
	IsVisibleInCommonSettingPage bool   `gorm:"column:is_visible_in_common_setting_page;not null"`
}

func (AppSetting) TableName() string {
	return "public.app_settings"
}
