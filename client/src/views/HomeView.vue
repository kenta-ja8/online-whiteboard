<script setup lang="ts">
import { computed, onMounted, reactive, ref, ssrContextKey } from 'vue'
import { UserApi, type GetUserListOutput } from "@/gen/api"
import StickyNote from "@/components/stickyNote.vue"


const state = reactive({
  userList: [] as Array<Exclude<GetUserListOutput['user_list'], undefined>[number]>,
  xpos: 0,
  ypos: 0,
  width: 0,
  height: 0,
})
const load = async () => {
  state.userList = (await (new UserApi()).getUserList()).data.user_list || []
}
load()


state.xpos = 100
state.ypos = 100
state.width = 100
state.height = 100

let pointerXOnSticky = 0
let pointerYOnSticky = 0

const canvas = ref<HTMLDivElement>()
const oldcanvas = ref<HTMLDivElement>()
let isClickedSticky = false
const onMouseDownSticky = (e: any) => {
  console.log('onMouseDownSticky')
  isClickedSticky = true
  pointerXOnSticky = e.offsetX
  pointerYOnSticky = e.offsetY
}


let isClickedScaleIcon = false
const onMouseDownScaleIcon = () => { isClickedScaleIcon = true }
const onMouseMove = (e: any) => {
  if (isClickedScaleIcon) {
    console.log('onMouseMoveScaleIcon', e)
    const absoluteCanvas = canvas.value?.getBoundingClientRect()
    console.log('return vue', canvas)
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


const selectedStickyNote = {
  onMouseMove: () => { },
  onMouseUp: () => { },
}
const resister = (mouseEvent: typeof selectedStickyNote) => {
  console.log('resister')
  selectedStickyNote.onMouseMove = mouseEvent.onMouseMove
  selectedStickyNote.onMouseUp = mouseEvent.onMouseUp
}


</script>

<template>
  <div>
    <div class="mb-3 xl:w-96">
      <select
        class="form-select form-select-lg mb-3 appearance-none block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none">
        <option value></option>
        <template v-for="user of state.userList">
          <option :value="user.id">{{ user.last_name }}</option>
        </template>
      </select>
    </div>
    {{ state.xpos }}
    {{ state.ypos }}
    <div>
      <div class="canvas relative w-[1000px] h-[600px] bg-white border-2" ref="oldcanvas" @mousemove="onMouseMove"
        @mouseup="onMouseUp">
        <div ref="sticky" class="flex left-2 bg-green-200 shadow-lg"
          :style="{ top: state.ypos + 'px', left: state.xpos + 'px', width: state.width + 'px', height: state.height + 'px' }">
          <textarea class="w-full h-full resize-none bg-transparent" @mousedown="onMouseDownSticky" />
          <span class="scale-icon absolute right-0 bottom-0 bg-black w-3 h-3 cursor-nwse-resize"
            @mousedown.stop="onMouseDownScaleIcon" />
        </div>
      </div>
    </div>
    <textarea />
    <div>
      <div class="canvas relative w-[1000px] h-[600px] bg-white border-2" ref="canvas"
        @mousemove="selectedStickyNote.onMouseMove" @mouseup="selectedStickyNote.onMouseUp">
        <StickyNote @selectStickyNote="resister" :xpos="2" :ypos="2" :init-width="3" :init-height="3"
          :canvas="canvas" />
        <StickyNote @selectStickyNote="resister" :xpos="2" :ypos="2" :init-width="3" :init-height="3"
          :canvas="canvas" />
      </div>
    </div>
  </div>
</template>
