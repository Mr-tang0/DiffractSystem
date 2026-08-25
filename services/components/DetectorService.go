package components

import (
	sdk "Diffract/services/CSDK"
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DetectorService struct {
	ctx      context.Context
	detector *sdk.NetCom
}

func NewDetectorService() *DetectorService {
	return &DetectorService{}
}

func (this *DetectorService) Startup(ctx context.Context) error {
	this.ctx = ctx
	return nil
}

func (this *DetectorService) Init() {
	fmt.Println("初始化探测器设备")
	this.detector.RegisterCallback()     // 先注册回调
	this.detector.COM_Init()             // 再初始化SDK
	this.detector.COM_StartNet()         // 启动网络监听
	this.detector.COM_SetCalibMode(0x06) // 设置校准模式 IMG_CALIB_GAIN | IMG_CALIB_DEFECT

	//监听raw事件
	runtime.EventsOn(this.ctx, "raw", func(raw16Data ...interface{}) {
		fmt.Println("接收到探测器原始数据")
		if len(raw16Data) > 0 {
			this.handleDetectorImageEvent(raw16Data[0])
		}
	})
}

// handleDetectorImageEvent 处理raw事件数据
func (this *DetectorService) handleDetectorImageEvent(data interface{}) {

}

// DetectorConnect 连接探测器
func (this *DetectorService) DetectorConnect() error {
	fmt.Println("连接探测器")
	flag := this.detector.COM_Open(0)
	if !flag {
		fmt.Println("无法连接探测器: ")
		return fmt.Errorf("无法连接探测器")
	}
	return nil
}

// DetectorDisconnect 断开连接探测器
func (this *DetectorService) DetectorDisconnect() error {
	fmt.Println("断开连接探测器")
	this.detector.COM_StopNet()
	flag := this.detector.COM_Close()
	if !flag {
		fmt.Println("无法断开连接")
		return fmt.Errorf("无法断开连接")
	}
	return nil
}

func (this *DetectorService) heartbeat() {

}

func (this *DetectorService) DetectorSetMode(mode string) error {
	switch mode {
	case "DST":
		this.detector.COM_Dst()
		time.Sleep(50 * time.Millisecond)
		t := this.detector.COM_GetExposeTime()
		fmt.Printf("[HST] 曝光时间: %d\n", t)
		this.detector.COM_Dprep()
		time.Sleep(time.Duration(t+3500) * time.Millisecond) //等待拍背底完成
	default:
		return fmt.Errorf("不支持的校准模式")
	}
	return nil
}

func (this *DetectorService) DetectorTrigger() {
	this.detector.COM_Dacq()
}
