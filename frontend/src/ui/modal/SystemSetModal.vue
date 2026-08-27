<template>
    <teleport to="body">
        <transition name="modal">
            <div v-if="visible" class="modal-overlay" @click.self="handleCancel">
                <div class="modal-container">
                    <div class="modal-header">
                        <span class="modal-title">系统设置</span>
                        <button class="close-btn" @click="handleCancel">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>

                    <div class="modal-body">
                        <div class="table-head">
                            <span class="col-axis">轴</span>
                            <span class="col-speed">速度</span>
                            <span class="col-res">分辨率</span>
                        </div>
                        <div
                            v-for="axis in axes"
                            :key="axis.key"
                            class="axis-row"
                        >
                            <span class="axis-name">{{ axis.label }}</span>
                            <div class="form-control">
                                <input
                                    class="form-input"
                                    type="number"
                                    min="0"
                                    step="0.01"
                                    v-model.number="form[axis.key].speed"
                                />
                                <span class="form-unit">mm/s</span>
                            </div>
                            <div class="form-control">
                                <input
                                    class="form-input"
                                    type="number"
                                    min="0"
                                    step="0.0001"
                                    v-model.number="form[axis.key].resolution"
                                />
                                <span class="form-unit">pulse/mm</span>
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
import { StageSetSpeed, StageSetResolution, StageGetSpeed, StageGetResolution } from '../../../wailsjs/go/components/StageService'

const props = defineProps({
    visible: { type: Boolean, default: false }
})

const emit = defineEmits(['close', 'saved'])

const axes = [
    { key: 'X',  label: 'X' },
    { key: 'Y',  label: 'Y' },
    { key: 'Z',  label: 'Z' },
    { key: 'R',  label: 'R' },
    { key: 'XX', label: 'XX' }
]

const form = reactive({
    X:  { speed: 0, resolution: 0 },
    Y:  { speed: 0, resolution: 0 },
    Z:  { speed: 0, resolution: 0 },
    R:  { speed: 0, resolution: 0 },
    XX: { speed: 0, resolution: 0 }
})

const saving = ref(false)

const loading = ref(false)

watch(
    () => props.visible,
    async (val) => {
        if (val) {
            saving.value = false
            loading.value = true
            try {
                for (const axis of axes) {
                    const [speed, resolution] = await Promise.all([
                        StageGetSpeed(axis.key).catch(() => 0),
                        StageGetResolution(axis.key).catch(() => 0)
                    ])
                    form[axis.key].speed = Number(speed) || 0
                    form[axis.key].resolution = Number(resolution) || 0
                }
            } catch (err) {
                console.error('[SystemSetModal] 读取参数失败:', err)
            } finally {
                loading.value = false
            }
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
        for (const axis of axes) {
            const a = form[axis.key]
            await StageSetSpeed(axis.key, Number(a.speed) || 0)
            await StageSetResolution(axis.key, Number(a.resolution) || 0)
        }
        emit('saved', { ...form })
        emit('close')
    } catch (err) {
        console.error('[SystemSetModal] 保存失败:', err)
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
    width: 560px;
    max-height: calc(100vh - 80px);
    background: #fff;
    border-radius: 10px;
    border: 1px solid #e2e8f0;
    box-shadow: 0 20px 50px rgba(15, 23, 42, 0.22), 0 4px 12px rgba(15, 23, 42, 0.1);
    overflow: hidden;
    display: flex;
    flex-direction: column;
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
    padding: 14px 16px;
    display: flex;
    flex-direction: column;
    gap: 8px;
    overflow-y: auto;
}

.table-head {
    display: grid;
    grid-template-columns: 50px 1fr 1fr;
    gap: 12px;
    align-items: center;
    padding: 0 0 4px 2px;
    font-size: 11px;
    font-weight: 600;
    color: #64748b;
    letter-spacing: 0.04em;
    border-bottom: 1px solid #e2e8f0;
}
.col-axis { text-align: left; }
.col-speed,
.col-res { text-align: left; padding-left: 10px; }

.axis-row {
    display: grid;
    grid-template-columns: 50px 1fr 1fr;
    gap: 12px;
    align-items: center;
}
.axis-name {
    font-size: 13px;
    font-weight: 700;
    color: #0f172a;
    letter-spacing: 0.04em;
}

.form-control {
    display: flex;
    align-items: center;
    gap: 6px;
}
.form-input {
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
    min-width: 0;
    transition: border-color 0.15s ease, box-shadow 0.15s ease;
}
.form-input:focus {
    border-color: #0284c7;
    box-shadow: 0 0 0 2px rgba(14, 165, 233, 0.15);
}
.form-input::-webkit-inner-spin-button,
.form-input::-webkit-outer-spin-button {
    opacity: 0.4;
}
.form-unit {
    font-size: 11px;
    color: #94a3b8;
    font-weight: 500;
    flex-shrink: 0;
    min-width: 28px;
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
