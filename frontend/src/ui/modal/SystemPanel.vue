<template>
    <div class="system-panel">
        <div class="panel-header">
            <span class="panel-title">系统配置</span>
            <button class="close-btn" @click="emit('close')">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
            </button>
        </div>
        
        <div class="panel-content">
            <!-- 位移台配置 -->
            <div class="param-card">
                <div class="card-title">位移台信息</div>
                
                <!-- X轴配置 -->
                <div class="axis-row">
                    <span class="axis-label">X</span>
                    <span class="axis-text">轴号</span>
                    <select v-model.number="systemData.stageConfig.xAxis.axisNumber" class="axis-input small-input">
                        <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
                    </select>
                    <span class="axis-text">分辨率</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.xAxis.resolution" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="pulse/mm"
                    />
                    <span class="axis-text">速度</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.xAxis.speed" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="mm/s"
                    />
                </div>

                <!-- Y轴配置 -->
                <div class="axis-row">
                    <span class="axis-label">Y</span>
                    <span class="axis-text">轴号</span>
                    <select v-model.number="systemData.stageConfig.yAxis.axisNumber" class="axis-input small-input">
                        <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
                    </select>
                    <span class="axis-text">分辨率</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.yAxis.resolution" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="pulse/mm"
                    />
                    <span class="axis-text">速度</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.yAxis.speed" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="mm/s"
                    />
                </div>

                <!-- Z轴配置 -->
                <div class="axis-row">
                    <span class="axis-label">Z</span>
                    <span class="axis-text">轴号</span>
                    <select v-model.number="systemData.stageConfig.zAxis.axisNumber" class="axis-input small-input">
                        <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
                    </select>
                    <span class="axis-text">分辨率</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.zAxis.resolution" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="pulse/mm"
                    />
                    <span class="axis-text">速度</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.zAxis.speed" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="mm/s"
                    />
                </div>

                <!-- R轴配置 -->
                <div class="axis-row">
                    <span class="axis-label">R</span>
                    <span class="axis-text">轴号</span>
                    <select v-model.number="systemData.stageConfig.rAxis.axisNumber" class="axis-input small-input">
                        <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
                    </select>
                    <span class="axis-text">分辨率</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.rAxis.resolution" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="pulse/°"
                    />
                    <span class="axis-text">速度</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.rAxis.speed" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="°/s"
                    />
                </div>

                <!-- L轴配置 -->
                <div class="axis-row">
                    <span class="axis-label">L</span>
                    <span class="axis-text">轴号</span>
                    <select v-model.number="systemData.stageConfig.lAxis.axisNumber" class="axis-input small-input">
                        <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
                    </select>
                    <span class="axis-text">分辨率</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.lAxis.resolution" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="pulse/°"
                    />
                    <span class="axis-text">速度</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.stageConfig.lAxis.speed" 
                        class="axis-input"
                        min="0.001" 
                        step="0.001"
                        placeholder="°/s"
                    />
                </div>
            </div>

            <!-- 图像保存配置 -->
            <div class="param-card">
                <div class="card-title">图像保存信息</div>
                
                <div class="simple-row">
                    <span class="simple-label">保存格式</span>
                    <select v-model="systemData.imageConfig.saveFormat" class="axis-input small-input">
                        <option value="raw">RAW</option>
                        <option value="tiff">TIFF</option>
                    </select>
                    <span class="simple-label">像素大小 (mm/pixel)</span>
                    <input 
                        type="number" 
                        v-model.number="systemData.imageConfig.pixelSize" 
                        class="axis-input small-input"
                        min="0.001" 
                        max="99"
                        step="0.001"
                    />
                    <span class="simple-label">旋转角度</span>
                    <select v-model.number="systemData.imageConfig.rotateAngle" class="axis-input small-input">
                        <option :value="0">0°</option>
                        <option :value="90">90°</option>
                        <option :value="-90">-90°</option>
                        <option :value="180">180°</option>
                    </select>
                </div>
                
            </div>

            <!-- 其他配置 -->
            <!-- <div class="param-card">
                <div class="card-title">其他配置</div>
                
                <div class="simple-row">
                    <span class="simple-label">日志级别</span>
                    <select v-model="systemData.logLevel" class="axis-input small-input">
                        <option value="debug">Debug</option>
                        <option value="info">Info</option>
                        <option value="warn">Warn</option>
                        <option value="error">Error</option>
                    </select>
                </div>
            </div> -->

            <!-- 按钮组 -->
            <div class="btn-group">
                <button class="action-btn primary-btn" @click="saveSystem">
                    保存配置
                </button>
                <button class="action-btn secondary-btn" @click="emit('close')">
                    取消配置
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { reactive, onMounted } from 'vue';
import { LoadHistorySystem, SaveHistorySystem } from '../../wailsjs/go/services/System';
import { DiffractSetAxisSpeed } from '../../wailsjs/go/services/DiffractService'

const emit = defineEmits(['close', 'save']);

