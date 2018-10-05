import { Component, OnInit } from '@angular/core';
import { ModalController, ActionSheetController, NavParams } from '@ionic/angular';
import { ObjsPlusPage } from '../../api/objs.puls.page';
import { Customer } from '../customer';
import { CustomerService } from '../customer.service';
import { TagService } from '../../tags/tag.service';
import { CustomerPage } from '../customer/customer.page';

@Component({
  selector: 'app-select',
  templateUrl: './select.page.html',
  styleUrls: ['./select.page.scss'],
})
export class SelectPage extends ObjsPlusPage<Customer>{
  customerService: CustomerService
  cid: string = ''
  constructor(
    protected tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
    private navParams: NavParams
  ) {
    super(tagService, modalCtrl, actionSheetCtrl)
    this.customerService = this.navParams.get('service')
  }

  get tagKey() { return 'tag-C' };
  get service() { return this.customerService }
  set service(s) { }
  get page() { return {} }
  get newObj() { return {} }
  get title() { return '客户' }

  cancel() {
    this.modalCtrl.dismiss();
  }
  ok() {
    this.modalCtrl.dismiss(this.cid)
  }
}
