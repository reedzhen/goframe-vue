<!-- 性别选择下拉框 -->
<template>
  <a-select
    show-search
    optionFilterProp="label"
    :options="data"
    allow-clear
    :value="value"
    :placeholder="placeholder"
    @update:value="updateValue"
    @blur="onBlur"
  />
</template>

<script setup>
  import { ref } from 'vue';
  import { message } from 'ant-design-vue';
  import { listDictionaryData } from '@/api/system/dictionary-data';
  import type { SelectProps } from 'ant-design-vue';

  const emit = defineEmits<{
    (e: 'update:value', value: number): void;
    (e: 'blur'): void;
  }>();

  withDefaults(
    defineProps<{
      value?: number;
      placeholder?: string;
    }>(),
    {
      placeholder: '请选择性别'
    }
  );

  // 字典数据
  const data = ref<SelectProps['options']>([]);

  /* 更新选中数据 */
  const updateValue = (value: number) => {
    emit('update:value', value);
  };

  /* 获取字典数据 */
  listDictionaryData({ dict_code: 'sex' })
    .then((list) => {
      data.value = list.map((d) => {
        return {
          value: parseInt(d.code ?? ''),
          label: d.name
        };
      });
    })
    .catch((e) => {
      message.error(e.message);
    });

  /* 失去焦点 */
  const onBlur = () => {
    emit('blur');
  };
</script>
