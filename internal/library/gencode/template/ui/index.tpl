<template>
  <div class="ele-body">
    <a-card :bordered="false">
      <!-- 搜索表单 -->
      <${.table.short_name|CaseCamel}Search :where="defaultWhere" @search="reload" />
      <!-- 表格 -->
      <ele-pro-table ref="tableRef" row-key="id" :columns="columns" :datasource="datasource"
        :scroll="{ x: 800 }" cache-key="pro${.table.name|CaseCamel}SearchTable" resizable :where="defaultWhere">
        <template #toolbar>
          <a-space>
            <a-button type="primary" class="ele-btn-icon" @click="openEdit()">
              <template #icon>
                <plus-outlined />
              </template>
              <span>新建</span>
            </a-button>
          </a-space>
        </template>
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'action'">
            <a-space>
              <a @click="openEdit(record)">修改</a>
              <a-divider type="vertical" />
              <a-popconfirm placement="topRight" title="确定要删除此${.table.comment}吗？" @confirm="remove(record)">
                <a class="ele-text-danger">删除</a>
              </a-popconfirm>
            </a-space>
          </template>
        </template>
      </ele-pro-table>
    </a-card>
    <!-- 编辑弹窗 -->
    <${.table.short_name|CaseCamel}Edit v-model:visible="showEdit" :data="current" @done="reload" />
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { message, Modal } from 'ant-design-vue/es';
import { PlusOutlined } from '@ant-design/icons-vue';
import { messageLoading, toDateString } from 'ele-admin-pro/es';
import ${.table.short_name|CaseCamel}Search from './components/${.table.short_name|CaseKebab}-search.vue';
import ${.table.short_name|CaseCamel}Edit from './components/${.table.short_name|CaseKebab}-edit.vue';
import { page${.table.short_name|CaseCamel}, remove${.table.short_name|CaseCamel} } from '@/api/${.table.short_name|CaseKebab}';

// 表格实例
const tableRef = ref(null);

// 默认搜索条件
const defaultWhere = reactive({
  sort: 'id',
  order: 'desc'
});

// 表格列配置
const columns = ref([
   ${- range $index, $elem := .table.fields}
      {
        title: '${$elem.comment_short}',
        dataIndex: '${$elem.name}',
        width: 90,
        align: 'left',
        ellipsis: true,
        resizable: true
      },
   ${- end}
  {
    title: '创建时间',
    dataIndex: 'created_at',
    width:140,
    sorter: true,
    showSorterTooltip: false,
    ellipsis: true,
    customRender: ({ text }) => toDateString(text, 'YYYY-MM-DD HH:mm')
  },
  {
    title: '操作',
    key: 'action',
    width: 120,
    align: 'center'
  }
]);

// 当前编辑数据
const current = ref(null);

// 是否显示编辑弹窗
const showEdit = ref(false);

// 表格数据源
const datasource = ({ page, limit, where, orders }) => {
  return page${.table.short_name|CaseCamel}({ ...where, ...orders, page, limit });
};

/* 搜索 */
const reload = (where) => {
  tableRef?.value?.reload({ page: 1, where });
};

/* 打开编辑弹窗 */
const openEdit = (row) => {
  current.value = row ?? null;
  showEdit.value = true;
};

/* 删除单个 */
const remove = (row) => {
  const hide = messageLoading('请求中..', 0);
  remove${.table.short_name|CaseCamel}(row.id)
    .then((msg) => {
      hide();
      message.success(msg);
      reload();
    })
    .catch((e) => {
      hide();
      message.error(e.message);
    });
};
</script>

<script>
export default {
  name: '${.table.name|CaseCamel}'
};
</script>
