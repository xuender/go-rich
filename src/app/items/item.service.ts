import { HttpClient } from '@angular/common/http';
import { Injectable, NgZone, EventEmitter } from '@angular/core';

import { Item } from './item';
import { ObjService } from '../api/obj.service';

@Injectable({
  providedIn: 'root'
})
export class ItemService extends ObjService<Item>{
  itemsReload: EventEmitter<boolean>
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
  }

  itemsReset() {
    if (this.itemsReload) {
      this.itemsReload.emit(true)
    }
  }

  path() { return '/api/items' }
}
