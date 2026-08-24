<template>
    <div class="project-panel">
        <div class="panel-header">
            <span class="panel-title">项目配置</span>
            <button class="close-btn" @click="emit('close')">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
            </button>
        </div>
        
        <div class="panel-content">
            <!-- 基本信息 -->
            <div class="param-card">
                <div class="card-title">基本信息</div>
                
                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">项目ID</label>
                        <div class="param-input-wrap">
                            <input 
                                type="text" 
                                v-model="projectData.projectId" 
                                class="param-input readonly"
                                readonly
                                placeholder="项目ID"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">使用时间</label>
                        <div class="param-input-wrap">
                            <input 
                                type="text" 
                                v-model="projectData.createTime" 
                                class="param-input readonly"
                                readonly
                                placeholder="使用时间"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">实验人</label>
                        <div class="param-input-wrap">
                            <input 
                                type="text" 
                                v-model="projectData.userName" 
                                class="param-input"
                                placeholder="请输入实验人姓名"
                            />
                        </div>
                    </div>
                </div>
            </div>

            <!-- 样品信息 -->
            <div class="param-card">
                <div class="card-title">样品信息</div>
                
                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">样品名称</label>
                        <div class="param-input-wrap">
                            <input 
                                type="text" 
                                v-model="projectData.sampleName" 
                                class="param-input"
                                placeholder="请输入样品名称"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">备注信息</label>
                        <div class="param-input-wrap">
                            <input 
                                type="text" 
                                v-model="projectData.remark" 
                                class="param-input"
                                placeholder="请输入备注信息"
                            />
                        </div>
                    </div>
                </div>
            </div>

            <!-- 采集信息 -->
            <div class="param-card">
                <div class="card-title">采集信息</div>
                
                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">CT扫描角度(°)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="projectData.ctProject.totalScanAngle" 
                                class="param-input"
                                min="1" 
                                max="720"
                                placeholder="扫描角度"
                                @input="calculateStep"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">采集张数</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="projectData.ctProject.acquisitionNum" 
                                class="param-input"
                                min="1" 
                                max="10000"
                                placeholder="采集张数"
                                @input="calculateStep"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">采集步长(°)</label>
                        <div class="param-input-wrap">
                            <input 
                                type="number" 
                                v-model.number="projectData.ctProject.angularStep" 
                                class="param-input readonly"
                                readonly
                                step="0.001"
                                placeholder="采集步长"
                            />
                        </div>
                    </div>
                </div>
            </div>

            <!-- 文件信息 -->
            <div class="param-card">
                <div class="card-title">文件信息</div>
                
                <div class="param-row">
                    <div class="param-item">
                        <label class="param-label">文件名</label>
                        <div class="param-input-wrap">
                            <input 
                                type="text" 
                                v-model="projectData.fileName" 
                                class="param-input"
                                placeholder="文件名"
                            />
                        </div>
                    </div>
                    <div class="param-item">
                        <label class="param-label">文件路径</label>
                        <div class="param-input-wrap">
                            <input 
                                type="text" 
                                v-model="projectData.filePath" 
                                class="param-input"
                                placeholder="文件路径"
                            />
                            <button class="path-select-btn" @click="selectFilePath">
                                ...
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <!-- 按钮组 -->
            <div class="btn-group">
                <button class="action-btn primary-btn" @click="saveProject">
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
import { LoadHistoryProject, SaveHistoryProject, SelectDirectory } from '../../wailsjs/go/services/Project';

const emit = defineEmits(['close', 'save']);

// 项目数据
const projectData = reactive({
    userName: 'user',
    projectId: '',
    createTime: '',
    sampleName: '',
    remark: '',
    ctProject: {
        totalScanAngle: 360,
        angularStep: 1,
        acquisitionNum: 360
    },
    fileName: '',
    filePath: './projects/'
});