// 系统配置数据
const systemData = reactive({
    stageConfig: {
        xAxis: {
            axisNumber: 1,
            resolution: 1000.0,
            speed: 10.0
        },
        yAxis: {
            axisNumber: 2,
            resolution: 1000.0,
            speed: 10.0
        },
        zAxis: {
            axisNumber: 3,
            resolution: 1000.0,
            speed: 5.0
        },
        rAxis: {
            axisNumber: 4,
            resolution: 3600.0,
            speed: 180.0
        },
        lAxis: {
            axisNumber: 5,
            resolution: 1600.0,
            speed: 5.0
        }
    },
    imageConfig: {
        saveFormat: 'tiff',
        pixelSize: 0.05,
        rotateAngle: 90
    },
    logLevel: 'info'
});


// 加载系统配置
const loadSystem = async () => {
    try {
        const system = await LoadHistorySystem();
        if (system.stage_config) {
            systemData.stageConfig.xAxis.axisNumber = system.stage_config.x_axis?.axis_number || 1;
            systemData.stageConfig.xAxis.resolution = system.stage_config.x_axis?.resolution || 1000.0;
            systemData.stageConfig.xAxis.speed = system.stage_config.x_axis?.speed || 10.0;
            
            systemData.stageConfig.yAxis.axisNumber = system.stage_config.y_axis?.axis_number || 2;
            systemData.stageConfig.yAxis.resolution = system.stage_config.y_axis?.resolution || 1000.0;
            systemData.stageConfig.yAxis.speed = system.stage_config.y_axis?.speed || 10.0;
            
            systemData.stageConfig.zAxis.axisNumber = system.stage_config.z_axis?.axis_number || 3;
            systemData.stageConfig.zAxis.resolution = system.stage_config.z_axis?.resolution || 1000.0;
            systemData.stageConfig.zAxis.speed = system.stage_config.z_axis?.speed || 5.0;
            
            systemData.stageConfig.rAxis.axisNumber = system.stage_config.r_axis?.axis_number || 4;
            systemData.stageConfig.rAxis.resolution = system.stage_config.r_axis?.resolution || 3600.0;
            systemData.stageConfig.rAxis.speed = system.stage_config.r_axis?.speed || 180.0;
            
            systemData.stageConfig.lAxis.axisNumber = system.stage_config.l_axis?.axis_number || 5;
            systemData.stageConfig.lAxis.resolution = system.stage_config.l_axis?.resolution || 1600.0;
            systemData.stageConfig.lAxis.speed = system.stage_config.l_axis?.speed || 5.0;
        }
        
        if (system.image_config) {
            systemData.imageConfig.saveFormat = system.image_config.save_format || 'tiff';
            systemData.imageConfig.pixelSize = system.image_config.pixel_size || 0.05;
            systemData.imageConfig.rotateAngle = system.image_config.rotate_angle || 90;
        }
        
        systemData.logLevel = system.log_level || 'info';
        
    } catch (err) {
        console.log('加载系统配置失败:', err);
    }
};

// 验证轴号唯一性
const validateAxisNumbers = () => {
    const axisNumbers = [
        systemData.stageConfig.xAxis.axisNumber,
        systemData.stageConfig.yAxis.axisNumber,
        systemData.stageConfig.zAxis.axisNumber,
        systemData.stageConfig.rAxis.axisNumber,
        systemData.stageConfig.lAxis.axisNumber
    ];
    
    const seen = new Set();
    for (const num of axisNumbers) {
        if (num < 1 || num > 5) {
            alert('轴号必须在1-5之间');
            return false;
        }
        if (seen.has(num)) {
            alert('轴号不能重复');
            return false;
        }
        seen.add(num);
    }
    return true;
};

// 保存系统配置
const saveSystem = async () => {
    if (!validateAxisNumbers()) {
        return;
    }
    
    try {
        const system = {
            stage_config: {
                x_axis: {
                    axis_number: systemData.stageConfig.xAxis.axisNumber,
                    resolution: systemData.stageConfig.xAxis.resolution,
                    speed: systemData.stageConfig.xAxis.speed
                },
                y_axis: {
                    axis_number: systemData.stageConfig.yAxis.axisNumber,
                    resolution: systemData.stageConfig.yAxis.resolution,
                    speed: systemData.stageConfig.yAxis.speed
                },
                z_axis: {
                    axis_number: systemData.stageConfig.zAxis.axisNumber,
                    resolution: systemData.stageConfig.zAxis.resolution,
                    speed: systemData.stageConfig.zAxis.speed
                },
                r_axis: {
                    axis_number: systemData.stageConfig.rAxis.axisNumber,
                    resolution: systemData.stageConfig.rAxis.resolution,
                    speed: systemData.stageConfig.rAxis.speed
                },
                l_axis: {
                    axis_number: systemData.stageConfig.lAxis.axisNumber,
                    resolution: systemData.stageConfig.lAxis.resolution,
                    speed: systemData.stageConfig.lAxis.speed
                }
            },
            image_config: {
                save_format: systemData.imageConfig.saveFormat,
                pixel_size: systemData.imageConfig.pixelSize,
                rotate_angle: systemData.imageConfig.rotateAngle
            },
            log_level: systemData.logLevel
        };
        
        console.log('[SystemPanel] 发送的数据:', JSON.stringify(system, null, 2));
        await SaveHistorySystem(system);
        
        try {
            await DiffractSetAxisSpeed('X', systemData.stageConfig.xAxis.speed);
            await DiffractSetAxisSpeed('Y', systemData.stageConfig.yAxis.speed);
            await DiffractSetAxisSpeed('Z', systemData.stageConfig.zAxis.speed);
            await DiffractSetAxisSpeed('R', systemData.stageConfig.rAxis.speed);
            await DiffractSetAxisSpeed('L', systemData.stageConfig.lAxis.speed);
            console.log('[SystemPanel] 设置电机速度成功');
        } catch (err) {
            console.log('[SystemPanel] 设置电机速度失败:', err);
        }
        
        emit('save', system, systemData.imageConfig.rotateAngle);
    } catch (err) {
        console.log('保存系统配置失败:', err);
    }
};

