<template>
  <div class="life-page">
    <h1 class="section-title">其它功能</h1>

    <!-- 英雄选择 -->
    <div class="section">
      <a-typography-title :heading="5">英雄选择</a-typography-title>
      <a-collapse :default-active-key="[]" expand-icon-position="right" class="collapse-wrapper">
        <!-- 自动接受对局 -->
        <a-collapse-item header="自动接受对局" class="collapse-item">
          <div class="item-desc">在你设置的倒数之后自动接受对局匹配</div>
          <div class="collapse-body">
            <a-form layout="vertical">
              <a-form-item label="在对局找到后延迟接受对局的秒数：">
                <a-input-number v-model="form.autoAcceptDelay" :min="0" :max="20" />
              </a-form-item>
              <a-form-item label="是否启用">
                <a-switch v-model="form.autoAcceptEnabled" />
              </a-form-item>
            </a-form>
          </div>
        </a-collapse-item>

        <!-- 自动接受交换请求 -->
        <a-collapse-item header="自动接受交换请求" class="collapse-item">
          <div class="item-desc">自动接受队友发起的英雄交换请求</div>
          <div class="collapse-body">
            <a-form layout="vertical">
              <a-form-item label="自动接受模糊层交换请求：">
                <a-switch v-model="form.autoAcceptBlindExchange" />
              </a-form-item>
              <a-form-item label="自动接受明确交换请求：">
                <a-switch v-model="form.autoAcceptExplicitExchange" />
              </a-form-item>
            </a-form>
          </div>
        </a-collapse-item>

        <!-- 自动亮起英雄 -->
        <a-collapse-item header="自动亮起英雄" class="collapse-item">
          <div class="item-desc">在你的选定阶段时自动亮起预设英雄</div>
          <div class="collapse-body">
            <a-form layout="vertical">
              <a-form-item label="默认设置">
                <a-space align="center">
                  <a-input v-model="form.defaultChampion" placeholder="默认亮起英雄" disabled />
                  <a-button type="outline">选择</a-button>
                </a-space>
              </a-form-item>

              <a-form-item label="按位置设置">
                <a-row :gutter="12">
                  <a-col :span="12">
                    <a-form-item label="上路：">
                      <a-space>
                        <a-input v-model="form.top" disabled />
                        <a-button type="outline">选择</a-button>
                      </a-space>
                    </a-form-item>
                  </a-col>
                  <a-col :span="12">
                    <a-form-item label="打野：">
                      <a-space>
                        <a-input v-model="form.jungle" disabled />
                        <a-button type="outline">选择</a-button>
                      </a-space>
                    </a-form-item>
                  </a-col>
                </a-row>

                <a-row :gutter="12">
                  <a-col :span="12">
                    <a-form-item label="中路：">
                      <a-space>
                        <a-input v-model="form.mid" disabled />
                        <a-button type="outline">选择</a-button>
                      </a-space>
                    </a-form-item>
                  </a-col>
                  <a-col :span="12">
                    <a-form-item label="下路：">
                      <a-space>
                        <a-input v-model="form.bottom" disabled />
                        <a-button type="outline">选择</a-button>
                      </a-space>
                    </a-form-item>
                  </a-col>
                </a-row>

                <a-form-item label="辅助：">
                  <a-space>
                    <a-input v-model="form.support" disabled />
                    <a-button type="outline">选择</a-button>
                  </a-space>
                </a-form-item>
              </a-form-item>

              <a-form-item label="启用自动亮起">
                <a-switch v-model="form.enableAutoPick" />
              </a-form-item>
            </a-form>
          </div>
        </a-collapse-item>

        <!-- 自动禁用英雄 -->
        <a-collapse-item header="自动禁用英雄" class="collapse-item">
          <div class="item-desc">在你的禁用环节时自动禁用预设英雄</div>
          <div class="collapse-body">
            <a-space align="center" justify="space-between">
              <a-tag color="gray">未启用</a-tag>
              <a-button type="outline">恢复默认</a-button>
            </a-space>
          </div>
        </a-collapse-item>
      </a-collapse>
    </div>

    <!-- 游戏模块 -->
    <div class="section">
      <a-typography-title :heading="5">游戏</a-typography-title>
      <a-collapse :default-active-key="[]" expand-icon-position="right" class="collapse-wrapper">

        <!-- 自动重连 -->
        <a-collapse-item header="自动重连" class="collapse-item">
          <div class="item-desc">当你掉线退出游戏时自动重新连接</div>
          <div class="collapse-body">
            <a-switch v-model="form.autoReconnect" />
          </div>
        </a-collapse-item>

        <!-- 创建练习模式 -->
        <a-collapse-item header="创建 5v5 练习模式" class="collapse-item">
          <div class="item-desc">只能添加人机玩家</div>
          <div class="collapse-body">
            <a-form layout="vertical">
              <a-form-item label="房间名：（不可为空）">
                <a-input v-model="form.practiceRoomName" placeholder="请输入房间名" />
              </a-form-item>
              <a-form-item label="房间密码：（若留空则不设密码）">
                <a-input v-model="form.practiceRoomPassword" placeholder="请输入房间密码" />
              </a-form-item>
              <a-button type="primary">创建</a-button>
            </a-form>
          </div>
        </a-collapse-item>

        <!-- 观战 -->
        <a-collapse-item header="观战" class="collapse-item">
          <div class="item-desc">观战同大区玩家正在进行的游戏</div>
          <div class="collapse-body">
            <a-form layout="vertical">
              <a-form-item label="你想观战的同大区召唤师名及编号：">
                <a-input v-model="form.spectateSummoner" placeholder="请输入召唤师名" />
              </a-form-item>
              <a-form-item label="观战启动方式：">
                <a-select v-model="form.spectateMode">
                  <a-option value="lcu">LCU API</a-option>
                  <a-option value="riot">Riot Client</a-option>
                </a-select>
              </a-form-item>
              <a-button type="primary">观战</a-button>
            </a-form>
          </div>
        </a-collapse-item>

        <!-- 锁定设置 -->
        <a-collapse-item header="锁定游戏设置" class="collapse-item">
          <div class="item-desc">让你的游戏设置不会因更新或异常写回默认</div>
          <div class="collapse-body">
            <a-switch v-model="form.lockGameSettings" />
          </div>
        </a-collapse-item>
      </a-collapse>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

