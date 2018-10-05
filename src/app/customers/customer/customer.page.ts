import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Customer } from '../customer';
import { ObjPage } from '../../api/obj.page';
import { CustomerService } from '../customer.service';
import { TradeService } from '../../trades/trade.service';
import { TradePage } from '../../trades/trade/trade.page';

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

  async trade() {
    const modal = await this.modalCtrl.create({
      component: TradePage,
      componentProps: { obj: { cid: this.obj.id } }
    })
    return await modal.present()
  }
}
