<template>
    <main-container> 
        <!-- 菜单栏 -->
        <div class='top-bar' @dblclick="handleTitleBarDblClick">
            <div class='logo-section'>
                <logo>DIFFRACT CONTROL</logo>
            </div>
            <menu-bar>
                <!-- 项目 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('项目')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>
                    项目
                </button>

                <!-- 设备 -->
                 <button class="menu-item" type="button" @click="onMenuItemClick('设备')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"></path></svg>
                    设备
                </button>

                <!-- 设置 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('设置')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="3"></circle><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"></path></svg>
                    设置
                </button>

                <!-- 帮助 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('帮助')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
                    帮助
                </button>

                <!-- 关于 -->
                <button class="menu-item" type="button" @click="onMenuItemClick('关于')">
                    <svg class="menu-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>
                    关于
                </button>
            </menu-bar>

            <div class="top-right-controls">
                <user-header>
                    <button class="user-avatar">
                        <svg viewBox="0 0 24 24" fill="currentColor" width="30" height="30">
                            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08 1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z"/>
                        </svg>
                    </button>
                </user-header>
                <div class='window-actions'>    
                    <button class="window-btn minimize-btn" type="button" title="最小化" @click="minimize">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <line x1="4" y1="12" x2="22" y2="12"/>
                        </svg>
                    </button>
                    <button class="window-btn maximize-btn" type="button" title="最大化" @click="toggleMaximize">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <rect x="5" y="5" width="17" height="17" rx="2"/>
                        </svg>
                    </button>
                    <button class="window-btn close-btn" type="button" title="关闭" @click="closeApp">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <line x1="20" y1="4" x2="4" y2="20"/>
                            <line x1="4" y1="4" x2="20" y2="20"/>
                        </svg>
                    </button>
                </div>
            </div>
        </div>

        <contollr-container>
            <!-- 左控制面板 -->
            <left-panel>
                <div class="status-group">
                    <div class="status-card power-card">
                        <div class="card-header">
                            <span class="card-title">射线源状态</span>
                            <span class="card-status status-running">{{ Status.Power.runing ? '运行中' : '未运行' }}</span>
                        </div>
                        <div class="card-body">
                            <div class="info-row">
                                <span class="info-label">电源电压</span>
                                <span class="info-value highlight">{{ Status.Power.HV.toFixed(2) }} kV</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">电源电流</span>
                                <span class="info-value highlight">{{ Status.Power.HI.toFixed(2) }} uA</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">灯丝预热电压</span>
                                <span class="info-value highlight">{{ Status.Power.FV.toFixed(2) }} V</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">灯丝电流限制</span>
                                <span class="info-value highlight">{{ Status.Power.FI.toFixed(2) }} A</span>
                            </div>
                        </div>
                    </div>

                    <div class="status-card detector-card">
                        <div class="card-header">
                            <span class="card-title">样品台状态</span>
                            <span class="card-status status-ready">{{ Status.Stage.runing ? '运行中' : '未就绪' }}</span>
                        </div>
                        <div class="card-body">
                            <div class="info-row">
                                <span class="info-label">X当前坐标</span>
                                <span class="info-value highlight" >{{ Status.Stage.X.toFixed(4) }} mm</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">Y当前坐标</span>
                                <span class="info-value highlight" >{{ Status.Stage.Y.toFixed(4) }} mm</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">Z当前坐标</span>
                                <span class="info-value highlight" >{{ Status.Stage.Z.toFixed(4) }} mm</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">R当前角度</span>
                                <span class="info-value highlight" >{{ Status.Stage.R.toFixed(4) }} °</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">XX当前坐标</span>
                                <span class="info-value highlight" >{{ Status.Stage.XX.toFixed(4) }} mm</span>
                            </div>
                        </div>
                    </div>

                    <div class="status-card network-card">
                        <div class="card-header">
                            <span class="card-title">探测器状态</span>
                            <span class="card-status status-connected">{{ Status.Detector.runing ? '运行中' : '未连接' }}</span>
                        </div>
                        <div class="card-body">
                            <div class="info-row">
                                <span class="info-label">探测器SN</span>
                                <span class="info-value highlight" >{{ Status.Detector.sn }}</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">当前模式</span>
                                <span class="info-value highlight" >{{ Status.Detector.mode }}</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">表面温度</span>
                                <span class="info-value highlight" >{{ Status.Detector.tempreture.toFixed(2) }} °C</span>
                            </div>
                            <div class="info-row">
                                <span class="info-label">表面湿度</span>
                                <span class="info-value highlight" >{{ Status.Detector.humidity.toFixed(2) }}%</span>
                            </div>
                        </div>
                    </div>
                </div>
            </left-panel>

            <!-- 中间显示面板 -->
            <div class='mid-panel-wrapper'>
                <!-- 左侧控制面板按钮 -->
                <button 
                    class="panel-toggle-btn" 
                    :class="{ 'expanded': showDetectorPanel }"
                    @click="toggleDetectorPanel"
                >{{ showDetectorPanel ? '《' : '》' }}</button>
                
                <!-- 左侧探测器控制面板 -->
                <transition name="slide" >
                    <div v-if="showDetectorPanel" class="detector-panel-container" >
                        <DetectorPanel ref="detectorPanelRef" @close="showDetectorPanel = false" @apply="handleApplyParams" />
                    </div>
                </transition>
                
                <!-- 中间显示区域 -->
                <div class='mid-panel'>
                    <div class='detector-container'>
                        <div class='detector-content-wrapper'>
                            <!-- 纵向直方图显示 -->
                            <div class='histogram-container' v-if="Status.CurrentImage.total > 0">
                                <div class='histogram-axis'>
                                    <span class='axis-label-top'>10000</span>
                                    <span class='axis-label-btm'>0</span>
                                </div>
                                <div class='histogram-bars'>
                                    <div 
                                        v-for="(count, index) in Status.CurrentImage.histogramData" 
                                        :key="index"
                                        class='histogram-bar'
                                        :style="{ opacity: Math.max(0.1, count / maxHistogramCount) }"
                                        :title="'值: ' + (index * 100) + '-' + ((index + 1) * 100) + ', 像素: ' + count"
                                    ></div>
                                </div>
                            </div>
                            <!-- 探测器显示 -->
                            <div class='detector-display'>
                                <!-- 无图像显示 -->
                                <div v-if="!Status.CurrentImage.data" class='detector-placeholder'>
                                    <div class='placeholder-icon'>
                                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                                            <rect x="3" y="3" width="18" height="18" rx="2"/>
                                            <circle cx="12" cy="12" r="6"/>
                                            <circle cx="12" cy="12" r="2"/>
                                        </svg>
                                    </div>
                                    <div class='placeholder-text'>等待图像数据...</div>
                                </div>

                                <!-- 图像显示 -->
                                <img v-else :src="Status.CurrentImage.data" class='detector-image' alt='探测器图像' :style="{ transform: `rotate(${Status.CurrentImage.rotate}deg)` }" @mousemove="handleImageMouseMove"/>
                            </div>
                            
                            <!-- 双滑块控制条 -->
                             <div v-if="Status.CurrentImage.total > 0" class='intensity-slider-container'>
                                <div class='slider-track-wrapper'>
                                    <div class='slider-track'>
                                        <div class='slider-fill' :style="{ top: (100 - contrastMax) + '%', height: (contrastMax - contrastMin) + '%' }"></div>
                                        <!-- 最大值滑块 -->
                                        <div 
                                            class='slider-thumb max-thumb'
                                            :style="{ top: Math.max(0, Math.min(100, 100 - contrastMax)) + '%' }"
                                            @mousedown="startDrag('max', $event)"
                                            @touchstart="startDrag('max', $event)"
                                        ></div>
                                        <!-- 最小值滑块 -->
                                        <div 
                                            class='slider-thumb min-thumb'
                                            :style="{ top: Math.max(0, Math.min(100, 100 - contrastMin)) + '%' }"
                                            @mousedown="startDrag('min', $event)"
                                            @touchstart="startDrag('min', $event)"
                                        ></div>
                                    </div>
                                </div>
                                <div class='slider-labels'>
                                    <span class='label-max'></span>
                                    <span class='label-min'></span>
                                </div>
                            </div>

                            
                        </div>
                            
                        <!-- 探测器监控面板 -->
                        <DetectorMonitor 
                            :angle="Status.Stage.R" 
                            :imageWidth="Status.CurrentImage.width" 
                            :imageHeight="Status.CurrentImage.height"  
                            :imageValue="Status.CurrentImage.value"
                            :status="Status.DiffractScan.runing ? '运行中' : '未运行'"
                            />
                    </div>
                    
                    <!-- 图像切换滑块（在mid-panel下方） -->
                    <div v-if="Status.CurrentImage.total > 0" class='image-slider-container'>
                        <span class='image-counter'>{{ Status.CurrentImage.index }} / {{ Status.CurrentImage.total }}</span>
                        <input 
                            type="range" 
                            :min="1" 
                            :max="Status.CurrentImage.total" 
                            v-model.number="Status.CurrentImage.index"
                            class='image-slider'
                            @input="handleImageSliderChange"
                        />
                        <button class='clear-btn' @click="handleClearImages" title="清除所有图片">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M3 6h18"/>
                                <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/>
                                <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/>
                            </svg>
                        </button>
                    </div>
                </div>
            </div>

            <!-- 右控制面板 -->
            <div class='right-panel'>
                <!-- 射线控制模块 -->
                <div class="control-section">
                    <div class="section-title">射线控制</div>
                    <div class="toggle-group">
                        <div class="toggle-item">
                            <span class="toggle-label">电源开关</span>
                            <button 
                                class="toggle-switch" 
                                :class="{ 'toggle-on': Status.Power.powerSwitch }"
                                @click="handlePowerSwitch"
                            >
                                <span class="toggle-thumb"></span>
                            </button>
                        </div>
                        <div class="toggle-item">
                            <span class="toggle-label">灯丝开关</span>
                            <button 
                                class="toggle-switch" 
                                :class="{ 'toggle-on': Status.Power.filamentSwitch }"
                                @click="handleFilamentSwitch"
                            >
                                <span class="toggle-thumb"></span>
                            </button>
                        </div>
                    </div>
                </div>

                <!-- 采集控制按钮 -->
                 <div class="control-section">
                    <div class="section-title">探测器采集控制</div>
                    <div class="button-grid">
                        <button  class="action-btn" :class="{ 'primary-btn': !Status.DiffractScan.runing, 'danger-btn': Status.DiffractScan.runing }">
                            <svg class="btn-svg" viewBox="0 0 24 24" fill="currentColor">
                                <path v-if="!Status.DiffractScan.runing" d="M6 3v18l15-9z"/>
                                <rect v-else x="6" y="6" width="30" height="30" rx="1"/>
                            </svg>
                            <span class="btn-text">{{ Status.DiffractScan.runing ? '停止采集' : '连续采集' }}</span>
                        </button>
                        <button class="action-btn info-btn" @click="handleSingleScan" @dblclick="handleContinuousScan">
                            <svg class="btn-svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                                <circle cx="12" cy="13" r="4"/>
                            </svg>
                            <span class="btn-text">单帧采集</span>
                        </button>
                    </div>
                </div>


                <!-- 2. 新增：定位运动控制（上侧） -->
            <div class="control-section motion-section">
                <div class="section-title">样品台定位运动</div>
                
                <!-- 模式切换 -->
                <div class="mode-selector">
                    <button 
                        type="button" 
                        class="mode-tab" 
                        :class="{ active: motionTargets.mode === 'relative' }" 
                        @click="motionTargets.mode = 'relative'"
                    >
                        相对运动
                    </button>
                    <button 
                        type="button" 
                        class="mode-tab" 
                        :class="{ active: motionTargets.mode === 'absolute' }" 
                        @click="motionTargets.mode = 'absolute'">
                        绝对运动
                    </button>
                   
                </div>

                <!-- 四轴数值输入 -->
                <div class="axis-inputs-grid">
                    <div class="axis-input-item">
                        <span class="axis-input-label">X</span>
                        <div class="axis-input-wrapper">
                            <input 
                                type="number" 
                                step="0.001" 
                                v-model.number="motionTargets.X" 
                                @keyup.enter="handleExecuteMotion('X', motionTargets.X)"
                                :class="{ 'highlight-input': isInputModified('X') }"
                            />
                            <span class="axis-unit">mm</span>
                        </div>
                    </div>
                    <div class="axis-input-item">
                        <span class="axis-input-label">Y</span>
                        <div class="axis-input-wrapper">
                            <input 
                                type="number" 
                                step="0.001" 
                                v-model.number="motionTargets.Y" 
                                @keyup.enter="handleExecuteMotion('Y', motionTargets.Y)"
                                :class="{ 'highlight-input': isInputModified('Y') }"
                            />
                            <span class="axis-unit">mm</span>
                        </div>
                    </div>
                    <div class="axis-input-item">
                        <span class="axis-input-label">Z</span>
                        <div class="axis-input-wrapper">
                            <input 
                                type="number" 
                                step="0.001" 
                                v-model.number="motionTargets.Z" 
                                @keyup.enter="handleExecuteMotion('Z', motionTargets.Z)"
                                :class="{ 'highlight-input': isInputModified('Z') }"
                            />
                            <span class="axis-unit">mm</span>
                        </div>
                    </div>
                    <div class="axis-input-item">
                        <span class="axis-input-label">R</span>
                        <div class="axis-input-wrapper">
                            <input 
                                type="number" 
                                step="0.001" 
                                v-model.number="motionTargets.R" 
                                @keyup.enter="handleExecuteMotion('R', motionTargets.R)"
                                :class="{ 'highlight-input': isInputModified('R') }"
                            />
                            <span class="axis-unit">°</span>
                        </div>
                    </div>
                    <div class="axis-input-item">
                        <span class="axis-input-label">XX</span>
                        <div class="axis-input-wrapper">
                            <input 
                                type="number" 
                                step="0.001" 
                                v-model.number="motionTargets.XX" 
                                @keyup.enter="handleExecuteMotion('XX', motionTargets.XX)"
                                :class="{ 'highlight-input': isInputModified('XX') }"
                            />
                            <span class="axis-unit">mm</span>
                        </div>
                    </div>
                </div>
            </div>
                <!-- 位移台点动按钮 -->
                <div class="sector-wrapper">
                    <div class="sector-container">
                        <div class="bigbox s5" 
                        @mousedown="handleSectorPress('Y','CW')" 
                        @mouseup="handleSectorRelease('Y')">
                            <span class="sector-text">前</span>
                        </div>
                        <div class="bigbox s6"
                        @mousedown="handleSectorPress('R','CW')" 
                        @mouseup="handleSectorRelease('R')">
                            <span class="sector-text">顺</span>
                        </div>
                        <div class="bigbox s7" 
                        @mousedown="handleSectorPress('R','CCW')" 
                        @mouseup="handleSectorRelease('R')">
                            <span class="sector-text">逆</span>
                        </div>
                        <div class="bigbox s8" 
                        @mousedown="handleSectorPress('Y','CCW')" 
                        @mouseup="handleSectorRelease('Y')">
                            <span class="sector-text">后</span>
                        </div>
                        <div class="sector">
                            <div class="box s1" 
                            @mousedown="handleSectorPress('Z','CW')" 
                            @mouseup="handleSectorRelease('Z')">
                                <span class="sector-text">上</span>
                            </div>
                            <div class="box s2" 
                            @mousedown="handleSectorPress('X','CW')" 
                            @mouseup="handleSectorRelease('X')">
                                <span class="sector-text">右</span>
                            </div>
                            <div class="box s3" 
                            @mousedown="handleSectorPress('Z','CCW')" 
                            @mouseup="handleSectorRelease('Z')">
                                <span class="sector-text">下</span>
                            </div>
                            <div class="box s4" 
                            @mousedown="handleSectorPress('X','CCW')" 
                            @mouseup="handleSectorRelease('X')">
                                <span class="sector-text">左</span>
                            </div>
                            <div class="center" @click="handleStop">
                                <span class="sector-text">STOP </span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </contollr-container>

        <!-- 底部栏 -->
        <div class='bottom-bar'>
            <div class="bottom-left">
                <span class="status-item">系统状态: <span class="status-normal">{{Status.SystemStatus}}</span></span>
                <span class="status-item">运行时间: <span class="time-display">{{Status.RunTime}}</span></span>
            </div>
        </div>
    </main-container>
    
    <!-- 设备连接模态框 -->
    <DeviceConnectModal :visible="showDeviceModal" @close="showDeviceModal = false" />
    <div v-if="showProjectModal" class="modal-overlay" @click.self="showProjectModal = false">
        <div class="modal-container" >
            <ProjectPanel @close="showProjectModal = false" @save="showProjectModal = false"/>
        </div>
    </div>

    <!-- 系统设置模态框 -->
    <div v-if="showSettingModal" class="modal-overlay" @click.self="showSettingModal = false">
        <div class="modal-container" >
            <SystemPanel @close="showSettingModal = false" @save="handleSystemSave"/>
        </div>
    </div>

    <!-- 帮助模态框 -->
    <div v-if="showHelpModal" class="modal-overlay" @click.self="showHelpModal = false">
        <div class="modal-container" >
            <HelpPanel @close="showHelpModal = false"/>
        </div>
    </div>

    <!-- 关于模态框 -->
    <div v-if="showAboutModal" class="modal-overlay" @click.self="showAboutModal = false">
        <div class="modal-container" >
            <AboutPanel @close="showAboutModal = false"/>
        </div>
    </div>

    <!-- 更新提示模态框 -->
    <teleport to="body">
        <transition name="modal">
            <div v-if="showUpdateModal" class="modal-overlay" @click.self="showUpdateModal = false">
                <div class="modal-container update-modal">
                    <div class="modal-header">
                        <h3 class="modal-title">发现新版本</h3>
                        <button class="modal-close" @click="showUpdateModal = false">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    <div class="modal-body">
                        <div class="update-content">
                            <div class="update-icon">
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/>
                                    <path d="M3 3v5h5"/>
                                    <path d="M3 16a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"/>
                                    <path d="M16 21h5v-5"/>
                                </svg>
                            </div>
                            <div class="update-info">
                                <p class="update-version">最新版本: <strong>{{ updateInfo.tagName }}</strong></p>
                                <p class="update-current-version">当前版本: <strong>{{ updateInfo.currentVersion }}</strong></p>
                                <p class="update-desc">{{ updateInfo.body }}</p>
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button class="btn btn-secondary" @click="showUpdateModal = false">
                            稍后更新
                        </button>
                        <button class="btn btn-primary" @click="handleUpdate('github')">
                            更新(Github)
                        </button>
                        <button class="btn btn-primary" @click="handleUpdate('accelerate')">
                            更新(加速)
                        </button>
                    </div>
                </div>
            </div>
        </transition>
    </teleport>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, watch, computed, nextTick } from 'vue';

