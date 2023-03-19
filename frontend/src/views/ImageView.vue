<template>
    <a-spin :spinning="spinning">
        <div>
            <div class="table-operations">
                <a-button @click="handleRefresh">刷新镜像</a-button>
                <a-popconfirm :title="'确定删除已选的 ' + selectedRows.length + ' 个镜像?'" ok-text="确定" cancel-text="取消"
                    @confirm="handleDelete">
                    <a-button type="danger">删除镜像</a-button>
                </a-popconfirm>
            </div>
            <a-table :row-selection="{ onChange: (selectedRowKeys, _selectedRows) => selectedRows = _selectedRows }"
                row-key="ID" :columns="columns" :data-source="data" :scroll="{ x: 'max-content' }">
            </a-table>
        </div>
    </a-spin>
</template>
<script setup>
import { ref } from 'vue';
import { ImageList, ImageDelete } from "../../wailsjs/go/main/App";
import { message } from "ant-design-vue";

const spinning = ref(false);

const columns = [{
    title: 'REPOSITORY',
    dataIndex: 'Repository',
    fixed: 'left',
}, {
    title: 'TAG',
    dataIndex: 'Tag',
    fixed: 'left',
}, {
    title: 'IMAGE ID',
    dataIndex: 'ID',
}, {
    title: 'CREATED',
    dataIndex: 'Created',
}, {
    title: 'SIZE',
    dataIndex: 'Size',
}];

const data = ref([]);
const selectedRows = ref([]);

const handleRefresh = () => {
    ImageList()
        .then((images) => {
            data.value = images;
        })
        .catch(e => message.error(e));
};

const handleDelete = () => {
    spinning.value = true;
    let promises = [];
    selectedRows.value.forEach((row) => {
        let p = ImageDelete(row.ID)
        p.then(() => { message.success(`镜像${row.Repository}:${row.Tag}删除成功!`); }).catch((e) => message.error(`${row.Repository}:${row.Tag}失败: ${e}!`));
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