package components

import (
	"context"
	"errors"
	"fmt"

	XinJie "github.com/Mr-tang0/PIMSGoMod/protocol/XinJie"
)

const debug = true

type StageService struct {
	motors MotorDetails
	plc    XinJie.XinjieClient
	ctx    context.Context
}

func NewStageService() *StageService {
	return &StageService{
		plc: *XinJie.NewXinjieClient(""),
		motors: MotorDetails{
			X:  Motor{Status: "Stopped"},
			Y:  Motor{Status: "Stopped"},
			Z:  Motor{Status: "Stopped"},
			R:  Motor{Status: "Stopped"},
			XX: Motor{Status: "Stopped"},
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
	go this.heartbeat()
	return nil
}

// StagesDisconnect 断开 diffract 设备连接
func (this *StageService) StagesDisconnect() error {
	if debug {
		fmt.Println("StagesDisconnect")
	}
	this.plc.Close()
	return nil
}

func (this *StageService) heartbeat() {

}

func (this *StageService) StageStop(Axis string) error {
	if debug {
		fmt.Println("StageStop", Axis)
	}
	switch Axis {
	case "X", "x":
		return this.plc.WriteCoil(XinJie.M(X_STOP_M), true)
	case "Y", "y":
		return this.plc.WriteCoil(XinJie.M(Y_STOP_M), true)
	case "Z", "z":
		return this.plc.WriteCoil(XinJie.M(Z_STOP_M), true)
	case "R", "r":
		return this.plc.WriteCoil(XinJie.M(R_STOP_M), true)
	case "XX", "xx":
		return this.plc.WriteCoil(XinJie.M(XX_STOP_M), true)
	default:
		return errors.New("Axis is not supported")
	}
}

func (this *StageService) StageRelMove(Axis string, pos float32) error {
	if debug {
		fmt.Println("StagesStop", Axis)
	}
	switch Axis {
	case "X", "x":
		err := this.plc.WriteRegister(XinJie.HD(X_POS_HD), pos, XinJie.Float32, false)
		if err != nil {
			return err
		}
		return this.plc.WriteCoil(XinJie.M(X_MOVE_M), true)
	case "Y", "y":
		err := this.plc.WriteRegister(XinJie.HD(Y_POS_HD), pos, XinJie.Float32, false)
		if err != nil {
			return err
		}
		return this.plc.WriteCoil(XinJie.M(Y_MOVE_M), true)
	case "Z", "z":
		err := this.plc.WriteRegister(XinJie.HD(Z_POS_HD), pos, XinJie.Float32, false)
		if err != nil {
			return err
		}
		return this.plc.WriteCoil(XinJie.M(Z_MOVE_M), true)
	case "R", "r":
		err := this.plc.WriteRegister(XinJie.HD(R_POS_HD), pos, XinJie.Float32, false)
		if err != nil {
			return err
		}
		return this.plc.WriteCoil(XinJie.M(R_MOVE_M), true)
	case "XX", "xx":
		err := this.plc.WriteRegister(XinJie.HD(XX_POS_HD), pos, XinJie.Float32, false)
		if err != nil {
			return err
		}
		return this.plc.WriteCoil(XinJie.M(XX_MOVE_M), true)
	default:
		return errors.New("Axis is not supported")
	}
}

// // GetStagesDetails 获取 diffract 设备详情
// func (this *StageService) GetStagesDetails() error {
// 	if debug {
// 		fmt.Println("GetStagesDetails")
// 	}
// 	details, err := this.plc.ReadHDRegisters(0, 25, XinjieClient.Float32)
// 	if err != nil {
// 		return err
// 	}
// 	this.motors.X.Speed = details[X_SPEED_HD/2].(float32)
// 	this.motors.Y.Speed = details[Y_SPEED_HD/2].(float32)
// 	this.motors.Z.Speed = details[Z_SPEED_HD/2].(float32)
// 	this.motors.R.Speed = details[R_SPEED_HD/2].(float32)
// 	this.motors.XX.Speed = details[XX_SPEED_HD/2].(float32)

// 	this.motors.X.Resolution = details[X_RESOLUTION_HD/2].(float32)
// 	this.motors.Y.Resolution = details[Y_RESOLUTION_HD/2].(float32)
// 	this.motors.Z.Resolution = details[Z_RESOLUTION_HD/2].(float32)
// 	this.motors.R.Resolution = details[R_RESOLUTION_HD/2].(float32)
// 	this.motors.XX.Resolution = details[XX_RESOLUTION_HD/2].(float32)

// 	if debug {
// 		fmt.Println(this.motors)
// 	}
// 	return nil
// }

// // heartbeat 心跳
// func (this *StageService) heartbeat() {
// 	if debug {
// 		fmt.Println("heartbeat")
// 	}
// 	for {
// 		this.updateAxis(&this.motors.X, X_POS_HD)
// 		this.updateAxis(&this.motors.Y, Y_POS_HD)
// 		this.updateAxis(&this.motors.Z, Z_POS_HD)
// 		this.updateAxis(&this.motors.R, R_POS_HD)
// 		this.updateAxis(&this.motors.XX, XX_POS_HD)

// 		if debug {
// 			fmt.Println(this.motors.X.Position, this.motors.Y.Position, this.motors.Z.Position, this.motors.R.Position, this.motors.XX.Position)
// 		}
// 		runtime.EventsEmit(this.ctx, "motor_heartbeat", this.motors)
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }

// // updateAxis 更新轴状态
// func (this *StageService) updateAxis(motor *Motor, addr uint16) {
// 	pos, _ := this.plc.ReadRegister(XinJie.HD(addr), XinJie.Float32, false)
// 	newPos := pos.(float32)
// 	if newPos == motor.Position {
// 		motor.Status = "Stopped"
// 	} else {
// 		motor.Status = "Moving"
// 	}
// 	motor.Position = newPos
// }

// // SetAxisSpeed 设置 diffract 设备轴速度
// func (this *StageService) SetAxisSpeed(Axis string, speed float32) error {
// 	if debug {
// 		fmt.Println("SetAxisSpeed", Axis, speed)
// 	}
// 	if Axis == "" || speed == 0 {
// 		return errors.New("Axis or speed is empty")
// 	}
// 	switch Axis {
// 	case "X":
// 		return this.plc.WriteRegister(X_JOG_M, speed, XinJie.Float32, false)
// 	case "Y":
// 		return this.plc.WriteRegister(Y_JOG_M, speed, XinJie.Float32, false)
// 	case "Z":
// 		return this.plc.WriteRegister(Z_JOG_M, speed, XinJie.Float32, false)
// 	case "R":
// 		return this.plc.WriteRegister(R_JOG_M, speed, XinJie.Float32, false)
// 	case "XX":
// 		return this.plc.WriteRegister(XX_JOG_M, speed, XinJie.Float32, false)
// 	default:
// 		return errors.New("Axis is not supported")
// 	}
// }

// // SetAxisResolution 设置 diffract 设备轴分辨率
// func (this *StageService) SetAxisResolution(Axis string, resolution float32) error {
// 	if debug {
// 		fmt.Println("SetAxisResolution", Axis, resolution)
// 	}
// 	if Axis == "" || resolution == 0 {
// 		return errors.New("Axis or resolution is empty")
// 	}
// 	switch Axis {
// 	case "X":
// 		return this.plc.WriteRegister(X_RESOLUTION_HD, resolution, XinJie.Float32, false)
// 	case "Y":
// 		return this.plc.WriteRegister(Y_RESOLUTION_HD, resolution, XinJie.Float32, false)
// 	case "Z":
// 		return this.plc.WriteRegister(Z_RESOLUTION_HD, resolution, XinJie.Float32, false)
// 	case "R":
// 		return this.plc.WriteRegister(R_RESOLUTION_HD, resolution, XinJie.Float32, false)
// 	case "XX":
// 		return this.plc.WriteHDRegister(XX_RESOLUTION_HD, resolution, XinjieClient.Float32)
// 	default:
// 		return errors.New("Axis is not supported")
// 	}
// }

// // StagesJOGMove 运动 diffract 设备轴
// func (this *StageService) StagesJOGMove(Axis string, dir int) error {
// 	if debug {
// 		fmt.Println("StagesJOGMove", Axis, dir, dir)
// 	}

// 	var M uint16 = 0
// 	var Speed float32 = 0.0

// 	switch Axis {
// 	case "X":
// 		M = X_JOG_M
// 		Speed = this.motors.X.Speed * float32(dir)
// 	case "Y":
// 		M = Y_JOG_M
// 		Speed = this.motors.Y.Speed * float32(dir)
// 	case "Z":
// 		M = Z_JOG_M
// 		Speed = this.motors.Z.Speed * float32(dir)
// 	case "R":
// 		M = R_JOG_M
// 		Speed = this.motors.R.Speed * float32(dir)
// 	case "XX":
// 		M = XX_JOG_M
// 		Speed = this.motors.XX.Speed * float32(dir)
// 	default:
// 		return errors.New("Axis is not supported")
// 	}
// 	if err := this.SetAxisSpeed(Axis, Speed); err != nil {
// 		return err
// 	}
// 	return this.plc.WriteMCoil(M, true)
// }

// // StagesSTOPMove 停止 diffract 设备轴
// func (this *StageService) StagesSTOPMove(Axis string) error {
// 	if debug {
// 		fmt.Println("StagesSTOPMove", Axis)
// 	}
// 	if Axis == "" {
// 		return errors.New("Axis is empty")
// 	}
// 	switch Axis {
// 	case "X":
// 		return this.plc.WriteCoil(XinJie.M(X_STOP_M), true)
// 	case "Y":
// 		return this.plc.WriteCoil(XinJie.M(Y_STOP_M), true)
// 	case "Z":
// 		return this.plc.WriteCoil(XinJie.M(Z_STOP_M), true)
// 	case "R":
// 		return this.plc.WriteCoil(XinJie.M(R_STOP_M), true)
// 	case "XX":
// 		return this.plc.WriteCoil(XinJie.M(XX_STOP_M), true)
// 	default:
// 		return errors.New("Axis is not supported")
// 	}
// }

// // StagesRELMove 相对运动 diffract 设备轴
// func (this *StageService) StagesRELMove(Axis string, rel_len float32) error {
// 	if debug {
// 		fmt.Println("StagesRELMove", Axis, rel_len)
// 	}
// 	if Axis == "" || rel_len == 0 {
// 		return errors.New("Axis or rel_len is empty")
// 	}
// 	switch Axis {
// 	case "X":
// 		this.plc.WriteHDRegister(X_MOVE_M, rel_len, XinjieClient.Float32)
// 		return this.plc.WriteMCoil(X_MOVE_M, true)
// 	case "Y":
// 		this.plc.WriteHDRegister(Y_MOVE_M, rel_len, XinjieClient.Float32)
// 		return this.plc.WriteMCoil(Y_MOVE_M, true)
// 	case "Z":
// 		this.plc.WriteHDRegister(Z_MOVE_M, rel_len, XinjieClient.Float32)
// 		return this.plc.WriteMCoil(Z_MOVE_M, true)
// 	case "R":
// 		this.plc.WriteHDRegister(R_MOVE_M, rel_len, XinjieClient.Float32)
// 		return this.plc.WriteMCoil(R_MOVE_M, true)
// 	case "XX":
// 		this.plc.WriteHDRegister(XX_MOVE_M, rel_len, XinjieClient.Float32)
// 		return this.plc.WriteMCoil(XX_MOVE_M, true)
// 	default:
// 		return errors.New("Axis is not supported")
// 	}
// }

// // StagesABSMove 绝对运动 diffract 设备轴
// func (this *StageService) StagesABSMove(axis string, abs_len float32) error {
// 	if debug {
// 		fmt.Println("StagesABSMove", axis, abs_len)
// 	}
// 	rel_len := float32(0.0)
// 	switch axis {
// 	case "X":
// 		rel_len = abs_len - this.motors.X.Position
// 	case "Y":
// 		rel_len = abs_len - this.motors.Y.Position
// 	case "Z":
// 		rel_len = abs_len - this.motors.Z.Position
// 	case "R":
// 		rel_len = abs_len - this.motors.R.Position
// 	case "XX":
// 		rel_len = abs_len - this.motors.XX.Position
// 	default:
// 		return errors.New("Axis is not supported")
// 	}
// 	if err := this.StagesRELMove(axis, rel_len); err != nil {
// 		return err
// 	}

// 	return nil
// }
