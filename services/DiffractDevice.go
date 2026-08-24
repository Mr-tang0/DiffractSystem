package services

// 探测器工作模式
type CTMODE int

const (
	IDLE    = 0
	HST     = 2
	DST     = 9
	AED1    = 3
	AED2    = 4
	RECOVER = 9
	ERR     = 99
)

type DiffractDeviceStatus struct {
	Mode          CTMODE // 当前工作模式
	ExposeTime    int    // 曝光时间
	Binning       string // 当前binning
	TriggerRepeat int    // 手动触发图片张数
	TriggerDelay  int    // 手动触发图片间隔
}

type DiffractDevice struct {
}

// NewDiffractDevice 创建一个新探测器设备
func NewDiffractDevice() *DiffractDevice {
	return &DiffractDevice{}
}

// // SetContent 设置上下文
// func (this *DiffractDevice) SetContent(ctx context.Context) {
// 	this.ctx = ctx
// 	Cnet := sdk.NewNetCom(ctx)
// 	this.CT = Cnet
// }

// // SetStage 设置电机设备
// func (this *DiffractDevice) SetStage(stage *components.StageService) {
// 	this.Stage = stage
// }

// // 设置project
// func (this *DiffractDevice) SetProject(project *Project) {
// 	this.project = project
// }

// // SetSystem 设置系统配置
// func (this *DiffractDevice) SetSystem(system *System) {
// 	this.system = system
// }

// // InitCT 初始化探测器设备，自动初始化SDK, 并注册事件回调函数
// func (this *DiffractDevice) InitCT() {
// 	fmt.Println("初始化探测器设备")
// 	this.CT.RegisterCallback()     // 先注册回调
// 	this.CT.COM_Init()             // 再初始化SDK
// 	this.CT.COM_StartNet()         // 启动网络监听
// 	this.CT.COM_SetCalibMode(0x06) // 设置校准模式 IMG_CALIB_GAIN | IMG_CALIB_DEFECT

// 	//监听ct_raw事件
// 	runtime.EventsOn(this.ctx, "ct_raw", func(raw16Data ...interface{}) {
// 		fmt.Println("接收到CT原始数据")
// 		if len(raw16Data) > 0 {
// 			this.handleCTImageEvent(raw16Data[0])
// 		}
// 	})
// }

// // DetectorConnect 连接探测器
// func (this *DiffractDevice) DetectorConnect() error {
// 	fmt.Println("连接探测器")
// 	flag := this.CT.COM_Open(0)
// 	if flag {
// 		this.IsConnecting = true
// 		return nil
// 	} else {
// 		fmt.Println("无法连接探测器: ")
// 		return fmt.Errorf("无法连接探测器")
// 	}
// }

// // DetectorDisconnect 断开连接探测器
// func (this *DiffractDevice) DetectorDisconnect() error {
// 	fmt.Println("断开连接器")
// 	flag := this.CT.COM_Close()
// 	if flag {
// 		this.IsConnecting = false
// 		return nil
// 	} else {
// 		fmt.Println("无法断开连接")
// 		return fmt.Errorf("无法断开连接")
// 	}
// }

// // DetectorIsConnected 检查连接状态
// func (this *DiffractDevice) DetectorIsConnected() bool {
// 	return this.IsConnecting
// }

// // SingleImageTrigger 单次采集：扫描一次并保存图片为backup
// func (this *DiffractDevice) SingleImageTrigger() error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	this.CollectMode = 1 // 设置为单次采集模式
// 	return this.DetectorSingleScan()
// }

// // DetectorSingleScan 单次扫描
// func (this *DiffractDevice) DetectorSingleScan() error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	if this.status.Mode != DST {
// 		this.CT.COM_Dst()
// 		time.Sleep(3000 * time.Millisecond)
// 		this.CT.COM_Dprep()
// 		time.Sleep(3000 * time.Millisecond) //等待拍背底完成
// 		this.status.Mode = DST
// 		fmt.Println("[DST] 启动DST")
// 	}
// 	this.CT.COM_Dacq()

// 	return nil
// }

