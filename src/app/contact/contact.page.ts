import { Component } from '@angular/core';
import { ModalController } from '@ionic/angular'
import { CustomerService } from '../api/customer.service';
import { Customer } from '../api/customer';
import { CustomerPage } from '../customer/customer.page';
import { UploadPage } from './upload/upload.page';

@Component({
  selector: 'app-contact',
  templateUrl: 'contact.page.html',
  styleUrls: ['contact.page.scss']
})
export class ContactPage {
  constructor(
    public customer: CustomerService,
    public modalCtrl: ModalController
  ) { }
  /**
   * 编辑客户
   * @param c 客户
   */
  async open(c: Customer) {
    const modal = await this.modalCtrl.create({
      component: CustomerPage,
      componentProps: { customer: Object.assign({}, c) },
    });
    modal.onDidDismiss(d => {
      if (d.data) {
        Object.assign(c, d.data)
      }
    })
    return await modal.present()
  }
  /**
   * 上传客户xslx
   */
  async upload() {
    const modal = await this.modalCtrl.create({
      component: UploadPage,
    });
    modal.onDidDismiss(d => {
      if (d.data) {
      }
    })
    return await modal.present()
  }
}