const Status = reactive({
    Power:{
        powerSwitch: false,
        filamentSwitch: false,
        HV: 0.00,
        HI: 0.00,
        FI: 0.00,
        FV: 0.00,
        runing: false
    },
    Stage:{
        X:0.000,
        Y:0.000,
        Z:0.000,
        R:0.000,
        XX:0.000,
        runing: false,
    },
    Detector:{
        runing: false,
        sn:'unknown',
        width: 0,
        height: 0,
        tempreture: 0.0,
        humidity: 0.0,
        exposureTime: 0.0,
        mode:'null',
    },
    CurrentImage:{
        data: null,
        value: 0.0,
        width: 0,
        height: 0,
        index: 0,
        total: 0,
        rotate: 0.0,
        histogramData: [],
    },
    DiffractScan:{
        runing: false,
    },
    SystemStatus: '正常',
    RunTime: '00:00:00',
    StartTime: '',
})



// 射线源控制
import { HVPSSourceOpen, HVPSSetFilamentOpen } from '../../wailsjs/go/components/HVPSService'
const handlePowerSwitch = async () => { 
    try {
        await HVPSSourceOpen(!Status.Power.powerSwitch);
        Status.Power.powerSwitch = !Status.Power.powerSwitch;
    } catch (error) {
        console.error('射线源开关失败:', error);
    }
    
};
const handleFilamentSwitch = async () => {
    try {
        await HVPSSetFilamentOpen(!Status.Power.filamentSwitch);
        Status.Power.filamentSwitch = !Status.Power.filamentSwitch;
    } catch (error) {
        console.error('灯丝开关失败:', error);
    }
};