// 组件挂载时加载配置
onMounted(() => {
    loadSystem();
});
</script>

<style scoped>
.system-panel {
    width: 680px;
    background: linear-gradient(180deg, rgba(15, 23, 42, 0.95) 0%, rgba(30, 41, 59, 0.95) 100%);
    border-radius: 12px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
    display: flex;
    flex-direction: column;
    max-height: 90vh;
    backdrop-filter: blur(10px);
    border: 1px solid rgba(56, 189, 248, 0.2);
}

.panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 16px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.2);
    background: rgba(56, 189, 248, 0.1);
    border-radius: 12px 12px 0 0;
}

.panel-title {
    font-size: 14px;
    font-weight: 600;
    color: #38bdf8;
}

.close-btn {
    width: 26px;
    height: 26px;
    border: none;
    background: rgba(239, 68, 68, 0.2);
    border-radius: 6px;
    color: #ef4444;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
}

.close-btn:hover {
    background: rgba(239, 68, 68, 0.3);
    transform: rotate(90deg);
}

.close-btn svg {
    width: 14px;
    height: 14px;
}

.panel-content {
    flex: 1;
    padding: 10px;
    overflow-y: auto;
}

/* 参数卡片样式 */
.param-card {
    background: rgba(15, 23, 42, 0.6);
    border: 1px solid rgba(56, 189, 248, 0.15);
    border-radius: 8px;
    padding: 10px;
    margin-bottom: 8px;
}

.card-title {
    font-size: 11px;
    font-weight: 600;
    color: #38bdf8;
    margin-bottom: 8px;
    padding-bottom: 4px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.2);
}

/* 轴配置行 */
.axis-row {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 8px 0;
    border-bottom: 1px solid rgba(56, 189, 248, 0.08);
}

.axis-row:last-child {
    border-bottom: none;
}

.axis-label {
    width: 20px;
    height: 26px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    color: #fff;
    font-size: 12px;
    font-weight: 600;
    border-radius: 4px;
}

.axis-text {
    font-size: 11px;
    color: #94a3b8;
    min-width: 36px;
}

.axis-input {
    height: 26px;
    padding: 0 8px;
    background: rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(56, 189, 248, 0.2);
    border-radius: 4px;
    color: #ffffff;
    font-size: 12px;
    outline: none;
    transition: all 0.2s ease;
    box-sizing: border-box;
}

.axis-input.small-input {
    width: 70px;
}

.axis-input:not(.small-input) {
    width: 110px;
}

.axis-input:focus {
    border-color: #38bdf8;
    box-shadow: 0 0 0 2px rgba(56, 189, 248, 0.2);
}

.axis-input::placeholder {
    color: #475569;
    font-size: 10px;
}

/* 简单配置行 */
.simple-row {
    display: flex;
    align-items: center;
    gap: 16px;
    padding: 8px 0;
}

.simple-label {
    font-size: 11px;
    color: #94a3b8;
    min-width: 80px;
}

/* 按钮组 */
.btn-group {
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-top: 10px;
}

.action-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px 18px;
    border: none;
    border-radius: 5px;
    font-size: 11px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s ease;
    height: 32px;
}

.primary-btn {
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    color: #ffffff;
    box-shadow: 0 2px 8px rgba(56, 189, 248, 0.3);
}

.primary-btn:hover {
    transform: translateY(-1px);
    box-shadow: 0 3px 12px rgba(56, 189, 248, 0.5);
}

.secondary-btn {
    background: rgba(71, 85, 105, 0.3);
    border: 1px solid rgba(148, 163, 184, 0.3);
    color: #94a3b8;
}

.secondary-btn:hover {
    background: rgba(71, 85, 105, 0.5);
    color: #cbd5e1;
}

/* 滚动条样式 */
.panel-content::-webkit-scrollbar {
    width: 5px;
}

.panel-content::-webkit-scrollbar-track {
    background: rgba(0, 0, 0, 0.2);
    border-radius: 3px;
}

.panel-content::-webkit-scrollbar-thumb {
    background: rgba(56, 189, 248, 0.3);
    border-radius: 3px;
}

.panel-content::-webkit-scrollbar-thumb:hover {
    background: rgba(56, 189, 248, 0.5);
}
</style>
