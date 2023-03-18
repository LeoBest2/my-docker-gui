<template>
    <a-spin :spinning="spinning">
        <div>
            <div class="table-operations">
                <a-button @click="handleRefresh">刷新容器</a-button>
                <a-popconfirm :title="'确定启动已选的 ' + selectedRows.length + ' 个容器?'" ok-text="确定" cancel-text="取消"
                    @confirm="handleStart">
                    <a-button>启动容器</a-button>
                </a-popconfirm>

                <a-popconfirm :title="'确定停止已选的 ' + selectedRows.length + ' 个容器?'" ok-text="确定" cancel-text="取消"
                    @confirm="handleStop">
                    <a-button>停止容器</a-button>
                </a-popconfirm>

                <a-popconfirm :title="'确定删除已选的 ' + selectedRows.length + ' 个容器?'" ok-text="确定" cancel-text="取消"
                    @confirm="handleDelete">
                    <a-button type="danger">删除容器</a-button>
                </a-popconfirm>
            </div>
            <a-table :row-selection="{ onChange: (selectedRowKeys, _selectedRows) => selectedRows = _selectedRows }"
                row-key="ID" :columns="columns" :data-source="data" :scroll="{ x: 'max-content' }">
                <template #bodyCell="{ column, record }">
                    <template v-if="column.dataIndex === 'Names'">
                        <a-tag v-for="name in record.Names" color="geekblue">{{ name.substring(1) }}</a-tag>
                    </template>
                    <template v-else-if="column.dataIndex === 'State'">
                        <a-tag :color="record.State == 'running' ? 'green' : 'volcano'">{{ record.State }}</a-tag>
                    </template>
                    <template v-else-if="column.dataIndex === 'Ports'">
                        <a-tag v-for="port in record.Ports">
                            {{ port.IP }}:{{ port.PublicPort }}->{{ port.PrivatePort }}/{{ port.Type }}
                        </a-tag>
                    </template>
                </template>
            </a-table>
        </div>
    </a-spin>
</template>
<script setup>
import { ref } from 'vue';
import { ContainerList, ContainerStop, ContainerStart, ContainerDelete } from "../../wailsjs/go/main/App";
import { message } from "ant-design-vue";

const spinning = ref(false);

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
    title: 'IMAGE',
    dataIndex: 'Image',
}, {
    title: 'COMMAND',
    dataIndex: 'Command',
}, {
    title: 'CREATED',
    dataIndex: 'Created',
}
    , {
    title: 'STATUS',
    dataIndex: 'Status',
}, {
    title: 'PORTS',
    dataIndex: 'Ports',
}];

const data = ref([]);
const selectedRows = ref([]);

const handleRefresh = () => {
    ContainerList()
        .then((containers) => {
            data.value = containers;
        })
        .catch(e => message.error(e));
};

const handleStart = () => {
    spinning.value = true;
    let promises = [];
    selectedRows.value.forEach((row) => {
        let p = ContainerStart(row.ID)
        p.then(() => message.success(`容器${row.Names[0].substring(1)}启动成功!`)).catch((e) => message.error(`容器${row.Names[0].substring(1)}启动失败: ${e}!`));
        promises.push(p);
    });
    Promise.all(promises).finally(() => { spinning.value = false; handleRefresh() });
};

const handleStop = () => {
    spinning.value = true;
    let promises = [];
    selectedRows.value.forEach((row) => {
        let p = ContainerStop(row.ID)
        p.then(() => { message.success(`容器${row.Names[0].substring(1)}停止成功!`); }).catch((e) => message.error(`容器${row.Names[0].substring(1)}停止失败: ${e}!`));
        promises.push(p);
    });
    Promise.all(promises).finally(() => { spinning.value = false; handleRefresh() });
};

const handleDelete = () => {
    spinning.value = true;
    let promises = [];
    selectedRows.value.forEach((row) => {
        let p = ContainerDelete(row.ID)
        p.then(() => { message.success(`容器${row.Names[0].substring(1)}停止成功!`); }).catch((e) => message.error(`容器${row.Names[0].substring(1)}停止失败: ${e}!`));
        promises.push(p);
    });
    Promise.all(promises).finally(() => { spinning.value = false; handleRefresh() });
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