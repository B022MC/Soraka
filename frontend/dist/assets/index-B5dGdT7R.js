import { d as defineComponent, A as reactive, j as createElementBlock, i as createBaseVNode, h as createVNode, b as withCtx, g as resolveComponent, o as openBlock, p as createTextVNode } from "./index-BYmKbRsy.js";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
const _hoisted_1 = { class: "life-page" };
const _hoisted_2 = { class: "section" };
const _hoisted_3 = { class: "collapse-body" };
const _hoisted_4 = { class: "collapse-body" };
const _hoisted_5 = { class: "collapse-body" };
const _hoisted_6 = { class: "collapse-body" };
const _hoisted_7 = { class: "section" };
const _hoisted_8 = { class: "collapse-body" };
const _hoisted_9 = { class: "collapse-body" };
const _hoisted_10 = { class: "collapse-body" };
const _hoisted_11 = { class: "collapse-body" };
const _sfc_main = /* @__PURE__ */ defineComponent({
  __name: "index",
  setup(__props) {
    const form = reactive({
      autoAcceptDelay: 0,
      autoAcceptEnabled: false,
      autoAcceptBlindExchange: false,
      autoAcceptExplicitExchange: false,
      defaultChampion: "",
      top: "",
      jungle: "",
      mid: "",
      bottom: "",
      support: "",
      enableAutoPick: false,
      autoReconnect: false,
      practiceRoomName: "",
      practiceRoomPassword: "",
      spectateSummoner: "",
      spectateMode: "lcu",
      lockGameSettings: false
    });
    return (_ctx, _cache) => {
      const _component_a_typography_title = resolveComponent("a-typography-title");
      const _component_a_input_number = resolveComponent("a-input-number");
      const _component_a_form_item = resolveComponent("a-form-item");
      const _component_a_switch = resolveComponent("a-switch");
      const _component_a_form = resolveComponent("a-form");
      const _component_a_collapse_item = resolveComponent("a-collapse-item");
      const _component_a_input = resolveComponent("a-input");
      const _component_a_button = resolveComponent("a-button");
      const _component_a_space = resolveComponent("a-space");
      const _component_a_col = resolveComponent("a-col");
      const _component_a_row = resolveComponent("a-row");
      const _component_a_tag = resolveComponent("a-tag");
      const _component_a_collapse = resolveComponent("a-collapse");
      const _component_a_option = resolveComponent("a-option");
      const _component_a_select = resolveComponent("a-select");
      return openBlock(), createElementBlock("div", _hoisted_1, [
        _cache[39] || (_cache[39] = createBaseVNode("h1", { class: "section-title" }, "其它功能", -1)),
        createBaseVNode("div", _hoisted_2, [
          createVNode(_component_a_typography_title, { heading: 5 }, {
            default: withCtx(() => _cache[17] || (_cache[17] = [
              createTextVNode("英雄选择")
            ])),
            _: 1
          }),
          createVNode(_component_a_collapse, {
            "default-active-key": [],
            "expand-icon-position": "right",
            class: "collapse-wrapper"
          }, {
            default: withCtx(() => [
              createVNode(_component_a_collapse_item, {
                header: "自动接受对局",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[18] || (_cache[18] = createBaseVNode("div", { class: "item-desc" }, "在你设置的倒数之后自动接受对局匹配", -1)),
                  createBaseVNode("div", _hoisted_3, [
                    createVNode(_component_a_form, { layout: "vertical" }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, { label: "在对局找到后延迟接受对局的秒数：" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_input_number, {
                              modelValue: form.autoAcceptDelay,
                              "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => form.autoAcceptDelay = $event),
                              min: 0,
                              max: 20
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        }),
                        createVNode(_component_a_form_item, { label: "是否启用" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_switch, {
                              modelValue: form.autoAcceptEnabled,
                              "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => form.autoAcceptEnabled = $event)
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    })
                  ])
                ]),
                _: 1
              }),
              createVNode(_component_a_collapse_item, {
                header: "自动接受交换请求",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[19] || (_cache[19] = createBaseVNode("div", { class: "item-desc" }, "自动接受队友发起的英雄交换请求", -1)),
                  createBaseVNode("div", _hoisted_4, [
                    createVNode(_component_a_form, { layout: "vertical" }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, { label: "自动接受模糊层交换请求：" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_switch, {
                              modelValue: form.autoAcceptBlindExchange,
                              "onUpdate:modelValue": _cache[2] || (_cache[2] = ($event) => form.autoAcceptBlindExchange = $event)
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        }),
                        createVNode(_component_a_form_item, { label: "自动接受明确交换请求：" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_switch, {
                              modelValue: form.autoAcceptExplicitExchange,
                              "onUpdate:modelValue": _cache[3] || (_cache[3] = ($event) => form.autoAcceptExplicitExchange = $event)
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    })
                  ])
                ]),
                _: 1
              }),
              createVNode(_component_a_collapse_item, {
                header: "自动亮起英雄",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[26] || (_cache[26] = createBaseVNode("div", { class: "item-desc" }, "在你的选定阶段时自动亮起预设英雄", -1)),
                  createBaseVNode("div", _hoisted_5, [
                    createVNode(_component_a_form, { layout: "vertical" }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, { label: "默认设置" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_space, { align: "center" }, {
                              default: withCtx(() => [
                                createVNode(_component_a_input, {
                                  modelValue: form.defaultChampion,
                                  "onUpdate:modelValue": _cache[4] || (_cache[4] = ($event) => form.defaultChampion = $event),
                                  placeholder: "默认亮起英雄",
                                  disabled: ""
                                }, null, 8, ["modelValue"]),
                                createVNode(_component_a_button, { type: "outline" }, {
                                  default: withCtx(() => _cache[20] || (_cache[20] = [
                                    createTextVNode("选择")
                                  ])),
                                  _: 1
                                })
                              ]),
                              _: 1
                            })
                          ]),
                          _: 1
                        }),
                        createVNode(_component_a_form_item, { label: "按位置设置" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_row, { gutter: 12 }, {
                              default: withCtx(() => [
                                createVNode(_component_a_col, { span: 12 }, {
                                  default: withCtx(() => [
                                    createVNode(_component_a_form_item, { label: "上路：" }, {
                                      default: withCtx(() => [
                                        createVNode(_component_a_space, null, {
                                          default: withCtx(() => [
                                            createVNode(_component_a_input, {
                                              modelValue: form.top,
                                              "onUpdate:modelValue": _cache[5] || (_cache[5] = ($event) => form.top = $event),
                                              disabled: ""
                                            }, null, 8, ["modelValue"]),
                                            createVNode(_component_a_button, { type: "outline" }, {
                                              default: withCtx(() => _cache[21] || (_cache[21] = [
                                                createTextVNode("选择")
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
                                }),
                                createVNode(_component_a_col, { span: 12 }, {
                                  default: withCtx(() => [
                                    createVNode(_component_a_form_item, { label: "打野：" }, {
                                      default: withCtx(() => [
                                        createVNode(_component_a_space, null, {
                                          default: withCtx(() => [
                                            createVNode(_component_a_input, {
                                              modelValue: form.jungle,
                                              "onUpdate:modelValue": _cache[6] || (_cache[6] = ($event) => form.jungle = $event),
                                              disabled: ""
                                            }, null, 8, ["modelValue"]),
                                            createVNode(_component_a_button, { type: "outline" }, {
                                              default: withCtx(() => _cache[22] || (_cache[22] = [
                                                createTextVNode("选择")
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
                            }),
                            createVNode(_component_a_row, { gutter: 12 }, {
                              default: withCtx(() => [
                                createVNode(_component_a_col, { span: 12 }, {
                                  default: withCtx(() => [
                                    createVNode(_component_a_form_item, { label: "中路：" }, {
                                      default: withCtx(() => [
                                        createVNode(_component_a_space, null, {
                                          default: withCtx(() => [
                                            createVNode(_component_a_input, {
                                              modelValue: form.mid,
                                              "onUpdate:modelValue": _cache[7] || (_cache[7] = ($event) => form.mid = $event),
                                              disabled: ""
                                            }, null, 8, ["modelValue"]),
                                            createVNode(_component_a_button, { type: "outline" }, {
                                              default: withCtx(() => _cache[23] || (_cache[23] = [
                                                createTextVNode("选择")
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
                                }),
                                createVNode(_component_a_col, { span: 12 }, {
                                  default: withCtx(() => [
                                    createVNode(_component_a_form_item, { label: "下路：" }, {
                                      default: withCtx(() => [
                                        createVNode(_component_a_space, null, {
                                          default: withCtx(() => [
                                            createVNode(_component_a_input, {
                                              modelValue: form.bottom,
                                              "onUpdate:modelValue": _cache[8] || (_cache[8] = ($event) => form.bottom = $event),
                                              disabled: ""
                                            }, null, 8, ["modelValue"]),
                                            createVNode(_component_a_button, { type: "outline" }, {
                                              default: withCtx(() => _cache[24] || (_cache[24] = [
                                                createTextVNode("选择")
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
                            }),
                            createVNode(_component_a_form_item, { label: "辅助：" }, {
                              default: withCtx(() => [
                                createVNode(_component_a_space, null, {
                                  default: withCtx(() => [
                                    createVNode(_component_a_input, {
                                      modelValue: form.support,
                                      "onUpdate:modelValue": _cache[9] || (_cache[9] = ($event) => form.support = $event),
                                      disabled: ""
                                    }, null, 8, ["modelValue"]),
                                    createVNode(_component_a_button, { type: "outline" }, {
                                      default: withCtx(() => _cache[25] || (_cache[25] = [
                                        createTextVNode("选择")
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
                        }),
                        createVNode(_component_a_form_item, { label: "启用自动亮起" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_switch, {
                              modelValue: form.enableAutoPick,
                              "onUpdate:modelValue": _cache[10] || (_cache[10] = ($event) => form.enableAutoPick = $event)
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        })
                      ]),
                      _: 1
                    })
                  ])
                ]),
                _: 1
              }),
              createVNode(_component_a_collapse_item, {
                header: "自动禁用英雄",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[29] || (_cache[29] = createBaseVNode("div", { class: "item-desc" }, "在你的禁用环节时自动禁用预设英雄", -1)),
                  createBaseVNode("div", _hoisted_6, [
                    createVNode(_component_a_space, {
                      align: "center",
                      justify: "space-between"
                    }, {
                      default: withCtx(() => [
                        createVNode(_component_a_tag, { color: "gray" }, {
                          default: withCtx(() => _cache[27] || (_cache[27] = [
                            createTextVNode("未启用")
                          ])),
                          _: 1
                        }),
                        createVNode(_component_a_button, { type: "outline" }, {
                          default: withCtx(() => _cache[28] || (_cache[28] = [
                            createTextVNode("恢复默认")
                          ])),
                          _: 1
                        })
                      ]),
                      _: 1
                    })
                  ])
                ]),
                _: 1
              })
            ]),
            _: 1
          })
        ]),
        createBaseVNode("div", _hoisted_7, [
          createVNode(_component_a_typography_title, { heading: 5 }, {
            default: withCtx(() => _cache[30] || (_cache[30] = [
              createTextVNode("游戏")
            ])),
            _: 1
          }),
          createVNode(_component_a_collapse, {
            "default-active-key": [],
            "expand-icon-position": "right",
            class: "collapse-wrapper"
          }, {
            default: withCtx(() => [
              createVNode(_component_a_collapse_item, {
                header: "自动重连",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[31] || (_cache[31] = createBaseVNode("div", { class: "item-desc" }, "当你掉线退出游戏时自动重新连接", -1)),
                  createBaseVNode("div", _hoisted_8, [
                    createVNode(_component_a_switch, {
                      modelValue: form.autoReconnect,
                      "onUpdate:modelValue": _cache[11] || (_cache[11] = ($event) => form.autoReconnect = $event)
                    }, null, 8, ["modelValue"])
                  ])
                ]),
                _: 1
              }),
              createVNode(_component_a_collapse_item, {
                header: "创建 5v5 练习模式",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[33] || (_cache[33] = createBaseVNode("div", { class: "item-desc" }, "只能添加人机玩家", -1)),
                  createBaseVNode("div", _hoisted_9, [
                    createVNode(_component_a_form, { layout: "vertical" }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, { label: "房间名：（不可为空）" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_input, {
                              modelValue: form.practiceRoomName,
                              "onUpdate:modelValue": _cache[12] || (_cache[12] = ($event) => form.practiceRoomName = $event),
                              placeholder: "请输入房间名"
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        }),
                        createVNode(_component_a_form_item, { label: "房间密码：（若留空则不设密码）" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_input, {
                              modelValue: form.practiceRoomPassword,
                              "onUpdate:modelValue": _cache[13] || (_cache[13] = ($event) => form.practiceRoomPassword = $event),
                              placeholder: "请输入房间密码"
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        }),
                        createVNode(_component_a_button, { type: "primary" }, {
                          default: withCtx(() => _cache[32] || (_cache[32] = [
                            createTextVNode("创建")
                          ])),
                          _: 1
                        })
                      ]),
                      _: 1
                    })
                  ])
                ]),
                _: 1
              }),
              createVNode(_component_a_collapse_item, {
                header: "观战",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[37] || (_cache[37] = createBaseVNode("div", { class: "item-desc" }, "观战同大区玩家正在进行的游戏", -1)),
                  createBaseVNode("div", _hoisted_10, [
                    createVNode(_component_a_form, { layout: "vertical" }, {
                      default: withCtx(() => [
                        createVNode(_component_a_form_item, { label: "你想观战的同大区召唤师名及编号：" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_input, {
                              modelValue: form.spectateSummoner,
                              "onUpdate:modelValue": _cache[14] || (_cache[14] = ($event) => form.spectateSummoner = $event),
                              placeholder: "请输入召唤师名"
                            }, null, 8, ["modelValue"])
                          ]),
                          _: 1
                        }),
                        createVNode(_component_a_form_item, { label: "观战启动方式：" }, {
                          default: withCtx(() => [
                            createVNode(_component_a_select, {
                              modelValue: form.spectateMode,
                              "onUpdate:modelValue": _cache[15] || (_cache[15] = ($event) => form.spectateMode = $event)
                            }, {
                              default: withCtx(() => [
                                createVNode(_component_a_option, { value: "lcu" }, {
                                  default: withCtx(() => _cache[34] || (_cache[34] = [
                                    createTextVNode("LCU API")
                                  ])),
                                  _: 1
                                }),
                                createVNode(_component_a_option, { value: "riot" }, {
                                  default: withCtx(() => _cache[35] || (_cache[35] = [
                                    createTextVNode("Riot Client")
                                  ])),
                                  _: 1
                                })
                              ]),
                              _: 1
                            }, 8, ["modelValue"])
                          ]),
                          _: 1
                        }),
                        createVNode(_component_a_button, { type: "primary" }, {
                          default: withCtx(() => _cache[36] || (_cache[36] = [
                            createTextVNode("观战")
                          ])),
                          _: 1
                        })
                      ]),
                      _: 1
                    })
                  ])
                ]),
                _: 1
              }),
              createVNode(_component_a_collapse_item, {
                header: "锁定游戏设置",
                class: "collapse-item"
              }, {
                default: withCtx(() => [
                  _cache[38] || (_cache[38] = createBaseVNode("div", { class: "item-desc" }, "让你的游戏设置不会因更新或异常写回默认", -1)),
                  createBaseVNode("div", _hoisted_11, [
                    createVNode(_component_a_switch, {
                      modelValue: form.lockGameSettings,
                      "onUpdate:modelValue": _cache[16] || (_cache[16] = ($event) => form.lockGameSettings = $event)
                    }, null, 8, ["modelValue"])
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
});
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-36b89bf0"]]);
export {
  index as default
};
