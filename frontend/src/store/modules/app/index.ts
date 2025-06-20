import { defineStore } from "pinia";
import { AppState } from "./types";

const useAppStore = defineStore("app", {
  state: (): AppState => {
    return {
      sysTime: "",
      theme: "light",
      clientPath: "",
    };
  },
  getters: {
    //获取主题
    getTheme(): string {
      return this.theme;
    },
  },

  actions: {
    //设置主题
    toggleTheme(dark: boolean) {
      if (dark) {
        this.theme = "dark";
        document.body.setAttribute("arco-theme", "dark");
      } else {
        this.theme = "light";
        document.body.removeAttribute("arco-theme");
      }

      // ✅ 刷新图标
      setTimeout(() => {
        window.dispatchEvent(new Event("resize")); // 或强制刷新组件
      }, 0);
    },
  },
});

export default useAppStore;
