import { Injectable, NgZone } from '@angular/core'
import { HttpClient } from '@angular/common/http';

import { Customer } from './customer'
import { ObjService } from '../api/obj.service';

@Injectable({
  providedIn: 'root'
})
export class CustomerService extends ObjService<Customer> {
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
  }
  get isPaging() { return true }

  path() { return '/api/customers' }
}
