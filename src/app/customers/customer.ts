import { Obj } from "../api/obj";

export interface Customer extends Obj {
  // 扩展属性
  extend: any
  // 电话
  phone?: string
  // 消费记录
  trades?: string[]
  // 标签
  tags: string[]
}
