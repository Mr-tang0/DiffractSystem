<template>
    <div class="detector-panel" @click.self="emit('close')">
        <div class="panel-header">
            <span class="panel-title">采集设置</span>
        </div>
        
        
        <div class="panel-content">

            <!-- 第一组：电子枪参数 -->
            <div class="param-card">
                <div class="card-title">射线源参数</div>
                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">电源高压(KV)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="Params.Power.HV" 
                                class="param-input"
                                min="0" 
                                max="100"
                                placeholder="电压"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">电源电流(uA)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="Params.Power.HI" 
                                class="param-input"
                                min="0" 
                                max="2000"
                                placeholder="电流"
                            />
                        </div>
                    </div>
                </div>

                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">灯丝电压(V)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="Params.Power.FV" 
                                class="param-input"
                                min="0" 
                                max="5.5"
                                placeholder="电压"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">灯丝电流(A)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="Params.Power.FI" 
                                class="param-input"
                                min="0" 
                                max="3.6"
                                placeholder="电流"
                            />
                        </div>
                    </div>
                </div>
            </div>

            <!-- 第二组：加速管参数 -->
            <!-- <div class="param-card">
                <div class="card-title">加速管参数</div>
                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">电压(kV)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="acceleratorVoltage" 
                                class="param-input"
                                min="0" 
                                max="200"
                                placeholder="电压"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">电流(mA)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="acceleratorCurrent" 
                                class="param-input"
                                min="0" 
                                max="2000"
                                placeholder="电流"
                            />
                        </div>
                    </div>
                </div>
            </div> -->

            <!-- 第三组：探测器参数 -->
            <div class="param-card">
                <div class="card-title">探测器参数</div>

                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">曝光时间(ms)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="Params.Detector.ExposureTime" 
                                class="param-input"
                                min="1" 
                                max="50000"
                                placeholder="曝光时间"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">增益</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="Params.Detector.Gain" 
                                class="param-input"
                                min="1" 
                                max="100"
                                placeholder="增益"
                            />
                        </div>
                    </div>
                </div>
                
                <!-- Binning模式 -->
                <div class="param-group">
                    <label class="param-label">Binning</label>
                    <div class="binning-options">
                        <button 
                            v-for="bin in binningOptions" 
                            :key="bin.value"
                            class="binning-btn"
                            :class="{ active: Params.Detector.Binning === bin.value }"
                            @click="Params.Detector.Binning = bin.value"
                        >
                            {{ bin.label }}
                        </button>
                    </div>
                </div>

            </div>

            

            <!-- 应用按钮 -->
            <div class="action-buttons">
                <button class="get-btn" @click="getParams">
                    获取参数
                </button>
                <button class="apply-btn" @click="applyParams">
                    应用参数
                </button>
            </div>
            
        </div>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { DetectorSetExposeTime, DetectorGetExposeTime, 
    DetectorSetBinning, DetectorGetBinning, 
    DetectorSetGain, DetectorGetGain } from '../../wailsjs/go/services/DiffractDevice'

import { HVPSGetSetpointInfo,HVPSSetHV,HVPSSetHI,HVPSSetFilamentPreheat,HVPSSetFilamentLimit } from '../../wailsjs/go/services/HVPSDevice'


const emit = defineEmits(['close', 'apply']);

// 参数状态
const Params = reactive({
    Power:{
        HV: 0.00,
        HI: 0.00,
        FI: 0.00,
        FV: 0.00,
    },
    Detector:{
        ExposureTime: 50,
        Binning: '1x1',
        Gain: 30,
    }
});


const binningOptions = [
    { label: '1x1', value: '1x1' },
    { label: '2x2', value: '2x2' },
    { label: '3x3', value: '3x3' },
    { label: '4x4', value: '4x4' },
];

const getParams = async() => {
    console.log('getParams');

    // 获取HVPS参数
    let hvpsError = null;
    try {
        const pointInfo = await HVPSGetSetpointInfo();
        Params.Power.HV = pointInfo.HV;
        Params.Power.HI = pointInfo.HI;
        Params.Power.FV = pointInfo.FV;
        Params.Power.FI = pointInfo.FI;
    } catch (error) {
        hvpsError = error;
        console.error('获取HVPS参数失败:', error);
    }

    // 获取探测器参数 - 独立于HVPS，无论上面是否出错都会执行
    let detectorError = null;
    try {
        const exposeTime = await DetectorGetExposeTime();
        Params.Detector.ExposureTime = exposeTime;

        const binning = await DetectorGetBinning();
        Params.Detector.Binning = binning;

        const gain = await DetectorGetGain();
        Params.Detector.Gain = gain;
    } catch (error) {
        detectorError = error;
        console.error('获取探测器参数失败:', error);
    }
}

