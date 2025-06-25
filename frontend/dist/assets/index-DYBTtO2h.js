import { d as defineComponent, u as useUserStore, r as ref, s as onMounted, O as On, h as resolveComponent, o as openBlock, a as createBlock, e as withCtx, i as createVNode, j as createBaseVNode, p as unref, t as toDisplayString, q as createTextVNode, k as createElementBlock, m as createCommentVNode, R as Reload, l as renderList, F as Fragment, N as Notification, M as Message } from "./index-BT_C9izJ.js";
import { G as GetClientPath, S as StartClient } from "./lcuapiwails-IFZpWFl2.js";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
function goodTimeText() {
  const time = /* @__PURE__ */ new Date();
  const hour = time.getHours();
  return hour < 9 ? "早上好" : hour <= 11 ? "上午好" : hour <= 13 ? "中午好" : hour <= 18 ? "下午好" : "晚上好";
}
const _hoisted_1$2 = ["src"];
const _hoisted_2$2 = { class: "welcome" };
const _hoisted_3$2 = { class: "hello" };
const _hoisted_4$1 = { class: "datetime" };
const _hoisted_5$1 = { class: "status" };
const _hoisted_6 = {
  key: 1,
  class: "waiting"
};
const _hoisted_7 = { key: 0 };
const _sfc_main$2 = /* @__PURE__ */ defineComponent({
  __name: "Welcome",
  setup(__props) {
    const userStore = useUserStore();
    const sysTime = ref("");
    const clientPath = ref("");
    const lcuOnline = ref(false);
    const lcuPort = ref("");
    const lcuToken = ref("");
    const defaultAvatar = new URL("/assets/logo-C9STPBS-.png", import.meta.url).href;
    onMounted(() => {
      On("time", (time) => {
        sysTime.value = time.data;
      });
      GetClientPath().then((p) => {
        clientPath.value = p;
      }).catch(() => {
        clientPath.value = "";
      });
      On("lcuStatus", (d) => {
        const payload = Array.isArray(d.data) ? d.data[0] : d.data;
        console.log("[Event] lcuStatus 收到:", payload);
        if (typeof (payload == null ? void 0 : payload.status) === "boolean") {
          lcuOnline.value = payload.status;
        }
      });
      On("lcuCreds", (d) => {
        const payload = Array.isArray(d.data) ? d.data[0] : d.data;
        console.log("[Event] lcuCreds 收到:", payload);
        if (typeof (payload == null ? void 0 : payload.port) === "number" && typeof (payload == null ? void 0 : payload.token) === "string") {
          lcuPort.value = payload.port.toString();
          lcuToken.value = payload.token;
        }
      });
      On("summonerInfo", (d) => {
        console.log("[Event] summonerInfo 原始:", d);
        const info = Array.isArray(d.data) ? d.data[0] : d.data;
        console.log("[Event] 解析后的 info:", info, "类型是", typeof info);
        if (info && typeof info === "object") {
          console.log("[Event] summonerInfo 收到:", info);
          userStore.setInfo({
            accountId: info.accountId,
            summonerId: info.summonerId,
            puuid: info.puuid,
            nickname: info.gameName || info.displayName,
            avatar: info.avatarUrl,
            region: info.region || "",
            // 如果后端没传，补空字符串
            tag: info.tagLine,
            rank: info.rank || "",
            // 可选数据
            winRate: info.winRate || 0,
            // 可选数据
            wins: info.wins || 0,
            losses: info.losses || 0,
            totalGames: info.totalGames || 0,
            createtime: info.createtime || "",
            level: info.summonerLevel,
            xpSinceLastLevel: info.xpSinceLastLevel,
            xpUntilNextLevel: info.xpUntilNextLevel,
            percentCompleteForNextLevel: info.percentCompleteForNextLevel,
            privacy: info.privacy
          });
        } else {
          console.warn("[Event] summonerInfo 无效数据:", info);
        }
        console.log("userStore.avatar:", userStore.avatar);
      });
    });
    return (_ctx, _cache) => {
      const _component_a_avatar = resolveComponent("a-avatar");
      const _component_a_tag = resolveComponent("a-tag");
      const _component_a_space = resolveComponent("a-space");
      const _component_a_row = resolveComponent("a-row");
      const _component_a_card = resolveComponent("a-card");
      return openBlock(), createBlock(_component_a_card, {
        class: "card",
        bordered: false
      }, {
        default: withCtx(() => [
          createVNode(_component_a_row, {
            wrap: "",
            gutter: [{ xs: 0, sm: 14, md: 14, lg: 14, xl: 14, xxl: 14 }, 16],
            class: "content"
          }, {
            default: withCtx(() => [
              createVNode(_component_a_space, { size: "medium" }, {
                default: withCtx(() => [
                  createVNode(_component_a_avatar, {
                    style: { "padding": "2px" },
                    size: 68
                  }, {
                    default: withCtx(() => [
                      createBaseVNode("img", {
                        src: unref(userStore).avatar || unref(defaultAvatar)
                      }, null, 8, _hoisted_1$2)
                    ]),
                    _: 1
                  }),
                  createBaseVNode("div", _hoisted_2$2, [
                    createBaseVNode("p", _hoisted_3$2, toDisplayString(unref(goodTimeText)()) + "！" + toDisplayString(unref(userStore).nickname), 1),
                    createBaseVNode("p", null, [
                      _cache[0] || (_cache[0] = createTextVNode(" 当前时间：")),
                      createBaseVNode("span", _hoisted_4$1, toDisplayString(sysTime.value ? sysTime.value[0] : "---"), 1)
                    ]),
                    createBaseVNode("p", null, [
                      createTextVNode(" 客户端路径：" + toDisplayString(clientPath.value || "未找到") + " ", 1),
                      createBaseVNode("span", _hoisted_5$1, [
                        lcuOnline.value ? (openBlock(), createBlock(_component_a_tag, {
                          key: 0,
                          color: "green"
                        }, {
                          default: withCtx(() => _cache[1] || (_cache[1] = [
                            createTextVNode("已连接LCU")
                          ])),
                          _: 1
                        })) : (openBlock(), createElementBlock("span", _hoisted_6, "等待连接"))
                      ])
                    ]),
                    lcuPort.value ? (openBlock(), createElementBlock("p", _hoisted_7, "端口：" + toDisplayString(lcuPort.value) + " Token：" + toDisplayString(lcuToken.value), 1)) : createCommentVNode("", true)
                  ])
                ]),
                _: 1
              })
            ]),
            _: 1
          })
        ]),
        _: 1
      });
    };
  }
});
const Welcome = /* @__PURE__ */ _export_sfc(_sfc_main$2, [["__scopeId", "data-v-148647c7"]]);
const _hoisted_1$1 = { class: "icon" };
const _hoisted_2$1 = ["href"];
const _hoisted_3$1 = { class: "icon" };
const _sfc_main$1 = /* @__PURE__ */ defineComponent({
  __name: "quick-operation",
  setup(__props) {
    onMounted(async () => {
      Reload();
    });
    const links = [
      // { text: '通知提醒框', icon: 'icon-filled',type:"button",value:"notification" },
      // { text: '默认浏览器', icon: 'icon-wangye',type:"browser",value:"https://Sorakas.cn"},
      // { text: 'Webview', icon: 'icon-wangye',type:"a",value:"https://v3alpha.wails.io"},
      { text: "启动LOL", icon: "icon-Game", type: "button", value: "startclient" }
      // { text: 'workplace.onlinePromotion', icon: 'icon-mobile' },
      // { text: 'workplace.contentPutIn', icon: 'icon-fire' },
    ];
    const handleOpen = (val) => {
      if (val == "notification") {
        Notification.info({
          title: "Notification",
          content: "This is a notification!",
          position: "bottomRight",
          closable: true
        });
      } else if (val == "startclient") {
        Message.loading({
          content: "启动中，请稍后",
          id: "startclient",
          duration: 0
        });
        StartClient().then(() => {
          Message.success({
            content: "启动成功",
            id: "startclient",
            duration: 2e3
          });
        }).catch((e) => {
          Notification.error({
            title: "启动失败",
            content: String(e),
            position: "bottomRight"
          });
          Message.error({
            content: "启动失败",
            id: "startclient",
            duration: 2e3
          });
        });
      }
    };
    return (_ctx, _cache) => {
      const _component_icon_font = resolveComponent("icon-font");
      const _component_a_typography_paragraph = resolveComponent("a-typography-paragraph");
      const _component_a_col = resolveComponent("a-col");
      const _component_a_row = resolveComponent("a-row");
      const _component_a_card = resolveComponent("a-card");
      return openBlock(), createBlock(_component_a_card, {
        class: "general-card",
        title: "快捷操作",
        "header-style": { paddingBottom: "0" },
        "body-style": { padding: "24px 20px 0 20px" }
      }, {
        default: withCtx(() => [
          createVNode(_component_a_row, { gutter: 8 }, {
            default: withCtx(() => [
              (openBlock(), createElementBlock(Fragment, null, renderList(links, (link) => {
                return openBlock(), createElementBlock(Fragment, null, [
                  link.type == "button" ? (openBlock(), createBlock(_component_a_col, {
                    key: 0,
                    span: 8,
                    class: "wrapper",
                    onClick: ($event) => handleOpen(link.value)
                  }, {
                    default: withCtx(() => [
                      createBaseVNode("div", _hoisted_1$1, [
                        createVNode(_component_icon_font, {
                          type: link.icon,
                          class: "icon",
                          size: "26"
                        }, null, 8, ["type"])
                      ]),
                      createVNode(_component_a_typography_paragraph, { class: "text" }, {
                        default: withCtx(() => [
                          createTextVNode(toDisplayString(link.text), 1)
                        ]),
                        _: 2
                      }, 1024)
                    ]),
                    _: 2
                  }, 1032, ["onClick"])) : link.type == "a" ? (openBlock(), createBlock(_component_a_col, {
                    key: 1,
                    span: 8,
                    class: "wrapper"
                  }, {
                    default: withCtx(() => [
                      createBaseVNode("a", {
                        href: link.value,
                        target: "_blank"
                      }, [
                        createBaseVNode("div", _hoisted_3$1, [
                          createVNode(_component_icon_font, {
                            type: link.icon,
                            class: "icon",
                            size: "26"
                          }, null, 8, ["type"])
                        ]),
                        createVNode(_component_a_typography_paragraph, { class: "text" }, {
                          default: withCtx(() => [
                            createTextVNode(toDisplayString(link.text), 1)
                          ]),
                          _: 2
                        }, 1024)
                      ], 8, _hoisted_2$1)
                    ]),
                    _: 2
                  }, 1024)) : createCommentVNode("", true)
                ], 64);
              }), 64))
            ]),
            _: 1
          })
        ]),
        _: 1
      });
    };
  }
});
const QuickOperation = /* @__PURE__ */ _export_sfc(_sfc_main$1, [["__scopeId", "data-v-5d28ec59"]]);
const _hoisted_1 = { class: "container" };
const _hoisted_2 = { class: "left-side" };
const _hoisted_3 = { class: "panel" };
const _hoisted_4 = { class: "right-side" };
const _hoisted_5 = { class: "panel moduler-wrap" };
const _sfc_main = {
  __name: "index",
  setup(__props) {
    return (_ctx, _cache) => {
      const _component_a_grid_item = resolveComponent("a-grid-item");
      const _component_a_grid = resolveComponent("a-grid");
      return openBlock(), createElementBlock("div", _hoisted_1, [
        createBaseVNode("div", _hoisted_2, [
          createBaseVNode("div", _hoisted_3, [
            createVNode(Welcome)
          ])
        ]),
        createBaseVNode("div", _hoisted_4, [
          createVNode(_component_a_grid, {
            cols: 24,
            "row-gap": 16
          }, {
            default: withCtx(() => [
              createVNode(_component_a_grid_item, { span: 24 }, {
                default: withCtx(() => [
                  createBaseVNode("div", _hoisted_5, [
                    createVNode(QuickOperation)
                  ])
                ]),
                _: 1
              })
            ]),
            _: 1
          })
        ])
      ]);
    };
  }
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-6283aa77"]]);
export {
  index as default
};
