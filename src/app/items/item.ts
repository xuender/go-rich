import { Obj } from "../api/obj";

export interface Item extends Obj {
  // 价格
  price: number
  // 创建时间
  extend: any
  // 标签
  tags: string[]
  // 成本
  cost?: number
  // 批次
  batchs?: Array<Batch>
}

export interface Batch {
  // 成本
  cost: number
  costMoney?:number
  // 总数
  total: number
  // 库存
  inventory: number
}