const form = reactive({
  autoAcceptDelay: 0,
  autoAcceptEnabled: false,
  autoAcceptBlindExchange: false,
  autoAcceptExplicitExchange: false,
  defaultChampion: '',
  top: '',
  jungle: '',
  mid: '',
  bottom: '',
  support: '',
  enableAutoPick: false,

  autoReconnect: false,
  practiceRoomName: '',
  practiceRoomPassword: '',
  spectateSummoner: '',
  spectateMode: 'lcu',
  lockGameSettings: false
})
</script>

<style scoped lang="less">
.life-page {
  padding: 24px;
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4f8, #ffffff);
  color: var(--color-text);
}

.section {
  margin-bottom: 40px;
}

.section-title {
  margin-bottom: 16px;
  font-weight: bold;
  font-size: 20px;
}

.item-desc {
  font-size: 13px;
  color: var(--color-text-2);
  margin-bottom: 12px;
}

.collapse-wrapper {
  border: none;
  background: transparent;
}

.collapse-item {
  background: rgba(255, 255, 255, 0.4);
  border: 1px solid rgba(180, 180, 180, 0.35);
  border-radius: 12px;
  backdrop-filter: blur(12px);
  margin-bottom: 12px;
  overflow: hidden;

  .arco-collapse-item-header {
    padding: 16px;
    font-weight: 500;
    font-size: 14px;
  }

  .collapse-body {
    padding: 16px;
  }
}
</style>