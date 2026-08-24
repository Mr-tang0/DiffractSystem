<template>
    <div class="help-panel">
        <div class="panel-header">
            <span class="panel-title">帮助文档</span>
            <button class="close-btn" @click="emit('close')">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
            </button>
        </div>
        
        <div class="panel-content">
            <div class="help-section">
                <h3 class="section-title">快速开始</h3>
                <ul class="help-list">
                    <li>1. 连接设备并配置系统参数</li>
                    <li>2. 创建新项目并设置扫描参数</li>
                    <li>3. 使用样品台定位功能调整位置</li>
                    <li>4. 开始CT扫描采集数据</li>
                </ul>
            </div>

            <div class="help-section">
                <h3 class="section-title">常见问题</h3>
                <div class="faq-item">
                    <p class="faq-q">Q: 如何连接设备？</p>
                    <p class="faq-a">A: 点击菜单栏"设备"，打开设备连接面板进行连接。</p>
                </div>
                <div class="faq-item">
                    <p class="faq-q">Q: 如何调整样品位置？</p>
                    <p class="faq-a">A: 使用右侧面板的样品台控制区域，可进行相对或绝对运动。</p>
                </div>
            </div>

            <div class="help-section">
                <h3 class="section-title">快捷键</h3>
                <div class="shortcuts">
                    <div class="shortcut-row">
                        <span class="shortcut-key">快捷键说明</span>
                        <span class="shortcut-desc">功能</span>
                    </div>
                </div>
            </div>
        </div>

        <div class="panel-footer">
            <button class="action-btn secondary-btn" @click="emit('close')">
                关闭
            </button>
        </div>
    </div>
</template>

<script setup>
const emit = defineEmits(['close']);
</script>

<style scoped>
.help-panel {
    width: 450px;
    background: linear-gradient(180deg, rgba(15, 23, 42, 0.95) 0%, rgba(30, 41, 59, 0.95) 100%);
    border-radius: 12px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
    display: flex;
    flex-direction: column;
    max-height: 80vh;
    backdrop-filter: blur(10px);
    border: 1px solid rgba(56, 189, 248, 0.2);
}

.panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 14px 16px;
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
    padding: 16px;
    overflow-y: auto;
}

.help-section {
    margin-bottom: 20px;
}

.help-section:last-child {
    margin-bottom: 0;
}

.section-title {
    font-size: 13px;
    font-weight: 600;
    color: #38bdf8;
    margin-bottom: 10px;
    padding-bottom: 6px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.2);
}

.help-list {
    list-style: none;
    padding: 0;
    margin: 0;
}

.help-list li {
    font-size: 12px;
    color: #cbd5e1;
    padding: 6px 0;
    line-height: 1.5;
}

.faq-item {
    margin-bottom: 12px;
}

.faq-q {
    font-size: 12px;
    color: #94a3b8;
    margin-bottom: 4px;
}

.faq-a {
    font-size: 12px;
    color: #cbd5e1;
    padding-left: 12px;
}

.shortcuts {
    background: rgba(0, 0, 0, 0.2);
    border-radius: 6px;
    padding: 8px;
}

.shortcut-row {
    display: flex;
    justify-content: space-between;
    padding: 6px 0;
    font-size: 12px;
}

.shortcut-key {
    color: #64748b;
}

.shortcut-desc {
    color: #cbd5e1;
}

.panel-footer {
    display: flex;
    justify-content: center;
    padding: 12px 16px;
    border-top: 1px solid rgba(56, 189, 248, 0.2);
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

.secondary-btn {
    background: rgba(71, 85, 105, 0.3);
    border: 1px solid rgba(148, 163, 184, 0.3);
    color: #94a3b8;
}

.secondary-btn:hover {
    background: rgba(71, 85, 105, 0.5);
    color: #cbd5e1;
}

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
</style>
