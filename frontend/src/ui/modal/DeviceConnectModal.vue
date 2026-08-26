<template>
    <teleport to="body">
        <transition name="modal" @click.self="handleClose">
            <div v-if="visible" class="modal-overlay">
                <div class="modal-container">
                    <div class="modal-header">
                        <span class="modal-title">设备连接</span>
                        <button class="close-btn" @click="handleClose">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>
                    
                    <div class="modal-content">
                        <!-- 位移台 -->
                        <div class="device-item">
                            <label class="device-label">位移台</label>
                            <div class="device-input-group">
                                <input 
                                    v-model="devices.stage.ip" 
                                    type="text" 
                                    class="device-input" 
                                    placeholder="输入位移台IP地址"
                                />
                                <div class="device-buttons">
                                    <button 
                                        class="conn-btn" 
                                        :class="{ 'connected': devices.stage.connected }"
                                        @click="toggleDevice('stage')"
                                    >
                                        {{ devices.stage.connected ? '断开' : '连接' }}
                                    </button>
                                </div>
                            </div>
                            <div class="device-status" :class="{ 'connected': devices.stage.connected }">
                                <span class="status-dot"></span>
                                <span class="status-text">{{ devices.stage.connected ? '已连接' : '未连接' }}</span>
                            </div>
                        </div>

                        <!-- 高压电源 -->
                        <div class="device-item">
                            <label class="device-label">高压电源</label>
                            <div class="device-input-group">
                                <input 
                                    v-model="devices.power.ip" 
                                    type="text" 
                                    class="device-input" 
                                    placeholder="输入高压电源IP地址"
                                />
                                <div class="device-buttons">
                                    <button 
                                        class="conn-btn" 
                                        :class="{ 'connected': devices.power.connected }"
                                        @click="toggleDevice('power')"
                                    >
                                        {{ devices.power.connected ? '断开' : '连接' }}
                                    </button>
                                </div>
                            </div>
                            <div class="device-status" :class="{ 'connected': devices.power.connected }">
                                <span class="status-dot"></span>
                                <span class="status-text">{{ devices.power.connected ? '已连接' : '未连接' }}</span>
                            </div>
                        </div>

                        <!-- 探测器 -->
                        <div class="device-item">
                            <label class="device-label">探测器</label>
                            <div class="device-input-group">
                                <select v-model="devices.detector.selected" class="device-select">
                                    <option value="">请选择探测器</option>
                                    <option v-for="det in detectorList" :key="det.id" :value="det.id">
                                        {{ det.name }} ({{ det.serial }})
                                    </option>
                                </select>
                                <div class="device-buttons">
                                    <button class="refresh-btn" @click="refreshDetectors">
                                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                            <polyline points="23 4 23 10 17 10"/>
                                            <polyline points="1 20 1 14 7 14"/>
                                            <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15"/>
                                        </svg>
                                    </button>
                                    <button 
                                        class="conn-btn" 
                                        :class="{ 'connected': devices.detector.connected }"
                                        @click="toggleDevice('detector')"
                                    >
                                        {{ devices.detector.connected ? '断开' : '连接' }}
                                    </button>
                                </div>
                            </div>
                            <div class="device-status" :class="{ 'connected': devices.detector.connected }">
                                <span class="status-dot"></span>
                                <span class="status-text">{{ devices.detector.connected ? '已连接' : '未连接' }}</span>
                            </div>
                        </div>
                    </div>

                    <!-- <div class="modal-footer">
                        <button class="action-btn primary-btn" @click="connectAll">
                            一键连接
                        </button>
                        <button class="action-btn danger-btn" @click="disconnectAll">
                            断开全部
                        </button>
                    </div> -->
                </div>
            </div>
        </transition>
    </teleport>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { HighVoltageConnect, HighVoltageDisconnect} from '../../../wailsjs/go/components/HVPSService'
import { DetectorConnect, DetectorDisconnect,} from '../../../wailsjs/go/components/DetectorService'
import { StagesConnect, StagesDisconnect } from '../../../wailsjs/go/components/StageService'

