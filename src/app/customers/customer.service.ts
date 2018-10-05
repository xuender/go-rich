import { Injectable, NgZone } from '@angular/core'
import { HttpClient } from '@angular/common/http';

import { URL } from '../api/init'
import { Customer } from './customer'
import { ObjService } from '../api/obj.service';
import { ModalController } from '@ionic/angular';
import { SelectPage } from './select/select.page';

@Injectable({
  providedIn: 'root'
})
export class CustomerService extends ObjService<Customer> {
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
    protected modalCtrl: ModalController,
  ) {
    super(http, ngZone)
  }

  path() { return '/api/customers' }

  download(params: any) {
    return this.http.get(`${URL}/api/customers`, { params: params, responseType: 'blob', })
  }

  async doSelect() {
    const modal = await this.modalCtrl.create({
      component: SelectPage,
      componentProps: { service: this }
    })
    await modal.present()
    return await modal.onDidDismiss()
  }
}
