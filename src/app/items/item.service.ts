import { HttpClient } from '@angular/common/http';
import { Injectable, NgZone } from '@angular/core';

import { Item } from './item';
import { ObjService } from '../api/obj.service';

@Injectable({
  providedIn: 'root'
})
export class ItemService extends ObjService<Item>{
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
  }

  path() { return '/api/items' }
}