// 运动控制
import {StagesJOGMove, StagesSTOPMove, StagesRELMove, StagesABSMove } from '../../wailsjs/go/components/StageService'
const motionTargets = reactive({ 
    X: 0, 
    Y: 0, 
    Z: 0, 
    R: 0, 
    XX: 0,

    MotionX: 0,
    MotionY: 0,
    MotionZ: 0,
    MotionR: 0,
    MotionXX: 0,

    mode:'relative',
});

const isInputModified = (axis) => {
    console.log(motionTargets[axis])
    return motionTargets[`Motion${axis}`] !== motionTargets[axis]
}

const handleSectorPress = async (axis, dir) => {
    console.log('[Debug] handleSectorPress', axis, dir);
    try {
        await StagesJOGMove(axis, dir)
    } catch (error) {
        console.error('运动失败:', error);
    }
};

const handleSectorRelease = async (axis) => {
    console.log('[Debug] handleSectorRelease', axis); 
    try {
        await StagesSTOPMove(axis)

    } catch (error) {
        console.error('停止失败:', error);
    }
};
const handleStop = async () => { 
    console.log('[Debug] handleStop'); 
    try {
        await StagesSTOPMove('X')
        await StagesSTOPMove('Y')
        await StagesSTOPMove('Z')
        await StagesSTOPMove('R')
        await StagesSTOPMove('XX')
    } catch (error) {
        console.error('停止失败:', error);
    }
};

