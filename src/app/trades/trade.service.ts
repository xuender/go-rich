import { Injectable, NgZone } from '@angular/core';
import { Trade } from './trade';
import { ObjService } from '../api/obj.service';
import { HttpClient } from '@angular/common/http';
import { DatePipe } from '@angular/common';
import { includes } from 'lodash'
import { ModalController } from '@ionic/angular';
import { TradePage } from './trade/trade.page';
import { CustomerService } from '../customers/customer.service';

@Injectable({
  providedIn: 'root'
})
export class TradeService extends ObjService<Trade>{
  constructor(
    private customerService: CustomerService,
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
  }
  path() { return '/api/trades' }

  async days(now: string) {
    const days = await this.http.get<string[]>(this.url).toPromise()
    if (!includes(days, now)) {
      days.push(now)
    }
    return days.sort()
  }

  getTrades(day: string) {
    return this.http.get<Trade[]>(`${this.url}?day=${day}`)
  }

  getTrade$(id: string) {
    return this.http.get<Trade>(`${this.url}/${id}`)
  }

  async save(o: Trade) {
    const t = await super.save(o)
    if (t.cid) {
      const c = await this.customerService.obj$(t.cid).toPromise()
      this.customerService.reset(c)
    }
    if (o.cid && o.cid != t.cid) {
      const c = await this.customerService.obj$(o.cid).toPromise()
      this.customerService.reset(c)
    }
    return t
  }

  async delete(o: Trade) {
    await super.delete(o)
    if (o.cid) {
      const c = await this.customerService.obj$(o.cid).toPromise()
      this.customerService.reset(c)
    }
  }
}
