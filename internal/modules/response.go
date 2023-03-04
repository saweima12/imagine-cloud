package modules

type LoginResponse struct {
	Token string `json:"token"`
}

type DashBoardResponse struct {
	Disk *DiskStatus `json:"disk"`
}
