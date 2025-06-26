import { defineStore } from "pinia";
import type { AppState } from "./types";

const useAppStore = defineStore("app", {
  state: (): AppState => ({
    system: {
      theme: "light",
      sysTime: "",
    },
    client: {
      clientPath: "",
    },
    lcu: {
      online: false,
      port: "",
      token: "",
    },
  }),

  getters: {
    getTheme: (state) => state.system.theme,
    getSysTime: (state) => state.system.sysTime,
    getClientPath: (state) => state.client.clientPath,
    getLcuOnline: (state) => state.lcu.online,
    getLcuPort: (state) => state.lcu.port,
    getLcuToken: (state) => state.lcu.token,
  },

  actions: {
    toggleTheme(dark: boolean) {
      if (dark) {
        this.system.theme = "dark";
        document.body.setAttribute("arco-theme", "dark");
      } else {
        this.system.theme = "light";
        document.body.removeAttribute("arco-theme");
      }
      setTimeout(() => {
        window.dispatchEvent(new Event("resize"));
      }, 0);
    },

    setSysTime(time: string) {
      this.system.sysTime = time;
    },

    setClientPath(path: string) {
      this.client.clientPath = path;
    },

    setLcuStatus(online: boolean) {
      this.lcu.online = online;
    },

    setLcuCreds(port: string, token: string) {
      this.lcu.port = port;
      this.lcu.token = token;
    },
  },
});

export default useAppStore;
