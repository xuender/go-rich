import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { URL } from './init'
import { Xlsx } from './xlsx';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class XlsxService {
  constructor(
    private http: HttpClient,
  ) {
  }
  get xlsesx(){
    return this.http.get<Xlsx[]>(`${URL}/api/xlsxes`)
  }
  post(x: Xlsx) {
      return this.http.post<Xlsx>(`${URL}/api/xlsxes`, x)
  }
  put(x: Xlsx) {
    return this.http.put<Xlsx>(`${URL}/api/xlsxes/${x.id}`, x)
  }
}