// // DetectorContinuousTrigger 连续采集：调用DetectorSingleScan，等采集获取图片后再调用一次直到传入false，不保存图片
// func (this *DiffractDevice) DetectorContinuousTrigger(enable bool) error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	fmt.Println("连续触发", enable)
// 	if enable {
// 		this.CollectMode = 2 // 设置为连续采集模式
// 		return this.DetectorSingleScan()
// 	} else {
// 		this.CollectMode = 0 // 停止连续采集
// 		return nil
// 	}
// }

// // ****************************************   CT联测   ***************************************
// // StartCTScan CT采集：先调用轴运动，再调用DetectorSingleScan获取一张图片，等图片获取后再调用一次运动和采集，对采集的图片进行计数直到采集完成或调用StopCTScan，需要保存图片
// func (this *DiffractDevice) StartCTScan() error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	fmt.Printf("[CT扫描] 开始扫描到路径: %s\n", this.project.FilePath)
// 	// 确保输出目录存在
// 	if err := os.MkdirAll(this.project.FilePath, 0755); err != nil {
// 		fmt.Printf("[CT扫描] 创建输出目录失败: %v\n", err)
// 		return fmt.Errorf("创建输出目录失败: %v", err)
// 	}

// 	if this.CollectMode == 3 {
// 		return fmt.Errorf("CT扫描已经在运行中")
// 	}

// 	this.CollectMode = 3 // 设置为CT采集模式
// 	this.CTImageCount = 0

// 	runtime.EventsEmit(this.ctx, "CTScanStatus", "Running")

// 	this.runCTScanCycle()

// 	return nil
// }

// // runCTScanCycle CT扫描循环：运动 -> 采集
// func (this *DiffractDevice) runCTScanCycle() {
// 	if !this.DetectorIsConnected() {
// 		return
// 	}

// 	// 检查是否停止或完成
// 	if this.CollectMode != 3 {
// 		fmt.Printf("[CT扫描] 扫描已停止")
// 		runtime.EventsEmit(this.ctx, "CTScanStatus", "Paused")
// 		return
// 	}
// 	if this.CTImageCount >= this.project.CTProject.AcquisitionNum {
// 		fmt.Printf("[CT扫描] 已扫描完成")
// 		runtime.EventsEmit(this.ctx, "CTScanStatus", "Completed")
// 		this.CollectMode = 0
// 		return
// 	}

// 	err := this.Stage.DiffractStagesMOVE("R", this.project.CTProject.AngularStep)
// 	if err != nil {
// 		fmt.Printf("[CT扫描] 运动R轴失败: %v\n", err)
// 		return
// 	}

// 	waitTime := this.project.CTProject.AngularStep / float32(this.system.StageConfig.RAxis.Speed)
// 	fmt.Printf("[CT扫描] 等待R轴运动完成, 预计时间: %.5f秒\n", waitTime)
// 	time.Sleep(time.Duration(waitTime)*time.Second + 30*time.Millisecond) // 等待R轴运动完成,50ms等待

// 	fmt.Printf("[CT扫描] 运动R轴完成, 当前R位置: %.2f°, 已采集: %d/%d\n", this.Stage.motors.RAxis.Position, this.CTImageCount, this.project.CTProject.AcquisitionNum)

// 	// 执行单次扫描
// 	err = this.DetectorSingleScan()
// 	if err != nil {
// 		fmt.Printf("[CT扫描] 扫描失败: %v\n", err)
// 		return
// 	}
// }

// // StopCTScan 停止CT扫描
// func (this *DiffractDevice) StopCTScan() error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	this.CollectMode = 0 // 停止所有采集模式
// 	fmt.Printf("[CT扫描] 停止扫描, 总扫描次数: %d\n", this.CTImageCount)
// 	return nil
// }

// // DetectorSetExposeTime 设置曝光时间
// func (this *DiffractDevice) DetectorSetExposeTime(exposeTime int) error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	fmt.Println("设置曝光时间", exposeTime)
// 	if !this.CT.COM_SetExposeTime(exposeTime) {
// 		fmt.Println("无法设置曝光时间")
// 		return fmt.Errorf("无法设置曝光时间")
// 	}
// 	this.status.ExposeTime = exposeTime
// 	fmt.Println("设置曝光时间成功")
// 	return nil
// }

