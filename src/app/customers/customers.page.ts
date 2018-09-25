import { Component } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular'

import { Customer } from './customer'
import { TagService } from '../tags/tag.service';
import { UploadPage } from './upload/upload.page';
import { ObjsPlusPage } from '../api/objs.puls.page';
import { CustomerService } from './customer.service';
import { CustomerPage } from './customer/customer.page';

@Component({
  selector: 'app-customers',
  templateUrl: 'customers.page.html',
  styleUrls: ['customers.page.scss']
})
export class CustomersPage extends ObjsPlusPage<Customer>{
  constructor(
    public customerService: CustomerService,
    protected tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(tagService, modalCtrl, actionSheetCtrl)
  }

  get tagKey() { return 'tag-C' };
  get service() { return this.customerService }
  get page() { return CustomerPage }
  get newObj() { return { name: '', extend: {}, tags: [] } }
  get title() { return '客户' }

  // 上传客户xslx
  async upload() {
    const modal = await this.modalCtrl.create({
      component: UploadPage,
    });
    await modal.present()
    await modal.onDidDismiss()
    await this.reset()
  }
}
