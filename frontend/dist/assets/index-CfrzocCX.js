import { z as defineStore } from "./index-BYmKbRsy.js";
const TOKEN_KEY = "biztoken";
const getToken = () => {
  return localStorage.getItem(TOKEN_KEY);
};
const setToken = (token) => {
  if (token) {
    localStorage.setItem(TOKEN_KEY, token);
  } else {
    localStorage.removeItem(TOKEN_KEY);
  }
};
const clearToken = () => {
  localStorage.removeItem(TOKEN_KEY);
};
const useUserStore = defineStore("user", {
  state: () => ({
    nickname: void 0,
    avatar: void 0,
    region: "",
    tag: "",
    rank: "",
    winRate: void 0,
    server: "",
    wins: void 0,
    losses: void 0,
    totalGames: void 0,
    createtime: "",
    level: void 0,
    xpSinceLastLevel: void 0,
    xpUntilNextLevel: void 0
  }),
  getters: {
    userInfo(state) {
      return { ...state };
    }
  },
  actions: {
    setInfo(partial) {
      this.$patch(partial);
    },
    resetInfo() {
      this.$reset();
    },
    async setTokenArr(token) {
      setToken(token);
    },
    async info() {
    },
    async login() {
      setToken(void 0);
    },
    async logout(goLogin = false) {
    },
    clearloginfo() {
      clearToken();
    }
  }
});
export {
  getToken as g,
  useUserStore as u
};