// // DetectorGetExposeTime 获取曝光时间
// func (this *DiffractDevice) DetectorGetExposeTime() int {
// 	if !this.DetectorIsConnected() {
// 		return 0
// 	}

// 	fmt.Println("获取曝光时间")
// 	this.status.ExposeTime = this.CT.COM_GetExposeTime()
// 	return this.status.ExposeTime
// }

// // DetectorSetBinning 设置像素数
// func (this *DiffractDevice) DetectorSetBinning(binning string) error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	fmt.Println("设置像素数", binning)
// 	if !this.CT.COM_SetBinning(binning) {
// 		fmt.Println("无法设置像素数", binning)
// 		return fmt.Errorf("无法设置像素数")
// 	}
// 	this.status.Binning = binning
// 	fmt.Println("设置像素数成功")
// 	return nil
// }

// // DetectorGetBinning 获取像素数
// func (this *DiffractDevice) DetectorGetBinning() string {
// 	if !this.DetectorIsConnected() {
// 		return ""
// 	}

// 	fmt.Println("获取像素数")
// 	this.status.Binning = this.CT.COM_GetBinning()
// 	return this.status.Binning
// }

// // DetectorSetRepeat 设置重复次数
// func (this *DiffractDevice) DetectorSetRepeat(repeat int) error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	return nil
// }

// // DetectorGetRepeat 获取重复次数
// func (this *DiffractDevice) DetectorGetRepeat() int {
// 	if !this.DetectorIsConnected() {
// 		return 0
// 	}
// 	return 0
// }

// // DetectorSetGain 设置增益
// func (this *DiffractDevice) DetectorSetGain(gain int) error {
// 	if !this.DetectorIsConnected() {
// 		return fmt.Errorf("未连接探测器")
// 	}

// 	return nil
// }

// // DetectorGetGain 获取增益
// func (this *DiffractDevice) DetectorGetGain() int {
// 	if !this.DetectorIsConnected() {
// 		return 0
// 	}
// 	return 0
// }

// // handleCTImageEvent 处理ct_raw事件数据
// func (this *DiffractDevice) handleCTImageEvent(data interface{}) {
// 	// 解析事件数据
// 	rawData, ok := data.(map[string]interface{})
// 	if !ok {
// 		fmt.Printf("[DiffractDevice] 数据类型转换失败，期望map[string]interface{}类型，实际类型: %T\n", data)
// 		return
// 	}

// 	// 保存原始数据到rawList
// 	this.rawList = append(this.rawList, rawData)
// 	this.currentImageIndex = len(this.rawList) - 1

// 	// 使用统一的归一化函数处理图片并发送
// 	normalizedImg, err := this.normalizeRawImage(rawData)
// 	if err != nil {
// 		fmt.Printf("[DiffractDevice] 归一化图片失败: %v\n", err)
// 		return
// 	}

// 	// 发送ct_image事件供前端显示
// 	runtime.EventsEmit(this.ctx, "ct_image", normalizedImg)

// 	switch this.CollectMode {
// 	case 1: // 单次采集：保存图片为backup，然后重置模式
// 		switch this.system.ImageConfig.SaveFormat {
// 		case "raw":
// 			this.SaveRawToLocal(rawData)
// 		case "tiff":
// 			this.SaveTiffToLocal(rawData)
// 		}
// 		this.CollectMode = 0 // 重置为空闲模式
// 	case 2: // 连续采集：不保存图片，触发下一次采集
// 		this.DetectorSingleScan()
// 	case 3: // CT采集：保存图片，计数，触发下一次运动+采集循环
// 		this.CTImageCount++
// 		switch this.system.ImageConfig.SaveFormat {
// 		case "raw":
// 			this.SaveRawToLocal(rawData)
// 		case "tiff":
// 			this.SaveTiffToLocal(rawData)
// 		}
// 		this.runCTScanCycle()
// 	}
// }

// func (this *DiffractDevice) normalizeRawImage(rawData map[string]interface{}) (map[string]interface{}, error) {
// 	imageData, ok := rawData["image"].([]byte)
// 	if !ok {
// 		return nil, fmt.Errorf("图片数据类型转换失败")
// 	}

