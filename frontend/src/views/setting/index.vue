<template>
  <div class="container">
    <div class="left-side">
      <div class="card">
        <a-form :model="formData" layout="vertical" size="large">
          <a-row :gutter="16">
            <a-col :span="16">
              <a-form-item field="theme" label="主题">
                <a-radio-group
                  v-model="formData.theme"
                  type="button"
                  @change="handleTheme"
                >
                  <a-radio value="light">亮色</a-radio>
                  <a-radio value="dark">黑色</a-radio>
                  <a-radio value="auto">跟随系统</a-radio>
                </a-radio-group>
              </a-form-item>
            </a-col>
            <a-col :span="8">
              <a-form-item field="concurrency" label="LCU并发数">
<!--                <a-input-number-->
<!--                  v-model="formData.concurrency"-->
<!--                  :min="1"-->
<!--                  :max="10"-->
<!--                  @change="saveConcurrency"-->
<!--                />-->
              </a-form-item>
            </a-col>
            <a-col :span="8">
              <a-form-item field="SetAlwaysOnTop" label="置顶窗口">
                <a-switch type="round" @change="changeSetAlwaysOnTop">
                  <template #checked> 开启 </template>
                  <template #unchecked> 关闭 </template>
                </a-switch>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <!--              <a-form-item field="SetFrameless" label="无边框窗口">-->
              <!--                <a-switch type="round" @change="changeSetFrameless">-->
              <!--                  <template #checked>-->
              <!--                    开启-->
              <!--                  </template>-->
              <!--                  <template #unchecked>-->
              <!--                    关闭-->
              <!--                  </template>-->
              <!--                </a-switch>-->
              <!--              </a-form-item>-->
            </a-col>
            <a-col :span="12">
              <a-form-item
                field="SetFullscreenButtonEnabled"
                label="窗口是否可调整大小"
              >
                <a-switch
                  type="round"
                  :default-checked="true"
                  @change="changeSetResizable"
                >
                  <template #checked> 可以 </template>
                  <template #unchecked> 不可 </template>
                </a-switch>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="Reload" label="重置窗口">
                <a-button
                  type="primary"
                  @click="() => Window.Reload()"
                  size="small"
                >
                  <template #icon>
                    <icon-translate />
                  </template>
                  <template #default>立即重置</template>
                </a-button>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="OpenDevTools" label="打开调试工具">
                <a-button
                  type="primary"
                  @click="() => Window.OpenDevTools()"
                  size="small"
                >
                  <template #icon>
                    <icon-launch />
                  </template>
                  <template #default>立即打开</template>
                </a-button>
              </a-form-item>
            </a-col>
          </a-row>
        </a-form>
      </div>
    </div>
    <div class="right-side">
      <div class="card">
        <a-descriptions title="系统信息" :column="1">
          <a-descriptions-item label="窗口name">
            {{ windowsname }}</a-descriptions-item
          >
          <a-descriptions-item label="系统(OS)">
            {{ Osinfo?.OS }}</a-descriptions-item
          >
          <a-descriptions-item label="系统品牌">
            {{ Osinfo?.OSInfo.Branding }}</a-descriptions-item
          >
          <a-descriptions-item label="系统版本">
            {{ Osinfo?.OSInfo.Version }}</a-descriptions-item
          >
          <a-descriptions-item label="Arch">
            {{ Osinfo?.Arch }}</a-descriptions-item
          >
          <a-descriptions-item label="调试状态">
            {{ Osinfo?.Debug }}</a-descriptions-item
          >
          <a-descriptions-item label="WebView2">
            {{ Osinfo?.PlatformInfo.WebView2 }}</a-descriptions-item
          >
        </a-descriptions>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from "vue";
import { useAppStore } from "@/store";
import { System, Window } from "@wailsio/runtime";
import { GreetSvcWails } from "/#/Soraka/internal/wails/greet";
import { Message } from "@arco-design/web-vue";
const appStore = useAppStore();
const theme = computed(() => {
  return appStore.theme;
});
const formData = ref({
  theme: theme.value,
  concurrency: 1,
});
const windowsname = ref("");
const IsWindows = ref(false);
const Osinfo = ref();
onMounted(async () => {
  windowsname.value = await Window.Name();
  IsWindows.value = await System.IsWindows();
  Osinfo.value = await System.Environment();
  console.log(Osinfo.value);
  // formData.value.concurrency = await ClientService.GetConcurrency();
});
//切换主题
const handleTheme = async () => {
  var isdark = false;
  if (formData.value.theme == "dark") {
    isdark = true;
  } else if (formData.value.theme == "auto") {
    //自动
    const data = await System.IsDarkMode();
    if (data) {
      isdark = true;
    } else {
      isdark = false;
    }
  } else {
    isdark = false;
  }
  appStore.toggleTheme(isdark);
  if (isdark) GreetSvcWails.SetTheme();
};
// 将窗口设置为始终位于顶部。
const changeSetAlwaysOnTop = (value: any) => {
  Window.SetAlwaysOnTop(value);
  Message.success({
    content: value ? "窗口已位于顶部" : "窗口已取消置顶",
    id: "setting",
  });
};
// 卸载窗框和标题栏。
const changeSetFrameless = (value: any) => {
  Window.SetFrameless(value);
  Message.success({
    content: value ? "窗口已无边框" : "窗口已取消无边框",
    id: "setting",
  });
};
//  设置窗口是否可调整大小。
const changeSetResizable = (value: any) => {
  Window.SetResizable(value);
  Message.success({
    content: value ? "窗口可调整大小" : "窗口不可调整大小",
    id: "setting",
  });
};

const saveConcurrency = async (value: number) => {
  // await ClientService.SaveConcurrency(value);
};
</script>

<style scoped lang="less">
.container {
  padding: 8px;
  display: flex;
  .left-side {
    flex: 1;
  }
  .right-side {
    width: 280px;
    margin-left: 16px;
  }
  .card {
    padding: 8px;
    border-radius: 5px;
    background: var(--color-bg-1);
    box-shadow:
      0 10px 11px rgb(var(--arcoblue-3), 0.08),
      0 6px 4px rgb(var(--arcoblue-3), 0.06),
      0 0 0 1px rgb(var(--arcoblue-3), 0.05),
      0 2.89797px 2.12518px rgb(var(--arcoblue-3), 0.04),
      0 1.87823px 1.37737px rgb(var(--arcoblue-3), 0.03),
      0 1.18233px 0.867039px rgb(var(--arcoblue-3), 0.02),
      0 0.67932px 0.498168px rgb(var(--arcoblue-3), 0.02),
      0 0.298986px 0.219257px rgb(var(--arcoblue-3), 0.01);
  }
  :deep(.arco-form-item) {
    margin-bottom: 10px;
  }
  :deep(.arco-form-item-label-col) {
    margin-bottom: 0px;
  }
}
</style>
