import { Obj } from "../api/obj";

// 扩展数据
export interface User extends Obj {
  phone?: string
  extend: any // 扩展属性
}
