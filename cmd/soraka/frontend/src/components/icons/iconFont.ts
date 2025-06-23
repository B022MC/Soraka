// iconFont.ts
import { h, defineComponent, computed } from 'vue'
import { Icon } from '@arco-design/web-vue'
import useAppStore from '@/store/modules/app'

const IconFont = Icon.addFromIconFontCn({
  src: 'https://at.alicdn.com/t/c/font_4954474_cghrq9lcye6.js'
})

export default defineComponent({
  name: 'icon-font',
  props: {
    type: { type: String, required: true },
    size: { type: Number, default: 18 }
  },
  setup(props) {
    const appStore = useAppStore()

    const dynamicType = computed(() => {
      if (props.type.endsWith('_white') || props.type.endsWith('_black')) {
        return props.type
      }
      const suffix = appStore.theme === 'dark' ? 'white' : 'black'
      return `${props.type}_${suffix}`
    })

    return () =>
        h(IconFont, {
          type: dynamicType.value,
          style: { fontSize: `${props.size}px` }
        })
  }
})
