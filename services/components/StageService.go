package components

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	XinJie "github.com/Mr-tang0/PIMSGoMod/protocol/XinJie"
)

const debug = true

type StageService struct {
	motors    map[string]Motor
	plc       XinJie.XinjieClient
	ctx       context.Context
	connected bool
}

func NewStageService() *StageService {
	return &StageService{
		plc: *XinJie.NewXinjieClient(""),
		motors: map[string]Motor{
			"Y": Motor{
				Status: "Stopped",
				ADDR: ADDR{
					LEN_HD: 0,
					POS_HD: 4,

					SPEED_HD:      X_SPEED_HD,
					RESOLUTION_HD: X_RESOLUTION_HD,

					MOVE_M: X_MOVE_M,
					JOG_M:  X_JOG_M,
					STOP_M: X_STOP_M,
				},
			},
			"R": Motor{
				Status: "Stopped",
				ADDR: ADDR{
					LEN_HD: 10,
					POS_HD: 14,

					SPEED_HD:      Y_SPEED_HD,
					RESOLUTION_HD: Y_RESOLUTION_HD,

					MOVE_M: Y_MOVE_M,
					JOG_M:  Y_JOG_M,
					STOP_M: Y_STOP_M,
				},
			},
			"Z": Motor{
				Status: "Stopped",
				ADDR: ADDR{
					LEN_HD: 20,
					POS_HD: 24,

					SPEED_HD:      Z_SPEED_HD,
					RESOLUTION_HD: Z_RESOLUTION_HD,

					MOVE_M: Z_MOVE_M,
					JOG_M:  Z_JOG_M,
					STOP_M: Z_STOP_M,
				},
			},
			"X": Motor{
				Status: "Stopped",
				ADDR: ADDR{
					LEN_HD: 30,
					POS_HD: 34,

					SPEED_HD:      R_SPEED_HD,
					RESOLUTION_HD: R_RESOLUTION_HD,

					MOVE_M: R_MOVE_M,
					JOG_M:  R_JOG_M,
					STOP_M: R_STOP_M,
				},
			},
			"XX": Motor{
				Status: "Stopped",
				ADDR: ADDR{
					LEN_HD: 40,
					POS_HD: 44,

					SPEED_HD:      XX_SPEED_HD,
					RESOLUTION_HD: XX_RESOLUTION_HD,

					MOVE_M: XX_MOVE_M,
					JOG_M:  XX_JOG_M,
					STOP_M: XX_STOP_M,
				},
			},
		},
	}
}

func (this *StageService) Startup(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx is nil")
	}
	this.ctx = ctx
	return nil
}

// StagesConnect 连接 diffract 设备
func (this *StageService) StagesConnect(ip string) error {
	if debug {
		fmt.Println("StagesConnect", ip)
	}
	if ip == "" {
		return errors.New("ip is empty")
	}
	err := this.plc.OpenTCP(ip, 1)
	if err != nil {
		return err
	}
	this.connected = true
	// 发送连接成功事件
	runtime.EventsEmit(this.ctx, "stage_running", true)

	go this.heartbeat()
	return nil
}

// StagesDisconnect 断开 diffract 设备连接
func (this *StageService) StagesDisconnect() error {
	if debug {
		fmt.Println("StagesDisconnect")
	}
	this.plc.Close()
	this.connected = false
	// 发送连接成功事件
	runtime.EventsEmit(this.ctx, "stage_running", false)
	return nil
}

func (this *StageService) heartbeat() {
	for {
		if !this.connected {
			return
		}

		for motor := range this.motors {
			m := this.motors[motor]
			val, err := this.plc.ReadRegister(XinJie.HD(m.ADDR.POS_HD), XinJie.Float32, false)
			if err != nil {
				fmt.Println("ReadRegister error:", err)
				continue
			}
			pos, ok := val.(float32)
			if !ok {
				fmt.Println("ReadRegister type error:", motor, val)
				continue
			}
			if pos == m.Position {
				m.Status = "Stopped"
			} else {
				m.Status = "Moving"
			}
			m.Position = pos
			this.motors[motor] = m
		}

		runtime.EventsEmit(this.ctx, "motor_heartbeat", this.motors)

		door, err := this.plc.ReadCoil(XinJie.X(43), true)
		if err != nil {
			fmt.Println("ReadCoil error:", err)
			continue
		}
		// fmt.Println("door:", door)

		runtime.EventsEmit(this.ctx, "door", door)

		time.Sleep(100 * time.Millisecond)

	}
}

func (this *StageService) StageStop(Axis string) error {
	if debug {
		fmt.Println("StageStop", Axis)
	}
	return this.plc.WriteCoil(XinJie.M(this.motors[Axis].ADDR.STOP_M), true)
}

func (this *StageService) StageRelMove(Axis string, pos float32) error {
	if debug {
		fmt.Println("StageRelMove", Axis, pos)
	}

	err := this.plc.WriteRegister(XinJie.HD(this.motors[Axis].ADDR.LEN_HD), pos, XinJie.Float32, false)
	if err != nil {
		return err
	}
	return this.plc.WriteCoil(XinJie.M(this.motors[Axis].ADDR.MOVE_M), true)
}

func (this *StageService) StageGetSpeed(Axis string) (float32, error) {
	if debug {
		fmt.Println("StageGetSpeed", Axis)
	}
	val, err := this.plc.ReadRegister(XinJie.HD(this.motors[Axis].ADDR.SPEED_HD), XinJie.Float32, false)
	if err != nil {
		return 0, err
	}
	speed, ok := val.(float32)
	if !ok {
		return 0, errors.New("ReadRegister type error: " + Axis)
	}
	return speed, nil
}

func (this *StageService) StageGetResolution(Axis string) (uint32, error) {
	if debug {
		fmt.Println("StageGetResolution", Axis)
	}
	val, err := this.plc.ReadRegister(XinJie.HD(this.motors[Axis].ADDR.RESOLUTION_HD), XinJie.UInt32, false)
	if err != nil {
		return 0, err
	}
	resolution, ok := val.(uint32)
	if !ok {
		return 0, errors.New("ReadRegister type error: " + Axis)
	}
	return resolution, nil
}

func (this *StageService) StageSetSpeed(Axis string, speed float32) error {
	if debug {
		fmt.Println("StageSetSpeed", Axis, speed)
	}
	if _, ok := this.motors[Axis]; !ok {
		return errors.New("unknown axis: " + Axis)
	}
	return this.plc.WriteRegister(XinJie.HD(this.motors[Axis].ADDR.SPEED_HD), speed, XinJie.Float32, false)
}

// StageSetResolution 设置指定轴的分辨率
func (this *StageService) StageSetResolution(Axis string, resolution uint32) error {
	if debug {
		fmt.Println("StageSetResolution", Axis, resolution)
	}
	if _, ok := this.motors[Axis]; !ok {
		return errors.New("unknown axis: " + Axis)
	}
	return this.plc.WriteRegister(XinJie.HD(this.motors[Axis].ADDR.RESOLUTION_HD), resolution, XinJie.UInt32, false)
}

func (this *StageService) SetAlarmLED(alarm bool) error {
	if debug {
		fmt.Println("SetAlarmLED", alarm)
	}
	return this.plc.WriteCoil(XinJie.Y(15), alarm)
}
