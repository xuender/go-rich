import { Ext } from "./ext";

// 扩展数据
export interface User {
  id?: string
  nick: string
  phone?: string
  note?: string
  ca?: Date
  extend: any // 扩展属性
}
