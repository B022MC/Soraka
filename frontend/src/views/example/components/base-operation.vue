<template>
  <a-card
    class="general-card"
    title="基础"
    :header-style="{ paddingBottom: '0' }"
    :body-style="{ padding: '18px 20px 0 20px' }"
  >
    <template #extra>
      <!-- <a-link>管理</a-link> -->
    </template>
    <a-row :gutter="8">
      <a-col v-for="link in links" :key="link.text" @click="handleEvent(link.event)" :span="8" class="wrapper">
        <div class="icon">
          <icon-font :type="link.icon" class="icon" size="26"/>
        </div>
        <a-typography-paragraph class="text">
          {{ link.text }}
        </a-typography-paragraph>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import {Dialogs} from "@wailsio/runtime";
  import {MessageService,GreetService, ExampleApi} from "/#/Soraka/service";
  import { Message,Notification } from '@arco-design/web-vue';
  const links = [
    { text: '提示弹框', icon: 'icon-filled',event:"info"},
    { text: '警告提示', icon: 'icon-warning',event:"warning"},
    { text: '问题弹框', icon: 'icon-dialog',event:"question"},
    { text: '选择文件', icon: 'icon-wenjian',event:"file"},
    { text: '系统提示', icon: 'icon-lirunfenxikuangjia',event:"sys"},
    { text: '调用函数', icon: 'icon-go',event:"service"},
    { text: '示例API', icon: 'icon-go',event:"wapi"},
    // { text: 'workplace.contentPutIn', icon: 'icon-fire' },
  ];
  //执行事件
  const handleEvent=async(event:string)=>{
    if(event=="info"){
       Dialogs.Info({Title:"提示标题",Message:"这里Dialog提示框内容。"})
    }else if(event=="warning"){
       Dialogs.Warning({Title:"提示标题",Message:"这里Dialog提示框内容。"})
    }else if(event=="question"){
       Dialogs.Question({Title:"删除确认",Message:"您确定删除当前数据吗？\n敏感数据谨慎操作，删除后将无法恢复。"})
    }else if(event=="file"){
       const filedata=await Dialogs.OpenFile({Title:"选择Soraka附件"})
       console.log("选择附件结果：",filedata)
    }else if(event=="sys"){
      MessageService.Dialog()
    }else if(event=="service"){
      Message.loading({content:"调用中·请稍后",id:"greet",duration:0})
      GreetService.Greet("Soraka，调用service函数了").then((result) => {
        Notification.info({
          title: '调用函数结果',
          content: result,
          position: 'bottomRight',
          closable: true,
        })
        Message.success({content:"调用成功",id:"greet",duration:2000})
      }).catch((err) => {
          Message.error({content:err,id:"greet",duration:2000})
          console.log("错误：",err);
      });
    }else if(event=="wapi"){
      Message.loading({content:"调用中·请稍后",id:"wapi",duration:0})
      ExampleApi.Hello("前端调用").then((result) => {
        Notification.info({
          title: '调用函数结果',
          content: result,
          position: 'bottomRight',
          closable: true,
        })
        Message.success({content:"调用成功",id:"wapi",duration:2000})
      }).catch((err) => {
          Message.error({content:err,id:"wapi",duration:2000})
      });
    }
  }
</script>

<style scoped lang="less">
.general-card{
  background: transparent;
  .wrapper{
    cursor: pointer;
    margin-bottom: 10px;
    padding: 5px 0px;
    border-radius: 3px;
    &:hover{
      background: var(--color-neutral-3);
    }
    .icon{
      text-align: center;
    }
    .text{
      text-align: center;
      margin-bottom: 0px;
    }
  }
}
</style>
