<template>
  <el-pagination
    :background="background"
    v-model:current-page="innerPageNo"
    v-model:page-size="innerPageSize"
    :layout="layout"
    :page-sizes="pageSizes"
    v-model:total="total"
    v-bind="$attrs"
    class="mt-10px float-right"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  />
</template>

<script lang="ts" setup>
import { computed } from 'vue';


const props = defineProps(
  {
    background: {
      type: Boolean,
      default: true,
    },
    layout: {
      type: String,
    //   default: 'total, sizes, prev, pager, next',
      default: 'total,prev,next, pager,jumper',
    },
    pageSizes: {
      type: Array,
      default: () => [5, 10],
    },
    total: {
      type: Number,
      default: 10,
    },
    pageSize: {
      type: Number,
      default: 5,
    },
    pageNo: {
      type: Number,
      default: 1,
    },
  }
)

const emits = defineEmits(['change'])
const innerPageSize = computed({
  get() {
    return props.pageSize
  },
  set(val) {
    emits('change', {
      pageSize: val,
      pageNo: props.pageNo,
    })
  },
})

const innerPageNo = computed({
  get() {
    return props.pageNo
  },
  set(val) {
    emits('change', {
      pageSize: props.pageSize,
      pageNo: val,
    })
  },
})

const handleCurrentChange = (val: number) => {
  innerPageNo.value = val
}

const handleSizeChange = (val: number) => {
  innerPageSize.value = val
}

</script>