import { Injectable } from '@angular/core'
import { ActionSheetController } from '@ionic/angular'
import { Customer } from './customer'
import { pull } from 'lodash'
import { HttpClient } from '@angular/common/http';
import { URL } from './init'
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CustomerService {
  customers: Customer[] = []

  groups: Observable<string[]>
  constructor(
    private http: HttpClient,
    private actionSheetCtrl: ActionSheetController
  ) {
    this.groups = this.http.get<string[]>(`${URL}/api/groups`)
  }
  // 加载客户
  load(g: string) {
    this.http.get<Customer[]>(`${URL}/api/groups/${g}`)
      .subscribe((cs: Customer[]) => {
        this.customers = cs
      })
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
          this.http.delete(`${URL}/api/customers/${c.id}`)
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
  // 创建客户
  post(c: Customer) {
    if (c) {
      this.http.post<Customer>(`${URL}/api/customers`, c)
        .subscribe((nc: Customer) => this.customers.push(nc))
    }
  }
  // 修改客户
  put(c: Customer) {
    if (c && c.id) {
      this.http.put<Customer>(`${URL}/api/customers/${c.id}`, c)
        .subscribe((nc: Customer) => Object.assign(c, nc))
    }
  }
}
