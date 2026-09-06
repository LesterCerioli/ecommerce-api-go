package core

type AppSettingDTO struct {
	ID                           string `json:"id"`
	Value                        string `json:"value"`
	Module                       string `json:"module"`
	IsVisibleInCommonSettingPage bool   `json:"is_visible_in_common_setting_page"`
}
