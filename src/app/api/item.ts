export interface Item {
  // 主键
  id: string
  // 名称
  name: string
  // 创建时间
  ca: Date
  // 扩展属性
  extend: any
  // 备注
  note?: string
}
