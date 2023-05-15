<template>
  <a-layout class="layout">
    <a-layout-header>
      <div class="logo">Docker Tool</div>
      <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="horizontal" :style="{ lineHeight: '64px' }">
        <a-menu-item key="1"><router-link to="/info">版本信息</router-link></a-menu-item>
        <a-menu-item key="2"><router-link to="/container">容器管理</router-link></a-menu-item>
        <a-menu-item key="3"><router-link to="/image">镜像管理</router-link></a-menu-item>
        <a-menu-item key="4"><router-link to="/volume">数据卷管理</router-link></a-menu-item>
        <a-menu-item key="5"><router-link to="/log">日志信息</router-link></a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content style="padding: 16px 50px">
      <router-view></router-view>
    </a-layout-content>
    <a-layout-footer style="text-align: center">
      Docker 管理工具 ©2023 Created by LeoForBest
    </a-layout-footer>
  </a-layout>
</template>

<script setup>
import { ref } from 'vue';
const selectedKeys = ref(['1']);
</script>

<style scoped>
.logo {
  float: left;
  width: 200px;
  font-size: 20px;
  color: yellowgreen;
}
</style>