// 	width, _ := rawData["width"].(int)
// 	height, _ := rawData["height"].(int)
// 	if width <= 0 || height <= 0 {
// 		return nil, fmt.Errorf("图片尺寸无效: width=%d, height=%d", width, height)
// 	}

// 	grayImg := image.NewGray(image.Rect(0, 0, width, height))

// 	rangeVal := this.CTImageMaxVal - this.CTImageMinVal
// 	if rangeVal == 0 {
// 		rangeVal = 1
// 	}

// 	// 计算直方图（0-10000范围，100个bin）
// 	const histogramBins = 100
// 	histogram := make([]int, histogramBins)

// 	for i := 0; i < width*height; i++ {
// 		val := uint32(imageData[i*2]) | (uint32(imageData[i*2+1]) << 8)
// 		var normalized uint32
// 		if val <= this.CTImageMinVal {
// 			normalized = 0
// 		} else if val >= this.CTImageMaxVal {
// 			normalized = 255
// 		} else {
// 			normalized = (val - this.CTImageMinVal) * 255 / rangeVal
// 		}
// 		grayImg.Pix[i] = uint8(normalized)

// 		// 统计直方图：将16位像素值映射到0-10000范围并分桶
// 		histVal := uint32(float64(val) / 65535.0 * 10000.0)
// 		if histVal >= 10000 {
// 			histVal = 9999
// 		}
// 		binIndex := int(histVal / 100)
// 		histogram[binIndex]++
// 	}

// 	buf := new(bytes.Buffer)
// 	if err := jpeg.Encode(buf, grayImg, &jpeg.Options{Quality: 60}); err != nil {
// 		return nil, err
// 	}

// 	encodedStr := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

// 	return map[string]interface{}{
// 		"image":     encodedStr,
// 		"width":     width,
// 		"height":    height,
// 		"histogram": histogram,
// 	}, nil
// }

// func (this *DiffractDevice) SetCTNormalizeRange(minVal uint32, maxVal uint32) {
// 	this.CTImageMinVal = uint32(float64(minVal) / 100.0 * 10000.0)
// 	this.CTImageMaxVal = uint32(float64(maxVal) / 100.0 * 10000.0)
// 	fmt.Printf("[DiffractDevice] 设置对比度范围: min=%d, max=%d\n", this.CTImageMinVal, this.CTImageMaxVal)

// 	// 只重新归一化当前显示的图片
// 	if len(this.rawList) > 0 && this.currentImageIndex >= 0 && this.currentImageIndex < len(this.rawList) {
// 		normalizedImg, err := this.normalizeRawImage(this.rawList[this.currentImageIndex])
// 		if err != nil {
// 			fmt.Printf("[DiffractDevice] 重新归一化当前图片失败: %v\n", err)
// 			return
// 		}
// 		runtime.EventsEmit(this.ctx, "ct_image_refresh", normalizedImg)
// 		fmt.Printf("[DiffractDevice] 已重新归一化第%d张图片并通知前端刷新\n", this.currentImageIndex)
// 	}
// }

// func (this *DiffractDevice) GetCTImageFromList(index int) {
// 	fmt.Println("获取CT图片", index)
// 	if index < 0 || index >= len(this.rawList) {
// 		fmt.Println("索引超出范围")
// 		return
// 	}
// 	this.currentImageIndex = index
// 	// 从rawList获取并归一化后发送
// 	normalizedImg, err := this.normalizeRawImage(this.rawList[index])
// 	if err != nil {
// 		fmt.Printf("[DiffractDevice] 归一化图片失败: %v\n", err)
// 		return
// 	}
// 	runtime.EventsEmit(this.ctx, "ct_image_list", normalizedImg)
// }

// func (this *DiffractDevice) ClearCTImageList() {
// 	this.rawList = []map[string]interface{}{}
// 	this.currentImageIndex = -1
// }

// func (this *DiffractDevice) GetPixelValue(imageIndex int, x int, y int) {
// 	if imageIndex < 0 || imageIndex >= len(this.rawList) {
// 		fmt.Println("[DiffractDevice] 图像索引超出范围")
// 		return
// 	}

// 	rawData := this.rawList[imageIndex]
// 	imageData, ok := rawData["image"].([]byte)
// 	if !ok {
// 		fmt.Println("[DiffractDevice] 图片数据类型转换失败")
// 		return
// 	}

