export interface Customer {
  id: string                      // 主键
  name: string                    // 姓名
  pinyin: string                  // 拼音
  extend: any // 扩展属性
  phone?: string                   // 电话
  ca: Date                        // 创建时间
  trades?: string[]                // 消费记录
  note?: string                    // 备注
}
