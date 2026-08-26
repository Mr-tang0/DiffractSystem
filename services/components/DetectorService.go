package components

import (
	sdk "Diffract/services/CSDK"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DetectorService struct {
	ctx          context.Context
	detector     *sdk.NetCom
	IsConnecting bool

	heartbeat DetectorHeartbeat
}

type DetectorHeartbeat struct {
	SN         string  `json:"sn"`
	Mode       string  `json:"mode"`
	Tempreture float32 `json:"tempreture"`
	Humidity   float32 `json:"humidity"`

	ExposeTime  uint16 `json:"expose_time"`
	Binning     string `json:"binning"`
	RepeatTimes uint16 `json:"repeat_times"`
	Gain        uint16 `json:"gain"`

	Width  int `json:"width"`
	Height int `json:"height"`

	// 图像数据
	ImageMinVal uint32 `json:"image_min_val"`
	ImageMaxVal uint32 `json:"image_max_val"`
}

func NewDetectorService() *DetectorService {
	return &DetectorService{
		heartbeat: DetectorHeartbeat{
			SN:          "",
			Mode:        "",
			Tempreture:  0.0,
			Humidity:    0.0,
			ImageMinVal: 0,
			ImageMaxVal: 255,
		},
	}
}

func (this *DetectorService) Startup(ctx context.Context) error {
	if ctx == nil {
		return fmt.Errorf("ctx is nil")
	}
	this.ctx = ctx
	this.detector = sdk.NewNetCom(ctx)
	return nil
}

// Init 初始化探测器设备，自动初始化SDK并注册事件回调
func (this *DetectorService) Init() {
	fmt.Println("初始化探测器设备")
	this.detector.RegisterCallback()     // 先注册回调
	this.detector.COM_Init()             // 再初始化SDK
	this.detector.COM_StartNet()         // 启动网络监听
	this.detector.COM_SetCalibMode(0x06) // 设置校准模式 IMG_CALIB_GAIN | IMG_CALIB_DEFECT

	//监听net_raw事件
	runtime.EventsOn(this.ctx, "net_raw", func(raw16Data ...interface{}) {
		fmt.Println("接收到NET原始数据")
		if len(raw16Data) > 0 {
			this.handleNETImageEvent(raw16Data[0])
		}
	})

	runtime.EventsOn(this.ctx, "net_heartbeat", func(net ...interface{}) {
		data, ok := net[0].(map[string]interface{})
		if !ok {
			return
		}

		if v, ok := data["sn"].(string); ok {
			this.heartbeat.SN = v
		}
		if v, ok := data["mode"].(string); ok {
			this.heartbeat.Mode = v
		}
		if v, ok := data["tempreture"].(float64); ok {
			this.heartbeat.Tempreture = float32(v)
		}
		if v, ok := data["humidity"].(float64); ok {
			this.heartbeat.Humidity = float32(v)
		}

		runtime.EventsEmit(this.ctx, "detector_heartbeat", this.heartbeat)
	})

	runtime.EventsOn(this.ctx, "net_linked", func(linked ...interface{}) {
		this.CTSetDstModel()
		dynamicPara := this.detector.COM_GetDynamicPara()
		if v, ok := dynamicPara["pxwin"].(uint32); ok {
			this.heartbeat.ExposeTime = uint16(v)
		}
		if v, ok := dynamicPara["pbinMode"].(string); ok {
			this.heartbeat.Binning = v
		}
		if v, ok := dynamicPara["prepeat"].(uint16); ok {
			this.heartbeat.RepeatTimes = v
		}
		fmt.Println("动态参数:", this.heartbeat)
	})

}

// DetectorConnect 连接探测器
func (this *DetectorService) DetectorConnect() error {
	fmt.Println("连接探测器")
	flag := this.detector.COM_Open(0)
	if !flag {
		return fmt.Errorf("无法连接探测器")
	}
	this.IsConnecting = true
	return nil
}

// DetectorDisconnect 断开连接探测器
func (this *DetectorService) DetectorDisconnect() error {
	fmt.Println("断开连接探测器")
	this.detector.COM_StopNet()
	flag := this.detector.COM_Close()
	if !flag {
		return fmt.Errorf("无法断开连接")
	}
	this.IsConnecting = false
	return nil
}

func (this *DetectorService) CTSetDstModel() {
	fmt.Println("[DST] 设置DST模式")

	this.detector.COM_Dst()
	time.Sleep(100 * time.Millisecond) //等待DST模式切换完成")

	t := this.detector.COM_GetExposeTime()
	fmt.Printf("[HST] 曝光时间: %d\n", t)

	this.detector.COM_Dprep()
	time.Sleep(time.Duration(t+3500) * time.Millisecond) //等待拍背底完成

	fmt.Println("[DST] 启动DST")
}

func (this *DetectorService) DetectorCapture() {
	this.detector.COM_Dacq()
}

func (this *DetectorService) handleNETImageEvent(data interface{}) {
	// 解析事件数据
	rawData, ok := data.(map[string]interface{})
	if !ok {
		fmt.Printf("[DetectorService] 数据类型转换失败，期望map[string]interface{}类型，实际类型: %T\n", data)
		return
	}

	// 使用统一的归一化函数处理图片
	normalizedImg, err := this.normalizeRawImage(rawData)
	if err != nil {
		fmt.Printf("[CTDevice] 归一化图片失败: %v\n", err)
		return
	}

	// 发送detector_image事件供前端显示
	runtime.EventsEmit(this.ctx, "detector_image", normalizedImg)
}

func (this *DetectorService) normalizeRawImage(rawData map[string]interface{}) (map[string]interface{}, error) {
	imageData, ok := rawData["image"].([]byte)
	if !ok {
		return nil, fmt.Errorf("图片数据类型转换失败")
	}

	width, _ := rawData["width"].(int)
	height, _ := rawData["height"].(int)
	this.heartbeat.Width = width
	this.heartbeat.Height = height

	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("图片尺寸无效: width=%d, height=%d", width, height)
	}

	grayImg := image.NewGray(image.Rect(0, 0, width, height))

	rangeVal := this.heartbeat.ImageMaxVal - this.heartbeat.ImageMinVal
	if rangeVal == 0 {
		rangeVal = 1
	}

	// 计算直方图（0-10000范围，100个bin）
	const histogramBins = 100
	histogram := make([]int, histogramBins)

	for i := 0; i < width*height; i++ {
		val := uint32(imageData[i*2]) | (uint32(imageData[i*2+1]) << 8)
		var normalized uint32
		if val <= this.heartbeat.ImageMinVal {
			normalized = 0
		} else if val >= this.heartbeat.ImageMaxVal {
			normalized = 255
		} else {
			normalized = (val - this.heartbeat.ImageMinVal) * 255 / rangeVal
		}
		grayImg.Pix[i] = uint8(normalized)

		// 统计直方图：将16位像素值映射到0-10000范围并分桶
		histVal := uint32(float64(val) / 65535.0 * 10000.0)
		if histVal >= 10000 {
			histVal = 9999
		}
		binIndex := int(histVal / 100)
		histogram[binIndex]++
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, grayImg, &jpeg.Options{Quality: 60}); err != nil {
		return nil, err
	}

	encodedStr := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

	return map[string]interface{}{
		"image":     encodedStr,
		"width":     width,
		"height":    height,
		"histogram": histogram,
	}, nil
}