// 计算采集步长
const calculateStep = () => {
    const angle = projectData.ctProject.totalScanAngle;
    const num = projectData.ctProject.acquisitionNum;
    if (angle > 0 && num > 0) {
        projectData.ctProject.angularStep = parseFloat((angle / num).toFixed(3));
    }
};

// 选择文件路径
const selectFilePath = async () => {
    try {
        const selectedPath = await SelectDirectory();
        if (selectedPath) {
            projectData.filePath = selectedPath;
        }
    } catch (err) {
        console.log('选择路径失败:', err);
    }
};

// 加载项目配置
const loadProject = async () => {
    try{
       const project = await LoadHistoryProject(); 
       projectData.userName = project.user_name || 'user';
       projectData.projectId = Date.now().toString();
       projectData.createTime = new Date().toLocaleString();
       projectData.sampleName = project.sample_name || '';
       projectData.remark = project.remark || '';
       projectData.ctProject.totalScanAngle = project.ct_project?.total_scan_angle || 360;
       projectData.ctProject.angularStep = project.ct_project?.angular_step || 1;
       projectData.ctProject.acquisitionNum = project.ct_project?.acquisition_num || 360;
       projectData.fileName = project.file_name || '';
       projectData.filePath = project.file_path || './projects/';
    }catch(err){
        console.log(err);
    }
};

// 保存项目配置
const saveProject = async () => {
    try{
        const project = {
            user_name: projectData.userName,
            project_id: projectData.projectId,
            create_time: projectData.createTime,
            sample_name: projectData.sampleName,
            remark: projectData.remark,
            ct_project: {
                total_scan_angle: projectData.ctProject.totalScanAngle,
                angular_step: projectData.ctProject.angularStep,
                acquisition_num: projectData.ctProject.acquisitionNum
            },
            file_name: projectData.fileName,
            file_path: projectData.filePath
        };
        console.log('[ProjectPanel] 发送的数据:', JSON.stringify(project, null, 2));
        await SaveHistoryProject(project);
        emit('save', project);
    }catch(err){
        console.log(err);
    }
};

// 组件挂载时加载配置
onMounted(() => {
    loadProject();
});
</script>

<style scoped>
.project-panel {
    width: 650px;
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
    padding: 16px 20px;
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

.close-btn svg {
    width: 16px;
    height: 16px;
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
    margin-bottom: 10px;
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
    margin-bottom: 10px;
}

.param-row:last-child {
    margin-bottom: 0;
}

.param-row .param-item {
    flex: 1;
    min-width: 0;
}

.param-label {
    display: block;
    font-size: 11px;
    color: #94a3b8;
    margin-bottom: 4px;
}

.param-input-wrap {
    position: relative;
}

.param-input {
    width: 100%;
    height: 32px;
    padding: 0 10px;
    background: rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(56, 189, 248, 0.2);
    border-radius: 6px;
    color: #ffffff;
    font-size: 13px;
    outline: none;
    transition: all 0.2s ease;
    box-sizing: border-box;
}

.param-input:focus {
    border-color: #38bdf8;
    box-shadow: 0 0 0 2px rgba(56, 189, 248, 0.2);
}

.param-input.readonly {
    background: rgba(0, 0, 0, 0.1);
    color: #64748b;
    cursor: not-allowed;
}

.param-input::placeholder {
    color: #475569;
}

/* 路径选择按钮 */
.path-select-btn {
    position: absolute;
    right: 2px;
    top: 50%;
    transform: translateY(-50%);
    width: 24px;
    height: 24px;
    background: rgba(56, 189, 248, 0.2);
    border: none;
    border-radius: 4px;
    color: #38bdf8;
    font-size: 12px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background 0.2s ease;
}

.path-select-btn:hover {
    background: rgba(56, 189, 248, 0.3);
}

/* 按钮组 */
.btn-group {
    display: flex;
    justify-content: center;
    gap: 12px;
    margin-top: 12px;
}

.action-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px 20px;
    border: none;
    border-radius: 6px;
    font-size: 12px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s ease;
    max-height: 40px;
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
    width: 6px;
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