const handleExecuteMotion = async (axis, length) => {
    console.log('[Debug] handleExecuteMotion', axis, length); 
    if (motionTargets.mode === 'relative') {
        try {
            await StagesRELMove(axis, length)
            motionTargets[`Motion${axis}`] = length
        } catch (error) {
            console.error('相对运动失败:', error);
        }
    } else if (motionTargets.mode === 'absolute') {
        try {
            await StagesABSMove(axis, length)
            motionTargets[`Motion${axis}`] = length
        } catch (error) {
            console.error('绝对运动失败:', error);
        }
    } else {
        console.error('未知的模式:', mode);
    }
};

// 探测器控制
import { } from '../../wailsjs/go/components/DetectorService'
const handleDetectorMode = async (mode) => {
    console.log('[Debug] handleDetectorMode', mode);
    
};

// 窗口控制
import { WindowMinimise, WindowToggleMaximise, Quit, WindowIsMaximised, WindowUnmaximise, EventsOn, EventsOff, BrowserOpenURL } from '../../wailsjs/runtime/runtime'
const handleTitleBarDblClick = async () => {
    const isMax = await WindowIsMaximised()
    if (isMax) {
        WindowUnmaximise()
    } else {
        WindowToggleMaximise()
    }
};
const minimize = () => WindowMinimise();
const toggleMaximize = () => WindowToggleMaximise();
const closeApp = async () => Quit();


// ========== 占位函数（TODO：待补充实现） ==========

// --- 以下 import 需要取消注释并补充实际的组件/API路径 ---

// import DetectorMonitor from './pannel/DetectorMonitor.vue/index.js';
// import DeviceConnectModal from './modal/DeviceConnectModal.vue/index.js';
// import ProjectPanel from './modal/ProjectPanel.vue/index.js';
// import SystemPanel from './modal/SystemPanel.vue/index.js';
// import HelpPanel from './modal/HelpPanel.vue/index.js';
// import AboutPanel from './modal/AboutPanel.vue/index.js';
// import { APIUpdate } from '../../wailsjs/go/main/App'
// import { DiffractStagesJOG, DiffractStagesSTOP, DiffractStagesMOVE } from '../../wailsjs/go/services/DiffractService'
// import { LoadHistorySystem } from '../../wailsjs/go/services/System'
// import { HVPSSourceOpen, HVPSSetFilamentOpen } from '../../wailsjs/go/services/HVPSDevice'

// 模态框状态
const showProjectModal = ref(false); // TODO
const showDeviceModal = ref(false);  // TODO
const showSettingModal = ref(false); // TODO
const showHelpModal = ref(false);    // TODO
const showAboutModal = ref(false);   // TODO
const showUpdateModal = ref(false);  // TODO

// 探测器面板
const showDetectorPanel = ref(false); // TODO
const detectorPanelRef = ref(null);   // TODO

// 对比度滑块
const contrastMax = ref(100); // TODO
const contrastMin = ref(0);   // TODO
const dragging = ref(null);   // TODO
const trackRect = ref(null);  // TODO
const getEventY = (e) => { /* TODO */ return e.clientY; };

// 直方图
const maxHistogramCount = computed(() => { /* TODO */ return 1; });

// 运动控制
const motionMode = ref('relative'); // TODO

const inputModified = reactive({ X: false, Y: false, Z: false, R: false, L: false }); // TODO
const isContinuousScanning = ref(false); // TODO
const updateInfo = ref({ tagName: '无', currentVersion: '无', downloadUrl: '', body: '无' }); // TODO





// 菜单
const onMenuItemClick = (option) => { /* TODO */ console.log('[TODO] onMenuItemClick', option); };

// 探测器面板
const toggleDetectorPanel = () => { /* TODO */ console.log('[TODO] toggleDetectorPanel'); };
const handleApplyParams = (params) => { /* TODO */ console.log('[TODO] handleApplyParams', params); };

// 图像处理
const handleImageMouseMove = (e) => { /* TODO */ };
const handleImageSliderChange = async () => { /* TODO */ console.log('[TODO] handleImageSliderChange'); };
const handleClearImages = async () => { /* TODO */ console.log('[TODO] handleClearImages'); };

// 对比度拖拽
const startDrag = (type, e) => { /* TODO */ };
const handleDrag = (e) => { /* TODO */ };
const stopDrag = async () => { /* TODO */ };

// 系统配置
const handleSystemSave = (system, rotateAngle) => { /* TODO */ };

// 更新
const handleUpdate = (type) => { /* TODO */ console.log('[TODO] handleUpdate', type); };



// CT扫描
const handleCTScan = async () => { /* TODO */ console.log('[TODO] handleCTScan'); };
const handleSingleScan = async () => { /* TODO */ console.log('[TODO] handleSingleScan'); };
const handleContinuousScan = async () => { /* TODO */ console.log('[TODO] handleContinuousScan'); };

// 生命周期
const timer = ref(null);

onMounted(() => {
    // TODO: 初始化设备事件监听、加载配置等
    console.log('[TODO] onMounted - 初始化');
});

onUnmounted(() => {
    // TODO: 清理事件监听
    console.log('[TODO] onUnmounted - 清理');
});
</script>

<style>
/* 容器重置 */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body, html { 
    margin: 0; 
    padding: 0; 
    height: 100%; 
    font-family: 'Segoe UI', 'Microsoft YaHei', -apple-system, BlinkMacSystemFont, sans-serif;
}

main-container { 
    display: flex;
    flex-direction: column;
    height: 100vh;
    background: linear-gradient(135deg, #f5f7fa 0%, #ffffff 50%, #f5f7fa 100%);
    color: #1f2937;
    overflow: hidden;
}

/* 顶部栏优化 */
.top-bar { 
    display: flex;
    align-items: center;
    gap: 24px;
    padding: 0 20px;
    height: 64px;
    background: linear-gradient(180deg, rgba(245, 247, 250, 0.95) 0%, rgba(255, 255, 255, 0.9) 100%);
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
    backdrop-filter: blur(10px);
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    --wails-draggable: drag;
}

.logo-section {
    display: flex;
    align-items: center;
    gap: 12px;
}

.logo-icon {
    width: 32px;
    height: 32px;
    color: #38bdf8;
    animation: rotate 20s linear infinite;
}

@keyframes rotate {
    from { transform: rotate(0deg); }
    to { transform: rotate(360deg); }
}

logo {
    padding: 8px 20px;
    border-radius: 8px;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.15) 0%, rgba(56, 189, 248, 0.05) 100%);
    color: #38bdf8;
    font-weight: 700;
    font-size: 18px;
    letter-spacing: 0.1em;
    text-shadow: 0 0 20px rgba(56, 189, 248, 0.3);
    border: 1px solid rgba(56, 189, 248, 0.2);
}

menu-bar { 
    display: flex; 
    align-items: center; 
    gap: 8px;
}

