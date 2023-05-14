<template>
    <a-spin :spinning="spinning">
        <div>
            <div class="table-operations">
                <a-button @click="handleRefresh">刷新容器</a-button>
            </div>
            <a-table row-key="ID" :columns="columns" :data-source="data" :scroll="{ x: 'max-content' }">
                <template #bodyCell="{ column, record }">
                    <template v-if="column.dataIndex === 'Names'">
                        <a-tag v-for="name in record.Names" color="geekblue">{{ name.substring(1) }}</a-tag>
                    </template>
                    <template v-else-if="column.dataIndex === 'State'">
                        <a-tag :color="record.State == 'running' ? 'green' : 'volcano'">{{ record.State }}</a-tag>
                    </template>
                    <template v-else-if="column.dataIndex === 'Ops'">
                        <a-button type="primary" @click="handleLog(record)">查看日志</a-button>
                    </template>
                </template>
            </a-table>
        </div>
        <a-drawer size="large" :title="drawerModel.title" :placement="right" :closable="false" v-model:visible="drawerModel.visible">
            <a-textarea :value="drawerModel.content"  autosize="true" readOnly />
        </a-drawer>
    </a-spin>
</template>
<script setup>
import { onBeforeMount, reactive, ref } from 'vue';
import { ContainerList, Log } from "../../wailsjs/go/main/App";
import { message } from "ant-design-vue";

onBeforeMount(() => handleRefresh());

const spinning = ref(false);
const drawerModel = reactive({
    title: "",
    content: "",
    visible: false
});

const columns = [{
    title: 'ID',
    dataIndex: 'ID',
    fixed: 'left',
}, {
    title: 'NAMES',
    dataIndex: 'Names',
    fixed: 'left',
}, {
    title: 'STATE',
    dataIndex: 'State',
    fixed: 'left',
}, {
    title: '操作',
    dataIndex: 'Ops',
},];

const data = ref([]);

const handleRefresh = () => {
    ContainerList()
        .then((containers) => {
            data.value = containers;
        })
        .catch(e => message.error(e));
};

const handleLog = (record) => {
    spinning.value = true;
    Log(record.ID)
        .then((log) => {
            drawerModel.title = `容器 - ${record.Names[0]} 日志`;
            drawerModel.content = log;
            drawerModel.visible = true;
        })
        .catch(e => message.error(e))
        .finally(() => { spinning.value = false; });
};

</script>

<style scoped>
.table-operations {
    margin-bottom: 16px;
}

.table-operations>button {
    margin-right: 8px;
}
</style>