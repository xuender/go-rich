import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Customer } from '../customer';
import { ObjPage } from '../../api/obj.page';
import { CustomerService } from '../customer.service';
import { TradeService } from '../../trades/trade.service';
import { TradePage } from '../../trades/trade/trade.page';
import { Observable } from 'rxjs';
import { Trade } from '../../trades/trade';

@Component({
  selector: 'app-customer',
  templateUrl: './customer.page.html',
  styleUrls: ['./customer.page.scss'],
})
export class CustomerPage extends ObjPage<Customer>{
  trades: Observable<Trade>[] = []
  constructor(
    private customerService: CustomerService,
    private tradeService: TradeService,
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    super(modalCtrl, navParams)
    if (!this.obj.extend) { this.obj.extend = {} }
    if (!this.obj.tags) { this.obj.tags = [] }
    if (!this.obj.trades) {
      this.obj.trades = []
    }
    for (const id of this.obj.trades) {
      this.trades.push(this.tradeService.getTrade$(id))
    }
  }

  get service() { return this.customerService }

  async trade(trade = { cid: this.obj.id }) {
    const isOld = 'id' in trade
    const modal = await this.modalCtrl.create({
      component: TradePage,
      componentProps: { obj: trade }
    })
    await modal.present()
    const r = await modal.onDidDismiss()
    if (r.data) {
      if (isOld) {
        Object.assign(trade, r.data)
      } else {
        this.trades.push(this.tradeService.getTrade$(r.data.id))
      }
    }
  }
}
