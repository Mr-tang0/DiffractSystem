<template>
    <div class="about-panel">
        <div class="panel-header">
            <span class="panel-title">关于</span>
            <button class="close-btn" @click="emit('close')">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
            </button>
        </div>
        
        <div class="panel-content">
            <div class="app-info">
                <div class="app-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                        <circle cx="12" cy="12" r="10"/>
                        <circle cx="12" cy="12" r="3"/>
                        <line x1="12" y1="2" x2="12" y2="5"/>
                        <line x1="12" y1="19" x2="12" y2="22"/>
                        <line x1="2" y1="12" x2="5" y2="12"/>
                        <line x1="19" y1="12" x2="22" y2="12"/>
                    </svg>
                </div>
                <h2 class="app-name">Diffract Control</h2>
                <p class="app-version">版本 1.0.0</p>
            </div>

            <div class="info-section">
                <p class="info-text">
                    Diffract Control 是一款专业的衍射CT控制系统，
                    用于材料科学、医疗检测等领域的三维成像研究。
                </p>
            </div>

            <div class="info-section">
                <h3 class="section-title">技术特性</h3>
                <ul class="feature-list">
                    <li>高精度样品台运动控制</li>
                    <li>实时图像采集与显示</li>
                    <li>自动化CT扫描流程</li>
                    <li>多格式图像保存支持</li>
                </ul>
            </div>

            <div class="info-section">
                <h3 class="section-title">开发信息</h3>
                <div class="info-row">
                    <span class="info-label">开发单位：</span>
                    <span class="info-value">宁波材料技术与工程研究所</span>
                </div>
                <div class="info-row">
                    <span class="info-label">框架支持：</span>
                    <span class="info-value">Wails + Vue 3</span>
                </div>
            </div>
        </div>

        <div class="panel-footer">
            <button class="action-btn primary-btn" @click="emit('close')">
                确定
            </button>
        </div>
    </div>
</template>

<script setup>
const emit = defineEmits(['close']);
</script>

<style scoped>
.about-panel {
    width: 380px;
    background: linear-gradient(180deg, rgba(255, 255, 255, 0.98) 0%, rgba(245, 247, 250, 0.98) 100%);
    border-radius: 12px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
    display: flex;
    flex-direction: column;
    max-height: 80vh;
    backdrop-filter: blur(10px);
    border: 1px solid rgba(56, 189, 248, 0.3);
}

.panel-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 14px 16px;
    border-bottom: 1px solid rgba(56, 189, 248, 0.3);
    background: rgba(56, 189, 248, 0.1);
    border-radius: 12px 12px 0 0;
}

.panel-title {
    font-size: 14px;
    font-weight: 600;
    color: #0284c7;
}

.close-btn {
    width: 26px;
    height: 26px;
    border: none;
    background: rgba(239, 68, 68, 0.15);
    border-radius: 6px;
    color: #dc2626;
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
    padding: 20px 16px;
    overflow-y: auto;
}

.app-info {
    text-align: center;
    margin-bottom: 20px;
}

.app-icon {
    width: 64px;
    height: 64px;
    margin: 0 auto 12px;
    color: #0284c7;
}

.app-icon svg {
    width: 100%;
    height: 100%;
}

.app-name {
    font-size: 18px;
    font-weight: 600;
    color: #1f2937;
    margin: 0 0 6px;
}

.app-version {
    font-size: 12px;
    color: #9ca3af;
    margin: 0;
}

.info-section {
    margin-bottom: 16px;
}

.section-title {
    font-size: 12px;
    font-weight: 600;
    color: #0284c7;
    margin-bottom: 8px;
}

.info-text {
    font-size: 12px;
    color: #6b7280;
    line-height: 1.6;
    text-align: center;
    margin: 0;
}

.feature-list {
    list-style: none;
    padding: 0;
    margin: 0;
}

.feature-list li {
    font-size: 12px;
    color: #4b5563;
    padding: 5px 0;
    padding-left: 16px;
    position: relative;
}

.feature-list li::before {
    content: '•';
    position: absolute;
    left: 0;
    color: #0284c7;
}

.info-row {
    display: flex;
    padding: 4px 0;
    font-size: 12px;
}

.info-label {
    color: #9ca3af;
    min-width: 70px;
}

.info-value {
    color: #4b5563;
}

.panel-footer {
    display: flex;
    justify-content: center;
    padding: 12px 16px;
    border-top: 1px solid rgba(56, 189, 248, 0.3);
}

.action-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px 24px;
    border: none;
    border-radius: 5px;
    font-size: 12px;
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

.panel-content::-webkit-scrollbar {
    width: 5px;
}

.panel-content::-webkit-scrollbar-track {
    background: rgba(0, 0, 0, 0.08);
    border-radius: 3px;
}

.panel-content::-webkit-scrollbar-thumb {
    background: rgba(56, 189, 248, 0.4);
    border-radius: 3px;
}
</style>
