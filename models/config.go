package models

// Config holds the go model for the config json file
type Config struct {
	Version       string `json:"VERSION"`
	ShareServer   string `json:"SHARE_SERVER"`
	ShareFolder   string `json:"SHARE_FOLDER"`
	ListingServer string `json:"LISTING_SERVER"`
	OsxMountPoint string `json:"OSX_MOUNT"`
	ServerPort    int    `json:"SERVER_PORT"`
}
