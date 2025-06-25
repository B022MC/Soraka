import { defineStore } from "pinia";
import { setToken, getToken, clearToken } from "@/utils/auth";

export interface UserState {
  accountId?: number;
  summonerId?: number;
  puuid?: string;
  nickname?: string;
  avatar?: string;
  region?: string;
  tag?: string;
  rank?: string;
  winRate?: number;
  wins?: number;
  losses?: number;
  totalGames?: number;
  createtime?: string;
  server?: string;
  level?: number;
  xpSinceLastLevel?: number;
  xpUntilNextLevel?: number;
  percentCompleteForNextLevel?: number;
  privacy?: string;
}

const useUserStore = defineStore("user", {
  state: (): UserState => ({
    accountId: undefined,
    summonerId: undefined,
    puuid: undefined,
    nickname: undefined,
    avatar: undefined,
    region: "",
    tag: "",
    rank: "",
    winRate: 0,
    wins: 0,
    losses: 0,
    totalGames: 0,
    createtime: "",
    server: "",
    level: undefined,
    xpSinceLastLevel: undefined,
    xpUntilNextLevel: undefined,
    percentCompleteForNextLevel: undefined,
    privacy: "",
  }),

  getters: {
    userInfo(state): UserState {
      return { ...state };
    },
  },

  actions: {
    setInfo(partial: Partial<UserState>) {
      this.$patch(partial);
    },

    resetInfo() {
      this.$reset();
    },

    async setTokenArr(token: string) {
      setToken(token);
    },

    async info() {
      // const res = await getUserInfo()
      // this.setInfo(res)
    },

    async login() {
      setToken(undefined);
    },

    async logout(goLogin = false) {
      // TODO: 清空信息，跳转登录页等
    },

    clearloginfo() {
      clearToken();
    },
  }
});

// 持久化配置
(useUserStore as any).persist = {
  enabled: true,
  strategies: [
    {
      key: "userStore",
      storage: localStorage,
      paths: [
        "accountId",
        "summonerId",
        "puuid",
        "nickname",
        "avatar",
        "region",
        "tag",
        "rank",
        "winRate",
        "wins",
        "losses",
        "totalGames",
        "createtime",
        "server",
        "level",
        "xpSinceLastLevel",
        "xpUntilNextLevel",
        "percentCompleteForNextLevel",
        "privacy"
      ]
    }
  ]
};

export default useUserStore;
