import { Injectable } from '@angular/core'
import { ActionSheetController } from '@ionic/angular'
import { Customer, Ext } from './customer'
import { pull } from 'lodash'
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CustomerService {
  static URL = ''
  // static URL = 'http://localhost:6181'
  customers: Customer[] = []
  groups: string[] = []

  constructor(
    private http: HttpClient,
    private actionSheetCtrl: ActionSheetController
  ) {
    this.loadGroup()
  }
  loadGroup() {
    this.http.get<string[]>(`${CustomerService.URL}/api/c/groups`)
      .subscribe(gs => this.groups = gs)
  }
  // 加载客户
  load(g: string) {
    this.http.get<Customer[]>(`${CustomerService.URL}/api/c/g/${g}`)
      .subscribe((cs: Customer[]) => {
        this.customers = cs
      })
  }
  // 上传定义
  ext(): Promise<any> {
    return new Promise<Ext[]>(resolve => {
      this.http.get<Ext[]>(`${CustomerService.URL}/api/c/ext`)
        .subscribe(es => {
          for (const e of es) {
            e.value = `${parseInt(`${e.value}`) + 1}`
          }
          resolve(es)
        })
    })
  }
  // 修改上传定义
  saveExt(es: Ext[]) {
    for (const e of es) {
      e.value = `${parseInt(`${e.value}`) - 1}`
    }
    return this.http.put<Ext[]>(`${CustomerService.URL}/api/c/ext`, es)
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
