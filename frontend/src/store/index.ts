import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';  // ✅ 新增
import useAppStore from './modules/app';
import useUserStore from './modules/user';

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);  // ✅ 新增

export { useAppStore, useUserStore };
export default pinia;
