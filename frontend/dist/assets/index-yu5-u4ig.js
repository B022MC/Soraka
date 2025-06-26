import { B as ByID, d as defineComponent, v as useAppStore, c as computed, r as ref, o as onMounted, y as thisWindow, I as IsWindows, E as Environment, l as createElementBlock, k as createBaseVNode, j as createVNode, e as withCtx, h as resolveComponent, i as openBlock, s as createTextVNode, q as unref, t as toDisplayString, z as IsDarkMode, M as Message } from "./index-BECB_qz3.js";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
function SetTheme() {
  let $resultPromise = ByID(3063170729);
  return $resultPromise;
}
const _hoisted_1 = { class: "container" };
const _hoisted_2 = { class: "left-side" };
const _hoisted_3 = { class: "card" };
const _hoisted_4 = { class: "right-side" };
const _hoisted_5 = { class: "card" };
const _sfc_main = /* @__PURE__ */ defineComponent({
  __name: "index",
  setup(__props) {
    const appStore = useAppStore();
    const theme = computed(() => {
      return appStore.theme;
    });
    const formData = ref({
      theme: theme.value,
      concurrency: 1
    });
    const windowsname = ref("");
    const IsWindows$1 = ref(false);
    const Osinfo = ref();
    onMounted(async () => {
      windowsname.value = await thisWindow.Name();
      IsWindows$1.value = await IsWindows();
      Osinfo.value = await Environment();
      console.log(Osinfo.value);
    });
    const handleTheme = async () => {
      var isdark = false;
      if (formData.value.theme == "dark") {
        isdark = true;
      } else if (formData.value.theme == "auto") {
        const data = await IsDarkMode();
        if (data) {
          isdark = true;
        } else {
          isdark = false;
        }
      } else {
        isdark = false;
      }
      appStore.toggleTheme(isdark);
      if (isdark) SetTheme();
    };
    const changeSetAlwaysOnTop = (value) => {
      thisWindow.SetAlwaysOnTop(value);
      Message.success({
        content: value ? "窗口已位于顶部" : "窗口已取消置顶",
        id: "setting"
      });
    };
    const changeSetResizable = (value) => {
      thisWindow.SetResizable(value);
      Message.success({
        content: value ? "窗口可调整大小" : "窗口不可调整大小",
        id: "setting"
      });
    };
    return (_ctx, _cache) => {
      const _component_a_radio = resolveComponent("a-radio");
      const _component_a_radio_group = resolveComponent("a-radio-group");
      const _component_a_form_item = resolveComponent("a-form-item");
      const _component_a_col = resolveComponent("a-col");
      const _component_a_switch = resolveComponent("a-switch");
      const _component_icon_translate = resolveComponent("icon-translate");
      const _component_a_button = resolveComponent("a-button");
      const _component_icon_launch = resolveComponent("icon-launch");
      const _component_a_row = resolveComponent("a-row");
      const _component_a_form = resolveComponent("a-form");
      const _component_a_descriptions_item = resolveComponent("a-descriptions-item");
      const _component_a_descriptions = resolveComponent("a-descriptions");
      return openBlock(), createElementBlock("div", _hoisted_1, [
        createBaseVNode("div", _hoisted_2, [
          createBaseVNode("div", _hoisted_3, [
            createVNode(_component_a_form, {
              model: formData.value,
              layout: "vertical",
              size: "large"
            }, {
              default: withCtx(() => [
                createVNode(_component_a_row, { gutter: 16 }, {
                  default: withCtx(() => [
                    createVNode(_component_a_col, { span: 16 }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, {
                          field: "theme",
                          label: "主题"
                        }, {
                          default: withCtx(() => [
                            createVNode(_component_a_radio_group, {
                              modelValue: formData.value.theme,
                              "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => formData.value.theme = $event),
                              type: "button",
                              onChange: handleTheme
                            }, {
                              default: withCtx(() => [
                                createVNode(_component_a_radio, { value: "light" }, {
                                  default: withCtx(() => _cache[3] || (_cache[3] = [
                                    createTextVNode("亮色")
                                  ])),
                                  _: 1
                                }),
                                createVNode(_component_a_radio, { value: "dark" }, {
                                  default: withCtx(() => _cache[4] || (_cache[4] = [
                                    createTextVNode("黑色")
                                  ])),
                                  _: 1
                                }),
                                createVNode(_component_a_radio, { value: "auto" }, {
                                  default: withCtx(() => _cache[5] || (_cache[5] = [
                                    createTextVNode("跟随系统")
                                  ])),
                                  _: 1
                                })
                              ]),
                              _: 1
                            }, 8, ["modelValue"])
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_col, { span: 8 }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, {
                          field: "concurrency",
                          label: "LCU并发数"
                        })
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_col, { span: 8 }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, {
                          field: "SetAlwaysOnTop",
                          label: "置顶窗口"
                        }, {
                          default: withCtx(() => [
                            createVNode(_component_a_switch, {
                              type: "round",
                              onChange: changeSetAlwaysOnTop
                            }, {
                              checked: withCtx(() => _cache[6] || (_cache[6] = [
                                createTextVNode(" 开启 ")
                              ])),
                              unchecked: withCtx(() => _cache[7] || (_cache[7] = [
                                createTextVNode(" 关闭 ")
                              ])),
                              _: 1
                            })
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_col, { span: 12 }),
                    createVNode(_component_a_col, { span: 12 }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, {
                          field: "SetFullscreenButtonEnabled",
                          label: "窗口是否可调整大小"
                        }, {
                          default: withCtx(() => [
                            createVNode(_component_a_switch, {
                              type: "round",
                              "default-checked": true,
                              onChange: changeSetResizable
                            }, {
                              checked: withCtx(() => _cache[8] || (_cache[8] = [
                                createTextVNode(" 可以 ")
                              ])),
                              unchecked: withCtx(() => _cache[9] || (_cache[9] = [
                                createTextVNode(" 不可 ")
                              ])),
                              _: 1
                            })
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_col, { span: 12 }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, {
                          field: "Reload",
                          label: "重置窗口"
                        }, {
                          default: withCtx(() => [
                            createVNode(_component_a_button, {
                              type: "primary",
                              onClick: _cache[1] || (_cache[1] = () => unref(thisWindow).Reload()),
                              size: "small"
                            }, {
                              icon: withCtx(() => [
                                createVNode(_component_icon_translate)
                              ]),
                              default: withCtx(() => _cache[10] || (_cache[10] = [
                                createTextVNode("立即重置")
                              ])),
                              _: 1
                            })
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    }),
                    createVNode(_component_a_col, { span: 12 }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, {
                          field: "OpenDevTools",
                          label: "打开调试工具"
                        }, {
                          default: withCtx(() => [
                            createVNode(_component_a_button, {
                              type: "primary",
                              onClick: _cache[2] || (_cache[2] = () => unref(thisWindow).OpenDevTools()),
                              size: "small"
                            }, {
                              icon: withCtx(() => [
                                createVNode(_component_icon_launch)
                              ]),
                              default: withCtx(() => _cache[11] || (_cache[11] = [
                                createTextVNode("立即打开")
                              ])),
                              _: 1
                            })
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    })
                  ]),
                  _: 1
                })
              ]),
              _: 1
            }, 8, ["model"])
          ])
        ]),
        createBaseVNode("div", _hoisted_4, [
          createBaseVNode("div", _hoisted_5, [
            createVNode(_component_a_descriptions, {
              title: "系统信息",
              column: 1
            }, {
              default: withCtx(() => [
                createVNode(_component_a_descriptions_item, { label: "窗口name" }, {
                  default: withCtx(() => [
                    createTextVNode(toDisplayString(windowsname.value), 1)
                  ]),
                  _: 1
                }),
                createVNode(_component_a_descriptions_item, { label: "系统(OS)" }, {
                  default: withCtx(() => {
                    var _a;
                    return [
                      createTextVNode(toDisplayString((_a = Osinfo.value) == null ? void 0 : _a.OS), 1)
                    ];
                  }),
                  _: 1
                }),
                createVNode(_component_a_descriptions_item, { label: "系统品牌" }, {
                  default: withCtx(() => {
                    var _a;
                    return [
                      createTextVNode(toDisplayString((_a = Osinfo.value) == null ? void 0 : _a.OSInfo.Branding), 1)
                    ];
                  }),
                  _: 1
                }),
                createVNode(_component_a_descriptions_item, { label: "系统版本" }, {
                  default: withCtx(() => {
                    var _a;
                    return [
                      createTextVNode(toDisplayString((_a = Osinfo.value) == null ? void 0 : _a.OSInfo.Version), 1)
                    ];
                  }),
                  _: 1
                }),
                createVNode(_component_a_descriptions_item, { label: "Arch" }, {
                  default: withCtx(() => {
                    var _a;
                    return [
                      createTextVNode(toDisplayString((_a = Osinfo.value) == null ? void 0 : _a.Arch), 1)
                    ];
                  }),
                  _: 1
                }),
                createVNode(_component_a_descriptions_item, { label: "调试状态" }, {
                  default: withCtx(() => {
                    var _a;
                    return [
                      createTextVNode(toDisplayString((_a = Osinfo.value) == null ? void 0 : _a.Debug), 1)
                    ];
                  }),
                  _: 1
                }),
                createVNode(_component_a_descriptions_item, { label: "WebView2" }, {
                  default: withCtx(() => {
                    var _a;
                    return [
                      createTextVNode(toDisplayString((_a = Osinfo.value) == null ? void 0 : _a.PlatformInfo.WebView2), 1)
                    ];
                  }),
                  _: 1
                })
              ]),
              _: 1
            })
          ])
        ])
      ]);
    };
  }
});
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-446eec0f"]]);
export {
  index as default
};
