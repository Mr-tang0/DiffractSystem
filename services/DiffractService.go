package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	XinjieClient "github.com/Mr-tang0/PIMSGoMod/protocol"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const debug = true

type Motor struct {
	Status     string  `json:"status"` // 状态: Moving, Stopped, Error
	Position   float32 `json:"position"`
	Speed      float32 `json:"speed"`
	Resolution float32 `json:"resolution"`
}

type MotorDetails struct {
	X  Motor `json:"x"`
	Y  Motor `json:"y"`
	Z  Motor `json:"z"`
	R  Motor `json:"r"`
	XX Motor `json:"xx"`
}

type DiffractService struct {
	MODE string // 模式:

	motors MotorDetails
	plc    XinjieClient.XinjieClient
	ctx    context.Context
}

func NewDiffractService() *DiffractService {
	return &DiffractService{
		plc: *XinjieClient.NewXinjieClient(),
		motors: MotorDetails{
			X:  Motor{Status: "Stopped"},
			Y:  Motor{Status: "Stopped"},
			Z:  Motor{Status: "Stopped"},
			R:  Motor{Status: "Stopped"},
			XX: Motor{Status: "Stopped"},
		},
	}
}

func (this *DiffractService) Startup(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx is nil")
	}
	this.ctx = ctx
	return nil
}

// DiffractStagesConnect 连接 diffract 设备
func (this *DiffractService) DiffractStagesConnect(ip string) error {
	if debug {
		fmt.Println("DiffractStagesConnect", ip)
	}
	if ip == "" {
		return errors.New("ip is empty")
	}
	err := this.plc.OpenTCP(ip, 1)
	if err != nil {
		return err
	}
	go this.heartbeat()
	return nil
}

// DiffractGetStagesDetails 获取 diffract 设备详情
func (this *DiffractService) DiffractGetStagesDetails() error {
	if debug {
		fmt.Println("DiffractGetStagesDetails")
	}
	details, err := this.plc.ReadHDRegisters(0, 25, XinjieClient.Float32)
	if err != nil {
		return err
	}
	this.motors.X.Speed = details[X_SPEED_HD/2].(float32)
	this.motors.Y.Speed = details[Y_SPEED_HD/2].(float32)
	this.motors.Z.Speed = details[Z_SPEED_HD/2].(float32)
	this.motors.R.Speed = details[R_SPEED_HD/2].(float32)
	this.motors.XX.Speed = details[XX_SPEED_HD/2].(float32)

	this.motors.X.Resolution = details[X_RESOLUTION_HD/2].(float32)
	this.motors.Y.Resolution = details[Y_RESOLUTION_HD/2].(float32)
	this.motors.Z.Resolution = details[Z_RESOLUTION_HD/2].(float32)
	this.motors.R.Resolution = details[R_RESOLUTION_HD/2].(float32)
	this.motors.XX.Resolution = details[XX_RESOLUTION_HD/2].(float32)

	if debug {
		fmt.Println(this.motors)
	}
	return nil
}

// heartbeat 心跳
func (this *DiffractService) heartbeat() {
	if debug {
		fmt.Println("heartbeat")
	}
	for {
		this.updateMotor(&this.motors.X, X_POS_HD)
		this.updateMotor(&this.motors.Y, Y_POS_HD)
		this.updateMotor(&this.motors.Z, Z_POS_HD)
		this.updateMotor(&this.motors.R, R_POS_HD)
		this.updateMotor(&this.motors.XX, XX_POS_HD)

		if debug {
			fmt.Println(this.motors.X.Position, this.motors.Y.Position, this.motors.Z.Position, this.motors.R.Position, this.motors.XX.Position)
		}
		runtime.EventsEmit(this.ctx, "motor_heartbeat", this.motors)
		time.Sleep(500 * time.Millisecond)
	}
}

// updateMotor 更新电机状态
func (this *DiffractService) updateMotor(motor *Motor, addr uint16) {
	pos, _ := this.plc.ReadHDRegister(addr, XinjieClient.Float32)
	newPos := pos.(float32)
	if newPos == motor.Position {
		motor.Status = "Stopped"
	} else {
		motor.Status = "Moving"
	}
	motor.Position = newPos
}

// DiffractStagesDisconnect 断开 diffract 设备连接
func (this *DiffractService) DiffractStagesDisconnect() error {
	if debug {
		fmt.Println("DiffractStagesDisconnect")
	}
	this.plc.Close()
	return nil
}