.menu-item {
    display: flex;
    align-items: center;
    gap: 8px;
    color: #1f2937;
    background: transparent;
    border: none;
    cursor: pointer;
    font-size: 14px;
    padding: 10px 16px;
    border-radius: 8px;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    position: relative;
    overflow: hidden;
}

.menu-item::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.1) 0%, transparent 100%);
    opacity: 0;
    transition: opacity 0.3s;
}

.menu-item:hover {
    color: #ffffff;
    background: rgba(56, 189, 248, 0.1);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(56, 189, 248, 0.2);
}

.menu-item:hover::before {
    opacity: 1;
}

.menu-item {
    display: flex;
    align-items: center;
}

.menu-svg {
    width: 20px;
    height: 20px;
    margin-right: 8px;
    color: #6b7280;
    transition: color 0.2s ease;
}

.menu-item:hover .menu-svg {
    color: #1d4ed8; 
}

.top-right-controls {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 20px;
}

.status-indicator {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 16px;
    background: rgba(16, 185, 129, 0.1);
    border-radius: 20px;
    border: 1px solid rgba(16, 185, 129, 0.2);
}

.status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    animation: pulse 2s infinite;
}

.status-online {
    background: #10b981;
    box-shadow: 0 0 10px rgba(16, 185, 129, 0.5);
}

@keyframes pulse {
    0%, 100% { opacity: 1; }
    50% { opacity: 0.5; }
}

.status-text {
    font-size: 12px;
    color: #10b981;
    font-weight: 600;
}

user-header { 
    display: flex; 
    align-items: center; 
}

.user-avatar {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    background: linear-gradient(135deg, #e5e7eb 0%, #ffffff 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    border: 2px solid rgba(56, 189, 248, 0.3);
    cursor: pointer;
    transition: all 0.3s;
    color: #38bdf8;
}

.user-avatar:hover {
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    color: #ffffff;
    transform: scale(1.1);
    box-shadow: 0 0 20px rgba(56, 189, 248, 0.4);
}

.window-actions { 
    display: flex; 
    align-items: center; 
    gap: 4px;
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
    background: rgba(56, 189, 248, 0.1);
    color: #38bdf8;
    transform: scale(1.1);
}

.window-btn.close-btn:hover {
    background: rgba(239, 68, 68, 0.2);
    color: #ef4444;
}

/* 中间区域布局 */
contollr-container {
    display: flex;
    flex: 1;
    overflow: hidden;
    gap: 16px;
    padding: 16px;
}

left-panel {
    width: 250px;
    background: linear-gradient(180deg, rgba(255, 255, 255, 0.8) 0%, rgba(245, 247, 250, 0.9) 100%);
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
}

.panel-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
    background: rgba(56, 189, 248, 0.05);
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}

.panel-title {
    font-size: 16px;
    font-weight: 600;
    color: #38bdf8;
}

.panel-icon {
    font-size: 20px;
}

.status-group {
    flex: 1;
    overflow-y: auto;
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.status-card {
    background: linear-gradient(135deg, rgba(245, 247, 250, 0.6) 0%, rgba(255, 255, 255, 0.4) 100%);
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 10px;
    overflow: hidden;
    transition: all 0.3s;
}

.status-card:hover {
    border-color: rgba(56, 189, 248, 0.3);
    box-shadow: 0 4px 20px rgba(56, 189, 248, 0.1);
    transform: translateY(-2px);
}

.card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 16px;
    background: rgba(56, 189, 248, 0.05);
    border-bottom: 1px solid rgba(0, 0, 0, 0.05);
}

.card-title {
    font-size: 13px;
    font-weight: 600;
    color: #1f2937;
}

.card-status {
    font-size: 11px;
    padding: 4px 10px;
    border-radius: 12px;
    font-weight: 600;
}

.status-running {
    background: rgba(16, 185, 129, 0.2);
    color: #10b981;
}

.status-ready {
    background: rgba(56, 189, 248, 0.2);
    color: #38bdf8;
}

.status-connected {
    background: rgba(168, 85, 247, 0.2);
    color: #a855f7;
}

.card-body {
    padding: 12px 16px;
}

.info-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 0;
    border-bottom: 1px solid rgba(0, 0, 0, 0.05);
}

.info-row:last-child {
    border-bottom: none;
}

.info-label {
    font-size: 12px;
    color: #6b7280;
}

.info-value {
    font-size: 13px;
    font-weight: 600;
    color: #1f2937;
}

.info-value.highlight {
    color: #38bdf8;
    font-family: 'Courier New', monospace;
}

.mid-panel-wrapper {
    flex: 1;
    display: flex;
    align-items: stretch;
    gap: 0;
    position: relative;
    overflow: hidden;
}

/* 左侧控制面板切换按钮 */
.panel-toggle-btn {
    position: absolute;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    z-index: 10;
    width: 28px;
    height: 160px;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.3) 0%, rgba(14, 165, 233, 0.3) 100%);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-left: none;
    border-radius: 0 8px 8px 0;
    color: #38bdf8;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    backdrop-filter: blur(10px);
}

.panel-toggle-btn:hover {
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.5) 0%, rgba(14, 165, 233, 0.5) 100%);
    border-color: rgba(56, 189, 248, 0.5);
    box-shadow: 2px 0 15px rgba(56, 189, 248, 0.3);
}

.panel-toggle-btn.expanded {
    left: 280px;
}

.panel-toggle-btn svg {
    width: 18px;
    height: 18px;
    transition: transform 0.3s ease;
}

.panel-toggle-btn.expanded svg {
    transform: rotate(180deg);
}

/* 左侧控制面板容器 */
.detector-panel-container {
    width: 280px;
    flex-shrink: 0;
    height: 100%;
    overflow: hidden;
}

/* 滑动过渡动画 */
.slide-enter-active,
.slide-leave-active {
    transition: transform 0.3s ease;
}

.slide-enter-from,
.slide-leave-to {
    transform: translateX(-100%);
}

.mid-panel {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 1px;
    gap: 0px;
    transition: margin-left 0.3s ease;
}

.detector-container {
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, rgba(245, 247, 250, 0.8) 0%, rgba(255, 255, 255, 0.6) 100%);
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
}

.detector-content-wrapper {
    flex: 1;
    display: flex;
    overflow: hidden;
    min-height: 0;
}

.detector-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
    background: rgba(56, 189, 248, 0.05);
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}

.detector-title {
    font-size: 16px;
    font-weight: 600;
    color: #38bdf8;
}

.detector-controls {
    display: flex;
    gap: 8px;
}

.control-btn {
    width: 32px;
    height: 32px;
    border: 1px solid rgba(0, 0, 0, 0.2);
    background: rgba(245, 247, 250, 0.5);
    border-radius: 6px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s;
    padding: 0;
}

.control-btn:hover {
    background: rgba(56, 189, 248, 0.2);
    border-color: rgba(56, 189, 248, 0.4);
    color: #38bdf8;
}

.control-btn svg {
    width: 16px;
    height: 16px;
}

.detector-display {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 20px;
    background: rgba(0, 0, 0, 0.3);
    position: relative;
    overflow: hidden;
}

/* 双滑块容器 */
.intensity-slider-container {
    width: 20px;
    background: rgba(245, 247, 250, 0.6);
    border-left: 1px solid rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    padding: 8px 0;
}

.slider-track-wrapper {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
}

