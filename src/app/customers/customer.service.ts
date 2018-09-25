import { Injectable, NgZone } from '@angular/core'
import { HttpClient } from '@angular/common/http';

import { URL } from '../api/init'
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

  path() { return '/api/customers' }

  download(params: any) {
    return this.http.get(`${URL}/api/customers`, { params: params, responseType: 'blob', })
  }
}
