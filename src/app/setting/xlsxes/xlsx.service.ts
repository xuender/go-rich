import { Injectable, NgZone } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Xlsx } from './xlsx';
import { ObjService } from '../../api/obj.service';

@Injectable({
  providedIn: 'root'
})
export class XlsxService extends ObjService<Xlsx> {
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
  }

  path() { return '/api/xlsxes' }
}
