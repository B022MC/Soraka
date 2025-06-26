// iconFont.ts
import { h, defineComponent, computed } from 'vue';
import { Icon } from '@arco-design/web-vue';
import useAppStore from '@/store/modules/app';

const IconFont = Icon.addFromIconFontCn({
  src: 'https://at.alicdn.com/t/c/font_4954474_cghrq9lcye6.js'
});

export default defineComponent({
  name: 'IconFont',
  props: {
    type: { type: String, required: true },
    size: { type: Number, default: 18 }
  },
  setup(props) {
    const appStore = useAppStore();

    const dynamicType = computed(() => {
      // 如果传入的 type 已带白天 / 夜间后缀，直接返回
      if (/_white$|_black$/.test(props.type)) {
        return props.type;
      }

      // 根据主题拼接后缀
      const suffix = appStore.getTheme === 'dark' ? 'white' : 'black';
      return `${props.type}_${suffix}`;
    });

    return () =>
        h(IconFont, {
          type: dynamicType.value,
          style: { fontSize: `${props.size}px` }
        });
  }
});
