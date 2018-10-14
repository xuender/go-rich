import { Obj } from "../api/obj";
import { Order } from "./order";

export interface Trade extends Obj {
  // 用户ID
  cid?: string
  // 状态
  status: number
  // 付款时间
  pa: Date
  // 总额， 单位分
  total: number
  // 成本
  cost?: number
  // 订单详情
  orders: Order[]
}
