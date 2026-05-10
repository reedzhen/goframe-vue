<!-- ${.table.comment}编辑弹窗 -->
<template>
  <ele-modal :width="460" :visible="visible" :title="isUpdate ? '修改${.table.comment}' : '新建${.table.comment}'" :body-style="{ paddingBottom: '8px' }"
      style="top: 20px" :maskClosable="false">
      <!-- <a-drawer :title="isUpdate ? '修改客户' : '新建客户'" width="700" :visible="visible" :body-style="{ padding: '20px' }" @update:visible="updateVisible"> -->
      <template #footer>
        <a-space>
          <a-button key="back" @click="updateVisible(false)">取消</a-button>
          <a-button key="submit" type="primary" :loading="loading" @click="save">确定</a-button>
        </a-space>
      </template>
     <a-form ref="formRef" :model="form" :rules="rules" :label-col="{ flex: '90px' }">
          ${- range $index, $elem := .table.fields}
          <a-form-item label="${$elem.comment_short}" name="${$elem.name}">
            ${if eq $elem.type_js "number"}
            <a-input-number :min="0" :precision="0" class="ele-fluid" placeholder="请输入${$elem.comment_short}" v-model:value="form.${$elem.name}" />
            ${end}
            ${if eq $elem.type_js "string"}
            <a-input allow-clear :maxlength="100" placeholder="请输入${$elem.comment_short}" v-model:value="form.${$elem.name}" />
            ${end}
            ${if eq $elem.type_js "datetime"}
            <a-date-picker class="ele-fluid" valueFormat="YYYY-MM-DD" v-model:value="form.${$elem.name}" />
            ${end}
          </a-form-item>
          ${- end}
     </a-form>
     <!-- </a-drawer> -->
  </ele-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { message } from 'ant-design-vue/es';
import useFormData from '@/utils/use-form-data';
import { create${.table.short_name|CaseCamel}, update${.table.short_name|CaseCamel} } from '@/api/${.table.short_name|CaseKebab}';

const emit = defineEmits(['done', 'update:visible']);

const props = defineProps({
  // 弹窗是否打开
  visible: Boolean,
  // 修改回显的数据
  data: Object
});

//
const formRef = ref(null);

// 是否是修改
const isUpdate = ref(false);

// 提交状态
const loading = ref(false);

// 表单数据
const { form, resetFields, assignFields } = useFormData({
  id: undefined,
  ${- range $index, $elem := .table.fields}
  ${$elem.name}: '',
  ${- end}
});

// 表单验证规则
const rules = reactive({
   ${- range $index, $elem := .table.fields}
    ${if eq $elem.type_js "number"}
      ${$elem.name}: [{ required: true, message: '请输入${$elem.comment_short}', type: 'number' }],
    ${else}
      ${$elem.name}: [{ required: true, message: '请输入${$elem.comment_short}', type: 'string' }],
    ${end}
   ${- end}
});

/* 保存编辑 */
const save = () => {
  if (!formRef.value) {
    return;
  }
  formRef.value
    .validate()
    .then(() => {
      loading.value = true;
      const saveOrUpdate = isUpdate.value ? update${.table.short_name|CaseCamel} : create${.table.short_name|CaseCamel};
      saveOrUpdate(form)
        .then((msg) => {
          loading.value = false;
          message.success(msg);
          updateVisible(false);
          emit('done');
        })
        .catch((e) => {
          loading.value = false;
          message.error(e.message);
        });
    })
    .catch(() => { });
};

/* 更新visible */
const updateVisible = (value) => {
  emit('update:visible', value);
};

watch(
  () => props.visible,
  (visible) => {
    if (visible) {
      if (props.data) {
        assignFields(props.data);
        isUpdate.value = true;
      } else {
        isUpdate.value = false;
      }
    } else {
      resetFields();
      formRef.value?.clearValidate();
    }
  }
);
</script>
