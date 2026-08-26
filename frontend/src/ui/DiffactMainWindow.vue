<template>
    <div class="main-window">
        <div class="title-bar">
            <div class="title">NIMTE 衍射仪作业系统</div>
            <div class='window-actions'>
                    <button class="window-btn connect-btn" type="button" title="设备连接" @click="deviceConnectVisible = true">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
                            <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
                        </svg>
                    </button>
                    <button class="window-btn settings-btn" type="button" title="系统设置" @click="systemSettingsVisible = true">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="3"/>
                            <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06A1.65 1.65 0 0 0 4.6 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06A1.65 1.65 0 0 0 9 4.6a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06A1.65 1.65 0 0 0 19.4 9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/>
                        </svg>
                    </button>
                    <span class="title-divider"></span>
                    <button class="window-btn minimize-btn" type="button" title="最小化" @click="WindowMinimise">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <line x1="4" y1="12" x2="22" y2="12"/>
                        </svg>
                    </button>
                    <button class="window-btn maximize-btn" type="button" title="最大化" @click="WindowToggleMaximise">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <rect x="5" y="5" width="17" height="17" rx="2"/>
                        </svg>
                    </button>
                    <button class="window-btn close-btn" type="button" title="关闭" @click="Quit">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <line x1="20" y1="4" x2="4" y2="20"/>
                            <line x1="4" y1="4" x2="20" y2="20"/>
                        </svg>
                    </button>
                </div>
        </div>
        <div class="content">
            <div class="left-panel">
                <!-- 卡片1：设备信息 -->
                <div class="info-card">
                    <div class="card-header">
                        <span class="card-title">设备信息</span>
                    </div>
                    <div class="card-body">
                        <div class="info-row">
                            <span class="info-label">控制主机</span>
                            <span class="info-value" :class="{ highlight: Status.Stage.running }">
                                {{ Status.Stage.running ? '已连接' : '未连接' }}
                            </span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">高压电源</span>
                            <span class="info-value" :class="{ highlight: Status.Power.running }">
                                {{ Status.Power.running ? '已开启' : '未开启' }}
                            </span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">探测器</span>
                            <span class="info-value" :class="{ highlight: Status.Detector.running }">
                                {{ Status.Detector.running ? '已连接' : '未连接' }}
                            </span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">系统状态</span>
                            <span class="info-value" :class="{ danger: Status.SystemStatus !='正常' }">
                                {{ Status.SystemStatus }}
                            </span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">运行时间</span>
                            <span class="info-value">{{ Status.RunTime }}</span>
                        </div>
                    </div>
                </div>

                <!-- 卡片2：放射源状态 -->
                <div class="info-card">
                    <div class="card-header">
                        <span class="card-title">放射源状态</span>
                        <!-- <svg class="radiation-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="2.5" fill="currentColor" stroke="none"/>
                            <path d="M12 9.5a2.5 2.5 0 0 0-2.17-3.75A7 7 0 0 0 5 11"/>
                            <path d="M12 14.5a2.5 2.5 0 0 1-2.17 3.75A7 7 0 0 0 12 19a7 7 0 0 0 2.17-.75A2.5 2.5 0 0 1 12 14.5z"/>
                            <path d="M14.17 12a2.5 2.5 0 0 0 1.08-3.25A7 7 0 0 1 19 11a7 7 0 0 1-.43 2.5A2.5 2.5 0 0 1 14.17 12z"/>
                            <path d="M9.83 12a2.5 2.5 0 0 1-1.08 3.25A7 7 0 0 1 5 13a7 7 0 0 1 .43-2.5A2.5 2.5 0 0 1 9.83 12z"/>
                        </svg> -->
                    </div>
                    <div class="card-body">
                        <div class="info-row">
                            <span class="info-label">灯丝电流</span>
                            <span class="info-value" :class="{ highlight: Status.Power.FI > 0.1, 'alarm-value': Status.Power.FI > 0.1 }">{{ Status.Power.FI.toFixed(2) }} A</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">灯丝预热电压</span>
                            <span class="info-value" :class="{ highlight: Status.Power.FV > 0.1, 'alarm-value': Status.Power.FV > 0.1 }">{{ Status.Power.FV.toFixed(2) }} V</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">高压电压</span>
                            <span class="info-value" :class="{ highlight: Status.Power.HV > 0.1, 'alarm-value': Status.Power.HV > 0.1 }">{{ Status.Power.HV.toFixed(2) }} kV</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">高压电流</span>
                            <span class="info-value" :class="{ highlight: Status.Power.HI > 0.1, 'alarm-value': Status.Power.HI > 0.1 }">{{ Status.Power.HI.toFixed(2) }} uA</span>
                        </div>
                    </div>
                </div>

                <!-- 卡片3：位移台信息 -->
                <div class="info-card">
                    <div class="card-header">
                        <span class="card-title">位移台信息</span>
                    </div>
                    <div class="card-body">
                        <div class="info-row">
                            <span class="info-label">X 位置</span>
                            <span class="info-value highlight" :class="{ moving: Status.Stage.X_Status === 'Moving' }">{{ Status.Stage.X.toFixed(4) }} mm</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Y 位置</span>
                            <span class="info-value highlight" :class="{ moving: Status.Stage.Y_Status === 'Moving' }">{{ Status.Stage.Y.toFixed(4) }} mm</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Z 位置</span>
                            <span class="info-value highlight" :class="{ moving: Status.Stage.Z_Status === 'Moving' }">{{ Status.Stage.Z.toFixed(4) }} mm</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">R 角度</span>
                            <span class="info-value highlight" :class="{ moving: Status.Stage.R_Status === 'Moving' }">{{ Status.Stage.R.toFixed(4) }} °</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">XX 位置</span>
                            <span class="info-value highlight" :class="{ moving: Status.Stage.XX_Status === 'Moving' }">{{ Status.Stage.XX.toFixed(4) }} mm</span>
                        </div>
                    </div>
                </div>

                <!-- 卡片4：探测器状态 -->
                <div class="info-card">
                    <div class="card-header">
                        <span class="card-title">探测器状态</span>
                    </div>
                    <div class="card-body">
                        <div class="info-row">
                            <span class="info-label">探测器 SN</span>
                            <span class="info-value">{{ Status.Detector.sn }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">当前模式</span>
                            <span class="info-value highlight">{{ Status.Detector.mode === 'null' ? '—' : Status.Detector.mode }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">图像尺寸</span>
                            <span class="info-value highlight">{{ Status.Detector.width }} × {{ Status.Detector.height }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">曝光时间</span>
                            <span class="info-value highlight">{{ Status.Detector.exposureTime}} ms</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">增益</span>
                            <span class="info-value highlight">{{ Status.Detector.gain }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Binning</span>
                            <span class="info-value highlight">{{ Status.Detector.binning }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">重复次数</span>
                            <span class="info-value highlight">{{ Status.Detector.repeatTimes === 0 ? '无限' : Status.Detector.repeatTimes }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">表面温度</span>
                            <span class="info-value">{{ Status.Detector.tempreture.toFixed(1) }} °C</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">表面湿度</span>
                            <span class="info-value">{{ Status.Detector.humidity.toFixed(1) }} %</span>
                        </div>
                    </div>
                </div>

            </div>

            <div class="middle-panel">
                <!-- 1. 图像操作栏（高 50px） -->
                <div class="image-toolbar">
                    <div class="toolbar-left">
                        <span class="toolbar-label">图像总数</span>
                        <span class="toolbar-count">{{ imagesCounts}}</span>
                        <span class="toolbar-label" style="margin-left: 14px;">图像 ID</span>
                        <select
                            class="toolbar-select"
                            v-model="selectedImageId"
                            @change="handleSelectImageByID"
                            :disabled="imagesCounts === 0"
                        >
                            <option
                                v-for="id in imagesCounts"
                                :key="id"
                                :value="id"
                            >#第{{ id }}张图片</option>
                        </select>
                    </div>
                    <div class="toolbar-right">
                        <button
                            class="toolbar-btn save-btn"
                            type="button"
                            @click="handleSaveImage"
                            :disabled="selectedImageId === 0"
                        >
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/>
                                <polyline points="17 21 17 13 7 13 7 21"/>
                                <polyline points="7 3 7 8 15 8"/>
                            </svg>
                            <span>保存</span>
                        </button>
                        <button
                            class="toolbar-btn delete-btn"
                            type="button"
                            @click="handleDeleteImage"
                            :disabled="selectedImageId === 0"
                        >
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <polyline points="3 6 5 6 21 6"/>
                                <path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"/>
                                <path d="M10 11v6M14 11v6"/>
                                <path d="M9 6V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2"/>
                            </svg>
                            <span>删除</span>
                        </button>
                    </div>
                </div>

                <!-- 2. 图像显示区域（flex:1） -->
                <div class="image-viewer">
                    <div v-if="!DiffractImage" class="empty-viewer">
                        <svg class="empty-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                            <rect x="3" y="3" width="18" height="18" rx="2"/>
                            <circle cx="12" cy="12" r="6"/>
                            <circle cx="12" cy="12" r="2"/>
                        </svg>
                        <div class="empty-text">等待图像数据...</div>
                    </div>
                    <img
                        v-else
                        class="diffract-image"
                        :src="DiffractImage"
                        alt="探测器图像"
                        :style="{ transform: `rotate(${imageRotateAngle}deg)` }"
                        @error="onImageError"
                    />
                </div>

                <!-- 3. 图像处理栏（高 100px，横向分栏目） -->
                <div class="process-toolbar">
                    <!-- 去模糊 -->
                    <div class="process-group">
                        <div class="group-title">去模糊</div>
                        <div class="group-btns">
                            <button class="proc-btn" type="button" @click="handleProcess('deblur', 'RL')">Richardson Lucy</button>
                            <button class="proc-btn" type="button" @click="handleProcess('deblur', 'Wiener')">Wiener</button>
                        </div>
                    </div>

                    <!-- 降噪 -->
                    <div class="process-group">
                        <div class="group-title">降噪</div>
                        <div class="group-btns">
                            <button class="proc-btn" type="button" @click="handleProcess('denoise', 'AMF')">AMF</button>
                            <button class="proc-btn" type="button" @click="handleProcess('denoise', 'BM3D')">BM3D</button>
                            <button class="proc-btn" type="button" @click="handleProcess('denoise', 'Bilateral')">Bilateral</button>
                            <button class="proc-btn" type="button" @click="handleProcess('denoise', 'TV')">TV</button>
                            <button class="proc-btn" type="button" @click="handleProcess('denoise', 'Wavelet')">Wavelet</button>
                            <button class="proc-btn" type="button" @click="handleProcess('denoise', 'NLM')">NLM</button>
                            <button class="proc-btn" type="button" @click="handleProcess('denoise', 'Filter')">Filter</button>
                        </div>
                    </div>

                    <!-- 锐化 -->
                    <div class="process-group">
                        <div class="group-title">锐化</div>
                        <div class="group-btns">
                            <button class="proc-btn" type="button" @click="handleProcess('sharpen', 'Laplace')">Laplace</button>
                            <button class="proc-btn" type="button" @click="handleProcess('sharpen', 'USM')">USM</button>
                        </div>
                    </div>

                    <!-- 图像调整 -->
                    <div class="process-group">
                        <div class="group-title">图像调整</div>
                        <div class="group-btns">
                            <button class="proc-btn" type="button" @click="handleProcess('adjust', 'Gamma')">Gamma变化</button>
                            <button class="proc-btn" type="button" @click="handleProcess('adjust', 'Log')">Log变换</button>
                            <button class="proc-btn" type="button" @click="handleProcess('adjust', 'Exp')">指数变换</button>
                        </div>
                    </div>

                    <!-- 图像归一化 -->
                    <div class="process-group">
                        <div class="group-title">图像归一化</div>
                        <div class="group-btns">
                            <input class="proc-btn" type="text" @keyup.enter="handleProcess('normalize', 'ImageMinVal')" placeholder="输入最小值">
                            <input class="proc-btn" type="text" @keyup.enter="handleProcess('normalize', 'ImageMaxVal')" placeholder="输入最大值">
                        </div>
                    </div>

                </div>

                
            </div>
            
            <div class="right-panel">
                <!-- 1. 放射源控制 -->
                <div class="ctrl-card">
                    <div class="ctrl-header">
                        <span class="ctrl-title">放射源控制</span>
                    </div>

                    
                    <div class="ctrl-body">
                        <div class="ctrl-row">
                            <span class="ctrl-label">灯丝开关</span>
                            <label class="switch">
                                <input
                                    type="checkbox"
                                    :checked="hvpsSet.FilamentSwitch"
                                    :disabled="hvpsBusy"
                                    @click.prevent="handleFilamentSwitch"
                                />
                                <span class="slider"></span>
                            </label>
                            <span class="ctrl-label">放射开关</span>
                            <label class="switch">
                                <input
                                    type="checkbox"
                                    :checked="hvpsSet.PowerSwitch"
                                    :disabled="hvpsBusy"
                                    @click.prevent="handlePowerSwitch"
                                />
                                <span class="slider"></span>
                            </label>
                        </div>
                        

                        <div class="ctrl-divider"></div>

                        <div class="ctrl-row">
                            <span class="ctrl-label">灯丝电流</span>
                            <div class="num-input-wrap">
                                <input
                                    class="num-input"
                                    type="number"
                                    step="0.01"
                                    min="0"
                                    :disabled="hvpsBusy"
                                    v-model.number="hvpsSet.SI"
                                    @keyup.enter="handleSetFilamentPreheat"
                                />
                                <span class="num-unit">A</span>
                            </div>
                        </div>
                        <div class="ctrl-row">
                            <span class="ctrl-label">灯丝电流限制</span>
                            <div class="num-input-wrap">
                                <input
                                    class="num-input"
                                    type="number"
                                    step="0.01"
                                    min="0"
                                    :disabled="hvpsBusy"
                                    v-model.number="hvpsSet.LI"
                                    @keyup.enter="handleSetFilamentLimit"
                                />
                                <span class="num-unit">A</span>
                            </div>
                        </div>

                        <div class="ctrl-divider"></div>

                        <div class="ctrl-row">
                            <span class="ctrl-label">高压电压</span>
                            <div class="num-input-wrap">
                                <input
                                    class="num-input"
                                    type="number"
                                    step="0.1"
                                    min="0"
                                    max="60"
                                    :disabled="hvpsBusy"
                                    v-model.number="hvpsSet.HV"
                                    @keyup.enter="handleHVPSSetHV"
                                />
                                <span class="num-unit">kV</span>
                            </div>
                        </div>
                        <div class="ctrl-row">
                            <span class="ctrl-label">高压电流</span>
                            <div class="num-input-wrap">
                                <input
                                    class="num-input"
                                    type="number"
                                    step="0.1"
                                    min="0"
                                    :disabled="hvpsBusy"
                                    v-model.number="hvpsSet.HI"
                                    @keyup.enter="handleHVPSSetHI"
                                />
                                <span class="num-unit">uA</span>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- 2. 运动控制 -->
                <div class="ctrl-card">
                    <div class="ctrl-header">
                        <span class="ctrl-title">运动控制</span>
                    </div>
                    <div class="ctrl-body">
                        <div
                            v-for="axis in ['X', 'Y', 'Z', 'R', 'XX']"
                            :key="axis"
                            class="axis-row"
                        >
                            <span class="axis-label">{{ axis }}</span>
                            <div class="axis-value-wrap">
                                <input
                                    class="num-input axis-input"
                                    type="number"
                                    :step="axis === 'R' ? 0.01 : 0.001"
                                    v-model.number="motionTargets[axis]"
                                />
                                <span class="num-unit">{{ axis === 'R' ? '°' : 'mm' }}</span>
                            </div>
                            <div class="axis-btns">
                                <button
                                    class="axis-btn cw"
                                    type="button"
                                    :disabled="stageBusy"
                                    @click="handleAxisCW(axis, motionTargets[axis])"
                                >CW</button>
                                <button
                                    class="axis-btn ccw"
                                    type="button"
                                    :disabled="stageBusy"
                                    @click="handleAxisCCW(axis, motionTargets[axis])"
                                >CCW</button>
                            </div>
                        </div>

                        <div class="stop-row">
                            <button
                                class="stop-btn"
                                type="button"
                                :disabled="stageBusy"
                                @click="handleAllStop"
                            >
                                <svg viewBox="0 0 24 24" fill="currentColor" width="14" height="14"><rect x="6" y="6" width="12" height="12" rx="1.5"/></svg>
                                <span>全部停止</span>
                            </button>
                        </div>
                    </div>
                </div>

                <!-- 3. 探测器控制 -->
                <div class="ctrl-card">
                    <div class="ctrl-header">
                        <span class="ctrl-title">探测器控制</span>
                    </div>
                    <div class="ctrl-body">
                        <div class="detector-row">
                            <button
                                class="det-btn params-btn"
                                type="button"
                                @click="handleDetectorParams"
                            >
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                    <circle cx="12" cy="12" r="3"/>
                                    <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06A1.65 1.65 0 0 0 4.6 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06A1.65 1.65 0 0 0 9 4.6a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06A1.65 1.65 0 0 0 19.4 9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/>
                                </svg>
                                <span>参数设置</span>
                            </button>
                        </div>
                        <div class="detector-row">
                            <button
                                class="det-btn capture-btn"
                                type="button"

                                @click="handleDetectorCapture"
                            >
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                    <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                                    <circle cx="12" cy="13" r="4"/>
                                </svg>
                                <span>拍摄</span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            
        </div>
        <div class="footer">
            <div class="footer-item-title">版本</div>
            <div class="footer-item-value">1.0.0</div>
            <div class="footer-item-title">作者</div>
            <div class="footer-item-value">TangLifa</div>
        </div>
    </div>
    
    <transition name="toast-fade">
        <div v-if="toastMsg" class="toast-hint">{{ toastMsg }}</div>
    </transition>

    <DeviceConnectModal :visible="deviceConnectVisible" @close="deviceConnectVisible = false" />

    <DetectorParamsModal
        :visible="detectorParamsVisible"
        :exposure-time="Status.Detector.exposureTime"
        :gain="Status.Detector.gain"
        :binning="Status.Detector.binning"
        :repeat-times="Status.Detector.repeatTimes"
        @close="detectorParamsVisible = false"
        @saved="(form) => {
            Status.Detector.exposureTime = form.exposure
            Status.Detector.gain = form.gain
            Status.Detector.binning = form.binning
            Status.Detector.repeatTimes = form.repeatTimes
            showToast('探测器参数已保存')
        }"
    />

    <SystemAlarmModal
        :visible="systemAlarmVisible"
        @close="systemAlarmVisible = false"
    />

</template>

<script setup>
import DeviceConnectModal from './modal/DeviceConnectModal.vue'
import DetectorParamsModal from './pannel/DetectorParamsModal.vue'
import SystemAlarmModal from './modal/SystemAlarmModal.vue'

import { reactive, ref, onMounted } from 'vue'
import { WindowMinimise, WindowToggleMaximise, Quit,EventsOn,} from '../../wailsjs/runtime/runtime'
import { HVPSSourceOpen, HVPSSetFilamentOpen, HVPSSetFilamentPreheat, HVPSSetFilamentLimit, HVPSSetHV, HVPSSetHI } from '../../wailsjs/go/components/HVPSService'
import { StageStop, StageRelMove } from '../../wailsjs/go/components/StageService'
import { DetectorCapture, CallImageByID, DelImageByID, SaveImageByID } from '../../wailsjs/go/components/DetectorService'


// 设备连接模态框显隐
const deviceConnectVisible = ref(true)
// 系统设置模态框显隐
const systemSettingsVisible = ref(false)
// 探测器参数模态框显隐
const detectorParamsVisible = ref(false)
// 系统报警模态框显隐
const systemAlarmVisible = ref(false)

const Status = reactive({
    Power: {
        running: false,
        HV: 0.00,
        HI: 0.00,
        FI: 0.00,
        FV: 0.00,
    },
    Stage: {
        running: false,
        X: 0.000,
        X_Status: 'Stopped',
        Y: 0.000,
        Y_Status: 'Stopped',
        Z: 0.000,
        Z_Status: 'Stopped',
        R: 0.000,
        R_Status: 'Stopped',
        XX: 0.000,
        XX_Status: 'Stopped',
    },
    Detector: {
        running: false,
        sn: 'unknown',
        width: 0,
        height: 0,
        tempreture: 0.0,
        humidity: 0.0,
        exposureTime: 0,
        gain: 0,
        binning: '',
        repeatTimes: -1,
        mode: 'null',
    },
    SystemStatus: '正常',
    RunTime: '00:00:00',
})
// 图像数据状态
const DiffractImage = ref('');
const imageRotateAngle = ref(0.0)
const selectedImageId = ref(0)
const imagesCounts = ref(0)



//===================== 系统状态更新 ======================
onMounted(async () =>{
    EventsOn('stage_running', (running) => {
        Status.Stage.running = running
    });
    EventsOn('detector_running', (running) => {
        Status.Detector.running = running
    });
    EventsOn('hvps_running', (running) => {
        Status.Power.running = running
    });

    EventsOn('motor_heartbeat', (motors) => {
        Status.Stage.X = motors.X.position
        Status.Stage.X_Status = motors.X.status
        Status.Stage.Y = motors.Y.position
        Status.Stage.Y_Status = motors.Y.status
        Status.Stage.Z = motors.Z.position
        Status.Stage.Z_Status = motors.Z.status
        Status.Stage.R = motors.R.position
        Status.Stage.R_Status = motors.R.status
        Status.Stage.XX = motors.XX.position
        Status.Stage.XX_Status = motors.XX.status
    });

    EventsOn('door', (door) => {
        if (door) {
            Status.SystemStatus = '正常'
            systemAlarmVisible.value = false
        }else{
            Status.SystemStatus = '报警'
            showToast('系统报警')
            systemAlarmVisible.value = true
        }
    });

    EventsOn('hvps_heartbeat', (data) => {
        Status.Power.HV = data.HV
        Status.Power.HI = data.HI
        Status.Power.FI = data.FI
        Status.Power.FV = data.FV
    });

    // 监听后端detector_image事件（新图像到达）
    EventsOn('detector_image', (data) => {
        console.log("收到新图像")
        DiffractImage.value = data.image;
    });

    // 监听后端detector_heartbeat事件（新心跳到达）
    EventsOn('detector_heartbeat', (data) => {
        Status.Detector.sn = data.sn
        Status.Detector.tempreture = data.tempreture
        Status.Detector.humidity = data.humidity
        Status.Detector.mode = data.mode

        Status.Detector.exposureTime = data.expose_time
        Status.Detector.binning = data.binning
        Status.Detector.repeatTimes = data.repeat_times
        Status.Detector.gain = data.gain
        Status.Detector.width = data.width
        Status.Detector.height = data.height

        imagesCounts.value = data.image_counts
    });

    //计时器，每秒更新一次运行时间：YYYY-MM-DD HH:mm:ss
    setInterval(() => {
        Status.RunTime = new Date().toLocaleString().substring(0, 19)
    }, 1000);
});



// ===================== 放射源控制 =====================
const hvpsBusy = ref(false)
const hvpsSet = reactive({
    SI: 1.7,
    LI: 1.7,
    HV: 40.0,
    HI: 200.0,
    PowerSwitch: false,
    FilamentSwitch: false,
})

async function handlePowerSwitch() {
    if (hvpsBusy.value) return
    hvpsBusy.value = true
    try {
        await HVPSSourceOpen(!hvpsSet.PowerSwitch)
        hvpsSet.PowerSwitch = !hvpsSet.PowerSwitch
        showToast(hvpsSet.PowerSwitch ? '放射源已开启' : '放射源已关闭')
    } catch (err) {
        console.error('handlePowerSwitch fail:', err, hvpsSet.PowerSwitch)
        showToast('放射源开关失败')
    } finally {
        hvpsBusy.value = false
    }
}

async function handleFilamentSwitch() {
    if (hvpsBusy.value) return
    hvpsBusy.value = true
    try {
        await HVPSSetFilamentOpen(!hvpsSet.FilamentSwitch)
        hvpsSet.FilamentSwitch = !hvpsSet.FilamentSwitch
        showToast(hvpsSet.FilamentSwitch ? '灯丝已开启' : '灯丝已关闭')
    } catch (err) {
        console.error('handleFilamentSwitch fail:', err)
        showToast('灯丝开关失败')
    } finally {
        hvpsBusy.value = false
    }
}

async function handleSetFilamentPreheat() {
    try {
        await HVPSSetFilamentPreheat(hvpsSet.SI)
        showToast(`灯丝预热 ${hvpsSet.SI}`)
    } catch (err) {
        console.error('handleSetFilamentPreheat fail:', err)
        showToast('灯丝预热失败')
    } finally {
        hvpsBusy.value = false
    }
}

async function handleSetFilamentLimit() {
    try {
        await HVPSSetFilamentLimit(hvpsSet.LI)
        showToast(`灯丝限制 ${hvpsSet.LI}`)
    } catch (err) {
        console.error('handleSetFilamentLimit fail:', err)
        showToast('灯丝限制失败')
    } finally {
        hvpsBusy.value = false
    }
}

async function handleHVPSSetHV() {
    if (hvpsBusy.value) return
    hvpsBusy.value = true
    try {
        await HVPSSetHV(hvpsSet.HV)
        showToast(`高压电源电压 ${hvpsSet.HV}`)
    } catch (err) {
        console.error('handleHVPSSetHV fail:', err)
        showToast('高压电源电压失败')   
    } finally {
        hvpsBusy.value = false
    }
}

async function handleHVPSSetHI() {
    if (hvpsBusy.value) return
    hvpsBusy.value = true
    try {
        await HVPSSetHI(hvpsSet.HI)
        showToast(`高压电源电流 ${hvpsSet.HI}`)
    } catch (err) {
        console.error('handleHVPSSetHI fail:', err)
        showToast('高压电源电流失败')
    } finally {
        hvpsBusy.value = false
    }
}

// ===================== 运动控制 =====================
const stageBusy = ref(false)
const motionTargets = reactive({
    X: 0.0,
    Y: 0.0,
    Z: 0.0,
    R: 0.0,
    XX: 0.0,
})

async function handleAxisCW(axis,motionTarget) {
    console.log('handleAxisCW', axis, motionTarget)
    if (stageBusy.value) return
    try {
        await StageRelMove(axis, motionTarget)
    } catch (err) {
        console.error('StagesRelMove fail', axis, motionTarget, err)
        showToast(`${axis} ${motionTarget} 启动失败`)
    }
}

async function handleAxisCCW(axis,motionTarget) {
    console.log('handleAxisCCW', axis, motionTarget)
    if (stageBusy.value) return
    try {
        await StageRelMove(axis, -motionTarget)
        Status.Stage.runing = true
    } catch (err) {
        console.error('StagesRelMove fail', axis, motionTarget, err)
        showToast(`${axis} ${motionTarget} 启动失败`)
    }
}

async function handleAllStop() {
    if (stageBusy.value) return
    stageBusy.value = true
    try {
        await Promise.all(['X', 'Y', 'Z', 'R', 'XX'].map(a => StageStop(a).catch(() => {})))
        Status.Stage.runing = false
        showToast('已全部停止')
    } finally {
        stageBusy.value = false
    }
}

// ===================== 探测器控制 =====================
function handleDetectorParams() {
    if (Status.Detector.mode === 'null' || Status.Detector.sn === 'unknown') {
        showToast('探测器未连接')
        return
    }
    detectorParamsVisible.value = true
}

async function handleDetectorCapture() {
    try {
        await DetectorCapture()
    } catch (err) {
        console.error('handleDetectorCapture fail:', err)
        showToast('拍摄失败')
    }
}


// ===================== Toast =====================

const toastMsg = ref('')
const toastTimer = ref(/** @type {any} */(null))

function showToast(msg) {
    toastMsg.value = msg
    if (toastTimer.value) clearTimeout(toastTimer.value)
    toastTimer.value = setTimeout(() => { toastMsg.value = '' }, 1800)
}

// ===================== 图像处理（调用入口占位）=====================
async function handleSelectImageByID() {
    try {
        await CallImageByID(selectedImageId.value)
        showToast(`已选择图像 ${selectedImageId.value}`)
    } catch (err) {
        console.error('handleSelectImageByID fail:', err)
        showToast('选择失败')
    }
}

async function handleSaveImage() {
    try {
        await SaveImageByID(selectedImageId.value)
        showToast('保存成功')
    } catch (err) {
        console.error('handleSaveImage fail:', err)
        showToast('保存失败')
    }
}

async function handleDeleteImage() {
    try {
        await DelImageByID(selectedImageId.value)
        showToast('删除成功')
    } catch (err) {
        console.error('handleDeleteImage fail:', err)
        showToast('删除失败')
    }
}


// ===================== Python图像处理 =====================
async function handleProcess(category, method) {
    if (!DiffractImage.value) {
        showToast('无图像可处理')
        return
    }
    if (processBusy.value) return
    processBusy.value = true
    showToast(`${method} 处理中...`)
    try {
        // TODO: 对接 services/PythonSDK/XRDBridge.go 对应方法：
        // 去模糊: [RichardsonLucy|Wiener]Denoise16bit 等（命名模式类似 BM3DDenoise16bit）
        // 降噪  : AMF / BM3D / Bilateral / TV(Denoise16bit) / WaveletDenoise16bit / NLMDenoise16bit / Filter16bit
        // 锐化  : Laplace / USM  (src/sharpen/)
        // 调整  : Gamma / Log / Exp  (src/transform/)
        await new Promise(r => setTimeout(r, 250))

        const newImg = {
            id: nextImageId(),
            name: `processed_${method}.tif`,
            url: DiffractImage.value, // TODO: 用实际处理输出的图像 URL / base64
            source: `${category}:${method}`,
        }
        images.value.push(newImg)
        selectedImageId.value = newImg.id
        showToast(`${method} 完成`)
    } catch (err) {
        console.error('process error', category, method, err)
        showToast(`${method} 失败`)
    } finally {
        processBusy.value = false
    }
}
</script>

<style scoped>
.main-window {
    display: flex;
    flex-direction: column;
    height: 100%;
    box-sizing: border-box;
    background-color: #d7e0ea;
    border-radius: 12px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    overflow: hidden;
}

.title-bar {
    display: flex;
    width: 100%;
    height: 60px;
    align-items: center;
    gap: 24px;
    padding: 0 20px;
    box-sizing: border-box;
    
    background-color: #b8bec5;
    --wails-draggable: drag;
}

.title {
    font-size: 18px;
    font-weight: 500;
    color: #617d88;
}

.title-divider {
    width: 1px;
    height: 18px;
    background: #94a3b8;
    margin: 0 6px;
    flex-shrink: 0;
}

.window-actions {
    display: flex;
    align-items: center;
    gap: 4px;
    justify-content: flex-end;
    margin-left: auto;
    --wails-draggable: no-drag;
}

.window-btn {
    width: 36px;
    height: 36px;
    border: none;
    background: transparent;
    color: #6b7280;
    cursor: pointer;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s;
    padding: 0;
}

.window-btn svg {
    width: 18px;
    height: 18px;
}

.window-btn:hover { 
    background: rgba(36, 48, 53, 0.1);
    color: #1a2e36;
    transform: scale(1.1);
}

.window-btn.close-btn:hover {
    background: rgba(80, 63, 63, 0.2);
    color: #ef4444;
}

.window-btn.connect-btn:hover {
    background: rgba(14, 165, 233, 0.15);
    color: #0284c7;
}

.window-btn.settings-btn:hover {
    background: rgba(100, 116, 139, 0.18);
    color: #0f172a;
}





.content {
    display: flex;
    flex: 1;
    gap: 12px;
    margin: 8px;
    overflow: hidden;
    min-height: 0;
}

.left-panel {
    width: 15%;
    min-width: 250px;
    background-color: transparent;
    border-radius: 8px;
    padding: 0;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    gap: 10px;
    max-height: 100%;
    min-height: 0;
    overflow-y: auto;
    overflow-x: hidden;
}

/* ====== 通用信息卡片（工业风，浅色主题，紧凑型） ====== */
.info-card {
    background: #ffffff;
    border-radius: 8px;
    border: 1px solid #e2e8f0;
    box-shadow: 0 1px 3px rgba(15, 23, 42, 0.06);
    display: flex;
    flex-direction: column;
    overflow: hidden;
}

.card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px 14px;
    background: linear-gradient(180deg, #f1f5f9 0%, #e8eef5 100%);
    border-bottom: 1px solid #e2e8f0;
}

.card-title {
    font-size: 13px;
    font-weight: 600;
    color: #0f172a;
    letter-spacing: 0.02em;
}

.card-status {
    font-size: 11px;
    font-weight: 600;
    padding: 3px 8px;
    border-radius: 4px;
    line-height: 1.2;
    white-space: nowrap;
}

.status-alarm {
    background: rgba(239, 68, 68, 0.12);
    color: #dc2626;
    border: 1px solid rgba(239, 68, 68, 0.3);
}

.status-ready {
    background: rgba(16, 185, 129, 0.12);
    color: #059669;
    border: 1px solid rgba(16, 185, 129, 0.3);
}

.status-partial {
    background: rgba(251, 191, 36, 0.12);
    color: #d97706;
    border: 1px solid rgba(251, 191, 36, 0.3);
}

.status-running {
    background: rgba(14, 165, 233, 0.12);
    color: #0284c7;
    border: 1px solid rgba(14, 165, 233, 0.3);
}

.status-idle {
    background: rgba(100, 116, 139, 0.1);
    color: #475569;
    border: 1px solid rgba(100, 116, 139, 0.2);
}

.status-off {
    background: rgba(148, 163, 184, 0.12);
    color: #64748b;
    border: 1px solid rgba(148, 163, 184, 0.25);
}

.card-body {
    padding: 10px 14px 12px;
    display: flex;
    flex-direction: column;
    gap: 6px;
}

.info-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 12px;
    line-height: 1.4;
}

.info-label {
    color: #64748b;
    font-weight: 500;
    white-space: nowrap;
}

.info-value {
    color: #334155;
    font-weight: 600;
    font-variant-numeric: tabular-nums;
    text-align: right;
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    margin-left: 8px;
}

.info-value.highlight {
    color: #0284c7;
}

.info-value.moving {
    color: #16a34a;
    animation: blink 0.8s ease-in-out infinite alternate;
}

@keyframes blink {
    from { opacity: 1; }
    to { opacity: 0.4; }
}

.info-value.danger {
    color: #dc2626;
}

.info-value.alarm-value {
    color: #dc2626;
    animation: blink 0.8s ease-in-out infinite alternate;
}

.radiation-icon {
    width: 18px;
    height: 18px;
    color: #eab308;
    flex-shrink: 0;
}

.middle-panel {
    flex: 1;
    min-width: 0;
    background-color: #fff;
    border-radius: 8px;
    border: 1px solid #e2e8f0;
    box-shadow: 0 1px 3px rgba(15, 23, 42, 0.06);
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    position: relative;
}

/* ====== 1. 顶部操作栏（高 50px） ====== */
.image-toolbar {
    flex: 0 0 50px;
    height: 30px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 14px;
    background: linear-gradient(180deg, #f1f5f9 0%, #e8eef5 100%);
    border-bottom: 1px solid #e2e8f0;
    box-sizing: border-box;
    gap: 12px;
}

.toolbar-left,
.toolbar-right {
    display: flex;
    align-items: center;
    gap: 8px;
}

.toolbar-label {
    font-size: 12px;
    color: #64748b;
    font-weight: 500;
    white-space: nowrap;
}

.toolbar-count {
    min-width: 28px;
    padding: 2px 8px;
    background: #fff;
    border: 1px solid #cbd5e1;
    color: #0f172a;
    border-radius: 4px;
    font-size: 12px;
    font-weight: 600;
    text-align: center;
    font-variant-numeric: tabular-nums;
}

.toolbar-select {
    height: 30px;
    padding: 0 10px;
    font-size: 12px;
    color: #0f172a;
    background: #fff;
    border: 1px solid #cbd5e1;
    border-radius: 4px;
    outline: none;
    cursor: pointer;
    min-width: 180px;
    max-width: 260px;
}
.toolbar-select:focus {
    border-color: #0284c7;
    box-shadow: 0 0 0 2px rgba(14, 165, 233, 0.15);
}
.toolbar-select:disabled {
    opacity: 0.55;
    cursor: not-allowed;
    background: #f1f5f9;
}

.toolbar-btn {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    height: 32px;
    padding: 0 12px;
    font-size: 12px;
    font-weight: 600;
    border-radius: 4px;
    border: 1px solid transparent;
    cursor: pointer;
    white-space: nowrap;
    transition: all 0.15s ease;
    color: #fff;
}
.toolbar-btn svg {
    width: 14px;
    height: 14px;
}
.toolbar-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}
.toolbar-btn:not(:disabled):active {
    transform: translateY(1px);
}
.save-btn {
    background: #0284c7;
    border-color: #0278b4;
}
.save-btn:not(:disabled):hover {
    background: #0278b4;
}
.delete-btn {
    background: #ef4444;
    border-color: #dc2626;
}
.delete-btn:not(:disabled):hover {
    background: #dc2626;
}

/* ====== 2. 图像显示区（flex:1） ====== */
.image-viewer {
    flex: 1;
    min-height: 0;
    background: #4d5158;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: auto;
    position: relative;
}

.diffract-image {
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
    display: block;
    user-select: none;
    -webkit-user-drag: none;
    image-rendering: auto;
}

.empty-viewer {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 10px;
    color: #94a3b8;
}

.empty-icon {
    width: 56px;
    height: 56px;
    opacity: 0.7;
}

.empty-text {
    font-size: 13px;
    font-weight: 500;
    letter-spacing: 0.04em;
}

/* ====== 3. 图像处理栏（高 100px，横向栏目） ====== */
.process-toolbar {
    flex: 0 0 120px;
    height: 120px;
    box-sizing: border-box;
    background: linear-gradient(180deg, #ffffff 0%, #f8fafc 100%);
    border-top: 1px solid #e2e8f0;
    display: flex;
    align-items: stretch;
    gap: 0;
    overflow-x: auto;
    overflow-y: hidden;
}

.process-group {
    flex: 1 1 0;
    min-width: 0;
    padding: 8px 10px;
    display: flex;
    flex-direction: column;
    gap: 6px;
    border-right: 1px solid #e2e8f0;
    box-sizing: border-box;
}
.process-group:last-child {
    border-right: none;
}

.group-title {
    font-size: 11px;
    font-weight: 600;
    color: #0f172a;
    letter-spacing: 0.04em;
    white-space: nowrap;
    display: inline-flex;
    align-items: center;
    gap: 6px;
}
.group-title::before {
    content: '';
    display: inline-block;
    width: 3px;
    height: 11px;
    border-radius: 2px;
    background: #0284c7;
}

.group-btns {
    flex: 1;
    display: grid;
    grid-auto-flow: column;
    grid-auto-columns: 1fr;
    grid-template-rows: repeat(2, minmax(0, 1fr));
    gap: 5px;
    align-content: stretch;
    align-items: stretch;
    min-height: 0;
}

.proc-btn {
    padding: 0 6px;
    font-size: 11px;
    font-weight: 600;
    color: #334155;
    background: #fff;
    border: 1px solid #cbd5e1;
    border-radius: 4px;
    cursor: pointer;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    line-height: 1.2;
    transition: all 0.15s ease;
    min-height: 0;
}
.proc-btn:hover {
    color: #0284c7;
    border-color: #0284c7;
    background: rgba(14, 165, 233, 0.06);
}
.proc-btn:active {
    transform: translateY(1px);
}

/* ====== Toast 轻提示（相对 middle-panel 绝对定位，美化版） ====== */
.toast-hint {
    position: absolute;
    top: 60px;
    left: 50%;
    transform: translateX(-50%);
    min-width: 140px;
    max-width: calc(100% - 40px);
    padding: 10px 18px;
    background: linear-gradient(135deg, #1e293b 0%, #0f172a 100%);
    color: #f1f5f9;
    font-size: 13px;
    font-weight: 500;
    border-radius: 10px;
    pointer-events: none;
    z-index: 99;
    letter-spacing: 0.02em;
    white-space: nowrap;
    text-align: center;
    border: 1px solid rgba(148, 163, 184, 0.22);
    box-shadow:
        0 10px 30px rgba(15, 23, 42, 0.28),
        0 2px 8px rgba(15, 23, 42, 0.18),
        inset 0 1px 0 rgba(255, 255, 255, 0.06);
    backdrop-filter: blur(8px);
    -webkit-backdrop-filter: blur(8px);
}
.toast-hint::before {
    content: '';
    position: absolute;
    top: -1px;
    left: 50%;
    transform: translateX(-50%);
    width: 42px;
    height: 3px;
    border-radius: 0 0 3px 3px;
    background: linear-gradient(90deg, #0ea5e9, #22d3ee);
    box-shadow: 0 0 8px rgba(14, 165, 233, 0.45);
}

.toast-fade-enter-active,
.toast-fade-leave-active {
    transition:
        opacity 0.22s cubic-bezier(0.22, 1, 0.36, 1),
        transform 0.22s cubic-bezier(0.22, 1, 0.36, 1),
        filter 0.22s ease;
}
.toast-fade-enter-from {
    opacity: 0;
    filter: blur(2px);
    transform: translate(-50%, -10px) scale(0.96);
}
.toast-fade-leave-to {
    opacity: 0;
    filter: blur(2px);
    transform: translate(-50%, -10px) scale(0.96);
}


.right-panel {
    width: 15%;
    min-width: 290px;
    background-color: transparent;
    border-radius: 8px;
    padding: 0;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    gap: 10px;
    max-height: 100%;
    min-height: 0;
    overflow-y: auto;
    overflow-x: hidden;
}

/* ====== 左侧/右侧面板自定义滚动条（美化 + 占位小） ====== */
.left-panel,
.right-panel {
    scrollbar-gutter: stable;
    scrollbar-width: thin;
    scrollbar-color: #cbd5e1 transparent;
}
.left-panel::-webkit-scrollbar,
.right-panel::-webkit-scrollbar {
    width: 8px;
}
.left-panel::-webkit-scrollbar-track,
.right-panel::-webkit-scrollbar-track {
    background: transparent;
}
.left-panel::-webkit-scrollbar-thumb,
.right-panel::-webkit-scrollbar-thumb {
    background: #cbd5e1;
    border-radius: 4px;
    border: 2px solid transparent;
    background-clip: padding-box;
}
.left-panel::-webkit-scrollbar-thumb:hover,
.right-panel::-webkit-scrollbar-thumb:hover {
    background: #94a3b8;
    background-clip: padding-box;
    border: 2px solid transparent;
}

/* ====== 右侧控制卡片通用 ====== */
.ctrl-card {
    background: #fff;
    border-radius: 8px;
    border: 1px solid #e2e8f0;
    box-shadow: 0 1px 3px rgba(15, 23, 42, 0.06);
    display: flex;
    flex-direction: column;
    overflow: hidden;
}

.ctrl-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px 14px;
    background: linear-gradient(180deg, #f1f5f9 0%, #e8eef5 100%);
    border-bottom: 1px solid #e2e8f0;
}

.ctrl-title {
    font-size: 13px;
    font-weight: 600;
    color: #0f172a;
    letter-spacing: 0.02em;
}

.ctrl-status {
    font-size: 11px;
    font-weight: 600;
    padding: 3px 8px;
    border-radius: 4px;
    line-height: 1.2;
    white-space: nowrap;
}

.ctrl-body {
    padding: 10px 14px 12px;
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.ctrl-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 10px;
    min-height: 28px;
}

.ctrl-label {
    font-size: 12px;
    color: #64748b;
    font-weight: 500;
    white-space: nowrap;
}

.ctrl-divider {
    height: 1px;
    background: #e2e8f0;
    margin: 4px 0;
}

/* ====== Switch 开关 ====== */
.switch {
    position: relative;
    display: inline-block;
    width: 40px;
    height: 22px;
    flex: 0 0 auto;
}
.switch input {
    opacity: 0;
    width: 0;
    height: 0;
}
.slider {
    position: absolute;
    inset: 0;
    cursor: pointer;
    background: #cbd5e1;
    border-radius: 22px;
    transition: 0.2s ease;
}
.slider::before {
    content: '';
    position: absolute;
    height: 16px;
    width: 16px;
    left: 3px;
    top: 3px;
    background: #fff;
    border-radius: 50%;
    transition: 0.2s ease;
    box-shadow: 0 1px 2px rgba(15, 23, 42, 0.2);
}
.switch input:checked + .slider {
    background: #0284c7;
}
.switch input:checked + .slider::before {
    transform: translateX(18px);
}
.switch input:disabled + .slider {
    opacity: 0.5;
    cursor: not-allowed;
}

/* ====== 数字输入 + 单位 ====== */
.num-input-wrap {
    display: inline-flex;
    align-items: center;
    flex: 1 1 auto;
    max-width: 140px;
    justify-content: flex-end;
    gap: 6px;
}
.num-input {
    height: 28px;
    padding: 0 8px;
    font-size: 12px;
    font-weight: 600;
    color: #0f172a;
    background: #fff;
    border: 1px solid #cbd5e1;
    border-radius: 4px;
    outline: none;
    width: 100%;
    max-width: 96px;
    text-align: right;
    font-variant-numeric: tabular-nums;
    transition: border-color 0.15s ease, box-shadow 0.15s ease;
}
.num-input:focus {
    border-color: #0284c7;
    box-shadow: 0 0 0 2px rgba(14, 165, 233, 0.15);
}
.num-input:disabled {
    background: #f1f5f9;
    color: #94a3b8;
    cursor: not-allowed;
}
.num-input::-webkit-outer-spin-button,
.num-input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}
.num-input[type=number] {
    -moz-appearance: textfield;
}
.num-unit {
    font-size: 11px;
    font-weight: 600;
    color: #64748b;
    min-width: 22px;
    text-align: left;
}

/* ====== 运动控制行（轴行）====== */
.axis-row {
    display: grid;
    grid-template-columns: 24px 1fr auto;
    align-items: center;
    gap: 8px;
    min-height: 30px;
}
.axis-label {
    font-size: 13px;
    font-weight: 700;
    color: #0f172a;
    text-align: center;
    background: #e2e8f0;
    border-radius: 4px;
    height: 28px;
    line-height: 28px;
    padding: 0 4px;
}
.axis-value-wrap {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    justify-content: flex-end;
}
.axis-input {
    max-width: 92px;
    height: 28px;
}
.axis-btns {
    display: inline-flex;
    gap: 4px;
}
.axis-btn {
    height: 28px;
    min-width: 40px;
    padding: 0 8px;
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.02em;
    border-radius: 4px;
    border: 1px solid #cbd5e1;
    background: #fff;
    cursor: pointer;
    transition: all 0.12s ease;
    color: #334155;
    user-select: none;
}
.axis-btn:hover:not(:disabled) {
    border-color: #0284c7;
    color: #0284c7;
    background: rgba(14, 165, 233, 0.05);
}
.axis-btn:active:not(:disabled) {
    transform: translateY(1px);
}
.axis-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}
.axis-btn.cw {
    background: #ecfeff;
    border-color: #67e8f9;
    color: #0e7490;
}
.axis-btn.cw:hover:not(:disabled) {
    background: #cffafe;
    border-color: #22d3ee;
}
.axis-btn.ccw {
    background: #eff6ff;
    border-color: #93c5fd;
    color: #1d4ed8;
}
.axis-btn.ccw:hover:not(:disabled) {
    background: #dbeafe;
    border-color: #3b82f6;
}
.axis-btn.go {
    background: #f0fdf4;
    border-color: #86efac;
    color: #15803d;
}
.axis-btn.go:hover:not(:disabled) {
    background: #dcfce7;
    border-color: #4ade80;
}

.stop-row {
    margin-top: 4px;
    display: flex;
    justify-content: center;
}
.stop-btn {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    height: 34px;
    padding: 0 16px;
    font-size: 12px;
    font-weight: 700;
    color: #fff;
    background: linear-gradient(180deg, #ef4444 0%, #dc2626 100%);
    border: 1px solid #b91c1c;
    border-radius: 4px;
    cursor: pointer;
    transition: all 0.15s ease;
    width: 100%;
    justify-content: center;
}
.stop-btn:hover:not(:disabled) {
    background: linear-gradient(180deg, #dc2626 0%, #b91c1c 100%);
}
.stop-btn:active:not(:disabled) {
    transform: translateY(1px);
}
.stop-btn:disabled {
    opacity: 0.55;
    cursor: not-allowed;
}

/* ====== 探测器控制按钮 ====== */
.detector-row {
    display: flex;
    width: 100%;
}
.det-btn {
    flex: 1;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    height: 38px;
    padding: 0 14px;
    font-size: 13px;
    font-weight: 700;
    border-radius: 4px;
    border: 1px solid transparent;
    cursor: pointer;
    transition: all 0.15s ease;
    color: #fff;
}
.det-btn svg {
    width: 16px;
    height: 16px;
}
.det-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}
.det-btn:not(:disabled):active {
    transform: translateY(1px);
}
.params-btn {
    background: linear-gradient(180deg, #64748b 0%, #475569 100%);
    border-color: #334155;
}
.params-btn:not(:disabled):hover {
    background: linear-gradient(180deg, #475569 0%, #334155 100%);
}
.capture-btn {
    background: linear-gradient(180deg, #0284c7 0%, #0369a1 100%);
    border-color: #075985;
}
.capture-btn:not(:disabled):hover {
    background: linear-gradient(180deg, #0369a1 0%, #075985 100%);
}



.footer {
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #b8bec5;
    gap: 12px;
}

</style>
