<template>
    <a-spin :spinning="spinning">
        <div>
            <div class="table-operations">
                <a-button @click="handleRefresh">刷新数据卷</a-button>
                <a-popconfirm :title="'确定删除已选的 ' + selectedRows.length + ' 个数据卷?'" ok-text="确定" cancel-text="取消"
                    @confirm="handleDelete">
                    <a-button type="danger">删除数据卷</a-button>
                </a-popconfirm>
            </div>
            <a-table :row-selection="{ onChange: (selectedRowKeys, _selectedRows) => selectedRows = _selectedRows }"
            row-key="Name" :columns="columns" :data-source="data" :pagination="{ pageSize: 10 }" :scroll="{ x: 'max-content' }">
            </a-table>
        </div>
    </a-spin>
</template>
<script setup>
import { onBeforeMount, ref } from 'vue';
import { VolumeList, VolumeDelete } from "../../wailsjs/go/main/App";
import { message } from "ant-design-vue";

onBeforeMount(() => handleRefresh());

const spinning = ref(false);

const columns = [{
    title: 'NAME',
    dataIndex: 'Name',
    fixed: 'left',
    sorter: (a, b) => a.Name > b.Name,
}, {
    title: 'DRIVER',
    dataIndex: 'Driver',
    fixed: 'left',
}, {
    title: 'CreatedAt',
    dataIndex: 'CreatedAt',
    sorter: (a, b) => a.CreatedAt > b.CreatedAt,
}, {
    title: 'Mountpoint',
    dataIndex: 'Mountpoint',
}];

const data = ref([]);
const selectedRows = ref([]);

const handleRefresh = () => {
    VolumeList()
        .then((volumes) => {
            data.value = volumes;
        })
        .catch(e => message.error(e));
};

const handleDelete = () => {
    spinning.value = true;
    let promises = [];
    selectedRows.value.forEach((row) => {
        console.log(row);
        let p = VolumeDelete(row.Name)
        p.then(() => { message.success(`数据卷${row.Name}删除成功!`); }).catch((e) => message.error(`删除${row.Name}失败: ${e}!`));
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