// 	width, _ := rawData["width"].(int)
// 	height, _ := rawData["height"].(int)

// 	if x < 0 || x >= width || y < 0 || y >= height {
// 		fmt.Printf("[DiffractDevice] 坐标超出范围: x=%d, y=%d, width=%d, height=%d\n", x, y, width, height)
// 		return
// 	}

// 	offset := (y*width + x) * 2
// 	if offset+1 >= len(imageData) {
// 		fmt.Println("[DiffractDevice] 数据偏移超出范围")
// 		return
// 	}

// 	pixelValue := uint32(imageData[offset]) | (uint32(imageData[offset+1]) << 8)

// 	runtime.EventsEmit(this.ctx, "ct_pixel_value", map[string]interface{}{
// 		"value": int(pixelValue),
// 		"x":     x,
// 		"y":     y,
// 	})
// }

// // rotateRawImage 旋转16位RAW图像
// func (this *DiffractDevice) rotateRawImage(rawData map[string]interface{}) (map[string]interface{}, error) {
// 	if this.system == nil || this.system.ImageConfig.RotateAngle == 0 {
// 		return rawData, nil
// 	}

// 	imageData, ok := rawData["image"].([]byte)
// 	if !ok {
// 		return nil, fmt.Errorf("图片数据类型转换失败")
// 	}

// 	width, _ := rawData["width"].(int)
// 	height, _ := rawData["height"].(int)
// 	if width <= 0 || height <= 0 {
// 		return nil, fmt.Errorf("图片尺寸无效")
// 	}

// 	rotateAngle := this.system.ImageConfig.RotateAngle
// 	var newWidth, newHeight int
// 	var rotatedData []byte

// 	switch rotateAngle {
// 	case -90:
// 		newWidth = height
// 		newHeight = width
// 		rotatedData = make([]byte, width*height*2)
// 		for y := 0; y < height; y++ {
// 			for x := 0; x < width; x++ {
// 				srcIdx := (y*width + x) * 2
// 				dstIdx := ((width-1-x)*height + y) * 2
// 				rotatedData[dstIdx] = imageData[srcIdx]
// 				rotatedData[dstIdx+1] = imageData[srcIdx+1]
// 			}
// 		}
// 	case 90:
// 		newWidth = height
// 		newHeight = width
// 		rotatedData = make([]byte, width*height*2)
// 		for y := 0; y < height; y++ {
// 			for x := 0; x < width; x++ {
// 				srcIdx := (y*width + x) * 2
// 				dstIdx := (x*height + (height - 1 - y)) * 2
// 				rotatedData[dstIdx] = imageData[srcIdx]
// 				rotatedData[dstIdx+1] = imageData[srcIdx+1]
// 			}
// 		}
// 	case 180:
// 		newWidth = width
// 		newHeight = height
// 		rotatedData = make([]byte, width*height*2)
// 		for i := 0; i < width*height; i++ {
// 			srcIdx := i * 2
// 			dstIdx := (width*height - 1 - i) * 2
// 			rotatedData[dstIdx] = imageData[srcIdx]
// 			rotatedData[dstIdx+1] = imageData[srcIdx+1]
// 		}
// 	default:
// 		return rawData, nil
// 	}

// 	return map[string]interface{}{
// 		"image":  rotatedData,
// 		"width":  newWidth,
// 		"height": newHeight,
// 	}, nil
// }

// // SaveRawToLocal 保存原始图片到本地
// func (this *DiffractDevice) SaveRawToLocal(data interface{}) error {
// 	fmt.Println("保存原始图片到本地")
// 	rawData, ok := data.(map[string]interface{})
// 	if !ok {
// 		fmt.Printf("[DiffractDevice] 数据类型转换失败，期望map[string]interface{}类型，实际类型: %T\n", data)
// 		return fmt.Errorf("数据类型转换失败")
// 	}

// 	// 检查是否需要旋转
// 	var err error
// 	rawData, err = this.rotateRawImage(rawData)
// 	if err != nil {
// 		fmt.Printf("[DiffractDevice] 旋转图片失败: %v\n", err)
// 		return err
// 	}

