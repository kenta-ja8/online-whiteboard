<script setup lang="ts">
import { toRefs, reactive, ref } from 'vue'

interface Emits {
  (e: "selectStickyNote", value: any): void;
}
const emit = defineEmits<Emits>();

interface Props {
  xpos: Number,
  ypos: Number,
  initWidth: Number,
  initHeight: Number,
  canvas: HTMLDivElement | undefined,
}
const props = defineProps<Props>()
const { xpos, canvas } = toRefs(props);


const state = reactive({
  xpos: 0,
  ypos: 0,
  width: 0,
  height: 0,
})

state.xpos = 100
state.ypos = 100
state.width = 100
state.height = 100

let pointerXOnSticky = 0
let pointerYOnSticky = 0

let isClickedSticky = false
const onMouseDownSticky = (e: any) => {
  console.log('onMouseDownSticky')
  isClickedSticky = true
  pointerXOnSticky = e.offsetX
  pointerYOnSticky = e.offsetY

  emit('selectStickyNote', {
    onMouseUp,
    onMouseMove,
  })

}


let isClickedScaleIcon = false
const onMouseDownScaleIcon = () => {
  isClickedScaleIcon = true
  emit('selectStickyNote', {
    onMouseUp,
    onMouseMove,
  })
}


const onMouseMove = (e: any) => {
  if (isClickedScaleIcon) {
    console.log('onMouseMoveScaleIconCom', e)

    const absoluteCanvas = canvas.value?.getBoundingClientRect()
    console.log('return', canvas)
    if (!absoluteCanvas) {
      return
    }
    // 描画エリアをはみ出す場合は処理を行わない
    if (absoluteCanvas.right < e.clientX || absoluteCanvas.bottom < e.clientY) return

    const newW = e.clientX - (absoluteCanvas.x + state.xpos)
    const newH = e.clientY - (absoluteCanvas.y + state.ypos)
    state.width = newW
    state.height = newH
    return
  }
  if (isClickedSticky) {
    console.log('onMouseMoveSticky', e.offsetX, e.offsetY, e)
    const absoluteCanvas = canvas.value?.getBoundingClientRect()
    if (!absoluteCanvas) return
    const absoluteNewX = e.clientX - pointerXOnSticky
    const absoluteNewY = e.clientY - pointerYOnSticky
    if (
      absoluteNewX < absoluteCanvas.left
      || absoluteNewY < absoluteCanvas.top
      || absoluteCanvas.right < (absoluteNewX + state.width)
      || absoluteCanvas.bottom < (absoluteNewY + state.height)
    ) return
    const newX = absoluteNewX - absoluteCanvas.x
    const newY = absoluteNewY - absoluteCanvas.y
    state.xpos = newX
    state.ypos = newY
    return
  }
}
const onMouseUp = () => { isClickedScaleIcon = isClickedSticky = false }



// export default defineComponent({
//   props: {
//     xpos: Number,
//     ypos: Number,
//     initWidth: Number,
//     initHeight: Number,
//     canvas: [HTMLCanvasElement, Object],
//     canvasTopLeft: Number,
//     canvasTopRight: Number,
//     canvasBottomRight: Number,
//     canvasBottomLeft: Number,
//   },
//   data() {
//     return {
//       width: this.initWidth,
//       height: this.initHeight
//     }
//   },
//   setup(props, context) {
//     const { xpos, canvas } = toRefs(props);


//     const state = reactive({
//       xpos: 0,
//       ypos: 0,
//       width: 0,
//       height: 0,
//     })

//     state.xpos = 100
//     state.ypos = 100
//     state.width = 100
//     state.height = 100

//     let pointerXOnSticky = 0
//     let pointerYOnSticky = 0

//     const canvasx = ref<HTMLDivElement>()
//     let isClickedSticky = false
//     const onMouseDownSticky = (e: any) => {
//       console.log('onMouseDownSticky')
//       isClickedSticky = true
//       pointerXOnSticky = e.offsetX
//       pointerYOnSticky = e.offsetY
//     }


//     let isClickedScaleIcon = false
//     const onMouseDownScaleIcon = () => {
//       isClickedScaleIcon = true
//       context.emit('selectStickyNote', {
//         onMouseUp,
//         onMouseMove,
//       })
//     }


//     const onMouseMove = (e: any) => {
//       if (isClickedScaleIcon) {
//         console.log('onMouseMoveScaleIconCom', e)

//         const absoluteCanvas = canvas.value?.getBoundingClientRect()
//         console.log('return', canvas)
//         if (!absoluteCanvas) {
//           return
//         }
//         // 描画エリアをはみ出す場合は処理を行わない
//         if (absoluteCanvas.right < e.clientX || absoluteCanvas.bottom < e.clientY) return

//         const newW = e.clientX - (absoluteCanvas.x + state.xpos)
//         const newH = e.clientY - (absoluteCanvas.y + state.ypos)
//         state.width = newW
//         state.height = newH
//         return
//       }
//       if (isClickedSticky) {
//         console.log('onMouseMoveSticky', e.offsetX, e.offsetY, e)
//         const absoluteCanvas = canvas.value?.getBoundingClientRect()
//         if (!absoluteCanvas) return
//         const absoluteNewX = e.clientX - pointerXOnSticky
//         const absoluteNewY = e.clientY - pointerYOnSticky
//         if (
//           absoluteNewX < absoluteCanvas.left
//           || absoluteNewY < absoluteCanvas.top
//           || absoluteCanvas.right < (absoluteNewX + state.width)
//           || absoluteCanvas.bottom < (absoluteNewY + state.height)
//         ) return
//         const newX = absoluteNewX - absoluteCanvas.x
//         const newY = absoluteNewY - absoluteCanvas.y
//         state.xpos = newX
//         state.ypos = newY
//         return
//       }
//     }
//     const onMouseUp = () => { isClickedScaleIcon = isClickedSticky = false }

//     return {
//       onMouseUp,
//       onMouseDownScaleIcon,
//       onMouseMove,
//       onMouseDownSticky,
//       state,
//       isClickedScaleIcon,
//     }




//   },
//   methods: {
//     test() {
//       this.xpos
//     }

//   }
// })
</script>

<template>
  <div ref="sticky" class="absolute flex left-2 bg-green-200 shadow-lg"
    :style="{ top: state.ypos + 'px', left: state.xpos + 'px', width: state.width + 'px', height: state.height + 'px' }">
    <textarea class="w-full h-full resize-none bg-transparent" @mousedown="onMouseDownSticky" />
    <span class="scale-icon absolute right-0 bottom-0 bg-black w-3 h-3 cursor-nwse-resize"
      @mousedown.stop="onMouseDownScaleIcon" />
  </div>
  {{ isClickedScaleIcon }}
</template>
