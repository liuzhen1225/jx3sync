<script lang="ts" setup>
import {Close, Minus, SwitchFilled} from "@element-plus/icons-vue";
import {Quit, WindowMinimise} from "../wailsjs/runtime";

</script>

<template>
  <div class="header">
    <div style="margin-left: 10px;">
      <el-icon size="21" style="margin-top: 13px;margin-right: 15px;"><SwitchFilled /></el-icon>
      <div class="name">角色同步</div>
    </div>
    <div v-if="true">
      <el-icon class="btn-sys" size="small" style="margin-right: 10px" @click="WindowMinimise()"><Minus /></el-icon>
      <el-icon class="btn-sys" size="small" style="margin-right: 10px" @click="Quit()"><Close /></el-icon>
    </div>
  </div>
  <router-view/>
</template>

<style lang="less">
@import "./assets/style/theme.less";
.header{
  background-color: @primary-color;
  --wails-draggable:drag;
  width: 100%;
  height: 48px;
  color: white;
  line-height: 48px;
  display: flex;
  justify-content: space-between;
}
.name{
  display: inline-block;
  font-size: 14px;
  user-select: none;
  vertical-align: top;
  margin-left: -10px;
}
.btn-sys{
  display: inline-block;
  color:white;
  :hover{
    cursor: pointer;
    color: black;
  }
}
</style>