// 	// 创建项目文件夹：FilePath/FileName/
// 	projectDir := this.project.FilePath + "/" + this.project.FileName
// 	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
// 		os.MkdirAll(projectDir, 0755)
// 	}

// 	// 在项目文件夹下保存文件
// 	filePath := projectDir + "/" + this.project.FileName + "_" + fmt.Sprintf("%d", this.CTImageCount) + ".raw"

// 	if this.CollectMode == 1 {
// 		filePath = projectDir + "/" + this.project.FileName + "_backup" + ".raw"
// 		suffix := 1
// 		for {
// 			if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 				break
// 			}
// 			filePath = projectDir + "/" + this.project.FileName + "_backup(" + fmt.Sprintf("%d", suffix) + ").raw"
// 			suffix++
// 		}
// 	}

// 	// 将原始16位图片数据写入文件
// 	err = os.WriteFile(filePath, rawData["image"].([]byte), 0644)
// 	if err != nil {
// 		fmt.Printf("[DiffractDevice] 写入文件失败: %v\n", err)
// 	} else {
// 		fmt.Printf("[DiffractDevice] 图片已写入 %s，大小: %d 字节\n", filePath, len(rawData["image"].([]byte)))
// 	}
// 	return nil
// }

// // SaveTiffToLocal 保存Tiff图片到本地
// func (this *DiffractDevice) SaveTiffToLocal(data interface{}) error {
// 	fmt.Println("保存Tiff图片到本地")
// 	tiffData, ok := data.(map[string]interface{})
// 	if !ok {
// 		fmt.Printf("[DiffractDevice] 数据类型转换失败，期望map[string]interface{}类型，实际类型: %T\n", data)
// 		return fmt.Errorf("数据类型转换失败")
// 	}

// 	// 检查是否需要旋转
// 	var err error
// 	tiffData, err = this.rotateRawImage(tiffData)
// 	if err != nil {
// 		fmt.Printf("[DiffractDevice] 旋转图片失败: %v\n", err)
// 		return err
// 	}

// 	// 创建项目文件夹：FilePath/FileName/
// 	projectDir := this.project.FilePath + "/" + this.project.FileName
// 	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
// 		os.MkdirAll(projectDir, 0755)
// 	}

// 	// 在项目文件夹下保存文件
// 	filePath := projectDir + "/" + this.project.FileName + "_" + fmt.Sprintf("%d", this.CTImageCount) + ".tiff"

// 	if this.CollectMode == 1 {
// 		filePath = projectDir + "/" + this.project.FileName + "_backup" + ".tiff"
// 		suffix := 1
// 		for {
// 			if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 				break
// 			}
// 			filePath = projectDir + "/" + this.project.FileName + "_backup(" + fmt.Sprintf("%d", suffix) + ").tiff"
// 			suffix++
// 		}
// 	}

// 	// 获取图像数据
// 	imgBytes := tiffData["image"].([]byte)
// 	width, _ := tiffData["width"].(int)
// 	height, _ := tiffData["height"].(int)

// 	// 将字节数据转换为 16 位灰度图像
// 	img := image.NewGray16(image.Rect(0, 0, width, height))
// 	for i := 0; i < len(imgBytes) && i < width*height*2; i += 2 {
// 		// 小端序：低字节在前，高字节在后
// 		pixel := uint16(imgBytes[i]) | uint16(imgBytes[i+1])<<8
// 		idx := i / 2
// 		img.SetGray16(idx%width, idx/width, color.Gray16{Y: pixel})
// 	}

// 	// 创建 TIFF 文件
// 	tiffFile, err := os.Create(filePath)
// 	if err != nil {
// 		fmt.Printf("[DiffractDevice] 创建 TIFF 文件失败: %v\n", err)
// 		return err
// 	}
// 	defer tiffFile.Close()

// 	// 编码为 TIFF 格式
// 	err = tiff.Encode(tiffFile, img, nil)
// 	if err != nil {
// 		fmt.Printf("[DiffractDevice] TIFF 编码失败: %v\n", err)
// 		return err
// 	}

// 	fmt.Printf("[DiffractDevice] TIFF 图片已写入 %s，大小: %dx%d\n", filePath, width, height)
// 	return nil
// }
