package components

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// HVPS协议常量
const HVPS_ADDRESS = 0xA1 // 设备地址（固定）

// HVPS协议命令码定义 (2字节)
const (
	CMD_FEEDBACK_REQ      = 0x6000 // Feedback Value Request
	CMD_FEEDBACK_RESP     = 0x600A // Feedback Value Response
	CMD_SETPOINTINFO_REQ  = 0x6020 // Setpoint Info Request
	CMD_SETPOINTINFO_RESP = 0x602A // Setpoint Info Response
	CMD_HV_SETTING        = 0x6100 // High Voltage Setting
	CMD_HV_SETTING_CONF   = 0x610A // High Voltage Setting Confirm
	CMD_CURRENT_SETTING   = 0x6200 // Current Setting
	CMD_CURRENT_CONF      = 0x620A // Current Setting Confirm
	CMD_FIL_PREHEAT       = 0x6300 // Filament Preheating Setting
	CMD_FIL_PREHEAT_CONF  = 0x630A // Filament Preheating Setting Confirm
	CMD_FIL_LIMIT         = 0x6400 // Filament Limit Setting
	CMD_FIL_LIMIT_CONF    = 0x640A // Filament Limit Setting Confirm
	CMD_LOCAL_REMOTE      = 0x6500 // Local Remote Switch
	CMD_LOCAL_REMOTE_CONF = 0x650A // Local Remote Switch Confirm
	CMD_HV_ONOFF          = 0x6900 // HV ON/OFF Setting
	CMD_HV_ONOFF_CONF     = 0x690A // HV ON/OFF Setting Confirm
	CMD_FIL_ONOFF         = 0x7000 // Filament ON/OFF Setting
	CMD_FIL_ONOFF_CONF    = 0x700A // Filament ON/OFF Setting Confirm
)

type HVPSSetpointInfo struct {
	HV float32 `json:"HV"`
	HI float32 `json:"HI"`
	FV float32 `json:"FV"`
	FI float32 `json:"FI"`
}

// HVPSService 高压电源网口通信类
type HVPSService struct {
	conn    net.Conn        //实际Socket连接
	ctx     context.Context // 上下文，用于取消操作
	timeout time.Duration   // 通信超时时间
	mu      sync.Mutex      // 互斥锁，防止sendFrame并发冲突
}

// NewHVPSService 创建高压电源设备实例
func NewHVPSService() *HVPSService {
	return &HVPSService{
		timeout: 1 * time.Second, // 默认1秒超时
	}
}

func (h *HVPSService) SetContent(ctx context.Context) {
	h.ctx = ctx
	if conn := ctx.Value("conn"); conn != nil {
		h.conn = conn.(net.Conn)
	}
}

// HighVoltageConnect 连接高压电源设备
func (h *HVPSService) HighVoltageConnect(HVPS_ip string, HVPS_port int) error {
	fmt.Println("连接高压电源设备: ", HVPS_ip)
	addr := net.JoinHostPort(HVPS_ip, fmt.Sprintf("%d", HVPS_port))
	conn, err := net.DialTimeout("tcp", addr, h.timeout)
	if err != nil {
		return err
	}
	h.conn = conn

	runtime.EventsEmit(h.ctx, "hvps_linked", map[string]bool{"hvps_linked": true})
	go h.GetHVPSDetails()
	return nil
}

// HighVoltageDisconnect 断开连接
func (h *HVPSService) HighVoltageDisconnect() error {
	if h.conn != nil {
		err := h.conn.Close()
		h.conn = nil
		return err
	}
	runtime.EventsEmit(h.ctx, "hvps_linked", map[string]bool{"hvps_linked": false})
	return errors.New("未连接")
}

func (h *HVPSService) HighVoltageIsConnected() bool {
	return h.conn != nil
}

