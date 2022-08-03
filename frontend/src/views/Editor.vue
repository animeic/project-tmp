<template>
  <md-editor-v3 v-model="data.text">
    <template #defToolbars>
      <modal-toolbar
        :visible="data.modalVisible"
        :is-fullscreen="data.modalFullscreen"
        show-adjust
        title="帮助"
        modal-title="编辑预览"
        width="870px"
        height="600px"
        @on-click="data.modalVisible = true"
        @on-close="data.modalVisible = false"
        @on-adjust="data.modalFullscreen = !data.modalFullscreen"
      >
        <span>内容</span>
        <template #trigger>
          <svg class="md-icon" aria-hidden="true">
            <use xlink:href="#icon-read"></use>
          </svg>
        </template>
      </modal-toolbar>
    </template>
  </md-editor-v3>
</template>

<script setup>
import { reactive } from 'vue';

const data = reactive({
  text: '',
  modalVisible: false,
  modalFullscreen: false
});
</script>