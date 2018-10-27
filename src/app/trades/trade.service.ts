import { includes } from 'lodash'
import { HttpClient } from '@angular/common/http';
import { Injectable, NgZone } from '@angular/core';

import { Trade } from './trade';
import { ObjService } from '../api/obj.service';
import { CustomerService } from '../customers/customer.service';
import { ItemService } from '../items/item.service';

@Injectable({
  providedIn: 'root'
})
export class TradeService extends ObjService<Trade>{
  constructor(
    private itemService: ItemService,
    private customerService: CustomerService,
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
  }
  path() { return '/api/trades' }

  async days(now: string) {
    let days = await this.http.get<string[]>(this.url).toPromise()
    if (!days) { days = [] }
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
    this.itemService.itemsReset()
    return t
  }

  async delete(o: Trade) {
    await super.delete(o)
    if (o.cid) {
      const c = await this.customerService.obj$(o.cid).toPromise()
      this.customerService.reset(c)
    }
    this.itemService.itemsReset()
  }
}
