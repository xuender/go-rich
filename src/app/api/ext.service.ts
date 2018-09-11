import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Ext } from './ext';
import { URL } from './init'
import { ModalController } from '@ionic/angular';
import { ExtPage } from '../setting/ext/ext.page';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ExtService {

  extsByCustomer: Observable<Ext[]>
  extsByItem: Observable<Ext[]>
  constructor(
    private modalCtrl: ModalController,
    private http: HttpClient
  ) {
    this.extsByCustomer = this.http.get<Ext[]>(`${URL}/api/exts/E-C`)
    this.extsByItem = this.http.get<Ext[]>(`${URL}/api/exts/E-I`)
  }
  /**
   * 客户扩展设置
   */
  async extCustomer() {
    const exts = await this.extsByCustomer.toPromise()
    const modal = await this.modalCtrl.create({
      component: ExtPage,
      componentProps: {
        exts: exts,
      },
    })
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.http.put<Ext[]>(`${URL}/api/exts/E-C`, d.data).subscribe()
      }
    })
    return await modal.present()
  }
}
