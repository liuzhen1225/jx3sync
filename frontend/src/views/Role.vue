<template>
  <div>
    <el-input v-model="path" readonly placeholder="示例目录：D:\SeasunGame\Game\JX3">
      <template #prepend><span @click="selectDir" style="cursor: pointer">修改游戏目录</span></template>
      <template #append><span @click="getRoles" style="cursor: pointer">刷新角色</span></template>
    </el-input>
    <div  style="margin-top: 5px;width: 100%;"><el-text class="mx-1" type="warning" size="small"><el-icon><InfoFilled /></el-icon>程序从注册表读取目录，若不对，点修改目录手动指定。<el-link type="danger" style="font-size: 12px;margin-top: -1.8px" @click="clearSet">重置目录</el-link>可恢复注册表读取</el-text></div>
    <el-empty v-if="roles.length<=0" description="未查询到角色" />

    <div v-if="roles.length>0" style="margin: 10px auto 0;color: black;width: 90%;display: flex;justify-content: space-between">
      <el-select v-model="source" filterable placeholder="来源角色" style="width: 48%" @change="changeSync">
        <el-option
            v-for="item in roles"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        >
          <template #default>
            <span style="float: left">{{ item.label }}</span>
            <span style="float: right;color: var(--el-text-color-secondary);font-size: 13px;">{{ item.server }}</span>
          </template>
        </el-option>
      </el-select>
      <el-select v-model="target" filterable placeholder="目标角色" style="width: 48%;" @change="changeSync">
        <el-option
            v-for="item in roles"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        >
          <template #default>
            <span style="float: left">{{ item.label }}</span>
            <span style="float: right;color: var(--el-text-color-secondary);font-size: 13px;">{{ item.server }}</span>
          </template>
        </el-option>
      </el-select>
    </div>
    <div v-if="roles.length>0" style="width: 90%;color: gray;margin: 10px auto 0;display: flex;justify-content: space-between;">
      <el-card v-if="roles.length>0" style="width: 48%;" shadow="hover" class="box-card">
        <div style="font-weight: bold;font-size: 18px">{{roles[source].label}}</div>
        <div style="font-size: 14px;margin-top: 5px">
          <el-tag class="ml-2" type="success" style="margin-right: 10px">{{roles[source].server}}</el-tag>
          <el-tag class="ml-2" type="success">{{roles[source].bigServer}}</el-tag>
        </div>
        <div style="font-size: 14px;margin-top: 5px"><el-tag class="ml-2" type="success">{{roles[source].account}}</el-tag></div>
        <div style="font-size: 14px;margin-top: 5px">
          <el-tag class="ml-2" type="success" style="margin-right: 10px">{{roles[source].my==""?"茗伊 X":"茗伊 √"}}</el-tag>
          <el-tag class="ml-2" type="success" >{{roles[source].jx==""?"剑心 X":"剑心 √"}}</el-tag>
        </div>
      </el-card>
      <span style="margin-top: 65px;color: red">→</span>
      <el-card v-if="roles.length>0" style="width: 48%;" shadow="hover" class="box-card">
        <div style="font-weight: bold;font-size: 18px">{{roles[target].label}}</div>
        <div style="font-size: 14px;margin-top: 5px">
          <el-tag class="ml-2" type="danger" style="margin-right: 10px">{{roles[target].server}}</el-tag>
          <el-tag class="ml-2" type="danger">{{roles[target].bigServer}}</el-tag>
        </div>
        <div style="font-size: 14px;margin-top: 5px"><el-tag class="ml-2" type="danger">{{roles[target].account}}</el-tag></div>
        <div style="font-size: 14px;margin-top: 5px">
          <el-tag class="ml-2" type="danger" style="margin-right: 10px">{{roles[target].my==""?"茗伊 X":"茗伊 √"}}</el-tag>
          <el-tag class="ml-2" type="danger" >{{roles[target].jx==""?"剑心 X":"剑心 √"}}</el-tag>
        </div>
      </el-card>
    </div>
    <div v-if="roles.length>0" style="width: 90%;margin: 5px auto 0;display: flex;justify-content: space-between;">
      <div style="width: 70%">
        <el-row>
          <el-checkbox v-model="isUI" label="界面设置" size="large" />
          <el-checkbox v-model="isUI" label="按键设置" size="large" />
        </el-row>
        <el-row>
          <el-checkbox v-model="isMY" label="茗伊插件" size="large" />
          <el-checkbox v-model="isJX" label="剑心喊话" size="large" />
        </el-row>
      </div>
      <div style="width: 200px;margin-top: 8px;font-size: 12px;color: gray">
        <div style="color: green">界面数据：可以同步</div>
        <div style="color: green">按键数据：可以同步</div>
        <div :style="{color:roles[source].my==''||roles[target].my==''?'red':'green'}">茗伊插件：{{roles[source].my==""||roles[target].my==""?"不可同步":"可以同步"}}</div>
        <div :style="{color:roles[source].jx==''?'red':'green'}">剑心喊话：{{roles[source].jx==""?"不可同步":"可以同步"}}</div>
      </div>
      <div style="width: 30%">
        <div class="btn" @click="sync">
          同步
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
  import {onMounted, ref} from "vue";
  import {GetPath, GetRoles, SelectPath, SyncRoleData} from "../../wailsjs/go/service/Role";
  import {ElLoading, ElMessageBox} from "element-plus";
  import {InfoFilled} from "@element-plus/icons-vue";

  const path = ref("")
  const roles = ref<RoleInterface[]>([])
  const source = ref(0)
  const target = ref(0)

  const isUI = ref(true)
  const isMY = ref(true)
  const isJX = ref(true)
  const selectDir = async () => {
    const p = await SelectPath();
    if(p == ""){
      return;
    }
    path.value = p;
    localStorage.setItem("path",p)
    await getPath()
  }

  const changeSync = ()=>{
    isUI.value = true
    isMY.value = !(roles.value[source.value].my == "" || roles.value[target.value].my == "")
    isJX.value = !(roles.value[source.value].jx == "")
  }

  const clearSet = ()=>{
    ElMessageBox.confirm(
        `确认清楚路径配置文件？`,
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          roundButton:true,
          closeOnClickModal:false
        }
    )
        .then(async () => {
          localStorage.clear()
          await getPath()
          await ElMessageBox.alert("重置成功")
        })
        .catch(() => {

        })

  }


  const sync = ()=>{
    ElMessageBox.confirm(
        `确认覆盖角色【${roles.value[target.value].label}】配置文件？`,
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          roundButton:true,
          closeOnClickModal:false
        }
    )
        .then(async () => {
          if(source.value == target.value){
            await ElMessageBox.alert("同一角色之间不可同步")
            return;
          }
          const res = await SyncRoleData(
              roles.value[source.value].label,
              roles.value[target.value].label,
              roles.value[source.value].path,
              roles.value[target.value].path,
              roles.value[source.value].my,
              roles.value[target.value].my,
              isUI.value,isMY.value,isJX.value
          )
          if(res == "success"){
            await ElMessageBox.alert("同步成功")
          } else {
            await ElMessageBox.alert(res)
          }

        })
        .catch(() => {

        })
  }

  const getRoles=async () => {
    const loading = ElLoading.service({
      lock: true,
      text: 'Loading',
      background: 'rgba(0, 0, 0, 0.7)',
    })
    let roleArr = await GetRoles()
    roles.value = []
    for (let i = 0; i < roleArr.length; i++) {
      const role = roleArr[i].split("|")
      roles.value.push({label: role[0], value: i, server: role[1], bigServer: role[2], account: role[3], path: role[4],my:role[5],jx:role[6]})
    }
    source.value = 0
    target.value = 0
    changeSync()
    loading.close()
  }
  const  getPath = async () => {
    let pathDir = localStorage.getItem("path") || ""
    path.value = await GetPath(pathDir)
    await getRoles()
  }
  onMounted(async () => {
    await getPath()
  })
</script>
<style lang="less">
@import "../assets/style/theme";
.btn{
  width: 80px;
  height:80px;
  border-radius: 50%;
  float:right;
  line-height: 80px;
  box-shadow: 0 0 10px rgba(0,0,0,0.3);
  cursor: pointer;
  user-select: none;
  color: @primary-color;
  font-size: 14px;
  font-weight: bold;
}
.btn:active{
  box-shadow: 0 0 2px rgba(0,0,0,0.3);
}
</style>