import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Ext } from './ext';
import { URL } from './init'
import { ModalController } from '@ionic/angular';
import { ExtPage } from '../setting/ext/ext.page';

@Injectable({
  providedIn: 'root'
})
export class ExtService {

  customerExts: Ext[] = []
  itemExts: Ext[] = []
  constructor(
    private modalCtrl: ModalController,
    private http: HttpClient
  ) {
    this.load()
  }
  async getCustomerExts() {
    this.customerExts = await this.http.get<Ext[]>(`${URL}/api/exts/E-C`).toPromise()
    return this.customerExts
  }
  /**
   * 修改
   */
  putCustomerExts(exts: Ext[]) {
    this.http.put<Ext[]>(`${URL}/api/exts/E-C`, exts)
      .subscribe((s: Ext[]) => {
        this.getCustomerExts()
      })
  }
  async load() {
    this.customerExts = await this.http.get<Ext[]>(`${URL}/api/exts/E-C`).toPromise()
    this.itemExts = await this.http.get<Ext[]>(`${URL}/api/exts/E-I`).toPromise()
    return true
  }
  /**
   * 客户扩展设置
   */
  async extCustomer() {
    await this.getCustomerExts()
    const modal = await this.modalCtrl.create({
      component: ExtPage,
      componentProps: { exts: Object.assign([], this.customerExts) }
    })
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.putCustomerExts(d.data)
      }
    })
    return await modal.present()
  }
}
