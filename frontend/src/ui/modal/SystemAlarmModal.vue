<template>
    <teleport to="body">
        <transition name="modal">
            <div v-if="visible" class="modal-overlay" @click.self="handleClose">
                <div class="modal-container">
                    <div class="modal-header">
                        <span class="modal-title">系统报警</span>
                        <button class="close-btn" @click="handleClose">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>

                    <div class="modal-body">
                        <div class="alarm-icon">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
                                <line x1="12" y1="9" x2="12" y2="13"/>
                                <line x1="12" y1="17" x2="12.01" y2="17"/>
                            </svg>
                        </div>
                        <ul class="alarm-list" style="text-align: left;">
                            <li>1、检查设备屏蔽门是否关闭</li>
                            <li>2、检查急停按钮是否被按下</li>
                            <li>3、检查门机联锁传感器</li>
                            <li>4、请联系管理员</li>
                        </ul>
                    </div>

                    <div class="modal-footer">
                        <button class="action-btn primary-btn" @click="handleClose">
                            确认
                        </button>
                    </div>
                </div>
            </div>
        </transition>
    </teleport>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'

defineProps({
    visible: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['close'])

const handleClose = () => {
    emit('close')
}
</script>

<style scoped>
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(127, 29, 29, 0.35);
    backdrop-filter: blur(4px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal-container {
    width: 420px;
    background: #ffffff;
    border-radius: 10px;
    border: 1px solid #fecaca;
    box-shadow: 0 20px 50px rgba(127, 29, 29, 0.28), 0 4px 12px rgba(127, 29, 29, 0.12);
    overflow: hidden;
}

.modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 16px;
    background: linear-gradient(180deg, #fef2f2 0%, #fee2e2 100%);
    border-bottom: 1px solid #fecaca;
}

.modal-title {
    font-size: 14px;
    font-weight: 600;
    color: #b91c1c;
    letter-spacing: 0.02em;
}

.close-btn {
    width: 26px;
    height: 26px;
    border: 1px solid #fecaca;
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
    background: #dc2626;
    border-color: #b91c1c;
    color: #fff;
    transform: rotate(90deg);
}
.close-btn svg {
    width: 14px;
    height: 14px;
}

.modal-body {
    padding: 18px 16px;
    display: flex;
    align-items: flex-start;
    gap: 14px;
}

.alarm-icon {
    flex: 0 0 auto;
    width: 42px;
    height: 42px;
    border-radius: 50%;
    background: #fef2f2;
    border: 1px solid #fecaca;
    color: #dc2626;
    display: flex;
    align-items: center;
    justify-content: center;
}
.alarm-icon svg {
    width: 24px;
    height: 24px;
}

.alarm-list {
    flex: 1;
    margin: 0;
    padding: 0;
    list-style: none;
}
.alarm-list li {
    font-size: 13px;
    color: #1f2937;
    line-height: 1.6;
    padding: 3px 0;
    font-weight: 500;
}
.alarm-list li:not(:last-child) {
    border-bottom: 1px dashed #e5e7eb;
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
    padding: 0 22px;
    min-width: 120px;
    border: 1px solid transparent;
    border-radius: 4px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.15s ease;
}

.primary-btn {
    background: linear-gradient(180deg, #dc2626 0%, #b91c1c 100%);
    border-color: #991b1b;
    color: #fff;
}
.primary-btn:hover {
    background: linear-gradient(180deg, #b91c1c 0%, #991b1b 100%);
}
.primary-btn:active {
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