func (h *HVPSService) GetHVPSDetails() {
	for {
		if !h.HighVoltageIsConnected() {
			fmt.Println("已经断开连接")
			return
		}
		frame := buildFrame(CMD_FEEDBACK_REQ, []byte{0x00, 0x00})
		command, data, err := h.sendFrame(frame)
		if err != nil {
			continue
		}
		if command == CMD_FEEDBACK_RESP {
			fmt.Println("获取反馈信息: ", data)
			//解算HVPS数据
			HVFeedback := uint16(data[2])<<8 | uint16(data[3])
			CurrentFeedback := uint16(data[4])<<8 | uint16(data[5])
			FilamentCurrentFeedback := uint16(data[6])<<8 | uint16(data[7])
			FilamentVoltageFeedback := uint16(data[8])<<8 | uint16(data[9])

			heartbeatData := map[string]float64{
				"HV": float64(HVFeedback) / 4095.0 * 50.0,
				"HI": float64(CurrentFeedback) / 4095.0 * 1000.0,
				"FI": float64(FilamentCurrentFeedback) / 4095.0 * 3.6,
				"FV": float64(FilamentVoltageFeedback) / 4095.0 * 5.5,
			}
			if h.ctx != nil {
				runtime.EventsEmit(h.ctx, "hvps_heartbeat", heartbeatData)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (h *HVPSService) HVPSGetSetpointInfo() HVPSSetpointInfo {
	fmt.Println("获取设置信息")
	frame := buildFrame(CMD_SETPOINTINFO_REQ, []byte{0x00, 0x00})
	command, data, err := h.sendFrame(frame)
	fmt.Println("获取设置信息: ", command)
	if err != nil {
		return HVPSSetpointInfo{}
	}
	if command == CMD_SETPOINTINFO_RESP {
		fmt.Println("获取设置信息成功: ", data)
		//解算HVPS数据
		HV := float32(uint16(data[0])<<8|uint16(data[1])) / 4095.0 * 50.0
		HI := float32(uint16(data[2])<<8|uint16(data[3])) / 4095.0 * 1000.0
		FV := float32(uint16(data[4])<<8|uint16(data[5])) / 4095.0 * 10.0
		FI := float32(uint16(data[6])<<8|uint16(data[7])) / 4095.0 * 10.0
		return HVPSSetpointInfo{HV: HV, HI: HI, FV: FV, FI: FI}
	}

	return HVPSSetpointInfo{}
}

func (h *HVPSService) HVPSSetRemote(remote bool) error {
	frame := []byte{}
	if remote {
		frame = buildFrame(CMD_LOCAL_REMOTE, []byte{0x01, 0x00})
	} else {
		frame = buildFrame(CMD_LOCAL_REMOTE, []byte{0x00, 0x00})
	}
	command, _, err := h.sendFrame(frame)
	if err != nil {
		return err
	}
	if command == CMD_LOCAL_REMOTE_CONF {
		fmt.Println("设置本地电源开关: ", remote)
		return nil
	}
	return errors.New("设置本地电源开关失败")
}

func (h *HVPSService) HVPSSetHV(KV float32) error {
	KV_val := int(KV*4095.0/50.0 + 0.5)
	frame := buildFrame(CMD_HV_SETTING, []byte{byte(KV_val >> 8), byte(KV_val & 0xFF)})

	command, _, err := h.sendFrame(frame)
	if err != nil {
		return err
	}
	if command == CMD_HV_SETTING_CONF {
		fmt.Println("设置高压: ", KV)
		return nil
	}
	return errors.New("设置高压失败")
}

func (h *HVPSService) HVPSSetHI(HI float32) error {
	HI_val := int(HI*4095.0/1000.0 + 0.5)
	frame := buildFrame(CMD_CURRENT_SETTING, []byte{byte(HI_val >> 8), byte(HI_val & 0xFF)})

	command, _, err := h.sendFrame(frame)
	if err != nil {
		return err
	}

	if command == CMD_CURRENT_CONF {
		fmt.Println("设置电流: ", HI)
		return nil
	}
	return errors.New("设置电流失败")
}

func (h *HVPSService) HVPSSourceOpen(open bool) error {
	fmt.Println("设置源开关: ", open)
	frame := []byte{}
	if open {
		frame = buildFrame(CMD_HV_ONOFF, []byte{0x01, 0x00})
	} else {
		frame = buildFrame(CMD_HV_ONOFF, []byte{0x00, 0x00})
	}
	command, _, err := h.sendFrame(frame)
	if err != nil {
		fmt.Println("设置高压电源开关失败: ", err)
		return err
	}

	if command == CMD_HV_ONOFF_CONF {
		fmt.Println("设置高压电源开关: ", open)
		return nil
	}
	return errors.New("设置高压电源开关失败")
}

func (h *HVPSService) HVPSSetFilamentPreheat(FilamentPreheat float32) error {
	FilamentPreheat_val := int(FilamentPreheat*4095.0/10.0 + 0.5)
	frame := buildFrame(CMD_FIL_PREHEAT, []byte{byte(FilamentPreheat_val >> 8), byte(FilamentPreheat_val & 0xFF)})

	command, _, err := h.sendFrame(frame)
	if err != nil {
		return err
	}

	if command == CMD_FIL_PREHEAT_CONF {
		fmt.Println("设置 filament 预热: ", FilamentPreheat)
		return nil
	}
	return errors.New("设置 filament 预热失败")
}
func (h *HVPSService) HVPSSetFilamentLimit(FilamentLimit float32) error {

	FilamentLimit_val := int(FilamentLimit*4095.0/10.0 + 0.5)
	frame := buildFrame(CMD_FIL_LIMIT, []byte{byte(FilamentLimit_val >> 8), byte(FilamentLimit_val & 0xFF)})

	command, _, err := h.sendFrame(frame)
	if err != nil {
		return err
	}

	if command == CMD_FIL_LIMIT_CONF {
		fmt.Println("设置 filament 限制: ", FilamentLimit)
		return nil
	}
	return errors.New("设置 filament 限制失败")
}

func (h *HVPSService) HVPSSetFilamentOpen(FilamentOpen bool) error {
	fmt.Println("设置 filament 开关: ", FilamentOpen)
	frame := []byte{}
	if FilamentOpen {
		frame = buildFrame(CMD_FIL_ONOFF, []byte{0x01, 0x00})
	} else {
		frame = buildFrame(CMD_FIL_ONOFF, []byte{0x00, 0x00})
	}
	command, _, err := h.sendFrame(frame)
	if err != nil {
		return err
	}

	if command == CMD_FIL_ONOFF_CONF {
		fmt.Println("设置 filament 开关: ", FilamentOpen)
		return nil
	}
	return errors.New("设置 filament 开关失败")
}

func (h *HVPSService) sendFrame(frame []byte) (command uint16, data []byte, err error) {
	if h.conn == nil {
		return 0, nil, errors.New("设备未连接")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	_, err = h.conn.Write(frame)
	if err != nil {
		return 0, nil, fmt.Errorf("发送失败: %v", err)
	}

	// 设置读取超时
	if err := h.conn.SetReadDeadline(time.Now().Add(h.timeout)); err != nil {
		return 0, nil, fmt.Errorf("设置超时失败: %v", err)
	}

	buffer := make([]byte, 1024)
	response := make([]byte, 0)

	for {
		n, err := h.conn.Read(buffer)
		if err != nil {
			return 0, nil, err
		}
		response = append(response, buffer[:n]...)
		if len(response) > 0 && response[len(response)-1] == 0x0D {
			break
		}
	}

	return parseResponse(response)
}

func parseResponse(response []byte) (command uint16, data []byte, err error) {
	// fmt.Println("response: ", response)
	byteLength := len(response)
	if response[byteLength-1] != 0x0D || response[0] != HVPS_ADDRESS {
		fmt.Println("接收回复格式错误")
		err = errors.New("接收回复格式错误")
		return 0, nil, err
	}

	// 校验和校验
	checksum := uint16(response[len(response)-4])<<8 | uint16(response[len(response)-3])
	checksumCalc := uint16(popcount(response[0 : byteLength-4]))
	if checksumCalc != checksum {
		fmt.Println("校验和校验失败")
		err = errors.New("校验和校验失败")
		return 0, nil, err
	}

	command = uint16(response[1])<<8 | uint16(response[2])
	data = response[3 : byteLength-4]
	return command, data, nil
}

// popcountByte 计算单个字节中1的个数
func popcountByte(b byte) int {
	count := 0
	for b != 0 {
		count++
		b &= b - 1
	}
	return count
}

// popcount 计算多个字节中1的个数之和
func popcount(data []byte) int {
	count := 0
	for _, b := range data {
		count += popcountByte(b)
	}
	return count
}

// buildFrame 构建HVPS协议发送帧
func buildFrame(command uint16, data []byte) []byte {
	// 计算帧长度: Address(1) + Command(2) + Reserved(2) + Data + Checksum(2) + End(2)
	frameLen := 1 + 2 + 2 + len(data) + 2 + 2
	frame := make([]byte, frameLen)

	// Address (1字节)
	frame[0] = HVPS_ADDRESS

	// Command (2字节, BigEndian)
	frame[1] = byte(command >> 8)
	frame[2] = byte(command & 0xFF)

	// Reserved (2字节)
	frame[3] = 0x00
	frame[4] = 0x00

	// Data Message (可变长度)
	copy(frame[5:5+len(data)], data)

	// 计算校验和: Address + Command + Data 的二进制1的个数之和
	checksumData := make([]byte, 0, 5+len(data))
	checksumData = append(checksumData, HVPS_ADDRESS)
	checksumData = append(checksumData, byte(command>>8))
	checksumData = append(checksumData, byte(command&0xFF))
	checksumData = append(checksumData, 0x00)
	checksumData = append(checksumData, 0x00)
	checksumData = append(checksumData, data...)
	checksum := popcount(checksumData)

	// Checksum (2字节, BigEndian)
	frame[5+len(data)] = byte(checksum >> 8)
	frame[6+len(data)] = byte(checksum & 0xFF)

	// End Character (2字节)
	frame[frameLen-2] = 0x00
	frame[frameLen-1] = 0x0D

	return frame
}