// 应用参数
const applyParams = async () => {
    console.log('applyParams');

    // try {
    //     await HVPSSetHV(Params.Power.HV);
    //     await HVPSSetHI(Params.Power.HI);
    //     await HVPSSetFilamentPreheat(Params.Power.FI);
    //     await HVPSSetFilamentLimit(Params.Power.FV);
    // } catch (error) {
    //     console.error('应用HVPS参数失败:', error);
    //     return;
    // }

    try {
        await DetectorSetExposeTime(Params.Detector.ExposureTime);
        await DetectorSetBinning(Params.Detector.Binning);
        await DetectorSetGain(Params.Detector.Gain);
    } catch (error) {
        console.error('应用探测器参数失败:', error);
        return;
    }
    emit('apply', params);
};

defineExpose({getParams});


</script>

<style scoped>
.detector-panel {
    width: 280px;
    background: linear-gradient(180deg, rgba(15, 23, 42, 0.95) 0%, rgba(30, 41, 59, 0.95) 100%);
    border-radius: 0 12px 12px 0;
    box-shadow: 4px 0 20px rgba(0, 0, 0, 0.3);
    display: flex;
    flex-direction: column;
    height: 100%;
    backdrop-filter: blur(10px);
    border-left: 1px solid rgba(56, 189, 248, 0.2);
}

.panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 20px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.2);
    background: rgba(56, 189, 248, 0.1);
}

.panel-title {
    font-size: 14px;
    font-weight: 600;
    color: #38bdf8;
}

.close-btn {
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
}

.close-btn:hover {
    background: rgba(239, 68, 68, 0.3);
    transform: rotate(90deg);
}

.panel-content {
    flex: 1;
    padding: 12px;
    overflow-y: auto;
}

/* 参数卡片样式 */
.param-card {
    background: rgba(15, 23, 42, 0.6);
    border: 1px solid rgba(56, 189, 248, 0.15);
    border-radius: 10px;
    padding: 12px;
    margin-bottom: 5px;
}

.card-title {
    font-size: 12px;
    font-weight: 600;
    color: #38bdf8;
    margin-bottom: 10px;
    padding-bottom: 6px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.2);
}

/* 参数行（一行两个参数） */
.param-row {
    display: flex;
    gap: 10px;
}

.param-row .param-item {
    flex: 1;
    min-width: 0;
}

.param-group {
    margin-bottom: 10px;
}

.param-group:last-child {
    margin-bottom: 0;
}

.param-label {
    display: block;
    font-size: 11px;
    color: #94a3b8;
    margin-bottom: 4px;
    font-weight: 500;
}

.param-input-wrap {
    display: flex;
    align-items: center;
    gap: 8px;
}

.param-input {
    flex: 1;
    height: 32px;
    padding: 0 10px;
    background: rgba(56, 189, 248, 0.1);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-radius: 6px;
    color: #ffffff;
    font-size: 13px;
    outline: none;
    transition: all 0.2s ease;
}

.param-input:focus {
    border-color: #38bdf8;
    box-shadow: 0 0 10px rgba(56, 189, 248, 0.3);
}

.param-unit {
    font-size: 12px;
    color: #64748b;
    min-width: 30px;
}


.binning-options {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 6px;
}

.binning-btn {
    padding: 8px;
    background: rgba(56, 189, 248, 0.1);
    border: 1px solid rgba(56, 189, 248, 0.3);
    border-radius: 6px;
    color: #94a3b8;
    font-size: 12px;
    cursor: pointer;
    transition: all 0.2s ease;
}

.binning-btn:hover {
    border-color: #38bdf8;
    color: #ffffff;
}

.binning-btn.active {
    background: rgba(56, 189, 248, 0.3);
    border-color: #38bdf8;
    color: #38bdf8;
}

.size-info {
    padding: 10px 12px;
    background: rgba(56, 189, 248, 0.1);
    border-radius: 8px;
    color: #94a3b8;
    font-size: 13px;
}

.action-buttons {
    display: flex;
    gap: 10px;
    margin-top: 24px;
}

.get-btn,
.apply-btn 
{
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    padding: 12px;
    border: none;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
}

.apply-btn {
    background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
    color: #ffffff;
}

.apply-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 15px rgba(56, 189, 248, 0.4);
}

.get-btn {
    background: rgba(148, 163, 184, 0.2);
    color: #94a3b8;
    border: 1px solid rgba(148, 163, 184, 0.3);
}

.get-btn:hover {
    background: rgba(148, 163, 184, 0.3);
    color: #ffffff;
}
</style>
