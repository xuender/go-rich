import { Injectable, NgZone } from '@angular/core';
import { Trade } from './trade';
import { ObjService } from '../api/obj.service';
import { HttpClient } from '@angular/common/http';
import { DatePipe } from '@angular/common';
import { includes } from 'lodash'
import { ModalController } from '@ionic/angular';
import { TradePage } from './trade/trade.page';

@Injectable({
  providedIn: 'root'
})
export class TradeService extends ObjService<Trade>{
  constructor(
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
    return this.http.get<Trade[]>(`${this.url}/${day}`)
  }
}
