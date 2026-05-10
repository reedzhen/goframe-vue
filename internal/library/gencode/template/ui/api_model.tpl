import type { PageParam } from '@/api';

/**
 * ${.table.comment}
 */
export interface ${.table.short_name|CaseCamel} {
  // ${.table.comment}id
  id?: number;
${- range $index, $elem := .table.fields}
  // ${$elem.comment}
  ${$elem.name}?: ${$elem.type_js};
${- end}
}

/**
 * ${.table.comment}搜索条件
 */
export interface ${.table.short_name|CaseCamel}Param extends PageParam {
${- range $index, $elem := .table.fields}${ if eq $elem.need_search "true"}
  // ${$elem.comment}
  ${$elem.name}?: ${$elem.type_js};
 ${end}${- end}
}