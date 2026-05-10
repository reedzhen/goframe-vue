<!-- 搜索表单 -->
<template>
  <a-form :label-col="{ flex: '90px' }" :wrapper-col="{ flex: '1' }" labelAlign="left">
    <a-row :gutter="8">
    ${- range $index, $elem := .table.fields}
      ${ if eq $elem.need_search "true"}
         <a-col :span="6">
           <a-form-item label="${$elem.comment_short}">
             <a-input v-model:value.trim="form.${$elem.name}" placeholder="请输入${$elem.comment_short}" allow-clear />
           </a-form-item>
         </a-col>
      ${end}
    ${- end}
      <a-col :span="6">
        <a-form-item class="ele-text-right" :wrapper-col="{ span: 24 }">
          <a-space>
            <a-button type="primary" @click="search">查询</a-button>
            <a-button @click="reset">重置</a-button>
          </a-space>
        </a-form-item>
      </a-col>
    </a-row>
  </a-form>
</template>

<script setup>
  import useFormData from '@/utils/use-form-data';

  const props = defineProps({
    // 默认搜索条件
    where: Object
  });

  const emit = defineEmits(['search']);

  // 表单数据
  const { form, resetFields } = useFormData({
     ${- range $index, $elem := .table.fields}
          ${ if eq $elem.need_search "true"}
            ${$elem.name}: '',
          ${end}
     ${- end}
     ...props.where
  });

  /* 搜索 */
  const search = () => {
    emit('search', form);
  };

  /*  重置 */
  const reset = () => {
    resetFields();
    search();
  };
</script>