defineProps({
    visible: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['close']);

const detectorList = ref([
    { id: 'ct-1', name: 'CT', serial: 'AGAB1U10R1' },
]);

const devices = reactive({
    stage: { ip: '192.168.11.6', connected: false },
    power: { ip: '192.168.11.4', connected: false },
    detector: { selected: 'ct-1', connected: false }
});



const handleClose = () => {emit('close');};

const toggleDevice = async (deviceType) => {
    const device = devices[deviceType];
    
    if (deviceType === 'detector') {
        if (!device.selected) {
            console.log('请选择探测器');
            return;
        }

        if (!device.connected) {
            try {
                await DetectorConnect();
                device.connected = true;
            } catch (error) {
                console.error('连接探测器失败:', error);
                return;
            }
        }else{
            try {
                await DetectorDisconnect();
                device.connected = false;
            } catch (error) {
                console.error('断开探测器失败:', error);
                return;
            }
        }

        return;
    }else if (deviceType === 'stage') {
        if (!device.connected) {
            try {
                await StagesConnect(device.ip);
                device.connected = true;
            } catch (error) {
                console.error('连接位移台失败:', error);
                return;
            }
        }else{
            try {
                await StagesDisconnect();
                device.connected = false;
            } catch (error) {
                console.error('断开位移台失败:', error);
            }
        }
    }else if (deviceType === 'power') {
        if (!device.connected) {
            try {
                await HighVoltageConnect(device.ip, 50001);
                device.connected = true;
            } catch (error) {
                console.error('连接 HVPS 失败:', error);
                return;
            }
        }else{
            try {
                await HighVoltageDisconnect();
                device.connected = false;
            } catch (error) {
                console.error('断开 HVPS 失败:', error);
                return;
            }
        }
    }

};

const refreshDetectors = () => {
    console.log('刷新探测器列表');
    // 模拟刷新探测器列表
};

</script>

<style scoped>
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(15, 23, 42, 0.45);
    backdrop-filter: blur(4px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal-container {
    width: 480px;
    background: #ffffff;
    border-radius: 10px;
    border: 1px solid #e2e8f0;
    box-shadow: 0 20px 50px rgba(15, 23, 42, 0.22), 0 4px 12px rgba(15, 23, 42, 0.1);
    overflow: hidden;
}

.modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 16px;
    background: linear-gradient(180deg, #f1f5f9 0%, #e8eef5 100%);
    border-bottom: 1px solid #e2e8f0;
}

.modal-title {
    font-size: 14px;
    font-weight: 600;
    color: #0f172a;
    letter-spacing: 0.02em;
}

.close-btn {
    width: 26px;
    height: 26px;
    border: 1px solid #e2e8f0;
    background: #fff;
    border-radius: 4px;
    color: #64748b;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.18s ease;
}
.close-btn:hover {
    background: #ef4444;
    border-color: #dc2626;
    color: #fff;
    transform: rotate(90deg);
}
.close-btn svg {
    width: 14px;
    height: 14px;
}

.modal-content {
    padding: 16px;
}

.device-item {
    margin-bottom: 16px;
    padding: 12px;
    background: #f8fafc;
    border: 1px solid #e2e8f0;
    border-radius: 6px;
}
.device-item:last-child {
    margin-bottom: 0;
}

.device-label {
    display: block;
    font-size: 12px;
    color: #0f172a;
    margin-bottom: 8px;
    font-weight: 600;
    letter-spacing: 0.02em;
}

.device-input-group {
    display: flex;
    align-items: center;
    gap: 8px;
}

.device-input,
.device-select {
    flex: 1;
    height: 30px;
    padding: 0 10px;
    background: #fff;
    border: 1px solid #cbd5e1;
    border-radius: 4px;
    color: #0f172a;
    font-size: 12px;
    font-weight: 500;
    outline: none;
    transition: border-color 0.15s ease, box-shadow 0.15s ease;
}
.device-input:focus,
.device-select:focus {
    border-color: #0284c7;
    box-shadow: 0 0 0 2px rgba(14, 165, 233, 0.15);
}
.device-input::placeholder {
    color: #94a3b8;
}

.device-select {
    cursor: pointer;
    appearance: none;
    -webkit-appearance: none;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%2364748b' stroke-width='2'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
    background-repeat: no-repeat;
    background-position: right 8px center;
    background-size: 14px;
    padding-right: 28px;
}
.device-select option {
    background: #fff;
    color: #0f172a;
}

.device-buttons {
    display: flex;
    gap: 6px;
}

.conn-btn {
    height: 30px;
    padding: 0 14px;
    background: #fff;
    border: 1px solid #0284c7;
    border-radius: 4px;
    color: #0284c7;
    font-size: 12px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.15s ease;
    letter-spacing: 0.02em;
}
.conn-btn:hover {
    background: #0284c7;
    color: #fff;
}
.conn-btn:active {
    transform: translateY(1px);
}
.conn-btn.connected {
    background: #f0fdf4;
    border-color: #15803d;
    color: #15803d;
}
.conn-btn.connected:hover {
    background: #15803d;
    color: #fff;
}

.refresh-btn {
    width: 30px;
    height: 30px;
    background: #fff;
    border: 1px solid #cbd5e1;
    border-radius: 4px;
    color: #64748b;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.15s ease;
}
.refresh-btn:hover {
    border-color: #0284c7;
    color: #0284c7;
}
.refresh-btn svg {
    width: 14px;
    height: 14px;
}

.device-status {
    display: flex;
    align-items: center;
    gap: 6px;
    margin-top: 8px;
    padding-left: 2px;
}

.status-dot {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background: #cbd5e1;
    transition: all 0.2s ease;
}
.device-status.connected .status-dot {
    background: #10b981;
    box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.18);
}

.status-text {
    font-size: 11px;
    color: #94a3b8;
    font-weight: 500;
}
.device-status.connected .status-text {
    color: #15803d;
}

.modal-footer {
    display: flex;
    justify-content: center;
    gap: 12px;
    padding: 12px 16px;
    background: #f8fafc;
    border-top: 1px solid #e2e8f0;
}

.action-btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    height: 34px;
    padding: 0 18px;
    min-width: 120px;
    border: 1px solid transparent;
    border-radius: 4px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.15s ease;
}
.action-btn svg {
    width: 14px;
    height: 14px;
}

.primary-btn {
    background: linear-gradient(180deg, #0284c7 0%, #0369a1 100%);
    border-color: #075985;
    color: #fff;
}
.primary-btn:hover {
    background: linear-gradient(180deg, #0369a1 0%, #075985 100%);
}
.primary-btn:active {
    transform: translateY(1px);
}

.danger-btn {
    background: #fff;
    border-color: #ef4444;
    color: #dc2626;
}
.danger-btn:hover {
    background: #ef4444;
    color: #fff;
}
.danger-btn:active {
    transform: translateY(1px);
}

/* 过渡动画 */
.modal-enter-active,
.modal-leave-active {
    transition: opacity 0.22s cubic-bezier(0.22, 1, 0.36, 1);
}
.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}
.modal-enter-active .modal-container,
.modal-leave-active .modal-container {
    transition: transform 0.22s cubic-bezier(0.22, 1, 0.36, 1);
}
.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
    transform: scale(0.96) translateY(-12px);
}
</style>
