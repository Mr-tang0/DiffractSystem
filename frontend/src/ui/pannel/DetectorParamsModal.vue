<template>
    <teleport to="body">
        <transition name="modal">
            <div v-if="visible" class="modal-overlay" @click.self="handleCancel">
                <div class="modal-container">
                    <div class="modal-header">
                        <span class="modal-title">探测器参数设置</span>
                        <button class="close-btn" @click="handleCancel">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>

                    <div class="modal-body">
                        <div class="form-row">
                            <label class="form-label">曝光时间</label>
                            <div class="form-control">
                                <input
                                    class="form-input"
                                    type="number"
                                    min="0"
                                    step="1"
                                    v-model.number="form.exposure"
                                />
                                <span class="form-unit">ms</span>
                            </div>
                        </div>

                        <div class="form-row">
                            <label class="form-label">增益 Gain</label>
                            <div class="form-control">
                                <input
                                    class="form-input"
                                    type="number"
                                    min="0"
                                    step="1"
                                    v-model.number="form.gain"
                                />
                            </div>
                        </div>

                        <div class="form-row">
                            <label class="form-label">像素合并 Binning</label>
                            <div class="form-control">
                                <select class="form-select" v-model="form.binning">
                                    <option value="1x1">1×1</option>
                                    <option value="2x2">2×2</option>
                                    <option value="4x4">4×4</option>
                                </select>
                            </div>
                        </div>

                        <div class="form-row">
                            <label class="form-label">重复次数</label>
                            <div class="form-control">
                                <input
                                    class="form-input"
                                    type="number"
                                    min="0"
                                    step="1"
                                    v-model.number="form.repeatTimes"
                                />
                                <span class="form-unit">次</span>
                            </div>
                        </div>
                    </div>

                    <div class="modal-footer">
                        <button class="action-btn cancel-btn" @click="handleCancel">
                            取消
                        </button>
                        <button
                            class="action-btn primary-btn"
                            :disabled="saving"
                            @click="handleSave"
                        >
                            {{ saving ? '保存中…' : '保存' }}
                        </button>
                    </div>
                </div>
            </div>
        </transition>
    </teleport>
</template>

<script setup>
import { reactive, ref, watch } from 'vue'
import { SetDynamicPara } from '../../../wailsjs/go/components/DetectorService'

const props = defineProps({
    visible: { type: Boolean, default: false },
    exposureTime: { type: Number, default: 0 },
    gain: { type: Number, default: 0 },
    binning: { type: String, default: '1x1' },
    repeatTimes: { type: Number, default: 0 }
})

const emit = defineEmits(['close', 'saved'])

const form = reactive({
    exposure: 0,
    gain: 0,
    binning: '1x1',
    repeatTimes: 0
})

const saving = ref(false)

// 打开时用传入的初始值填充表单
watch(
    () => props.visible,
    (val) => {
        if (val) {
            form.exposure = Number(props.exposureTime) || 0
            form.gain = Number(props.gain) || 0
            form.binning = props.binning || '1x1'
            form.repeatTimes = Number(props.repeatTimes) || 0
            saving.value = false
        }
    }
)

function handleCancel() {
    if (saving.value) return
    emit('close')
}

async function handleSave() {
    if (saving.value) return
    saving.value = true
    try {
        // SetDynamicPara(exposure int, binning string, repeatTimes uint16, gain uint16)
        await SetDynamicPara(form.exposure, form.binning, form.repeatTimes, form.gain)
        emit('saved', { ...form })
        emit('close')
    } catch (err) {
        console.error('[DetectorParamsModal] 保存失败:', err)
        saving.value = false
    }
}
</script>

<style scoped>
.modal-overlay {
    position: fixed;
    top: 0; left: 0; right: 0; bottom: 0;
    background: rgba(15, 23, 42, 0.45);
    backdrop-filter: blur(4px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal-container {
    width: 420px;
    background: #fff;
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
    width: 26px; height: 26px;
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
.close-btn svg { width: 14px; height: 14px; }

.modal-body {
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.form-row {
    display: flex;
    align-items: center;
    gap: 12px;
}
.form-label {
    width: 130px;
    flex-shrink: 0;
    font-size: 12px;
    font-weight: 600;
    color: #0f172a;
    letter-spacing: 0.02em;
}
.form-control {
    flex: 1;
    display: flex;
    align-items: center;
    gap: 8px;
}
.form-input,
.form-select {
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
.form-input:focus,
.form-select:focus {
    border-color: #0284c7;
    box-shadow: 0 0 0 2px rgba(14, 165, 233, 0.15);
}
.form-input::-webkit-inner-spin-button,
.form-input::-webkit-outer-spin-button {
    opacity: 0.4;
}
.form-select {
    cursor: pointer;
    appearance: none;
    -webkit-appearance: none;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%2364748b' stroke-width='2'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
    background-repeat: no-repeat;
    background-position: right 8px center;
    background-size: 14px;
    padding-right: 28px;
}
.form-unit {
    font-size: 11px;
    color: #94a3b8;
    font-weight: 500;
    flex-shrink: 0;
    min-width: 18px;
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
    height: 34px;
    padding: 0 24px;
    min-width: 100px;
    border: 1px solid transparent;
    border-radius: 4px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.15s ease;
    letter-spacing: 0.02em;
}
.cancel-btn {
    background: #fff;
    border-color: #cbd5e1;
    color: #475569;
}
.cancel-btn:hover {
    border-color: #94a3b8;
    color: #0f172a;
}
.primary-btn {
    background: linear-gradient(180deg, #0284c7 0%, #0369a1 100%);
    border-color: #075985;
    color: #fff;
}
.primary-btn:hover:not(:disabled) {
    background: linear-gradient(180deg, #0369a1 0%, #075985 100%);
}
.primary-btn:active:not(:disabled) {
    transform: translateY(1px);
}
.primary-btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

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
