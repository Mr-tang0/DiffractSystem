<template>
    <div class="detector-monitor">
        
        <div class="monitor-content">
            <!-- 当前北京时间 -->
            <div class="monitor-item">
                <span class="item-label">时间</span>
                <span class="item-value time">{{ monitorData.currentTime }}</span>
            </div>
            
            <!-- 当前拍摄角度 -->
            <div class="monitor-item">
                <span class="item-label">角度</span>
                <span class="item-value angle">{{ monitorData.angle.toFixed(2) }}°</span>
            </div>
            
            <!-- 图像尺寸 -->
            <div class="monitor-item">
                <span class="item-label">图像尺寸</span>
                <span class="item-value size">{{ monitorData.imageWidth }} × {{ monitorData.imageHeight }}</span>
            </div>

             <!-- 图像Value -->
            <div class="monitor-item">
                <span class="item-label">图像Value</span>
                <span class="item-value value">{{ monitorData.imageValue }}</span>
            </div>
            
            <!-- 当前采集状态 -->
            <div class="monitor-item status-item">
                <span class="item-label">状态</span>
                <span class="item-value">{{ monitorData.status }}</span>
            </div>
        </div>
    </div>
</template>

<script setup>
import { reactive, onMounted, onUnmounted, watch } from 'vue';

const props = defineProps({
    imageWidth: {
        type: Number,
        default: null
    },
    imageHeight: {
        type: Number,
        default: null
    },
    angle: {
        type: Number,
        default: 0.00
    },
    imageValue: {
        type: Number,
        default: null
    },
    status: {
        type: String,
        default: '未运行'
    },
});

const monitorData = reactive({
    currentTime: '',
    angle: props.angle,
    imageValue: props.imageValue,
    imageWidth: props.imageWidth,
    imageHeight: props.imageHeight,
    status: props.status,
});

// 监听props变化
watch(() => props.imageWidth, (newVal) => {monitorData.imageWidth = newVal;});
watch(() => props.imageHeight, (newVal) => {monitorData.imageHeight = newVal;});
watch(() => props.imageValue, (newVal) => {monitorData.imageValue = newVal;});
watch(() => props.angle, (newVal) => {monitorData.angle = newVal;});
watch(() => props.status, (newVal) => { monitorData.status = newVal; });



let timer = null;
const updateTime = () => {
    const now = new Date();
    monitorData.currentTime = now.toLocaleString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false
    });
};


onMounted(() => {
    updateTime();
    timer = setInterval(() => {updateTime();}, 1000);
});

onUnmounted(() => {
    if (timer) {
        clearInterval(timer);
    }
});
</script>

<style scoped>
.detector-monitor {
    position: absolute;
    right: 16px;
    bottom: 16px;
    width: 220px;
    background: rgba(15, 23, 42, 0.55);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(56, 189, 248, 0.2);
    border-radius: 12px;
    padding: 14px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
    z-index: 10;
}

.monitor-header {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 12px;
    padding-bottom: 10px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.15);
}

.monitor-icon {
    width: 16px;
    height: 16px;
    color: #38bdf8;
}

.monitor-title {
    font-size: 13px;
    font-weight: 600;
    color: #38bdf8;
}

.monitor-content {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.monitor-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.item-label {
    font-size: 11px;
    color: #64748b;
}

.item-value {
    font-size: 12px;
    color: #e2e8f0;
    font-family: 'Courier New', monospace;
}

.item-value.serial {
    color: #94a3b8;
    font-size: 11px;
}

.item-value.time {
    color: #38bdf8;
}

.item-value.angle {
    color: #fbbf24;
}

.item-value.size {
    color: #a78bfa;
}

.status-item .item-value {
    padding: 2px 8px;
    border-radius: 10px;
    font-size: 11px;
    font-weight: 500;
}

.status-idle {
    background: rgba(148, 163, 184, 0.2);
    color: #94a3b8;
}

.status-exposing {
    background: rgba(234, 88, 12, 0.2);
    color: #f97316;
    animation: pulse 1.5s infinite;
}

.status-success {
    background: rgba(16, 185, 129, 0.2);
    color: #10b981;
}

@keyframes pulse {
    0%, 100% {
        opacity: 1;
    }
    50% {
        opacity: 0.6;
    }
}
</style>
