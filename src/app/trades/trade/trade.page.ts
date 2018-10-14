import { pull } from 'lodash'
import { Observable } from 'rxjs';
import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Order } from '../order';
import { Trade } from '../trade';
import { ObjPage } from '../../api/obj.page';
import { TradeService } from '../trade.service';
import { Customer } from '../../customers/customer';
import { CustomerService } from '../../customers/customer.service';
import { SelectPage } from '../../items/select/select.page';
import { ItemService } from '../../items/item.service';

@Component({
  selector: 'app-trade',
  templateUrl: './trade.page.html',
  styleUrls: ['./trade.page.scss'],
})
export class TradePage extends ObjPage<Trade>{
  customer$: Observable<Customer>
  constructor(
    public tradeService: TradeService,
    private customerService: CustomerService,
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    super(modalCtrl, navParams)
    if (this.obj.cid) {
      this.customer$ = this.customerService.obj$(this.obj.cid)
    }
    if (!this.obj.orders) {
      this.obj.orders = []
    }
  }
  get service() { return this.tradeService }

  async selectCustomer() {
    const c = await this.customerService.doSelect()
    if (c.data) {
      this.obj.cid = c.data
      this.customer$ = this.customerService.obj$(c.data)
    }
  }

  async add() {
    const modal = await this.modalCtrl.create({
      component: SelectPage,
    })
    await modal.present()
    const r = await modal.onDidDismiss()
    if (r.data) {
      for (const i of r.data) {
        this.obj.orders.push({
          id: i.id,
          price: i.price,
          cost: i.cost,
          num: 1,
        })
      }
    }
  }

  get price() {
    let price = 0
    for (const o of this.obj.orders) {
      price += (o.price * o.num)
    }
    return price
  }

  get cost() {
    let cost = 0
    for (const o of this.obj.orders) {
      cost += ((o.cost ? o.cost : 0) * o.num)
    }
    this.obj.cost = cost
    return cost
  }

  puls(o: Order) {
    o.num += 1
  }

  remove(o: Order) {
    o.num -= 1
    if (o.num <= 0) {
      pull(this.obj.orders, o)
    }
  }

  get tradeTotal() {
    return this.obj.total / 100
  }

  set tradeTotal(total: number) {
    this.obj.total = Math.ceil(total * 100)
  }
}
