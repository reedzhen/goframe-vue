import request from '@/utils/request';

/**
 * ${.table.comment}分页
 */
export async function page${.table.short_name|CaseCamel}(params) {
  const res = await request.get('${if ne .table.addon_name ""}/${.table.addon_name}${end}/${.table.short_name|CaseKebab}/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * ${.table.comment}列表
 */
export async function list${.table.short_name|CaseCamel}(params) {
  const res = await request.get('${if ne .table.addon_name ""}/${.table.addon_name}${end}/${.table.short_name|CaseKebab}/list', { params });
  if (res.data.code === 0) {
    return res.data.data || [];
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * 新建${.table.comment}
 */
export async function create${.table.short_name|CaseCamel}(data) {
  const res = await request.post('${if ne .table.addon_name ""}/${.table.addon_name}${end}/${.table.short_name|CaseKebab}/create', data);
  if (res.data.code === 0) {
    return res.data.message || '操作成功';
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * 修改${.table.comment}
 */
export async function update${.table.short_name|CaseCamel}(data) {
  const res = await request.post('${if ne .table.addon_name ""}/${.table.addon_name}${end}/${.table.short_name|CaseKebab}/update', data);
  if (res.data.code === 0) {
    return res.data.message || '操作成功';
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * 删除${.table.comment}
 */
export async function remove${.table.short_name|CaseCamel}(id) {
  const res = await request.post('${if ne .table.addon_name ""}/${.table.addon_name}${end}/${.table.short_name|CaseKebab}/delete/' + id);
  if (res.data.code === 0) {
    return res.data.message || '操作成功';
  }
  return Promise.reject(new Error(res.data.message));
}