.slider-track {
    width: 8px;
    height: 100%;
    background: rgba(0, 0, 0, 0.15);
    border-radius: 4px;
    position: relative;
    cursor: pointer;
}

.slider-fill {
    position: absolute;
    left: 0;
    width: 100%;
    background: linear-gradient(180deg, rgba(56, 189, 248, 0.7) 0%, rgba(56, 189, 248, 0.4) 100%);
    border-radius: 4px;
    pointer-events: none;
    transition: height 0.1s ease, top 0.1s ease;
}

.slider-thumb {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    width: 16px;
    height: 16px;
    border-radius: 50%;
    cursor: grab;
    border: 2px solid #fff;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
    transition: transform 0.2s ease, box-shadow 0.2s ease;
    z-index: 10;
}

.slider-thumb:active {
    cursor: grabbing;
    transform: translateX(-50%) scale(1.15);
}

.slider-thumb:hover {
    transform: translateX(-50%) scale(1.1);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4);
}

.max-thumb {
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
}

.min-thumb {
    background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);
}

.slider-labels {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 4px;
    padding-top: 8px;
    padding-bottom: 4px;
    border-top: 1px solid rgba(0, 0, 0, 0.1);
}

.slider-labels span {
    font-size: 9px;
    color: #6b7280;
    white-space: nowrap;
}

.slider-labels .label-max {
    color: #38bdf8;
}

.slider-labels .label-min {
    color: #f97316;
}

/* 纵向直方图显示 */
.histogram-container {
    width: 70px;
    display: flex;
    flex-direction: row;
    align-items: stretch;
    padding: 8px 6px;
    border-left: 1px solid rgba(0, 0, 0, 0.1);
    border-right: 1px solid rgba(0, 0, 0, 0.1);
    background: rgba(245, 247, 250, 0.4);
    gap: 4px;
    box-sizing: border-box;
    height: 100%;
    min-height: 0;
}

.histogram-axis {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: flex-end;
    padding-right: 4px;
    min-width: 32px;
    height: 100%;
}

.histogram-axis .axis-label-top,
.histogram-axis .axis-label-mid,
.histogram-axis .axis-label-btm {
    font-size: 8px;
    color: rgba(0, 0, 0, 0.7);
    white-space: nowrap;
    line-height: 1;
}

.histogram-axis .axis-label-mid {
    font-size: 7px;
}

.histogram-bars {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-end;
    gap: 0;
    position: relative;
    height: 100%;
    min-height: 0;
    overflow: hidden;
}

.histogram-bar {
    width: 100%;
    flex: 1;
    background: linear-gradient(180deg, rgba(56, 189, 248, 0.6) 0%, rgba(14, 165, 233, 0.3) 100%);
    border-radius: 0;
    transition: background-color 0.15s ease;
    min-height: 1px;
}

.histogram-bar:hover {
    background: linear-gradient(180deg, rgba(56, 189, 248, 0.9) 0%, rgba(14, 165, 233, 0.5) 100%);
}

.image-slider-container {
    width: 100%;
    margin-top: 8px;
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 0 16px;
    box-sizing: border-box;
}

.clear-btn {
    width: 28px;
    height: 28px;
    border: none;
    background: rgba(239, 68, 68, 0.2);
    border-radius: 6px;
    color: #ef4444;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
    flex-shrink: 0;
}

.clear-btn:hover {
    background: rgba(239, 68, 68, 0.3);
    transform: scale(1.05);
}

.clear-btn svg {
    width: 14px;
    height: 14px;
}

.image-counter {
    font-size: 12px;
    color: #6b7280;
    text-align: right;
    font-weight: 500;
}

.image-slider {
    flex: 1;
    height: 6px;
    -webkit-appearance: none;
    appearance: none;
    background: rgba(0, 0, 0, 0.5);
    border-radius: 3px;
    outline: none;
    cursor: pointer;
}

.image-slider::-webkit-slider-thumb {
    -webkit-appearance: none;
    appearance: none;
    width: 16px;
    height: 16px;
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    border-radius: 50%;
    cursor: pointer;
    box-shadow: 0 2px 6px rgba(56, 189, 248, 0.4);
    transition: all 0.2s ease;
}

.image-slider::-webkit-slider-thumb:hover {
    transform: scale(1.2);
    box-shadow: 0 3px 10px rgba(56, 189, 248, 0.6);
}

.image-slider::-moz-range-thumb {
    width: 16px;
    height: 16px;
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    border-radius: 50%;
    cursor: pointer;
    border: none;
    box-shadow: 0 2px 6px rgba(56, 189, 248, 0.4);
}

.detector-placeholder {
    text-align: center;
    color: #9ca3af;
}

.placeholder-icon {
    width: 120px;
    height: 120px;
    margin: 0 auto 20px;
    opacity: 0.3;
    animation: float 3s ease-in-out infinite;
}

@keyframes float {
    0%, 100% { transform: translateY(0); }
    50% { transform: translateY(-10px); }
}

.placeholder-text {
    font-size: 14px;
    font-weight: 500;
}

.detector-image {
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
    background: #f5f7fa;
}

.detector-footer {
    padding: 12px 20px;
    height: 88px;
    background: rgba(56, 189, 248, 0.05);
    border-top: 1px solid rgba(0, 0, 0, 0.1);
}

.detector-info {
    display: flex;
    gap: 16px;
    flex-wrap: wrap;
}

.info-tag {
    font-size: 11px;
    color: #6b7280;
    padding: 4px 12px;
    background: rgba(245, 247, 250, 0.5);
    border-radius: 12px;
    border: 1px solid rgba(0, 0, 0, 0.1);
}

.right-panel {
    width: 280px;
    background: linear-gradient(180deg, rgba(255, 255, 255, 0.8) 0%, rgba(245, 247, 250, 0.9) 100%);
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
    gap: 10px;
}

.control-group {
    flex: 1;
    overflow-y: auto;
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.control-section {
    background: rgba(245, 247, 250, 0.4);
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 10px;
    padding: 16px;
}

.section-title {
    font-size: 14px;
    font-weight: 600;
    color: #38bdf8;
    margin-bottom: 16px;
    padding-bottom: 8px;
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}

.button-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
}

/* Toggle开关样式 */
.toggle-group {
    display: flex;
    flex-direction: vertical;
    gap: 12px;
}