// DiffractSetAxisSpeed 设置 diffract 设备轴速度
func (this *DiffractService) DiffractSetAxisSpeed(Axis string, speed float32) error {
	if debug {
		fmt.Println("DiffractSetAxisSpeed", Axis, speed)
	}
	if Axis == "" || speed == 0 {
		return errors.New("Axis or speed is empty")
	}
	switch Axis {
	case "X":
		return this.plc.WriteHDRegister(X_JOG_M, speed, XinjieClient.Float32)
	case "Y":
		return this.plc.WriteHDRegister(Y_JOG_M, speed, XinjieClient.Float32)
	case "Z":
		return this.plc.WriteHDRegister(Z_JOG_M, speed, XinjieClient.Float32)
	case "R":
		return this.plc.WriteHDRegister(R_JOG_M, speed, XinjieClient.Float32)
	case "XX":
		return this.plc.WriteHDRegister(XX_JOG_M, speed, XinjieClient.Float32)
	default:
		return errors.New("Axis is not supported")
	}
}

// DiffractSetAxisResolution 设置 diffract 设备轴分辨率
func (this *DiffractService) DiffractSetAxisResolution(Axis string, resolution float32) error {
	if debug {
		fmt.Println("DiffractSetAxisResolution", Axis, resolution)
	}
	if Axis == "" || resolution == 0 {
		return errors.New("Axis or resolution is empty")
	}
	switch Axis {
	case "X":
		return this.plc.WriteHDRegister(X_RESOLUTION_HD, resolution, XinjieClient.Float32)
	case "Y":
		return this.plc.WriteHDRegister(Y_RESOLUTION_HD, resolution, XinjieClient.Float32)
	case "Z":
		return this.plc.WriteHDRegister(Z_RESOLUTION_HD, resolution, XinjieClient.Float32)
	case "R":
		return this.plc.WriteHDRegister(R_RESOLUTION_HD, resolution, XinjieClient.Float32)
	case "XX":
		return this.plc.WriteHDRegister(XX_RESOLUTION_HD, resolution, XinjieClient.Float32)
	default:
		return errors.New("Axis is not supported")
	}
}

// DiffractStagesJOG 运动 diffract 设备轴
func (this *DiffractService) DiffractStagesJOG(Axis string, speed float32, jog bool) error {
	if debug {
		fmt.Println("DiffractStagesJOG", Axis, speed, jog)
	}
	if Axis == "" || speed == 0 {
		return errors.New("Axis or speed is empty")
	}

	if err := this.DiffractSetAxisSpeed(Axis, speed); err != nil {
		return err
	}
	switch Axis {
	case "X":
		return this.plc.WriteMCoil(X_JOG_M, jog)
	case "Y":
		return this.plc.WriteMCoil(Y_JOG_M, jog)
	case "Z":
		return this.plc.WriteMCoil(Z_JOG_M, jog)
	case "R":
		return this.plc.WriteMCoil(R_JOG_M, jog)
	case "XX":
		return this.plc.WriteMCoil(XX_JOG_M, jog)
	default:
		return errors.New("Axis is not supported")
	}
}

// DiffractStagesSTOP 停止 diffract 设备轴
func (this *DiffractService) DiffractStagesSTOP(Axis string) error {
	if debug {
		fmt.Println("DiffractStagesSTOP", Axis)
	}
	if Axis == "" {
		return errors.New("Axis is empty")
	}
	switch Axis {
	case "X":
		return this.plc.WriteMCoil(X_STOP_M, true)
	case "Y":
		return this.plc.WriteMCoil(Y_STOP_M, true)
	case "Z":
		return this.plc.WriteMCoil(Z_STOP_M, true)
	case "R":
		return this.plc.WriteMCoil(R_STOP_M, true)
	case "XX":
		return this.plc.WriteMCoil(XX_STOP_M, true)
	default:
		return errors.New("Axis is not supported")
	}
}

// DiffractStagesMOVE 移动 diffract 设备轴
func (this *DiffractService) DiffractStagesMOVE(Axis string, len float32) error {
	if debug {
		fmt.Println("DiffractStagesMOVE", Axis, len)
	}
	if Axis == "" || len == 0 {
		return errors.New("Axis or len is empty")
	}
	switch Axis {
	case "X":
		this.plc.WriteHDRegister(X_MOVE_M, len, XinjieClient.Float32)
		return this.plc.WriteMCoil(X_MOVE_M, true)
	case "Y":
		this.plc.WriteHDRegister(Y_MOVE_M, len, XinjieClient.Float32)
		return this.plc.WriteMCoil(Y_MOVE_M, true)
	case "Z":
		this.plc.WriteHDRegister(Z_MOVE_M, len, XinjieClient.Float32)
		return this.plc.WriteMCoil(Z_MOVE_M, true)
	case "R":
		this.plc.WriteHDRegister(R_MOVE_M, len, XinjieClient.Float32)
		return this.plc.WriteMCoil(R_MOVE_M, true)
	case "XX":
		this.plc.WriteHDRegister(XX_MOVE_M, len, XinjieClient.Float32)
		return this.plc.WriteMCoil(XX_MOVE_M, true)
	default:
		return errors.New("Axis is not supported")
	}
}
