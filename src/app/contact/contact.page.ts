import { Component } from '@angular/core';
import { ModalController } from '@ionic/angular'
import { CustomerService } from '../api/customer.service';
import { Customer } from '../api/customer';
import { CustomerPage } from './customer/customer.page';
import { UploadPage } from './upload/upload.page';
import { GroupPage } from './group/group.page';

@Component({
  selector: 'app-contact',
  templateUrl: 'contact.page.html',
  styleUrls: ['contact.page.scss']
})
export class ContactPage {
  title: string = ''
  constructor(
    public customer: CustomerService,
    public modalCtrl: ModalController
  ) { }
  /**
   * 编辑客户
   * @param c 客户
   */
  async update(c: Customer) {
    const modal = await this.modalCtrl.create({
      component: CustomerPage,
      componentProps: { customer: Object.assign({}, c) },
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.customer.put(d.data)
        Object.assign(c, d.data)
      }
    })
    return await modal.present()
  }
  async create() {
    const modal = await this.modalCtrl.create({
      component: CustomerPage,
      componentProps: {
        customer: {
          name: '',
          extend: {}
        }
      }
    });
    modal.onDidDismiss().then(d => this.customer.post(d.data))
    return await modal.present()
  }
  /**
   * 上传客户xslx
   */
  async upload() {
    const modal = await this.modalCtrl.create({
      component: UploadPage,
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
      }
    })
    return await modal.present()
  }
  /**
   * 分组
   */
  async group() {
    const modal = await this.modalCtrl.create({
      component: GroupPage,
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.title = d.data
        this.customer.load(d.data)
      }
    })
    return await modal.present()
  }
}
