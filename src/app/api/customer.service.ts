import { Injectable } from '@angular/core'
import { ActionSheetController } from '@ionic/angular'
import { Customer, Divider } from './customer'
import { sortBy, chain, invert, isNumber, pull, remove } from 'lodash-es'
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CustomerService {
  // static URL = 'http://localhost:6181'
  static URL = ''
  customers: Customer[] = []
  dividers: Divider[] = []

  constructor(
    private http: HttpClient,
    private actionSheetCtrl: ActionSheetController
  ) {
    this.load()
  }
  // 加载客户
  load() {
    this.http.get<Customer[]>(`${CustomerService.URL}/api/c/all`)
      .subscribe((cs: Customer[]) => {
        this.customers = cs
        this.group()
      })
  }
  // 客户分组
  group() {
    this.dividers = chain(this.customers)
      .groupBy((c: Customer) => c.pinyin ? c.pinyin[0] : '')
      .map((v, k) => {
        return {
          divide: k,
          customers: sortBy(v, 'pinyin'),
        }
      })
      .sortBy('divide')
      .value()
  }
  // 上传定义
  promap(): Promise<any> {
    return new Promise<any>(resolve => {
      this.http.get<any>(`${CustomerService.URL}/api/c/promap`)
        .subscribe(a => {
          const pro = invert(a)
          for (let k in pro) {
            if (!isNumber(pro[k])) {
              pro[k] = parseInt(pro[k])
            }
            pro[k] += 1
          }
          resolve(pro)
        })
    })
  }
  // 修改上传定义
  savePromap(pro: any) {
    for (let k in pro) {
      pro[k] -= 1
    }
    return this.http.post<any>(`${CustomerService.URL}/api/c/promap`, invert(pro))
  }
  // 删除客户
  async del(c: Customer) {
    const actionSheet = await this.actionSheetCtrl.create({
      header: `确认删除客户 [ ${c.name} ] ?`,
      buttons: [{
        text: '删除',
        role: 'destructive',
        icon: 'trash',
        handler: () => {
          this.http.delete(`${CustomerService.URL}/api/c/${c.id}`)
            .subscribe(r => {
              pull(this.customers, c)
              for (const d of this.dividers) {
                pull(d.customers, c)
              }
              remove(this.dividers, d => d.customers.length === 0)
            })
        }
      }, {
        text: '取消',
        icon: 'close',
        role: 'cancel',
      }]
    });
    await actionSheet.present();
  }
}