.toggle-item {
    gap: 8px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.toggle-label {
    font-size: 13px;
    color: #1f2937;
    font-weight: 500;
}

.toggle-switch {
    width: 50px;
    height: 28px;
    border-radius: 14px;
    background: rgba(0, 0, 0, 0.5);
    border: 1px solid rgba(0, 0, 0, 0.2);
    cursor: pointer;
    position: relative;
    transition: all 0.3s ease;
    padding: 0;
}

.toggle-switch:hover {
    background: rgba(0, 0, 0, 0.7);
}

.toggle-switch.toggle-on {
    background: linear-gradient(135deg, #0ea5e9 0%, #0284c7 100%);
    border-color: rgba(56, 189, 248, 0.5);
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.3);
}

.toggle-thumb {
    position: absolute;
    top: 3px;
    left: 3px;
    width: 20px;
    height: 20px;
    border-radius: 50%;
    background: #ffffff;
    transition: all 0.3s ease;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.toggle-on .toggle-thumb {
    left: 25px;
    background: #ffffff;
}

.action-btn {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    padding: 16px 12px;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s;
    font-size: 12px;
    font-weight: 600;
    height: 75px;
    width: 120px;
}

.action-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.primary-btn {
    background: rgba(14, 165, 233, 0.1);
    border: 1px solid rgba(14, 165, 233, 0.3);
    color: #0284c7;
}

.primary-btn:hover {
    background: linear-gradient(135deg, #0ea5e9 0%, #0284c7 100%);
    box-shadow: 0 4px 12px rgba(56, 189, 248, 0.4);
    color: #ffffff;
}

.secondary-btn {
    background: rgba(168, 85, 247, 0.1); border: 1px solid rgba(168, 85, 247, 0.25); color: #7c3aed;
}

.secondary-btn:hover {
    background: linear-gradient(135deg, #9333ea 0%, #7c3aed 100%);
    box-shadow: 0 4px 12px rgba(168, 85, 247, 0.4);
    color: #ffffff;
}

.danger-btn {
    background: rgba(239, 68, 68, 0.08); border: 1px solid rgba(239, 68, 68, 0.25); color: #dc2626;
}

.danger-btn:hover {
    background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
    box-shadow: 0 4px 12px rgba(239, 68, 68, 0.4);
    color: #ffffff;
}

.info-btn {
    background: rgba(16, 185, 129, 0.1); border: 1px solid rgba(16, 185, 129, 0.25); color: #059669;
}

.info-btn:hover {
    background: linear-gradient(135deg, #059669 0%, #047857 100%);
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4);
    color: #ffffff;
}

.btn-icon {
    font-size: 20px;
}

.param-item {
    margin-bottom: 16px;
}

.param-item:last-child {
    margin-bottom: 0;
}

.param-label {
    display: block;
    font-size: 12px;
    color: #6b7280;
    margin-bottom: 8px;
}

.param-input {
    display: flex;
    align-items: center;
    gap: 12px;
}

.param-slider {
    flex: 1;
    -webkit-appearance: none;
    height: 6px;
    background: rgba(0, 0, 0, 0.2);
    border-radius: 3px;
    outline: none;
}

.param-slider::-webkit-slider-thumb {
    -webkit-appearance: none;
    width: 16px;
    height: 16px;
    background: #38bdf8;
    border-radius: 50%;
    cursor: pointer;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.5);
}

.param-value {
    font-size: 12px;
    font-weight: 600;
    color: #38bdf8;
    min-width: 50px;
    text-align: right;
}

.quick-actions {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 8px;
}

.quick-btn {
    padding: 12px;
    background: rgba(245, 247, 250, 0.5);
    border: 1px solid rgba(0, 0, 0, 0.2);
    border-radius: 8px;
    color: #1f2937;
    font-size: 12px;
    cursor: pointer;
    transition: all 0.3s;
}

.quick-btn:hover {
    background: rgba(56, 189, 248, 0.1);
    border-color: rgba(56, 189, 248, 0.3);
    color: #38bdf8;
    transform: translateY(-2px);
}

/* 底部栏 */
.bottom-bar {
    height: 40px;
    background: linear-gradient(180deg, rgba(245, 247, 250, 0.95) 0%, rgba(255, 255, 255, 0.9) 100%);
    border-top: 1px solid rgba(0, 0, 0, 0.1);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 20px;
    backdrop-filter: blur(10px);
}

.bottom-left,
.bottom-right {
    display: flex;
    gap: 24px;
}

.status-item {
    font-size: 12px;
    color: #6b7280;
}

.status-normal {
    color: #10b981;
    font-weight: 600;
}

.time-display,
.memory-display,
.cpu-display {
    color: #38bdf8;
    font-weight: 600;
    font-family: 'Courier New', monospace;
}

/* 滚动条样式 */
::-webkit-scrollbar {
    width: 6px;
}

::-webkit-scrollbar-track {
    background: rgba(245, 247, 250, 0.5);
}

::-webkit-scrollbar-thumb {
    background: rgba(56, 189, 248, 0.3);
    border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
    background: rgba(56, 189, 248, 0.5);
}

.sector-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
  background: rgba(245, 247, 250, 0.5);
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 0;
}

.sector {
  position: relative;
  width: 200px;
  height: 200px;
  margin: 0 auto;
  border: 2px solid rgba(56, 189, 248, 0.3);
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(245, 247, 250, 0.8) 0%, rgba(255, 255, 255, 0.6) 100%);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.12), inset 0 0 20px rgba(56, 189, 248, 0.1);
}

.sector .box {
  position: absolute;
  width: 200px;
  height: 200px;
  border-radius: 50%;
  clip: rect(0px, 99px, 99px, 0);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
}

.sector .box,
.sector .center {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
}

/* 扇形文字样式 */
.sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  white-space: nowrap;
}

.sector .s1 {
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  transform: rotate(45deg);
}

.sector .s1 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(-45deg) translate(0px, -65px);
}

.sector .s1:hover,
.sector .s1.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .s2 {
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  transform: rotate(135deg);
}


.sector .s2 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(-135deg) translate(65px, 0px);
}

.sector .s2:hover,
.sector .s2.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .s3 {
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  transform: rotate(-135deg);
}

/* 给 s3 的文字也加上定位 */
.sector .s3 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(135deg) translate(0px, 65px);
}

.sector .s3:hover,
.sector .s3.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .s4 {
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  transform: rotate(-45deg);
}

/* 给 s4 的文字也加上定位 */
.sector .s4 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(45deg) translate(-65px, 0px);
}

.sector .s4:hover,
.sector .s4.hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.sector .center {
  width: 90px;
  height: 90px;
  position: absolute;
  left: 54px;
  top: 54px;
  border: 2px solid rgba(56, 189, 248, 1);
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(239, 68, 68, 1) 0%, rgba(220, 38, 38, 1) 100%);
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 600;
  text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  box-shadow: 0 4px 15px rgba(239, 68, 68, 0.3);
}

.sector .center:hover {
  background: linear-gradient(135deg, rgba(220, 38, 38, 0.9) 0%, rgba(185, 28, 28, 0.9) 100%);
  border-color: rgba(239, 68, 68, 0.6);
  box-shadow: 0 0 25px rgba(239, 68, 68, 0.6);
  transform: scale(1.05);
}

.sector-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  transform: scale(0.85);
  transform-origin: center top;
}

.l-axis-row {
  display: flex;
  gap: 10px;
  margin-top: 0px;
  width: 100%;
  padding: 0 10px;
  box-sizing: border-box;
}

.medium-box {
  flex: 1;
  height: 35px;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  border: 2px solid rgba(56, 189, 248, 0.3);
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.medium-box:hover {
  border-color: rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
  box-shadow: 0 0 15px rgba(56, 189, 248, 0.4);
  transform: scale(1.05);
}

.medium-text {
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
}


.bigbox{
  position: absolute;
  width: 115px;
  height: 115px;
  background: linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%);
  border: 2px solid rgba(56, 189, 248, 0.3);
  font-size: 18px;
  font-weight: 600;

  transform: translate(-50%, -50%); 
  border-radius: 5px;
  display: flex;
}

.bigbox:hover {
  border: 2px solid rgba(56, 189, 248, 1);
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
  box-shadow: 0 0 20px rgba(35, 43, 46, 0.6);
}

