<template>
  <a-card
    class="general-card"
    title="快捷操作"
    :header-style="{ paddingBottom: '0' }"
    :body-style="{ padding: '24px 20px 0 20px' }"
  >
    <!--    <template #extra>-->
    <!--      <a-link>管理</a-link>-->
    <!--    </template>-->
    <a-row :gutter="8">
      <template v-for="link in links">
        <a-col
          :span="8"
          v-if="link.type == 'button'"
          class="wrapper"
          @click="handleOpen(link.value)"
        >
          <div class="icon">
            <icon-font :type="link.icon" class="icon" size="26" />
          </div>
          <a-typography-paragraph class="text">
            {{ link.text }}
          </a-typography-paragraph>
        </a-col>
        <!--        <a-col :span="8" v-else-if="link.type=='browser'" class="wrapper">-->
        <!--            <a :wml-openURL="link.value">-->
        <!--              <div class="icon">-->
        <!--                <icon-font :type="link.icon" class="icon" size="26"/>-->
        <!--              </div>-->
        <!--              <a-typography-paragraph class="text">-->
        <!--                {{ link.text }}-->
        <!--              </a-typography-paragraph>-->
        <!--            </a>-->
        <!--        </a-col>-->
        <a-col :span="8" v-else-if="link.type == 'a'" class="wrapper">
          <a :href="link.value" target="_blank">
            <div class="icon">
              <icon-font :type="link.icon" class="icon" size="26" />
            </div>
            <a-typography-paragraph class="text">
              {{ link.text }}
            </a-typography-paragraph>
          </a>
        </a-col>
      </template>
    </a-row>
    <!-- <a-divider class="split-line" style="margin: 0" /> -->
  </a-card>
</template>

<script lang="ts" setup>
import { onMounted } from "vue";
import { Notification, Message } from "@arco-design/web-vue";
import { WML } from "@wailsio/runtime";
import { LcuApiWails } from "/#/Soraka/internal/wails/lcu";
onMounted(async () => {
  WML.Reload();
});
const links = [
  // { text: '通知提醒框', icon: 'icon-filled',type:"button",value:"notification" },
  // { text: '默认浏览器', icon: 'icon-wangye',type:"browser",value:"https://Sorakas.cn"},
  // { text: 'Webview', icon: 'icon-wangye',type:"a",value:"https://v3alpha.wails.io"},
  { text: "启动LOL", icon: "icon-Game", type: "button", value: "startclient" },
  // { text: 'workplace.onlinePromotion', icon: 'icon-mobile' },
  // { text: 'workplace.contentPutIn', icon: 'icon-fire' },
];
//执行UI组件
const handleOpen = (val: string) => {
  if (val == "notification") {
    Notification.info({
      title: "Notification",
      content: "This is a notification!",
      position: "bottomRight",
      closable: true,
    });
  } else if (val == "startclient") {
    Message.loading({
      content: "启动中，请稍后",
      id: "startclient",
      duration: 0,
    });
    LcuApiWails.StartClient()
      .then(() => {
        Message.success({
          content: "启动成功",
          id: "startclient",
          duration: 2000,
        });
      })
      .catch((e: any) => {
        Notification.error({
          title: "启动失败",
          content: String(e),
          position: "bottomRight",
        });
        Message.error({
          content: "启动失败",
          id: "startclient",
          duration: 2000,
        });
      });
  }
};
</script>

<style scoped lang="less">
.general-card {
  background: transparent;
  .wrapper {
    cursor: pointer;
    .icon {
      text-align: center;
    }
    .text {
      text-align: center;
    }
  }
}
</style>
