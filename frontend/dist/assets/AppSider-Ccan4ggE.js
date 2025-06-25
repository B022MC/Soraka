import { d as defineComponent, r as ref, u as useUserStore, c as computed, w as watch, a as createBlock, b as useRoute, e as withCtx, f as useRouter, g as constantRouterMap, h as resolveComponent, o as openBlock, i as createVNode, j as createBaseVNode, n as normalizeClass, k as createElementBlock, l as renderList, F as Fragment, t as toDisplayString, m as createCommentVNode, p as unref, q as createTextVNode } from "./index-BT_C9izJ.js";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
const _imports_0 = "/assets/logo-C9STPBS-.png";
const _imports_1 = "data:image/svg+xml,%3csvg%20width='24'%20height='24'%20viewBox='0%200%2024%2024'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3crect%20x='1'%20y='1.00002'%20width='22'%20height='22'%20rx='4'%20fill='url(%23paint0_linear_1889_7314)'%20/%3e%3cpath%20fill-rule='evenodd'%20clip-rule='evenodd'%20d='M8.45818%209.80289C7.45621%209.80289%206.64115%208.98701%206.64115%207.98419C6.64115%206.98137%207.45621%206.1653%208.45818%206.1653C9.45981%206.1653%2010.2747%206.98137%2010.2747%207.98419C10.2747%208.98701%209.45981%209.80289%208.45818%209.80289ZM8.45801%204.52266C6.55134%204.52266%205%206.07559%205%207.98402C5%209.89281%206.55134%2011.4456%208.45801%2011.4456C10.3645%2011.4456%2011.9157%209.89281%2011.9157%207.98402C11.9157%206.07559%2010.3645%204.52266%208.45801%204.52266ZM16.0061%208.19942H14.1265V6.12403H16.0061C16.6547%206.12403%2016.8523%206.73514%2016.8523%207.16182C16.8523%207.60133%2016.6547%208.19942%2016.0061%208.19942ZM18.4856%207.16182C18.4856%205.67347%2017.4701%204.5242%2016.0137%204.5242H12.5749V11.4398H14.1263L14.1265%209.79945H16.0137C17.4177%209.79945%2018.4856%208.65767%2018.4856%207.16182ZM8.59486%2015.3233H11.6751C11.6826%2015.4094%2011.6922%2015.5464%2011.699%2015.7383C11.7102%2016.0605%2011.6931%2016.2576%2011.6282%2016.6464C11.3665%2018.2193%2010.1501%2019.1965%208.4548%2019.1965C6.54992%2019.1965%205%2017.6453%205%2015.7383C5%2013.8315%206.54992%2012.2798%208.4548%2012.2798C9.47602%2012.2798%2010.4395%2012.7293%2011.0983%2013.5129L11.1956%2013.6289L11.0698%2013.7129L9.95233%2014.4585L9.85913%2014.5209L9.78055%2014.4405C9.42773%2014.0789%208.95677%2013.9111%208.4548%2013.9111C7.43126%2013.9111%206.59839%2014.7137%206.59839%2015.7383C6.59839%2016.7627%207.43126%2017.5963%208.4548%2017.5963C9.31992%2017.5963%209.89388%2017.1072%2010.0065%2016.5614H8.59486V15.3233ZM18.9305%2015.3241H15.8501V16.5622H17.2617C17.1493%2017.108%2016.5751%2017.5973%2015.7098%2017.5973C14.6865%2017.5973%2013.8536%2016.7636%2013.8536%2015.7392C13.8536%2014.7144%2014.6865%2013.9117%2015.7098%2013.9117C16.2118%2013.9117%2016.6828%2014.0798%2017.036%2014.4413L17.1145%2014.5218L17.2077%2014.4593L18.3252%2013.7137L18.451%2013.6296L18.3535%2013.5138C17.6947%2012.7302%2016.7312%2012.2807%2015.7098%2012.2807C13.805%2012.2807%2012.2554%2013.8322%2012.2554%2015.7392C12.2554%2017.6462%2013.805%2019.1973%2015.7098%2019.1973C17.4055%2019.1973%2018.6217%2018.2202%2018.8836%2016.6473C18.9485%2016.2586%2018.9654%2016.0612%2018.9542%2015.7392C18.9474%2015.5473%2018.9378%2015.4103%2018.9305%2015.3241ZM17.6367%2010.9312C17.6367%2010.5443%2017.9512%2010.2295%2018.3377%2010.2295C18.7244%2010.2295%2019.0389%2010.5443%2019.0389%2010.9312C19.0389%2011.3181%2018.7244%2011.6328%2018.3377%2011.6328C17.9512%2011.6328%2017.6367%2011.3181%2017.6367%2010.9312Z'%20fill='white'%20/%3e%3cdefs%3e%3clinearGradient%20id='paint0_linear_1889_7314'%20x1='1'%20y1='23'%20x2='23'%20y2='1'%20gradientUnits='userSpaceOnUse'%3e%3cstop%20stop-color='%236636CF'%20/%3e%3cstop%20offset='1'%20stop-color='%23636FF9'%20/%3e%3c/linearGradient%3e%3c/defs%3e%3c/svg%3e";
const _hoisted_1 = { class: "sider-menu" };
const _hoisted_2 = { class: "sider-header" };
const _hoisted_3 = { class: "menu-item-inner" };
const _hoisted_4 = { class: "menu-item-inner" };
const _hoisted_5 = { class: "menu-item-inner" };
const _hoisted_6 = { class: "menu-item-inner" };
const _hoisted_7 = { class: "menu-item-inner" };
const _hoisted_8 = { class: "footer" };
const _hoisted_9 = ["src"];
const _hoisted_10 = { class: "footer-name" };
const _hoisted_11 = {
  key: 0,
  class: "tag"
};
const _hoisted_12 = { class: "footer-rank" };
const _hoisted_13 = { key: 0 };
const _sfc_main = /* @__PURE__ */ defineComponent({
  __name: "AppSider",
  setup(__props) {
    const router = useRouter();
    const route = useRoute();
    const current = ref(route.name);
    const collapsed = ref(false);
    const userStore = useUserStore();
    const menulist = computed(
      () => {
        var _a;
        return ((_a = constantRouterMap.find((item) => item.path === "/")) == null ? void 0 : _a.children) || [];
      }
    );
    ref(null);
    const defaultAvatar = new URL("/assets/logo-C9STPBS-.png", import.meta.url).href;
    watch(
      () => route.name,
      (name) => current.value = name
    );
    const menuHandle = (key) => {
      current.value = key;
      const target = menulist.value.find((i) => i.name === key);
      if (target) {
        router.push({ name: target.name });
      }
    };
    const handleSetting = () => {
      current.value = "setting";
      router.push({ name: "setting" });
    };
    const handleTool = (type) => {
      console.log("工具点击", type);
    };
    return (_ctx, _cache) => {
      const _component_a_avatar = resolveComponent("a-avatar");
      const _component_icon_font = resolveComponent("icon-font");
      const _component_a_menu_item = resolveComponent("a-menu-item");
      const _component_a_menu = resolveComponent("a-menu");
      const _component_a_layout_sider = resolveComponent("a-layout-sider");
      const _component_router_view = resolveComponent("router-view");
      const _component_a_layout_content = resolveComponent("a-layout-content");
      const _component_a_layout = resolveComponent("a-layout");
      return openBlock(), createBlock(_component_a_layout, { id: "app-layout-sider" }, {
        default: withCtx(() => [
          createVNode(_component_a_layout_sider, {
            collapsed: collapsed.value,
            "onUpdate:collapsed": _cache[3] || (_cache[3] = ($event) => collapsed.value = $event),
            width: 220,
            "collapsed-width": 64,
            collapsible: "",
            theme: "light",
            class: normalizeClass(["layout-sider", { "collapsed-sider": collapsed.value }])
          }, {
            default: withCtx(() => [
              createBaseVNode("div", _hoisted_1, [
                createBaseVNode("div", _hoisted_2, [
                  createVNode(_component_a_avatar, {
                    size: 32,
                    class: "logo-avatar"
                  }, {
                    default: withCtx(() => _cache[4] || (_cache[4] = [
                      createBaseVNode("img", { src: _imports_0 }, null, -1)
                    ])),
                    _: 1
                  }),
                  createBaseVNode("span", {
                    class: normalizeClass(["logo-title", { hidden: collapsed.value }])
                  }, "Soraka", 2)
                ]),
                createVNode(_component_a_menu, {
                  class: "menu-item",
                  theme: "light",
                  mode: "vertical",
                  "selected-keys": [current.value],
                  onMenuItemClick: menuHandle
                }, {
                  default: withCtx(() => [
                    (openBlock(true), createElementBlock(Fragment, null, renderList(menulist.value, (menuInfo) => {
                      var _a;
                      return openBlock(), createElementBlock(Fragment, {
                        key: menuInfo.name
                      }, [
                        !((_a = menuInfo.meta) == null ? void 0 : _a.hideInMenu) ? (openBlock(), createBlock(_component_a_menu_item, {
                          key: menuInfo.name
                        }, {
                          default: withCtx(() => [
                            createBaseVNode("div", _hoisted_3, [
                              createVNode(_component_icon_font, {
                                type: menuInfo.meta.icon,
                                class: "menu-icon"
                              }, null, 8, ["type"]),
                              createBaseVNode("span", {
                                class: normalizeClass(["menu-text", { hidden: collapsed.value }])
                              }, toDisplayString(menuInfo.meta.title), 3)
                            ])
                          ]),
                          _: 2
                        }, 1024)) : createCommentVNode("", true)
                      ], 64);
                    }), 128))
                  ]),
                  _: 1
                }, 8, ["selected-keys"]),
                createVNode(_component_a_menu, {
                  class: "footer-tools",
                  theme: "light",
                  mode: "vertical",
                  "selected-keys": current.value === "setting" ? ["setting"] : []
                }, {
                  default: withCtx(() => [
                    createVNode(_component_a_menu_item, {
                      key: "opgg",
                      onClick: _cache[0] || (_cache[0] = ($event) => handleTool("opgg"))
                    }, {
                      default: withCtx(() => [
                        createBaseVNode("div", _hoisted_4, [
                          _cache[5] || (_cache[5] = createBaseVNode("img", {
                            src: _imports_1,
                            class: "menu-icon"
                          }, null, -1)),
                          createBaseVNode("span", {
                            class: normalizeClass(["menu-text", { hidden: collapsed.value }])
                          }, "OP.GG", 2)
                        ])
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_menu_item, {
                      key: "fix",
                      onClick: _cache[1] || (_cache[1] = ($event) => handleTool("fix"))
                    }, {
                      default: withCtx(() => [
                        createBaseVNode("div", _hoisted_5, [
                          createVNode(_component_icon_font, {
                            type: "icon-ArrowCircle",
                            class: "menu-icon"
                          }),
                          createBaseVNode("span", {
                            class: normalizeClass(["menu-text", { hidden: collapsed.value }])
                          }, "修复无限加载", 2)
                        ])
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_menu_item, {
                      key: "notice",
                      onClick: _cache[2] || (_cache[2] = ($event) => handleTool("notice"))
                    }, {
                      default: withCtx(() => [
                        createBaseVNode("div", _hoisted_6, [
                          createVNode(_component_icon_font, {
                            type: "icon-Alert",
                            class: "menu-icon"
                          }),
                          createBaseVNode("span", {
                            class: normalizeClass(["menu-text", { hidden: collapsed.value }])
                          }, "公告", 2)
                        ])
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_menu_item, {
                      key: "setting",
                      onClick: handleSetting
                    }, {
                      default: withCtx(() => [
                        createBaseVNode("div", _hoisted_7, [
                          createVNode(_component_icon_font, {
                            type: "icon-Setting",
                            class: "menu-icon"
                          }),
                          createBaseVNode("span", {
                            class: normalizeClass(["menu-text", { hidden: collapsed.value }])
                          }, "设置", 2)
                        ])
                      ]),
                      _: 1
                    })
                  ]),
                  _: 1
                }, 8, ["selected-keys"]),
                createBaseVNode("div", _hoisted_8, [
                  createBaseVNode("img", {
                    src: unref(userStore).avatar || unref(defaultAvatar),
                    style: { "width": "32px", "height": "32px", "border-radius": "50%" }
                  }, null, 8, _hoisted_9),
                  createBaseVNode("div", {
                    class: normalizeClass(["footer-text", { hidden: collapsed.value }])
                  }, [
                    createBaseVNode("div", _hoisted_10, [
                      createTextVNode(toDisplayString(unref(userStore).nickname || "未登录") + " ", 1),
                      unref(userStore).tag ? (openBlock(), createElementBlock("span", _hoisted_11, toDisplayString(unref(userStore).tag), 1)) : createCommentVNode("", true)
                    ]),
                    createBaseVNode("div", _hoisted_12, [
                      unref(userStore).server ? (openBlock(), createElementBlock("span", _hoisted_13, toDisplayString(unref(userStore).server), 1)) : createCommentVNode("", true)
                    ])
                  ], 2)
                ])
              ])
            ]),
            _: 1
          }, 8, ["collapsed", "class"]),
          createVNode(_component_a_layout, null, {
            default: withCtx(() => [
              createVNode(_component_a_layout_content, { class: "layout-content" }, {
                default: withCtx(() => [
                  createVNode(_component_router_view)
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
const AppSider = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-216e1df2"]]);
export {
  AppSider as default
};
