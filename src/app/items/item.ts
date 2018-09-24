import { Obj } from "../api/obj";

export interface Item extends Obj {
  // 价格
  price: number
  // 创建时间
  extend: any
  // 标签
  tags: string[]
}
