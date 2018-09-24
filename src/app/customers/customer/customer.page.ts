import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Customer } from '../customer';
import { ObjPage } from '../../api/obj.page';
import { CustomerService } from '../customer.service';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.page.html',
  styleUrls: ['./customer.page.scss'],
})
export class CustomerPage extends ObjPage<Customer>{
  constructor(
    private customerService: CustomerService,
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    super(modalCtrl, navParams)
    if (!this.obj.extend) { this.obj.extend = {} }
    if (!this.obj.tags) { this.obj.tags = [] }
  }

  get service() { return this.customerService }
}
