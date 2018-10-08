import { Observable } from 'rxjs';
import { Component } from '@angular/core';
import { DatePipe } from '@angular/common';
import { indexOf, lastIndexOf } from 'lodash'
import { ModalController, ActionSheetController } from '@ionic/angular';

import { Trade } from './trade';
import { ObjsPage } from '../api/objs.page';
import { TradePage } from './trade/trade.page';
import { TradeService } from './trade.service';

@Component({
  selector: 'app-trades',
  templateUrl: './trades.page.html',
  styleUrls: ['./trades.page.scss'],
})
export class TradesPage extends ObjsPage<Trade>{
  trades$: Observable<Trade[]>
  day: string
  days: string[] = []
  now: string
  constructor(
    public tradeService: TradeService,
    private datePipe: DatePipe,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
    this.now = this.datePipe.transform(new Date(), 'yyyy-MM-dd')
    this.day = this.now
    this.trades$ = this.tradeService.getTrades(this.day)
    this.tradeService.days(this.day)
      .then(days => this.days.push(...days))
  }
  get service() { return this.tradeService }
  get page() { return TradePage }
  get newObj() { return {} }
  get title() { return '订单' }

  get next() {
    const i = lastIndexOf(this.days, this.day)
    if (i >= 0 && i < this.days.length - 1) {
      return this.days[i + 1]
    }
    return ''
  }

  get previous() {
    const i = indexOf(this.days, this.day)
    if (i > 0 && i < this.days.length) {
      return this.days[i - 1]
    }
    return ''
  }

  sum(trades: Trade[]) {
    let total = 0
    for (const t of trades) {
      total += t.total
    }
    return total
  }

  goto(day: string) {
    this.day = day
    this.trades$ = this.tradeService.getTrades(this.day)
  }
}
