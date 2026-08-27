<template>
    <teleport to="body">
        <transition name="modal">
            <div v-if="visible" class="modal-overlay" @click.self="handleClose">
                <div class="modal-container">
                    <div class="modal-header">
                        <div class="header-left">
                            <svg class="header-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
                                <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
                            </svg>
                            <span class="modal-title">历史报警消息</span>
                            <span class="msg-count">{{ messages.length }} 条</span>
                        </div>
                        <button class="close-btn" @click="handleClose">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <line x1="18" y1="6" x2="6" y2="18"/>
                                <line x1="6" y1="6" x2="18" y2="18"/>
                            </svg>
                        </button>
                    </div>

                    <div class="modal-body">
                        <div v-if="messages.length === 0" class="empty-history">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
                                <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
                            </svg>
                            <span>暂无历史消息</span>
                        </div>
                        <ul v-else class="history-list">
                            <li v-for="(item, idx) in messages" :key="idx" class="history-item">
                                <span class="item-time">{{ item.time }}</span>
                                <span class="item-msg">{{ item.msg }}</span>
                            </li>
                        </ul>
                    </div>

                    <div class="modal-footer">
                        <button class="action-btn clear-btn" @click="emit('clear')">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <polyline points="3 6 5 6 21 6"/>
                                <path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"/>
                                <path d="M10 11v6M14 11v6"/>
                            </svg>
                            <span>清除消息</span>
                        </button>
                        <button class="action-btn export-btn" @click="emit('export')">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                                <polyline points="7 10 12 15 17 10"/>
                                <line x1="12" y1="15" x2="12" y2="3"/>
                            </svg>
                            <span>导出消息</span>
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
    },
    messages: {
        type: Array,
        default: () => []
    }
})

const emit = defineEmits(['close', 'clear', 'export'])

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
    background: rgba(15, 23, 42, 0.35);
    backdrop-filter: blur(4px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal-container {
    width: 480px;
    height: 520px;
    display: flex;
    flex-direction: column;
    background: #ffffff;
    border-radius: 10px;
    border: 1px solid #e2e8f0;
    box-shadow: 0 20px 50px rgba(15, 23, 42, 0.25), 0 4px 12px rgba(15, 23, 42, 0.1);
    overflow: hidden;
}

.modal-header {
    flex: 0 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 16px;
    background: linear-gradient(180deg, #f1f5f9 0%, #e8eef5 100%);
    border-bottom: 1px solid #e2e8f0;
}

.header-left {
    display: flex;
    align-items: center;
    gap: 8px;
}

.header-icon {
    width: 18px;
    height: 18px;
    color: #0284c7;
}

.modal-title {
    font-size: 14px;
    font-weight: 600;
    color: #0f172a;
    letter-spacing: 0.02em;
}

.msg-count {
    font-size: 11px;
    font-weight: 600;
    color: #0284c7;
    background: rgba(2, 132, 199, 0.1);
    border: 1px solid rgba(2, 132, 199, 0.25);
    padding: 1px 8px;
    border-radius: 10px;
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
    background: #0284c7;
    border-color: #0369a1;
    color: #fff;
    transform: rotate(90deg);
}
.close-btn svg {
    width: 14px;
    height: 14px;
}

.modal-body {
    flex: 1;
    min-height: 0;
    padding: 12px 14px;
    overflow: hidden;
    display: flex;
    flex-direction: column;
}

.empty-history {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 10px;
    color: #94a3b8;
    font-size: 13px;
}
.empty-history svg {
    width: 44px;
    height: 44px;
    opacity: 0.6;
}

.history-list {
    flex: 1;
    min-height: 0;
    margin: 0;
    padding: 4px 6px;
    list-style: none;
    overflow-y: auto;
    scrollbar-width: thin;
    scrollbar-color: #cbd5e1 transparent;
}
.history-list::-webkit-scrollbar {
    width: 8px;
}
.history-list::-webkit-scrollbar-track {
    background: transparent;
}
.history-list::-webkit-scrollbar-thumb {
    background: #cbd5e1;
    border-radius: 4px;
    border: 2px solid transparent;
    background-clip: padding-box;
}
.history-list::-webkit-scrollbar-thumb:hover {
    background: #94a3b8;
    background-clip: padding-box;
}

.history-item {
    display: flex;
    align-items: baseline;
    gap: 10px;
    padding: 8px 6px;
    border-bottom: 1px dashed #e2e8f0;
    font-size: 13px;
    line-height: 1.5;
}
.history-item:last-child {
    border-bottom: none;
}

.item-time {
    flex: 0 0 auto;
    font-size: 11px;
    font-weight: 600;
    font-variant-numeric: tabular-nums;
    color: #0284c7;
    background: rgba(2, 132, 199, 0.08);
    padding: 1px 6px;
    border-radius: 3px;
    white-space: nowrap;
}

.item-msg {
    color: #1e293b;
    font-weight: 500;
    word-break: break-all;
}

.modal-footer {
    flex: 0 0 auto;
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
    padding: 0 20px;
    min-width: 120px;
    border: 1px solid transparent;
    border-radius: 4px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.15s ease;
    color: #fff;
}
.action-btn svg {
    width: 14px;
    height: 14px;
}
.action-btn:active {
    transform: translateY(1px);
}

.clear-btn {
    background: linear-gradient(180deg, #ef4444 0%, #dc2626 100%);
    border-color: #b91c1c;
}
.clear-btn:hover {
    background: linear-gradient(180deg, #dc2626 0%, #b91c1c 100%);
}

.export-btn {
    background: linear-gradient(180deg, #0284c7 0%, #0369a1 100%);
    border-color: #075985;
}
.export-btn:hover {
    background: linear-gradient(180deg, #0369a1 0%, #075985 100%);
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
