package components

type ADDR struct {
	LEN_HD        uint16 `json:"len_hd"`
	SPEED_HD      uint16 `json:"speed_hd"`
	POS_HD        uint16 `json:"pos_hd"`
	RESOLUTION_HD uint16 `json:"resolution_hd"`
	MOVE_M        uint16 `json:"move_m"`
	JOG_M         uint16 `json:"jog_m"`
	STOP_M        uint16 `json:"stop_m"`
}

type Motor struct {
	Status     string  `json:"status"` // 状态: Moving, Stopped
	Position   float32 `json:"position"`
	Speed      float32 `json:"speed"`
	Resolution float32 `json:"resolution"`

	ADDR ADDR `json:"addr"`
}

// HD 地址
const (
	//X轴
	X_LEN_HD        uint16 = 0
	X_SPEED_HD      uint16 = 2
	X_POS_HD        uint16 = 4
	X_RESOLUTION_HD uint16 = 6
	//Y轴
	Y_LEN_HD        uint16 = 10
	Y_SPEED_HD      uint16 = 12
	Y_POS_HD        uint16 = 14
	Y_RESOLUTION_HD uint16 = 16
	//Z轴
	Z_LEN_HD        uint16 = 20
	Z_SPEED_HD      uint16 = 22
	Z_POS_HD        uint16 = 24
	Z_RESOLUTION_HD uint16 = 26
	//R轴
	R_LEN_HD        uint16 = 30
	R_SPEED_HD      uint16 = 32
	R_POS_HD        uint16 = 34
	R_RESOLUTION_HD uint16 = 36
	//XX轴
	XX_LEN_HD        uint16 = 40
	XX_SPEED_HD      uint16 = 42
	XX_POS_HD        uint16 = 44
	XX_RESOLUTION_HD uint16 = 46
)

// M 地址
const (
	//X轴
	X_MOVE_M uint16 = 0
	X_JOG_M  uint16 = 1
	X_STOP_M uint16 = 2
	//Y轴
	Y_MOVE_M uint16 = 10
	Y_JOG_M  uint16 = 11
	Y_STOP_M uint16 = 12
	//Z轴
	Z_MOVE_M uint16 = 20
	Z_JOG_M  uint16 = 21
	Z_STOP_M uint16 = 22
	//R轴
	R_MOVE_M uint16 = 30
	R_JOG_M  uint16 = 31
	R_STOP_M uint16 = 32
	//XX轴
	XX_MOVE_M uint16 = 40
	XX_JOG_M  uint16 = 41
	XX_STOP_M uint16 = 42
)
