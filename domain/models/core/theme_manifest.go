package core

type ThemeManifest struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	DisplayName string `json:"display_name"`
	Version     string `json:"version"`
}
