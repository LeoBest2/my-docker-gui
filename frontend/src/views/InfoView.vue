<template>
    <a-badge :count="info.Containers" color="green" show-zero>
        <a-tag color="success">
            总容器数
        </a-tag>
    </a-badge>
    <a-divider type="vertical" style="background-color: #7cb305; margin-left: 2rem; margin-right: 2rem; height: 2rem;" />
    <a-badge :count="info.ContainersRunning" color="green" show-zero>
        <a-tag color="success">
            运行的容器数
        </a-tag>
    </a-badge>
    <a-divider type="vertical" style="background-color: #7cb305; margin-left: 2rem; margin-right: 2rem; height: 2rem;" />
    <a-badge :count="info.ContainersPaused" color="green" show-zero>
        <a-tag color="success">
            暂停的容器数
        </a-tag>
    </a-badge>
    <a-divider type="vertical" style="background-color: #7cb305; margin-left: 2rem; margin-right: 2rem; height: 2rem;" />
    <a-badge :count="info.ContainersStopped" color="green" show-zero>
        <a-tag color="success">
            停止的容器数
        </a-tag>
    </a-badge>
    <a-divider type="vertical" style="background-color: #7cb305; margin-left: 2rem; margin-right: 2rem; height: 2rem;" />
    <a-badge :count="info.Images" color="green" show-zero>
        <a-tag color="success">
            镜像数量
        </a-tag>
    </a-badge>
    <a-divider />
    <a-list item-layout="horizontal">
        <a-list-item>
            <a-list-item-meta :description="info.IndexServerAddress">
                <template #title>Registry:</template>
            </a-list-item-meta>
        </a-list-item>
        <a-list-item v-for="v, i in info.RegistryConfig && info.RegistryConfig.Mirrors">
            <a-list-item-meta :description="v">
                <template #title>Registry Mirrors {{ i }}:</template>
            </a-list-item-meta>
        </a-list-item>
        <a-list-item>
            <a-list-item-meta :description="info.KernelVersion">
                <template #title>Kernel Version::</template>
            </a-list-item-meta>
        </a-list-item>
        <a-list-item>
            <a-list-item-meta :description="info.OperatingSystem">
                <template #title>Operating System:</template>
            </a-list-item-meta>
        </a-list-item>
        <a-list-item>
            <a-list-item-meta :description="info.NCPU">
                <template #title>CPUs:</template>
            </a-list-item-meta>
        </a-list-item>
        <a-list-item>
            <a-list-item-meta :description="(info.MemTotal / 1024 / 1024 / 1024).toFixed(2) + 'GB'">
                <template #title>Total Memory:</template>
            </a-list-item-meta>
        </a-list-item>
        <a-list-item>
            <a-list-item-meta :description="info.SystemTime">
                <template #title>当前时间:</template>
            </a-list-item-meta>
        </a-list-item>
    </a-list>
</template>
<script setup>
import { message } from "ant-design-vue";
import { Info } from "../../wailsjs/go/main/App";
import { onBeforeMount, ref } from "vue";

const updateInfo = () => {
    Info().then((d) => {
        info.value = d;
        console.log(d);
    }).catch((e) => {
        message.error(`获取信息错误: ${e}!`);
    });
};

const info = ref({});
onBeforeMount(() => { updateInfo() });
</script>