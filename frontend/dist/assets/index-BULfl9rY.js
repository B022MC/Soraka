import { d as defineComponent, u as useUserStore, c as computed, w as watch, h as resolveComponent, o as openBlock, a as createBlock, e as withCtx, j as createBaseVNode, i as createVNode, q as createTextVNode, p as unref, t as toDisplayString, m as createCommentVNode, k as createElementBlock, l as renderList, n as normalizeClass, F as Fragment, r as ref, s as onMounted } from "./index-BT_C9izJ.js";
import { _ as _export_sfc } from "./_plugin-vue_export-helper-1tPrXgE0.js";
import { L as ListRecentMatches } from "./lcuapiwails-IFZpWFl2.js";
const _hoisted_1$3 = { class: "actions" };
const _hoisted_2$3 = { class: "level-avatar" };
const _hoisted_3$2 = ["src"];
const _hoisted_4$2 = { class: "level-text" };
const _hoisted_5$2 = { class: "name-line" };
const _hoisted_6$2 = { class: "nickname" };
const _hoisted_7$2 = { class: "tag-line" };
const _sfc_main$3 = /* @__PURE__ */ defineComponent({
  __name: "SummonerHeader",
  setup(__props) {
    const userStore = useUserStore();
    const defaultAvatar = new URL("/assets/logo-C9STPBS-.png", import.meta.url).href;
    const expPercent = computed(() => {
      if (userStore.percentCompleteForNextLevel !== void 0 && userStore.percentCompleteForNextLevel !== null) {
        return Math.min(Math.max(userStore.percentCompleteForNextLevel, 0), 100);
      }
      const current = userStore.xpSinceLastLevel ?? 0;
      const next = userStore.xpUntilNextLevel ?? 0;
      const total = current + next;
      if (total > 0) {
        return Math.min(Math.max(current / total * 100, 0), 100);
      }
      return 0;
    });
    console.log("当前 userStore 全部数据:", { ...userStore });
    console.log("当前 privacy:", userStore.privacy);
    watch(
      () => userStore.privacy,
      (val) => {
        console.log("privacy 变化了:", val);
      },
      { immediate: true }
    );
    watch(
      () => [userStore.xpSinceLastLevel, userStore.xpUntilNextLevel, userStore.percentCompleteForNextLevel],
      () => {
        console.log("经验相关变化了:", {
          xpSinceLastLevel: userStore.xpSinceLastLevel,
          xpUntilNextLevel: userStore.xpUntilNextLevel,
          percentCompleteForNextLevel: userStore.percentCompleteForNextLevel,
          percent: expPercent.value
        });
      },
      { immediate: true }
    );
    const onRefresh = () => {
      console.log("点击刷新召唤师信息");
    };
    const onShowHistory = () => {
      console.log("点击查看历史战绩");
    };
    return (_ctx, _cache) => {
      const _component_a_button = resolveComponent("a-button");
      const _component_a_progress = resolveComponent("a-progress");
      const _component_icon_font = resolveComponent("icon-font");
      const _component_a_card = resolveComponent("a-card");
      return openBlock(), createBlock(_component_a_card, {
        class: "summoner-header",
        bordered: false
      }, {
        default: withCtx(() => [
          createBaseVNode("div", _hoisted_1$3, [
            createVNode(_component_a_button, {
              type: "outline",
              size: "small",
              onClick: onRefresh
            }, {
              default: withCtx(() => _cache[0] || (_cache[0] = [
                createTextVNode("刷新")
              ])),
              _: 1
            }),
            createVNode(_component_a_button, {
              type: "text",
              size: "small",
              onClick: onShowHistory
            }, {
              default: withCtx(() => _cache[1] || (_cache[1] = [
                createTextVNode("历史战绩")
              ])),
              _: 1
            })
          ]),
          createBaseVNode("div", _hoisted_2$3, [
            createVNode(_component_a_progress, {
              type: "circle",
              percent: expPercent.value / 100,
              width: 88,
              strokeColor: "#00b42a",
              format: () => "",
              class: "avatar-progress"
            }, null, 8, ["percent"]),
            createBaseVNode("img", {
              src: unref(userStore).avatar || unref(defaultAvatar),
              class: "avatar-img"
            }, null, 8, _hoisted_3$2),
            createBaseVNode("div", _hoisted_4$2, toDisplayString(unref(userStore).level ?? "--"), 1)
          ]),
          createBaseVNode("div", _hoisted_5$2, [
            createBaseVNode("span", _hoisted_6$2, toDisplayString(unref(userStore).nickname || "未知召唤师"), 1),
            unref(userStore).privacy === "PRIVATE" ? (openBlock(), createBlock(_component_icon_font, {
              key: 0,
              type: "icon-Lock",
              class: "lock-icon"
            })) : createCommentVNode("", true)
          ]),
          createBaseVNode("div", _hoisted_7$2, "#" + toDisplayString(unref(userStore).tag || "00000"), 1)
        ]),
        _: 1
      });
    };
  }
});
const SummonerHeader = /* @__PURE__ */ _export_sfc(_sfc_main$3, [["__scopeId", "data-v-2f7d7f5c"]]);
const _hoisted_1$2 = { class: "match-scroll-list" };
const _hoisted_2$2 = { class: "match-content" };
const _hoisted_3$1 = { class: "left" };
const _hoisted_4$1 = { class: "left-top" };
const _hoisted_5$1 = { class: "avatar-wrapper" };
const _hoisted_6$1 = ["src"];
const _hoisted_7$1 = { class: "level" };
const _hoisted_8 = { class: "info-wrapper" };
const _hoisted_9 = { class: "mode" };
const _hoisted_10 = { class: "spells" };
const _hoisted_11 = ["src"];
const _hoisted_12 = { class: "middle" };
const _hoisted_13 = { class: "middle-row" };
const _hoisted_14 = { class: "kda-line" };
const _hoisted_15 = { class: "deaths" };
const _hoisted_16 = { class: "cs" };
const _hoisted_17 = { class: "items" };
const _hoisted_18 = ["src"];
const _hoisted_19 = { class: "gold" };
const _hoisted_20 = { class: "right" };
const _hoisted_21 = { class: "meta" };
const _hoisted_22 = { class: "map" };
const _hoisted_23 = { class: "time" };
const _sfc_main$2 = /* @__PURE__ */ defineComponent({
  __name: "MatchList",
  props: {
    matches: {}
  },
  setup(__props) {
    return (_ctx, _cache) => {
      const _component_a_card = resolveComponent("a-card");
      return openBlock(), createBlock(_component_a_card, {
        class: "match-wrapper",
        bordered: false
      }, {
        default: withCtx(() => [
          createBaseVNode("div", _hoisted_1$2, [
            (openBlock(true), createElementBlock(Fragment, null, renderList(_ctx.matches, (match) => {
              return openBlock(), createBlock(_component_a_card, {
                key: match.id,
                class: normalizeClass(["match-item", { win: match.result === "win", lose: match.result === "lose" }]),
                bordered: false
              }, {
                default: withCtx(() => [
                  createBaseVNode("div", _hoisted_2$2, [
                    createBaseVNode("div", _hoisted_3$1, [
                      createBaseVNode("div", _hoisted_4$1, [
                        createBaseVNode("div", _hoisted_5$1, [
                          createBaseVNode("img", {
                            src: match.champion,
                            class: "champion-avatar"
                          }, null, 8, _hoisted_6$1),
                          createBaseVNode("span", _hoisted_7$1, toDisplayString(match.level), 1)
                        ]),
                        createBaseVNode("div", _hoisted_8, [
                          createBaseVNode("span", {
                            class: normalizeClass(["result", match.result])
                          }, toDisplayString(match.result === "win" ? "胜利" : "失败"), 3),
                          createBaseVNode("span", _hoisted_9, toDisplayString(match.mode), 1),
                          createBaseVNode("div", _hoisted_10, [
                            (openBlock(true), createElementBlock(Fragment, null, renderList(match.spells, (spell, idx) => {
                              return openBlock(), createElementBlock("img", {
                                key: idx,
                                src: spell,
                                class: "spell-icon"
                              }, null, 8, _hoisted_11);
                            }), 128))
                          ])
                        ])
                      ])
                    ]),
                    createBaseVNode("div", _hoisted_12, [
                      createBaseVNode("div", _hoisted_13, [
                        createBaseVNode("div", _hoisted_14, [
                          createBaseVNode("strong", null, toDisplayString(match.kills), 1),
                          _cache[0] || (_cache[0] = createTextVNode(" /")),
                          createBaseVNode("strong", _hoisted_15, toDisplayString(match.deaths), 1),
                          _cache[1] || (_cache[1] = createTextVNode(" /")),
                          createBaseVNode("strong", null, toDisplayString(match.assists), 1)
                        ]),
                        createBaseVNode("span", _hoisted_16, toDisplayString(match.cs) + " 补刀", 1),
                        createBaseVNode("div", _hoisted_17, [
                          (openBlock(true), createElementBlock(Fragment, null, renderList(match.items, (item, idx) => {
                            return openBlock(), createElementBlock("img", {
                              key: idx,
                              src: item,
                              class: "item-icon"
                            }, null, 8, _hoisted_18);
                          }), 128))
                        ]),
                        createBaseVNode("span", _hoisted_19, toDisplayString(match.gold.toLocaleString()) + " 金币", 1)
                      ])
                    ]),
                    createBaseVNode("div", _hoisted_20, [
                      createBaseVNode("div", _hoisted_21, [
                        createBaseVNode("span", _hoisted_22, toDisplayString(match.map), 1),
                        createBaseVNode("span", _hoisted_23, toDisplayString(match.time), 1)
                      ])
                    ])
                  ])
                ]),
                _: 2
              }, 1032, ["class"]);
            }), 128))
          ])
        ]),
        _: 1
      });
    };
  }
});
const MatchList = /* @__PURE__ */ _export_sfc(_sfc_main$2, [["__scopeId", "data-v-89fc28ef"]]);
const _hoisted_1$1 = { class: "filter-header" };
const _hoisted_2$1 = { class: "summary" };
const _hoisted_3 = { class: "summary-win" };
const _hoisted_4 = { class: "summary-lose" };
const _hoisted_5 = { class: "avatars" };
const _hoisted_6 = ["src"];
const _hoisted_7 = { class: "tools" };
const _sfc_main$1 = /* @__PURE__ */ defineComponent({
  __name: "MatchFilterHeader",
  props: {
    matches: {}
  },
  emits: ["filterChange"],
  setup(__props, { emit: __emit }) {
    const props = __props;
    const emit = __emit;
    const total = computed(() => props.matches.length);
    const winCount = computed(() => props.matches.filter((m) => m.result === "win").length);
    const loseCount = computed(() => props.matches.filter((m) => m.result === "lose").length);
    const kdaStr = computed(() => {
      const k = props.matches.reduce((sum, m) => sum + m.kills, 0);
      const d = props.matches.reduce((sum, m) => sum + m.deaths, 0);
      const a = props.matches.reduce((sum, m) => sum + m.assists, 0);
      const ratio = d === 0 ? k + a : ((k + a) / d).toFixed(1);
      return `${k} / ${d} / ${a} (${ratio})`;
    });
    const recentChampions = computed(() => {
      const seen = /* @__PURE__ */ new Set();
      return props.matches.map((m) => m.champion).filter((c) => {
        if (seen.has(c)) return false;
        seen.add(c);
        return true;
      }).slice(0, 10);
    });
    const filterModes = ["全部", "匹配模式", "极地大乱斗", "单 / 双排", "灵活排位"];
    return (_ctx, _cache) => {
      const _component_a_button = resolveComponent("a-button");
      const _component_a_doption = resolveComponent("a-doption");
      const _component_a_dropdown = resolveComponent("a-dropdown");
      return openBlock(), createElementBlock("div", _hoisted_1$1, [
        createBaseVNode("div", _hoisted_2$1, [
          createBaseVNode("span", null, "近期对局（最近 " + toDisplayString(total.value) + " 场）", 1),
          createBaseVNode("span", _hoisted_3, "胜：" + toDisplayString(winCount.value), 1),
          createBaseVNode("span", _hoisted_4, "负：" + toDisplayString(loseCount.value), 1),
          createBaseVNode("span", null, "KDA：" + toDisplayString(kdaStr.value), 1)
        ]),
        createBaseVNode("div", _hoisted_5, [
          (openBlock(true), createElementBlock(Fragment, null, renderList(recentChampions.value, (hero, index2) => {
            return openBlock(), createElementBlock("img", {
              key: index2,
              src: hero,
              class: "champion-icon"
            }, null, 8, _hoisted_6);
          }), 128))
        ]),
        createBaseVNode("div", _hoisted_7, [
          createVNode(_component_a_button, {
            size: "small",
            type: "outline"
          }, {
            default: withCtx(() => _cache[0] || (_cache[0] = [
              createTextVNode("最近队友")
            ])),
            _: 1
          }),
          createVNode(_component_a_dropdown, null, {
            content: withCtx(() => [
              (openBlock(), createElementBlock(Fragment, null, renderList(filterModes, (mode) => {
                return createVNode(_component_a_doption, {
                  key: mode,
                  onClick: () => emit("filterChange", mode)
                }, {
                  default: withCtx(() => [
                    createTextVNode(toDisplayString(mode), 1)
                  ]),
                  _: 2
                }, 1032, ["onClick"]);
              }), 64))
            ]),
            default: withCtx(() => [
              createVNode(_component_a_button, { size: "small" }, {
                default: withCtx(() => _cache[1] || (_cache[1] = [
                  createTextVNode("全部")
                ])),
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
const MatchFilterHeader = /* @__PURE__ */ _export_sfc(_sfc_main$1, [["__scopeId", "data-v-ece1dd5b"]]);
const _hoisted_1 = { class: "life-page" };
const _hoisted_2 = { class: "match-scroll-container" };
const _sfc_main = /* @__PURE__ */ defineComponent({
  __name: "index",
  setup(__props) {
    function mapResult(result) {
      return result === "胜利" ? "win" : "lose";
    }
    const matchList = ref([]);
    const selectedMode = ref("全部");
    const filteredMatches = computed(() => {
      if (selectedMode.value === "全部") return matchList.value;
      return matchList.value.filter((m) => m.mode === selectedMode.value);
    });
    function onFilterChange(mode) {
      selectedMode.value = mode;
    }
    onMounted(async () => {
      try {
        const p = ListRecentMatches(20);
        const data = await p;
        matchList.value = (data ?? []).filter((m) => m !== null).map((m) => ({
          ...m,
          result: mapResult(m.result)
        }));
      } catch (e) {
        console.error("获取比赛数据失败", e);
      }
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("div", _hoisted_1, [
        createVNode(SummonerHeader),
        createVNode(MatchFilterHeader, {
          matches: matchList.value,
          onFilterChange
        }, null, 8, ["matches"]),
        createBaseVNode("div", _hoisted_2, [
          createVNode(MatchList, { matches: filteredMatches.value }, null, 8, ["matches"])
        ])
      ]);
    };
  }
});
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-ed78a3a7"]]);
export {
  index as default
};