.bigbox.s5 {
  transform: rotate(0deg) translate(-65px, -65px);
}

/* 修正选择器：bigbox.s5 内部的 .sector-text */
.bigbox.s5 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(0deg) translate(10px, 10px);
}

.bigbox.s6 {
  transform: rotate(90deg) translate(-65px, -65px);
}
.bigbox.s6 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(-90deg) translate(-15px, 15px);
}

.bigbox.s7 {
  transform: rotate(180deg) translate(-65px, -65px);
}

.bigbox.s7 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(180deg) translate(-15px, -15px);
}



.bigbox.s8 {
  transform: rotate(270deg) translate(-65px, -65px);
}

.bigbox.s8 .sector-text {
  position: absolute;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  text-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
  pointer-events: none;
  transform: rotate(-270deg) translate(15px, -15px);
}

.btn-svg {
    width: 20px;
    height: 20px;
    color: #ffffff;
    flex-shrink: 0;
}

.btn-svg-sm {
    width: 14px;
    height: 14px;
    color: #ffffff;
    flex-shrink: 0;
}




/* 运动控制面板容器 */
.motion-section {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

/* 绝对/相对 模式选择器 */
.mode-selector {
    display: flex;
    background: rgba(245, 247, 250, 0.6);
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    padding: 3px;
}

.mode-tab {
    flex: 1;
    background: transparent;
    border: none;
    color: #6b7280;
    padding: 8px 0;
    font-size: 12px;
    font-weight: 600;
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.2s ease;
}

.mode-tab:hover {
    color: #1f2937;
}

.mode-tab.active {
    background: rgba(56, 189, 248, 0.15);
    color: #38bdf8;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.1);
}

/* 四轴参数输入网格 */
.axis-inputs-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
}

.axis-input-item {
    display: flex;
    align-items: center;
    background: rgba(245, 247, 250, 0.4);
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 6px;
    padding: 6px 10px;
    transition: all 0.3s;
}

.axis-input-item:focus-within {
    border-color: rgba(56, 189, 248, 0.5);
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.15);
}

.axis-input-label {
    font-size: 13px;
    font-weight: 700;
    color: #38bdf8;
    width: 18px;
}

.axis-input-wrapper {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.axis-input-wrapper input {
    width: 100%;
    background: transparent;
    border: none;
    color: #1f2937;
    font-family: 'Courier New', monospace;
    font-weight: 600;
    font-size: 12px;
    padding: 0 4px;
    outline: none;
    text-align: right;
}

/* 输入框黄色高亮样式 */
.axis-input-wrapper input.highlight-input {
    background: rgba(255, 193, 7, 0.3);
    color: #78350f;
    box-shadow: inset 0 0 0 1px rgba(255, 193, 7, 0.6);
}

/* 隐藏 Chrome 等浏览器 input[type=number] 的微调按钮 */
.axis-input-wrapper input::-webkit-outer-spin-button,
.axis-input-wrapper input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

.axis-unit {
    font-size: 11px;
    color: #9ca3af;
    margin-left: 6px;
    min-width: 22px;
    text-align: left;
}

/* 定位运动动作按钮 */
.motion-actions {
    display: grid;
    grid-template-columns: 1.2fr 1fr;
    gap: 10px;
    margin-top: 4px;
}

.motion-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    padding: 10px;
    border: none;
    border-radius: 6px;
    font-size: 12px;
    font-weight: 600;
    color: #ffffff;
    cursor: pointer;
    transition: all 0.3s;
}

.motion-btn:hover {
    transform: translateY(-1px);
}

/* 执行按钮 - 绿色 */
.exec-btn {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    box-shadow: 0 2px 8px rgba(16, 185, 129, 0.2);
}

.exec-btn:hover {
    background: linear-gradient(135deg, #059669 0%, #047857 100%);
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.4);
}

/* 停止按钮 - 红色 */
.stop-btn {
    background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
    box-shadow: 0 2px 8px rgba(239, 68, 68, 0.2);
}

.stop-btn:hover {
    background: linear-gradient(135deg, #dc2626 0%, #b91c1c 100%);
    box-shadow: 0 4px 12px rgba(239, 68, 68, 0.4);
}

/* 更新模态框样式 */
/* 遮罩层 */
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.6);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 9999;
    backdrop-filter: blur(4px);
}

/* 模态框容器 */
.modal-container {
    background: linear-gradient(180deg, rgba(255, 255, 255, 0.95) 0%, rgba(245, 247, 250, 0.98) 100%);
    border-radius: 16px;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
    border: 1px solid rgba(0, 0, 0, 0.1);
    overflow: hidden;
}

.modal-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
}

.modal-title {
    margin: 0;
    font-size: 15px;
    font-weight: 600;
    color: #4b5563;
}

.modal-close {
    background: transparent;
    border: none;
    color: #6b7280;
    cursor: pointer;
    padding: 4px;
    border-radius: 8px;
    transition: all 0.2s;
    width: 28px;
    height: 28px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.modal-close:hover {
    background: rgba(0, 0, 0, 0.06);
    color: #1f2937;
}

.modal-close svg {
    width: 18px;
    height: 18px;
}

.modal-body {
    padding: 20px;
}

.modal-footer {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    gap: 12px;
    padding: 16px 20px;
}

.btn {
    padding: 10px 20px;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    border: none;
    transition: all 0.2s;
}

.btn-primary {
    background: linear-gradient(135deg, rgba(56, 189, 248, 1) 0%, rgba(59, 130, 246, 1) 100%);
    color: white;
}

.btn-primary:hover {
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.9) 0%, rgba(59, 130, 246, 0.9) 100%);
    transform: translateY(-1px);
}

.btn-secondary {
    background: rgba(0, 0, 0, 0.05);
    color: #4b5563;
}

.btn-secondary:hover {
    background: rgba(0, 0, 0, 0.1);
}

/* 更新模态框内容 */
.update-modal {
    width: 420px;
    max-width: 90vw;
}

.update-modal .modal-header {
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.1) 0%, rgba(59, 130, 246, 0.1) 100%);
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}

.update-modal .modal-title {
    color: #1f2937;
    font-size: 16px;
    font-weight: 600;
}

.update-content {
    display: flex;
    align-items: center;
    gap: 20px;
    padding: 10px 0;
}

.update-icon {
    width: 64px;
    height: 64px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, rgba(56, 189, 248, 0.2) 0%, rgba(59, 130, 246, 0.2) 100%);
    border-radius: 16px;
    color: #38bdf8;
    flex-shrink: 0;
}

.update-icon svg {
    width: 32px;
    height: 32px;
}

.update-info {
    flex: 1;
}

.update-version {
    font-size: 20px;
    color: #1f2937;
    margin: 0 0 8px 0;
    font-weight: 500;
}

.update-version strong {
    color: #38bdf8;
}

.update-desc {
    font-size: 13px;
    color: #6b7280;
    margin: 0;
    line-height: 1.5;
}

.update-modal .modal-footer {
    border-top: 1px solid rgba(0, 0, 0, 0.1);
}

/* 模态框动画 */
.modal-enter-active,
.modal-leave-active {
    transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}

.modal-enter-active .modal-container,
.modal-leave-active .modal-container {
    transition: transform 0.3s ease, opacity 0.3s ease;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
    transform: scale(0.95);
    opacity: 0;
}

</style>
